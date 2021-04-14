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

func testTeamMatchups(t *testing.T) {
	t.Parallel()

	query := TeamMatchups()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTeamMatchupsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMatchup{}
	if err = randomize.Struct(seed, o, teamMatchupDBTypes, true, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
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

	count, err := TeamMatchups().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTeamMatchupsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMatchup{}
	if err = randomize.Struct(seed, o, teamMatchupDBTypes, true, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TeamMatchups().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TeamMatchups().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTeamMatchupsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMatchup{}
	if err = randomize.Struct(seed, o, teamMatchupDBTypes, true, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TeamMatchupSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TeamMatchups().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTeamMatchupsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMatchup{}
	if err = randomize.Struct(seed, o, teamMatchupDBTypes, true, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TeamMatchupExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if TeamMatchup exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TeamMatchupExists to return true, but got false.")
	}
}

func testTeamMatchupsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMatchup{}
	if err = randomize.Struct(seed, o, teamMatchupDBTypes, true, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	teamMatchupFound, err := FindTeamMatchup(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if teamMatchupFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTeamMatchupsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMatchup{}
	if err = randomize.Struct(seed, o, teamMatchupDBTypes, true, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TeamMatchups().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testTeamMatchupsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMatchup{}
	if err = randomize.Struct(seed, o, teamMatchupDBTypes, true, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TeamMatchups().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTeamMatchupsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	teamMatchupOne := &TeamMatchup{}
	teamMatchupTwo := &TeamMatchup{}
	if err = randomize.Struct(seed, teamMatchupOne, teamMatchupDBTypes, false, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}
	if err = randomize.Struct(seed, teamMatchupTwo, teamMatchupDBTypes, false, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = teamMatchupOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = teamMatchupTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TeamMatchups().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTeamMatchupsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	teamMatchupOne := &TeamMatchup{}
	teamMatchupTwo := &TeamMatchup{}
	if err = randomize.Struct(seed, teamMatchupOne, teamMatchupDBTypes, false, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}
	if err = randomize.Struct(seed, teamMatchupTwo, teamMatchupDBTypes, false, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = teamMatchupOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = teamMatchupTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TeamMatchups().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func teamMatchupBeforeInsertHook(e boil.Executor, o *TeamMatchup) error {
	*o = TeamMatchup{}
	return nil
}

func teamMatchupAfterInsertHook(e boil.Executor, o *TeamMatchup) error {
	*o = TeamMatchup{}
	return nil
}

func teamMatchupAfterSelectHook(e boil.Executor, o *TeamMatchup) error {
	*o = TeamMatchup{}
	return nil
}

func teamMatchupBeforeUpdateHook(e boil.Executor, o *TeamMatchup) error {
	*o = TeamMatchup{}
	return nil
}

func teamMatchupAfterUpdateHook(e boil.Executor, o *TeamMatchup) error {
	*o = TeamMatchup{}
	return nil
}

func teamMatchupBeforeDeleteHook(e boil.Executor, o *TeamMatchup) error {
	*o = TeamMatchup{}
	return nil
}

func teamMatchupAfterDeleteHook(e boil.Executor, o *TeamMatchup) error {
	*o = TeamMatchup{}
	return nil
}

func teamMatchupBeforeUpsertHook(e boil.Executor, o *TeamMatchup) error {
	*o = TeamMatchup{}
	return nil
}

func teamMatchupAfterUpsertHook(e boil.Executor, o *TeamMatchup) error {
	*o = TeamMatchup{}
	return nil
}

func testTeamMatchupsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &TeamMatchup{}
	o := &TeamMatchup{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, teamMatchupDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TeamMatchup object: %s", err)
	}

	AddTeamMatchupHook(boil.BeforeInsertHook, teamMatchupBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	teamMatchupBeforeInsertHooks = []TeamMatchupHook{}

	AddTeamMatchupHook(boil.AfterInsertHook, teamMatchupAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	teamMatchupAfterInsertHooks = []TeamMatchupHook{}

	AddTeamMatchupHook(boil.AfterSelectHook, teamMatchupAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	teamMatchupAfterSelectHooks = []TeamMatchupHook{}

	AddTeamMatchupHook(boil.BeforeUpdateHook, teamMatchupBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	teamMatchupBeforeUpdateHooks = []TeamMatchupHook{}

	AddTeamMatchupHook(boil.AfterUpdateHook, teamMatchupAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	teamMatchupAfterUpdateHooks = []TeamMatchupHook{}

	AddTeamMatchupHook(boil.BeforeDeleteHook, teamMatchupBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	teamMatchupBeforeDeleteHooks = []TeamMatchupHook{}

	AddTeamMatchupHook(boil.AfterDeleteHook, teamMatchupAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	teamMatchupAfterDeleteHooks = []TeamMatchupHook{}

	AddTeamMatchupHook(boil.BeforeUpsertHook, teamMatchupBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	teamMatchupBeforeUpsertHooks = []TeamMatchupHook{}

	AddTeamMatchupHook(boil.AfterUpsertHook, teamMatchupAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	teamMatchupAfterUpsertHooks = []TeamMatchupHook{}
}

func testTeamMatchupsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMatchup{}
	if err = randomize.Struct(seed, o, teamMatchupDBTypes, true, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TeamMatchups().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTeamMatchupsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMatchup{}
	if err = randomize.Struct(seed, o, teamMatchupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(teamMatchupColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TeamMatchups().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTeamMatchupToOneRoundUsingRound(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local TeamMatchup
	var foreign Round

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, teamMatchupDBTypes, true, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
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

	slice := TeamMatchupSlice{&local}
	if err = local.L.LoadRound(tx, false, (*[]*TeamMatchup)(&slice), nil); err != nil {
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

func testTeamMatchupToOneTeamUsingHomeTeamTeam(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local TeamMatchup
	var foreign Team

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, teamMatchupDBTypes, true, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, teamDBTypes, false, teamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Team struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.HomeTeam, foreign.ID)
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.HomeTeamTeam().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TeamMatchupSlice{&local}
	if err = local.L.LoadHomeTeamTeam(tx, false, (*[]*TeamMatchup)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.HomeTeamTeam == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.HomeTeamTeam = nil
	if err = local.L.LoadHomeTeamTeam(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.HomeTeamTeam == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTeamMatchupToOneTeamUsingAwayTeamTeam(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local TeamMatchup
	var foreign Team

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, teamMatchupDBTypes, true, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, teamDBTypes, false, teamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Team struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.AwayTeam, foreign.ID)
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.AwayTeamTeam().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TeamMatchupSlice{&local}
	if err = local.L.LoadAwayTeamTeam(tx, false, (*[]*TeamMatchup)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.AwayTeamTeam == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.AwayTeamTeam = nil
	if err = local.L.LoadAwayTeamTeam(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.AwayTeamTeam == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTeamMatchupToOneTeamUsingWinnerTeam(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local TeamMatchup
	var foreign Team

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, teamMatchupDBTypes, true, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, teamDBTypes, false, teamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Team struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.Winner, foreign.ID)
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.WinnerTeam().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TeamMatchupSlice{&local}
	if err = local.L.LoadWinnerTeam(tx, false, (*[]*TeamMatchup)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.WinnerTeam == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.WinnerTeam = nil
	if err = local.L.LoadWinnerTeam(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.WinnerTeam == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTeamMatchupToOneSetOpRoundUsingRound(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a TeamMatchup
	var b, c Round

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, teamMatchupDBTypes, false, strmangle.SetComplement(teamMatchupPrimaryKeyColumns, teamMatchupColumnsWithoutDefault)...); err != nil {
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

		if x.R.RoundTeamMatchups[0] != &a {
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

func testTeamMatchupToOneRemoveOpRoundUsingRound(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a TeamMatchup
	var b Round

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, teamMatchupDBTypes, false, strmangle.SetComplement(teamMatchupPrimaryKeyColumns, teamMatchupColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.RoundTeamMatchups) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testTeamMatchupToOneSetOpTeamUsingHomeTeamTeam(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a TeamMatchup
	var b, c Team

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, teamMatchupDBTypes, false, strmangle.SetComplement(teamMatchupPrimaryKeyColumns, teamMatchupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, teamDBTypes, false, strmangle.SetComplement(teamPrimaryKeyColumns, teamColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, teamDBTypes, false, strmangle.SetComplement(teamPrimaryKeyColumns, teamColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Team{&b, &c} {
		err = a.SetHomeTeamTeam(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.HomeTeamTeam != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.HomeTeamTeamMatchups[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.HomeTeam, x.ID) {
			t.Error("foreign key was wrong value", a.HomeTeam)
		}

		zero := reflect.Zero(reflect.TypeOf(a.HomeTeam))
		reflect.Indirect(reflect.ValueOf(&a.HomeTeam)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.HomeTeam, x.ID) {
			t.Error("foreign key was wrong value", a.HomeTeam, x.ID)
		}
	}
}

func testTeamMatchupToOneRemoveOpTeamUsingHomeTeamTeam(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a TeamMatchup
	var b Team

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, teamMatchupDBTypes, false, strmangle.SetComplement(teamMatchupPrimaryKeyColumns, teamMatchupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, teamDBTypes, false, strmangle.SetComplement(teamPrimaryKeyColumns, teamColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetHomeTeamTeam(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveHomeTeamTeam(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.HomeTeamTeam().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.HomeTeamTeam != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.HomeTeam) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.HomeTeamTeamMatchups) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testTeamMatchupToOneSetOpTeamUsingAwayTeamTeam(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a TeamMatchup
	var b, c Team

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, teamMatchupDBTypes, false, strmangle.SetComplement(teamMatchupPrimaryKeyColumns, teamMatchupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, teamDBTypes, false, strmangle.SetComplement(teamPrimaryKeyColumns, teamColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, teamDBTypes, false, strmangle.SetComplement(teamPrimaryKeyColumns, teamColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Team{&b, &c} {
		err = a.SetAwayTeamTeam(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.AwayTeamTeam != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AwayTeamTeamMatchups[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.AwayTeam, x.ID) {
			t.Error("foreign key was wrong value", a.AwayTeam)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AwayTeam))
		reflect.Indirect(reflect.ValueOf(&a.AwayTeam)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.AwayTeam, x.ID) {
			t.Error("foreign key was wrong value", a.AwayTeam, x.ID)
		}
	}
}

func testTeamMatchupToOneRemoveOpTeamUsingAwayTeamTeam(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a TeamMatchup
	var b Team

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, teamMatchupDBTypes, false, strmangle.SetComplement(teamMatchupPrimaryKeyColumns, teamMatchupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, teamDBTypes, false, strmangle.SetComplement(teamPrimaryKeyColumns, teamColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetAwayTeamTeam(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveAwayTeamTeam(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.AwayTeamTeam().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.AwayTeamTeam != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.AwayTeam) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.AwayTeamTeamMatchups) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testTeamMatchupToOneSetOpTeamUsingWinnerTeam(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a TeamMatchup
	var b, c Team

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, teamMatchupDBTypes, false, strmangle.SetComplement(teamMatchupPrimaryKeyColumns, teamMatchupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, teamDBTypes, false, strmangle.SetComplement(teamPrimaryKeyColumns, teamColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, teamDBTypes, false, strmangle.SetComplement(teamPrimaryKeyColumns, teamColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Team{&b, &c} {
		err = a.SetWinnerTeam(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.WinnerTeam != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.WinnerTeamMatchups[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.Winner, x.ID) {
			t.Error("foreign key was wrong value", a.Winner)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Winner))
		reflect.Indirect(reflect.ValueOf(&a.Winner)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.Winner, x.ID) {
			t.Error("foreign key was wrong value", a.Winner, x.ID)
		}
	}
}

func testTeamMatchupToOneRemoveOpTeamUsingWinnerTeam(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a TeamMatchup
	var b Team

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, teamMatchupDBTypes, false, strmangle.SetComplement(teamMatchupPrimaryKeyColumns, teamMatchupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, teamDBTypes, false, strmangle.SetComplement(teamPrimaryKeyColumns, teamColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetWinnerTeam(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveWinnerTeam(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.WinnerTeam().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.WinnerTeam != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.Winner) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.WinnerTeamMatchups) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testTeamMatchupsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMatchup{}
	if err = randomize.Struct(seed, o, teamMatchupDBTypes, true, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
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

func testTeamMatchupsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMatchup{}
	if err = randomize.Struct(seed, o, teamMatchupDBTypes, true, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TeamMatchupSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testTeamMatchupsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamMatchup{}
	if err = randomize.Struct(seed, o, teamMatchupDBTypes, true, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TeamMatchups().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	teamMatchupDBTypes = map[string]string{`ID`: `int`, `RoundID`: `int`, `HomeTeam`: `int`, `AwayTeam`: `int`, `HomeTeamScore`: `int`, `AwayTeamScore`: `int`, `Winner`: `int`, `Status`: `int`, `CreatedAt`: `timestamp`, `DeletedAt`: `timestamp`}
	_                  = bytes.MinRead
)

func testTeamMatchupsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(teamMatchupPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(teamMatchupAllColumns) == len(teamMatchupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TeamMatchup{}
	if err = randomize.Struct(seed, o, teamMatchupDBTypes, true, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TeamMatchups().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, teamMatchupDBTypes, true, teamMatchupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTeamMatchupsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(teamMatchupAllColumns) == len(teamMatchupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TeamMatchup{}
	if err = randomize.Struct(seed, o, teamMatchupDBTypes, true, teamMatchupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TeamMatchups().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, teamMatchupDBTypes, true, teamMatchupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(teamMatchupAllColumns, teamMatchupPrimaryKeyColumns) {
		fields = teamMatchupAllColumns
	} else {
		fields = strmangle.SetComplement(
			teamMatchupAllColumns,
			teamMatchupPrimaryKeyColumns,
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

	slice := TeamMatchupSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTeamMatchupsUpsert(t *testing.T) {
	t.Parallel()

	if len(teamMatchupAllColumns) == len(teamMatchupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLTeamMatchupUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := TeamMatchup{}
	if err = randomize.Struct(seed, &o, teamMatchupDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TeamMatchup: %s", err)
	}

	count, err := TeamMatchups().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, teamMatchupDBTypes, false, teamMatchupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TeamMatchup struct: %s", err)
	}

	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TeamMatchup: %s", err)
	}

	count, err = TeamMatchups().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}