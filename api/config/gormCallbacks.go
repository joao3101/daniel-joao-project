package config

import (
	"errors"
	"reflect"

	uuid "github.com/nu7hatch/gouuid"

	// Gorm
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

var (
	ErrHardDelete = errors.New("hard delete is disabled")
)

func RegisterCallbacks(db *gorm.DB) {
	db.Callback().Create().Before("gorm:create").Register("create_uuid", UUIDCallback)
	db.Callback().Delete().Replace("gorm:delete", DeleteCallback)
	db.Callback().Query().Before("gorm:query").Register("filter_deleted", FilterDeletedCallback)
	db.Callback().Update().Before("gorm:update").Register("before_update", BeforeUpdateCallback)
}

//set automatically UUID if is BLANK
func UUIDCallback(db *gorm.DB) {
	field := db.Statement.Schema.LookUpField("uuid")

	if field == nil {
		return
	}

	fieldValue, isZero := field.ValueOf(db.Statement.ReflectValue)
	fieldValueStr := fieldValue.(string)
	if isZero || fieldValueStr == "" {
		uuid, _ := uuid.NewV4()

		field.Set(db.Statement.ReflectValue, uuid.String())
	}
}

func DeleteCallback(db *gorm.DB) {
	if db.Error != nil {
		return
	}

	field := db.Statement.Schema.LookUpField("DeletedAt")

	if field == nil {
		db.AddError(ErrHardDelete)
		return
	}

	curTime := db.Statement.DB.NowFunc()
	db.Statement.AddClause(
		clause.Set{{Column: clause.Column{Name: field.DBName}, Value: curTime.Unix()}},
	)
	db.Statement.SetColumn(field.DBName, curTime, true)

	if db.Statement.Schema != nil {
		_, queryValues := schema.GetIdentityFieldValuesMap(db.Statement.ReflectValue, db.Statement.Schema.PrimaryFields)
		column, values := schema.ToQueryValues(db.Statement.Table, db.Statement.Schema.PrimaryFieldDBNames, queryValues)

		if len(values) > 0 {
			db.Statement.AddClause(clause.Where{Exprs: []clause.Expression{clause.IN{Column: column, Values: values}}})
		}

		if db.Statement.ReflectValue.CanAddr() && db.Statement.Dest != db.Statement.Model && db.Statement.Model != nil {
			_, queryValues = schema.GetIdentityFieldValuesMap(reflect.ValueOf(db.Statement.Model), db.Statement.Schema.PrimaryFields)
			column, values = schema.ToQueryValues(db.Statement.Table, db.Statement.Schema.PrimaryFieldDBNames, queryValues)

			if len(values) > 0 {
				db.Statement.AddClause(clause.Where{Exprs: []clause.Expression{clause.IN{Column: column, Values: values}}})
			}
		}
	}

	db.Statement.AddClauseIfNotExists(clause.Update{})
	db.Statement.Build("UPDATE", "SET", "WHERE")

	if _, ok := db.Statement.Clauses["WHERE"]; !db.AllowGlobalUpdate && !ok && db.Error == nil {
		db.AddError(gorm.ErrMissingWhereClause)
		return
	}

	if !db.DryRun && db.Error == nil {
		result, err := db.Statement.ConnPool.ExecContext(db.Statement.Context, db.Statement.SQL.String(), db.Statement.Vars...)

		if err == nil {
			db.RowsAffected, _ = result.RowsAffected()
		} else {
			db.AddError(err)
		}
	}
}

func FilterDeletedCallback(db *gorm.DB) {
	if db.Error != nil {
		return
	}

	if db.Statement.Schema != nil && !db.Statement.Unscoped {
		field := db.Statement.Schema.LookUpField("DeletedAt")

		if field == nil {
			return
		}

		db.Statement.AddClause(clause.Where{Exprs: []clause.Expression{
			clause.Eq{Column: clause.Column{Table: clause.CurrentTable, Name: field.DBName}, Value: nil},
		}})
	}
}

func BeforeUpdateCallback(db *gorm.DB) {
	if db.Error != nil {
		return
	}

	if db.Statement.SQL.String() != "" {
		return
	}

	if db.Statement.Schema == nil {
		return
	}

	primaryFields := db.Statement.Schema.PrimaryFields

	destReflectValue := reflect.ValueOf(db.Statement.Dest)
	for destReflectValue.Kind() == reflect.Ptr {
		destReflectValue = destReflectValue.Elem()
	}

	for _, primaryField := range primaryFields {
		primaryField.Updatable = false

		switch destReflectValue.Kind() {
		case reflect.Struct:
			fieldValue, isZero := primaryField.ValueOf(destReflectValue)
			if !isZero {
				db.Statement.AddClause(clause.Where{Exprs: []clause.Expression{clause.Eq{Column: primaryField.DBName, Value: fieldValue}}})
			}
		}
	}

	updatedTimeField := db.Statement.Schema.LookUpField("ModifiedAt")
	if updatedTimeField != nil {
		updatedTimeField.AutoUpdateTime = schema.UnixSecond
	}
}
