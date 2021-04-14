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

func testTeamRoundPlayers(t *testing.T) {
	t.Parallel()

	query := TeamRoundPlayers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTeamRoundPlayersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamRoundPlayer{}
	if err = randomize.Struct(seed, o, teamRoundPlayerDBTypes, true, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
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

	count, err := TeamRoundPlayers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTeamRoundPlayersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamRoundPlayer{}
	if err = randomize.Struct(seed, o, teamRoundPlayerDBTypes, true, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TeamRoundPlayers().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TeamRoundPlayers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTeamRoundPlayersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamRoundPlayer{}
	if err = randomize.Struct(seed, o, teamRoundPlayerDBTypes, true, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TeamRoundPlayerSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TeamRoundPlayers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTeamRoundPlayersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamRoundPlayer{}
	if err = randomize.Struct(seed, o, teamRoundPlayerDBTypes, true, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TeamRoundPlayerExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if TeamRoundPlayer exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TeamRoundPlayerExists to return true, but got false.")
	}
}

func testTeamRoundPlayersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamRoundPlayer{}
	if err = randomize.Struct(seed, o, teamRoundPlayerDBTypes, true, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	teamRoundPlayerFound, err := FindTeamRoundPlayer(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if teamRoundPlayerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTeamRoundPlayersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamRoundPlayer{}
	if err = randomize.Struct(seed, o, teamRoundPlayerDBTypes, true, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TeamRoundPlayers().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testTeamRoundPlayersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamRoundPlayer{}
	if err = randomize.Struct(seed, o, teamRoundPlayerDBTypes, true, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TeamRoundPlayers().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTeamRoundPlayersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	teamRoundPlayerOne := &TeamRoundPlayer{}
	teamRoundPlayerTwo := &TeamRoundPlayer{}
	if err = randomize.Struct(seed, teamRoundPlayerOne, teamRoundPlayerDBTypes, false, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}
	if err = randomize.Struct(seed, teamRoundPlayerTwo, teamRoundPlayerDBTypes, false, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = teamRoundPlayerOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = teamRoundPlayerTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TeamRoundPlayers().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTeamRoundPlayersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	teamRoundPlayerOne := &TeamRoundPlayer{}
	teamRoundPlayerTwo := &TeamRoundPlayer{}
	if err = randomize.Struct(seed, teamRoundPlayerOne, teamRoundPlayerDBTypes, false, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}
	if err = randomize.Struct(seed, teamRoundPlayerTwo, teamRoundPlayerDBTypes, false, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = teamRoundPlayerOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = teamRoundPlayerTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TeamRoundPlayers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func teamRoundPlayerBeforeInsertHook(e boil.Executor, o *TeamRoundPlayer) error {
	*o = TeamRoundPlayer{}
	return nil
}

func teamRoundPlayerAfterInsertHook(e boil.Executor, o *TeamRoundPlayer) error {
	*o = TeamRoundPlayer{}
	return nil
}

func teamRoundPlayerAfterSelectHook(e boil.Executor, o *TeamRoundPlayer) error {
	*o = TeamRoundPlayer{}
	return nil
}

func teamRoundPlayerBeforeUpdateHook(e boil.Executor, o *TeamRoundPlayer) error {
	*o = TeamRoundPlayer{}
	return nil
}

func teamRoundPlayerAfterUpdateHook(e boil.Executor, o *TeamRoundPlayer) error {
	*o = TeamRoundPlayer{}
	return nil
}

func teamRoundPlayerBeforeDeleteHook(e boil.Executor, o *TeamRoundPlayer) error {
	*o = TeamRoundPlayer{}
	return nil
}

func teamRoundPlayerAfterDeleteHook(e boil.Executor, o *TeamRoundPlayer) error {
	*o = TeamRoundPlayer{}
	return nil
}

func teamRoundPlayerBeforeUpsertHook(e boil.Executor, o *TeamRoundPlayer) error {
	*o = TeamRoundPlayer{}
	return nil
}

func teamRoundPlayerAfterUpsertHook(e boil.Executor, o *TeamRoundPlayer) error {
	*o = TeamRoundPlayer{}
	return nil
}

func testTeamRoundPlayersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &TeamRoundPlayer{}
	o := &TeamRoundPlayer{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, teamRoundPlayerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer object: %s", err)
	}

	AddTeamRoundPlayerHook(boil.BeforeInsertHook, teamRoundPlayerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	teamRoundPlayerBeforeInsertHooks = []TeamRoundPlayerHook{}

	AddTeamRoundPlayerHook(boil.AfterInsertHook, teamRoundPlayerAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	teamRoundPlayerAfterInsertHooks = []TeamRoundPlayerHook{}

	AddTeamRoundPlayerHook(boil.AfterSelectHook, teamRoundPlayerAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	teamRoundPlayerAfterSelectHooks = []TeamRoundPlayerHook{}

	AddTeamRoundPlayerHook(boil.BeforeUpdateHook, teamRoundPlayerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	teamRoundPlayerBeforeUpdateHooks = []TeamRoundPlayerHook{}

	AddTeamRoundPlayerHook(boil.AfterUpdateHook, teamRoundPlayerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	teamRoundPlayerAfterUpdateHooks = []TeamRoundPlayerHook{}

	AddTeamRoundPlayerHook(boil.BeforeDeleteHook, teamRoundPlayerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	teamRoundPlayerBeforeDeleteHooks = []TeamRoundPlayerHook{}

	AddTeamRoundPlayerHook(boil.AfterDeleteHook, teamRoundPlayerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	teamRoundPlayerAfterDeleteHooks = []TeamRoundPlayerHook{}

	AddTeamRoundPlayerHook(boil.BeforeUpsertHook, teamRoundPlayerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	teamRoundPlayerBeforeUpsertHooks = []TeamRoundPlayerHook{}

	AddTeamRoundPlayerHook(boil.AfterUpsertHook, teamRoundPlayerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	teamRoundPlayerAfterUpsertHooks = []TeamRoundPlayerHook{}
}

func testTeamRoundPlayersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamRoundPlayer{}
	if err = randomize.Struct(seed, o, teamRoundPlayerDBTypes, true, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TeamRoundPlayers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTeamRoundPlayersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamRoundPlayer{}
	if err = randomize.Struct(seed, o, teamRoundPlayerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(teamRoundPlayerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TeamRoundPlayers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTeamRoundPlayerToOneTeamRoundUsingTeamRound(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local TeamRoundPlayer
	var foreign TeamRound

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, teamRoundPlayerDBTypes, true, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, teamRoundDBTypes, false, teamRoundColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRound struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.TeamRoundID, foreign.ID)
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.TeamRound().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TeamRoundPlayerSlice{&local}
	if err = local.L.LoadTeamRound(tx, false, (*[]*TeamRoundPlayer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.TeamRound == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TeamRound = nil
	if err = local.L.LoadTeamRound(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.TeamRound == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTeamRoundPlayerToOneTeamPlayerUsingTeamPlayer(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local TeamRoundPlayer
	var foreign TeamPlayer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, teamRoundPlayerDBTypes, true, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, teamPlayerDBTypes, false, teamPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamPlayer struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.TeamPlayerID, foreign.ID)
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.TeamPlayer().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TeamRoundPlayerSlice{&local}
	if err = local.L.LoadTeamPlayer(tx, false, (*[]*TeamRoundPlayer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.TeamPlayer == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TeamPlayer = nil
	if err = local.L.LoadTeamPlayer(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.TeamPlayer == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTeamRoundPlayerToOneSetOpTeamRoundUsingTeamRound(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a TeamRoundPlayer
	var b, c TeamRound

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, teamRoundPlayerDBTypes, false, strmangle.SetComplement(teamRoundPlayerPrimaryKeyColumns, teamRoundPlayerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, teamRoundDBTypes, false, strmangle.SetComplement(teamRoundPrimaryKeyColumns, teamRoundColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, teamRoundDBTypes, false, strmangle.SetComplement(teamRoundPrimaryKeyColumns, teamRoundColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*TeamRound{&b, &c} {
		err = a.SetTeamRound(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TeamRound != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TeamRoundTeamRoundPlayers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.TeamRoundID, x.ID) {
			t.Error("foreign key was wrong value", a.TeamRoundID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TeamRoundID))
		reflect.Indirect(reflect.ValueOf(&a.TeamRoundID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.TeamRoundID, x.ID) {
			t.Error("foreign key was wrong value", a.TeamRoundID, x.ID)
		}
	}
}

func testTeamRoundPlayerToOneRemoveOpTeamRoundUsingTeamRound(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a TeamRoundPlayer
	var b TeamRound

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, teamRoundPlayerDBTypes, false, strmangle.SetComplement(teamRoundPlayerPrimaryKeyColumns, teamRoundPlayerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, teamRoundDBTypes, false, strmangle.SetComplement(teamRoundPrimaryKeyColumns, teamRoundColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetTeamRound(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveTeamRound(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.TeamRound().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.TeamRound != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.TeamRoundID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.TeamRoundTeamRoundPlayers) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testTeamRoundPlayerToOneSetOpTeamPlayerUsingTeamPlayer(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a TeamRoundPlayer
	var b, c TeamPlayer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, teamRoundPlayerDBTypes, false, strmangle.SetComplement(teamRoundPlayerPrimaryKeyColumns, teamRoundPlayerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, teamPlayerDBTypes, false, strmangle.SetComplement(teamPlayerPrimaryKeyColumns, teamPlayerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, teamPlayerDBTypes, false, strmangle.SetComplement(teamPlayerPrimaryKeyColumns, teamPlayerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*TeamPlayer{&b, &c} {
		err = a.SetTeamPlayer(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TeamPlayer != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TeamPlayerTeamRoundPlayers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.TeamPlayerID, x.ID) {
			t.Error("foreign key was wrong value", a.TeamPlayerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TeamPlayerID))
		reflect.Indirect(reflect.ValueOf(&a.TeamPlayerID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.TeamPlayerID, x.ID) {
			t.Error("foreign key was wrong value", a.TeamPlayerID, x.ID)
		}
	}
}

func testTeamRoundPlayerToOneRemoveOpTeamPlayerUsingTeamPlayer(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a TeamRoundPlayer
	var b TeamPlayer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, teamRoundPlayerDBTypes, false, strmangle.SetComplement(teamRoundPlayerPrimaryKeyColumns, teamRoundPlayerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, teamPlayerDBTypes, false, strmangle.SetComplement(teamPlayerPrimaryKeyColumns, teamPlayerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetTeamPlayer(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveTeamPlayer(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.TeamPlayer().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.TeamPlayer != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.TeamPlayerID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.TeamPlayerTeamRoundPlayers) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testTeamRoundPlayersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamRoundPlayer{}
	if err = randomize.Struct(seed, o, teamRoundPlayerDBTypes, true, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
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

func testTeamRoundPlayersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamRoundPlayer{}
	if err = randomize.Struct(seed, o, teamRoundPlayerDBTypes, true, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TeamRoundPlayerSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testTeamRoundPlayersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TeamRoundPlayer{}
	if err = randomize.Struct(seed, o, teamRoundPlayerDBTypes, true, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TeamRoundPlayers().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	teamRoundPlayerDBTypes = map[string]string{`ID`: `int`, `TeamRoundID`: `int`, `TeamPlayerID`: `int`, `Status`: `int`, `CreatedAt`: `timestamp`, `DeletedAt`: `timestamp`}
	_                      = bytes.MinRead
)

func testTeamRoundPlayersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(teamRoundPlayerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(teamRoundPlayerAllColumns) == len(teamRoundPlayerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TeamRoundPlayer{}
	if err = randomize.Struct(seed, o, teamRoundPlayerDBTypes, true, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TeamRoundPlayers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, teamRoundPlayerDBTypes, true, teamRoundPlayerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTeamRoundPlayersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(teamRoundPlayerAllColumns) == len(teamRoundPlayerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TeamRoundPlayer{}
	if err = randomize.Struct(seed, o, teamRoundPlayerDBTypes, true, teamRoundPlayerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TeamRoundPlayers().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, teamRoundPlayerDBTypes, true, teamRoundPlayerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(teamRoundPlayerAllColumns, teamRoundPlayerPrimaryKeyColumns) {
		fields = teamRoundPlayerAllColumns
	} else {
		fields = strmangle.SetComplement(
			teamRoundPlayerAllColumns,
			teamRoundPlayerPrimaryKeyColumns,
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

	slice := TeamRoundPlayerSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTeamRoundPlayersUpsert(t *testing.T) {
	t.Parallel()

	if len(teamRoundPlayerAllColumns) == len(teamRoundPlayerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLTeamRoundPlayerUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := TeamRoundPlayer{}
	if err = randomize.Struct(seed, &o, teamRoundPlayerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TeamRoundPlayer: %s", err)
	}

	count, err := TeamRoundPlayers().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, teamRoundPlayerDBTypes, false, teamRoundPlayerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TeamRoundPlayer struct: %s", err)
	}

	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TeamRoundPlayer: %s", err)
	}

	count, err = TeamRoundPlayers().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}