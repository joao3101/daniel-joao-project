// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlboiler

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPlayerRounds(t *testing.T) {
	t.Parallel()

	query := PlayerRounds()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPlayerRoundsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayerRound{}
	if err = randomize.Struct(seed, o, playerRoundDBTypes, true, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PlayerRounds().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPlayerRoundsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayerRound{}
	if err = randomize.Struct(seed, o, playerRoundDBTypes, true, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PlayerRounds().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PlayerRounds().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPlayerRoundsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayerRound{}
	if err = randomize.Struct(seed, o, playerRoundDBTypes, true, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PlayerRoundSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PlayerRounds().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPlayerRoundsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayerRound{}
	if err = randomize.Struct(seed, o, playerRoundDBTypes, true, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PlayerRoundExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if PlayerRound exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PlayerRoundExists to return true, but got false.")
	}
}

func testPlayerRoundsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayerRound{}
	if err = randomize.Struct(seed, o, playerRoundDBTypes, true, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	playerRoundFound, err := FindPlayerRound(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if playerRoundFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPlayerRoundsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayerRound{}
	if err = randomize.Struct(seed, o, playerRoundDBTypes, true, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PlayerRounds().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testPlayerRoundsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayerRound{}
	if err = randomize.Struct(seed, o, playerRoundDBTypes, true, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PlayerRounds().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPlayerRoundsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	playerRoundOne := &PlayerRound{}
	playerRoundTwo := &PlayerRound{}
	if err = randomize.Struct(seed, playerRoundOne, playerRoundDBTypes, false, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}
	if err = randomize.Struct(seed, playerRoundTwo, playerRoundDBTypes, false, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = playerRoundOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = playerRoundTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PlayerRounds().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPlayerRoundsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	playerRoundOne := &PlayerRound{}
	playerRoundTwo := &PlayerRound{}
	if err = randomize.Struct(seed, playerRoundOne, playerRoundDBTypes, false, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}
	if err = randomize.Struct(seed, playerRoundTwo, playerRoundDBTypes, false, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = playerRoundOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = playerRoundTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PlayerRounds().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func playerRoundBeforeInsertHook(e boil.Executor, o *PlayerRound) error {
	*o = PlayerRound{}
	return nil
}

func playerRoundAfterInsertHook(e boil.Executor, o *PlayerRound) error {
	*o = PlayerRound{}
	return nil
}

func playerRoundAfterSelectHook(e boil.Executor, o *PlayerRound) error {
	*o = PlayerRound{}
	return nil
}

func playerRoundBeforeUpdateHook(e boil.Executor, o *PlayerRound) error {
	*o = PlayerRound{}
	return nil
}

func playerRoundAfterUpdateHook(e boil.Executor, o *PlayerRound) error {
	*o = PlayerRound{}
	return nil
}

func playerRoundBeforeDeleteHook(e boil.Executor, o *PlayerRound) error {
	*o = PlayerRound{}
	return nil
}

func playerRoundAfterDeleteHook(e boil.Executor, o *PlayerRound) error {
	*o = PlayerRound{}
	return nil
}

func playerRoundBeforeUpsertHook(e boil.Executor, o *PlayerRound) error {
	*o = PlayerRound{}
	return nil
}

func playerRoundAfterUpsertHook(e boil.Executor, o *PlayerRound) error {
	*o = PlayerRound{}
	return nil
}

func testPlayerRoundsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &PlayerRound{}
	o := &PlayerRound{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, playerRoundDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PlayerRound object: %s", err)
	}

	AddPlayerRoundHook(boil.BeforeInsertHook, playerRoundBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	playerRoundBeforeInsertHooks = []PlayerRoundHook{}

	AddPlayerRoundHook(boil.AfterInsertHook, playerRoundAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	playerRoundAfterInsertHooks = []PlayerRoundHook{}

	AddPlayerRoundHook(boil.AfterSelectHook, playerRoundAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	playerRoundAfterSelectHooks = []PlayerRoundHook{}

	AddPlayerRoundHook(boil.BeforeUpdateHook, playerRoundBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	playerRoundBeforeUpdateHooks = []PlayerRoundHook{}

	AddPlayerRoundHook(boil.AfterUpdateHook, playerRoundAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	playerRoundAfterUpdateHooks = []PlayerRoundHook{}

	AddPlayerRoundHook(boil.BeforeDeleteHook, playerRoundBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	playerRoundBeforeDeleteHooks = []PlayerRoundHook{}

	AddPlayerRoundHook(boil.AfterDeleteHook, playerRoundAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	playerRoundAfterDeleteHooks = []PlayerRoundHook{}

	AddPlayerRoundHook(boil.BeforeUpsertHook, playerRoundBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	playerRoundBeforeUpsertHooks = []PlayerRoundHook{}

	AddPlayerRoundHook(boil.AfterUpsertHook, playerRoundAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	playerRoundAfterUpsertHooks = []PlayerRoundHook{}
}

func testPlayerRoundsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayerRound{}
	if err = randomize.Struct(seed, o, playerRoundDBTypes, true, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PlayerRounds().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPlayerRoundsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayerRound{}
	if err = randomize.Struct(seed, o, playerRoundDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(playerRoundColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PlayerRounds().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPlayerRoundToOnePlayerUsingPlayer(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local PlayerRound
	var foreign Player

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, playerRoundDBTypes, true, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, playerDBTypes, false, playerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Player struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.PlayerID, foreign.ID)
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Player().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := PlayerRoundSlice{&local}
	if err = local.L.LoadPlayer(tx, false, (*[]*PlayerRound)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Player == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Player = nil
	if err = local.L.LoadPlayer(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Player == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPlayerRoundToOneRoundUsingRound(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local PlayerRound
	var foreign Round

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, playerRoundDBTypes, true, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, roundDBTypes, false, roundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Round struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.RoundID, foreign.ID)
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Round().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := PlayerRoundSlice{&local}
	if err = local.L.LoadRound(tx, false, (*[]*PlayerRound)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Round == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Round = nil
	if err = local.L.LoadRound(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Round == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPlayerRoundToOneSetOpPlayerUsingPlayer(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a PlayerRound
	var b, c Player

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, playerRoundDBTypes, false, strmangle.SetComplement(playerRoundPrimaryKeyColumns, playerRoundColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, playerDBTypes, false, strmangle.SetComplement(playerPrimaryKeyColumns, playerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, playerDBTypes, false, strmangle.SetComplement(playerPrimaryKeyColumns, playerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Player{&b, &c} {
		err = a.SetPlayer(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Player != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PlayerPlayerRounds[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.PlayerID, x.ID) {
			t.Error("foreign key was wrong value", a.PlayerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PlayerID))
		reflect.Indirect(reflect.ValueOf(&a.PlayerID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.PlayerID, x.ID) {
			t.Error("foreign key was wrong value", a.PlayerID, x.ID)
		}
	}
}

func testPlayerRoundToOneRemoveOpPlayerUsingPlayer(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a PlayerRound
	var b Player

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, playerRoundDBTypes, false, strmangle.SetComplement(playerRoundPrimaryKeyColumns, playerRoundColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, playerDBTypes, false, strmangle.SetComplement(playerPrimaryKeyColumns, playerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetPlayer(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemovePlayer(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Player().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Player != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.PlayerID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.PlayerPlayerRounds) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPlayerRoundToOneSetOpRoundUsingRound(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a PlayerRound
	var b, c Round

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, playerRoundDBTypes, false, strmangle.SetComplement(playerRoundPrimaryKeyColumns, playerRoundColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, roundDBTypes, false, strmangle.SetComplement(roundPrimaryKeyColumns, roundColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, roundDBTypes, false, strmangle.SetComplement(roundPrimaryKeyColumns, roundColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Round{&b, &c} {
		err = a.SetRound(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Round != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RoundPlayerRounds[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.RoundID, x.ID) {
			t.Error("foreign key was wrong value", a.RoundID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RoundID))
		reflect.Indirect(reflect.ValueOf(&a.RoundID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.RoundID, x.ID) {
			t.Error("foreign key was wrong value", a.RoundID, x.ID)
		}
	}
}

func testPlayerRoundToOneRemoveOpRoundUsingRound(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a PlayerRound
	var b Round

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, playerRoundDBTypes, false, strmangle.SetComplement(playerRoundPrimaryKeyColumns, playerRoundColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, roundDBTypes, false, strmangle.SetComplement(roundPrimaryKeyColumns, roundColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetRound(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveRound(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Round().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Round != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.RoundID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.RoundPlayerRounds) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPlayerRoundsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayerRound{}
	if err = randomize.Struct(seed, o, playerRoundDBTypes, true, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testPlayerRoundsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayerRound{}
	if err = randomize.Struct(seed, o, playerRoundDBTypes, true, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PlayerRoundSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testPlayerRoundsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PlayerRound{}
	if err = randomize.Struct(seed, o, playerRoundDBTypes, true, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PlayerRounds().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	playerRoundDBTypes = map[string]string{`ID`: `integer`, `PlayerID`: `integer`, `RoundID`: `integer`, `Status`: `integer`, `Score`: `integer`, `CreatedAt`: `timestamp without time zone`, `DeletedAt`: `timestamp without time zone`}
	_                  = bytes.MinRead
)

func testPlayerRoundsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(playerRoundPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(playerRoundAllColumns) == len(playerRoundPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PlayerRound{}
	if err = randomize.Struct(seed, o, playerRoundDBTypes, true, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PlayerRounds().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, playerRoundDBTypes, true, playerRoundPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPlayerRoundsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(playerRoundAllColumns) == len(playerRoundPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PlayerRound{}
	if err = randomize.Struct(seed, o, playerRoundDBTypes, true, playerRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PlayerRounds().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, playerRoundDBTypes, true, playerRoundPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(playerRoundAllColumns, playerRoundPrimaryKeyColumns) {
		fields = playerRoundAllColumns
	} else {
		fields = strmangle.SetComplement(
			playerRoundAllColumns,
			playerRoundPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := PlayerRoundSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPlayerRoundsUpsert(t *testing.T) {
	t.Parallel()

	if len(playerRoundAllColumns) == len(playerRoundPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PlayerRound{}
	if err = randomize.Struct(seed, &o, playerRoundDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PlayerRound: %s", err)
	}

	count, err := PlayerRounds().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, playerRoundDBTypes, false, playerRoundPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PlayerRound struct: %s", err)
	}

	if err = o.Upsert(tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PlayerRound: %s", err)
	}

	count, err = PlayerRounds().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
