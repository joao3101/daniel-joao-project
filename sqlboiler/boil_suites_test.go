// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlboiler

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("ClubMatchups", testClubMatchups)
	t.Run("Clubs", testClubs)
	t.Run("Leagues", testLeagues)
	t.Run("PlayerRounds", testPlayerRounds)
	t.Run("Players", testPlayers)
	t.Run("Rounds", testRounds)
	t.Run("TeamMatchups", testTeamMatchups)
	t.Run("TeamPlayers", testTeamPlayers)
	t.Run("TeamRoundPlayers", testTeamRoundPlayers)
	t.Run("TeamRounds", testTeamRounds)
	t.Run("Teams", testTeams)
	t.Run("TradePlayers", testTradePlayers)
	t.Run("Trades", testTrades)
	t.Run("Users", testUsers)
	t.Run("Waivers", testWaivers)
}

func TestDelete(t *testing.T) {
	t.Run("ClubMatchups", testClubMatchupsDelete)
	t.Run("Clubs", testClubsDelete)
	t.Run("Leagues", testLeaguesDelete)
	t.Run("PlayerRounds", testPlayerRoundsDelete)
	t.Run("Players", testPlayersDelete)
	t.Run("Rounds", testRoundsDelete)
	t.Run("TeamMatchups", testTeamMatchupsDelete)
	t.Run("TeamPlayers", testTeamPlayersDelete)
	t.Run("TeamRoundPlayers", testTeamRoundPlayersDelete)
	t.Run("TeamRounds", testTeamRoundsDelete)
	t.Run("Teams", testTeamsDelete)
	t.Run("TradePlayers", testTradePlayersDelete)
	t.Run("Trades", testTradesDelete)
	t.Run("Users", testUsersDelete)
	t.Run("Waivers", testWaiversDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("ClubMatchups", testClubMatchupsQueryDeleteAll)
	t.Run("Clubs", testClubsQueryDeleteAll)
	t.Run("Leagues", testLeaguesQueryDeleteAll)
	t.Run("PlayerRounds", testPlayerRoundsQueryDeleteAll)
	t.Run("Players", testPlayersQueryDeleteAll)
	t.Run("Rounds", testRoundsQueryDeleteAll)
	t.Run("TeamMatchups", testTeamMatchupsQueryDeleteAll)
	t.Run("TeamPlayers", testTeamPlayersQueryDeleteAll)
	t.Run("TeamRoundPlayers", testTeamRoundPlayersQueryDeleteAll)
	t.Run("TeamRounds", testTeamRoundsQueryDeleteAll)
	t.Run("Teams", testTeamsQueryDeleteAll)
	t.Run("TradePlayers", testTradePlayersQueryDeleteAll)
	t.Run("Trades", testTradesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
	t.Run("Waivers", testWaiversQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("ClubMatchups", testClubMatchupsSliceDeleteAll)
	t.Run("Clubs", testClubsSliceDeleteAll)
	t.Run("Leagues", testLeaguesSliceDeleteAll)
	t.Run("PlayerRounds", testPlayerRoundsSliceDeleteAll)
	t.Run("Players", testPlayersSliceDeleteAll)
	t.Run("Rounds", testRoundsSliceDeleteAll)
	t.Run("TeamMatchups", testTeamMatchupsSliceDeleteAll)
	t.Run("TeamPlayers", testTeamPlayersSliceDeleteAll)
	t.Run("TeamRoundPlayers", testTeamRoundPlayersSliceDeleteAll)
	t.Run("TeamRounds", testTeamRoundsSliceDeleteAll)
	t.Run("Teams", testTeamsSliceDeleteAll)
	t.Run("TradePlayers", testTradePlayersSliceDeleteAll)
	t.Run("Trades", testTradesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
	t.Run("Waivers", testWaiversSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("ClubMatchups", testClubMatchupsExists)
	t.Run("Clubs", testClubsExists)
	t.Run("Leagues", testLeaguesExists)
	t.Run("PlayerRounds", testPlayerRoundsExists)
	t.Run("Players", testPlayersExists)
	t.Run("Rounds", testRoundsExists)
	t.Run("TeamMatchups", testTeamMatchupsExists)
	t.Run("TeamPlayers", testTeamPlayersExists)
	t.Run("TeamRoundPlayers", testTeamRoundPlayersExists)
	t.Run("TeamRounds", testTeamRoundsExists)
	t.Run("Teams", testTeamsExists)
	t.Run("TradePlayers", testTradePlayersExists)
	t.Run("Trades", testTradesExists)
	t.Run("Users", testUsersExists)
	t.Run("Waivers", testWaiversExists)
}

func TestFind(t *testing.T) {
	t.Run("ClubMatchups", testClubMatchupsFind)
	t.Run("Clubs", testClubsFind)
	t.Run("Leagues", testLeaguesFind)
	t.Run("PlayerRounds", testPlayerRoundsFind)
	t.Run("Players", testPlayersFind)
	t.Run("Rounds", testRoundsFind)
	t.Run("TeamMatchups", testTeamMatchupsFind)
	t.Run("TeamPlayers", testTeamPlayersFind)
	t.Run("TeamRoundPlayers", testTeamRoundPlayersFind)
	t.Run("TeamRounds", testTeamRoundsFind)
	t.Run("Teams", testTeamsFind)
	t.Run("TradePlayers", testTradePlayersFind)
	t.Run("Trades", testTradesFind)
	t.Run("Users", testUsersFind)
	t.Run("Waivers", testWaiversFind)
}

func TestBind(t *testing.T) {
	t.Run("ClubMatchups", testClubMatchupsBind)
	t.Run("Clubs", testClubsBind)
	t.Run("Leagues", testLeaguesBind)
	t.Run("PlayerRounds", testPlayerRoundsBind)
	t.Run("Players", testPlayersBind)
	t.Run("Rounds", testRoundsBind)
	t.Run("TeamMatchups", testTeamMatchupsBind)
	t.Run("TeamPlayers", testTeamPlayersBind)
	t.Run("TeamRoundPlayers", testTeamRoundPlayersBind)
	t.Run("TeamRounds", testTeamRoundsBind)
	t.Run("Teams", testTeamsBind)
	t.Run("TradePlayers", testTradePlayersBind)
	t.Run("Trades", testTradesBind)
	t.Run("Users", testUsersBind)
	t.Run("Waivers", testWaiversBind)
}

func TestOne(t *testing.T) {
	t.Run("ClubMatchups", testClubMatchupsOne)
	t.Run("Clubs", testClubsOne)
	t.Run("Leagues", testLeaguesOne)
	t.Run("PlayerRounds", testPlayerRoundsOne)
	t.Run("Players", testPlayersOne)
	t.Run("Rounds", testRoundsOne)
	t.Run("TeamMatchups", testTeamMatchupsOne)
	t.Run("TeamPlayers", testTeamPlayersOne)
	t.Run("TeamRoundPlayers", testTeamRoundPlayersOne)
	t.Run("TeamRounds", testTeamRoundsOne)
	t.Run("Teams", testTeamsOne)
	t.Run("TradePlayers", testTradePlayersOne)
	t.Run("Trades", testTradesOne)
	t.Run("Users", testUsersOne)
	t.Run("Waivers", testWaiversOne)
}

func TestAll(t *testing.T) {
	t.Run("ClubMatchups", testClubMatchupsAll)
	t.Run("Clubs", testClubsAll)
	t.Run("Leagues", testLeaguesAll)
	t.Run("PlayerRounds", testPlayerRoundsAll)
	t.Run("Players", testPlayersAll)
	t.Run("Rounds", testRoundsAll)
	t.Run("TeamMatchups", testTeamMatchupsAll)
	t.Run("TeamPlayers", testTeamPlayersAll)
	t.Run("TeamRoundPlayers", testTeamRoundPlayersAll)
	t.Run("TeamRounds", testTeamRoundsAll)
	t.Run("Teams", testTeamsAll)
	t.Run("TradePlayers", testTradePlayersAll)
	t.Run("Trades", testTradesAll)
	t.Run("Users", testUsersAll)
	t.Run("Waivers", testWaiversAll)
}

func TestCount(t *testing.T) {
	t.Run("ClubMatchups", testClubMatchupsCount)
	t.Run("Clubs", testClubsCount)
	t.Run("Leagues", testLeaguesCount)
	t.Run("PlayerRounds", testPlayerRoundsCount)
	t.Run("Players", testPlayersCount)
	t.Run("Rounds", testRoundsCount)
	t.Run("TeamMatchups", testTeamMatchupsCount)
	t.Run("TeamPlayers", testTeamPlayersCount)
	t.Run("TeamRoundPlayers", testTeamRoundPlayersCount)
	t.Run("TeamRounds", testTeamRoundsCount)
	t.Run("Teams", testTeamsCount)
	t.Run("TradePlayers", testTradePlayersCount)
	t.Run("Trades", testTradesCount)
	t.Run("Users", testUsersCount)
	t.Run("Waivers", testWaiversCount)
}

func TestHooks(t *testing.T) {
	t.Run("ClubMatchups", testClubMatchupsHooks)
	t.Run("Clubs", testClubsHooks)
	t.Run("Leagues", testLeaguesHooks)
	t.Run("PlayerRounds", testPlayerRoundsHooks)
	t.Run("Players", testPlayersHooks)
	t.Run("Rounds", testRoundsHooks)
	t.Run("TeamMatchups", testTeamMatchupsHooks)
	t.Run("TeamPlayers", testTeamPlayersHooks)
	t.Run("TeamRoundPlayers", testTeamRoundPlayersHooks)
	t.Run("TeamRounds", testTeamRoundsHooks)
	t.Run("Teams", testTeamsHooks)
	t.Run("TradePlayers", testTradePlayersHooks)
	t.Run("Trades", testTradesHooks)
	t.Run("Users", testUsersHooks)
	t.Run("Waivers", testWaiversHooks)
}

func TestInsert(t *testing.T) {
	t.Run("ClubMatchups", testClubMatchupsInsert)
	t.Run("ClubMatchups", testClubMatchupsInsertWhitelist)
	t.Run("Clubs", testClubsInsert)
	t.Run("Clubs", testClubsInsertWhitelist)
	t.Run("Leagues", testLeaguesInsert)
	t.Run("Leagues", testLeaguesInsertWhitelist)
	t.Run("PlayerRounds", testPlayerRoundsInsert)
	t.Run("PlayerRounds", testPlayerRoundsInsertWhitelist)
	t.Run("Players", testPlayersInsert)
	t.Run("Players", testPlayersInsertWhitelist)
	t.Run("Rounds", testRoundsInsert)
	t.Run("Rounds", testRoundsInsertWhitelist)
	t.Run("TeamMatchups", testTeamMatchupsInsert)
	t.Run("TeamMatchups", testTeamMatchupsInsertWhitelist)
	t.Run("TeamPlayers", testTeamPlayersInsert)
	t.Run("TeamPlayers", testTeamPlayersInsertWhitelist)
	t.Run("TeamRoundPlayers", testTeamRoundPlayersInsert)
	t.Run("TeamRoundPlayers", testTeamRoundPlayersInsertWhitelist)
	t.Run("TeamRounds", testTeamRoundsInsert)
	t.Run("TeamRounds", testTeamRoundsInsertWhitelist)
	t.Run("Teams", testTeamsInsert)
	t.Run("Teams", testTeamsInsertWhitelist)
	t.Run("TradePlayers", testTradePlayersInsert)
	t.Run("TradePlayers", testTradePlayersInsertWhitelist)
	t.Run("Trades", testTradesInsert)
	t.Run("Trades", testTradesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
	t.Run("Waivers", testWaiversInsert)
	t.Run("Waivers", testWaiversInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("ClubMatchupToRoundUsingRound", testClubMatchupToOneRoundUsingRound)
	t.Run("ClubMatchupToClubUsingHomeClubClub", testClubMatchupToOneClubUsingHomeClubClub)
	t.Run("ClubMatchupToClubUsingAwayClubClub", testClubMatchupToOneClubUsingAwayClubClub)
	t.Run("PlayerRoundToPlayerUsingPlayer", testPlayerRoundToOnePlayerUsingPlayer)
	t.Run("PlayerRoundToRoundUsingRound", testPlayerRoundToOneRoundUsingRound)
	t.Run("PlayerToClubUsingClub", testPlayerToOneClubUsingClub)
	t.Run("TeamMatchupToRoundUsingRound", testTeamMatchupToOneRoundUsingRound)
	t.Run("TeamMatchupToTeamUsingHomeTeamTeam", testTeamMatchupToOneTeamUsingHomeTeamTeam)
	t.Run("TeamMatchupToTeamUsingAwayTeamTeam", testTeamMatchupToOneTeamUsingAwayTeamTeam)
	t.Run("TeamMatchupToTeamUsingWinnerTeam", testTeamMatchupToOneTeamUsingWinnerTeam)
	t.Run("TeamPlayerToPlayerUsingPlayer", testTeamPlayerToOnePlayerUsingPlayer)
	t.Run("TeamPlayerToTeamUsingTeam", testTeamPlayerToOneTeamUsingTeam)
	t.Run("TeamRoundPlayerToTeamRoundUsingTeamRound", testTeamRoundPlayerToOneTeamRoundUsingTeamRound)
	t.Run("TeamRoundPlayerToTeamPlayerUsingTeamPlayer", testTeamRoundPlayerToOneTeamPlayerUsingTeamPlayer)
	t.Run("TeamRoundToRoundUsingRound", testTeamRoundToOneRoundUsingRound)
	t.Run("TeamRoundToTeamUsingTeam", testTeamRoundToOneTeamUsingTeam)
	t.Run("TeamToUserUsingUser", testTeamToOneUserUsingUser)
	t.Run("TeamToLeagueUsingLeague", testTeamToOneLeagueUsingLeague)
	t.Run("TradePlayerToTeamPlayerUsingTeamPlayer", testTradePlayerToOneTeamPlayerUsingTeamPlayer)
	t.Run("TradePlayerToTradeUsingTrade", testTradePlayerToOneTradeUsingTrade)
	t.Run("TradePlayerToTeamUsingCurrentTeamTeam", testTradePlayerToOneTeamUsingCurrentTeamTeam)
	t.Run("TradeToTeamUsingFromTeamTeam", testTradeToOneTeamUsingFromTeamTeam)
	t.Run("TradeToTeamUsingToTeamTeam", testTradeToOneTeamUsingToTeamTeam)
	t.Run("WaiverToTeamUsingFromTeamTeam", testWaiverToOneTeamUsingFromTeamTeam)
	t.Run("WaiverToTeamUsingToTeamTeam", testWaiverToOneTeamUsingToTeamTeam)
	t.Run("WaiverToPlayerUsingPlayer", testWaiverToOnePlayerUsingPlayer)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("ClubToHomeClubClubMatchups", testClubToManyHomeClubClubMatchups)
	t.Run("ClubToAwayClubClubMatchups", testClubToManyAwayClubClubMatchups)
	t.Run("ClubToClubPlayers", testClubToManyClubPlayers)
	t.Run("LeagueToLeagueTeams", testLeagueToManyLeagueTeams)
	t.Run("PlayerToPlayerPlayerRounds", testPlayerToManyPlayerPlayerRounds)
	t.Run("PlayerToPlayerTeamPlayers", testPlayerToManyPlayerTeamPlayers)
	t.Run("PlayerToPlayerWaivers", testPlayerToManyPlayerWaivers)
	t.Run("RoundToRoundClubMatchups", testRoundToManyRoundClubMatchups)
	t.Run("RoundToRoundPlayerRounds", testRoundToManyRoundPlayerRounds)
	t.Run("RoundToRoundTeamMatchups", testRoundToManyRoundTeamMatchups)
	t.Run("RoundToRoundTeamRounds", testRoundToManyRoundTeamRounds)
	t.Run("TeamPlayerToTeamPlayerTeamRoundPlayers", testTeamPlayerToManyTeamPlayerTeamRoundPlayers)
	t.Run("TeamPlayerToTeamPlayerTradePlayers", testTeamPlayerToManyTeamPlayerTradePlayers)
	t.Run("TeamRoundToTeamRoundTeamRoundPlayers", testTeamRoundToManyTeamRoundTeamRoundPlayers)
	t.Run("TeamToHomeTeamTeamMatchups", testTeamToManyHomeTeamTeamMatchups)
	t.Run("TeamToAwayTeamTeamMatchups", testTeamToManyAwayTeamTeamMatchups)
	t.Run("TeamToWinnerTeamMatchups", testTeamToManyWinnerTeamMatchups)
	t.Run("TeamToTeamTeamPlayers", testTeamToManyTeamTeamPlayers)
	t.Run("TeamToTeamTeamRounds", testTeamToManyTeamTeamRounds)
	t.Run("TeamToCurrentTeamTradePlayers", testTeamToManyCurrentTeamTradePlayers)
	t.Run("TeamToFromTeamTrades", testTeamToManyFromTeamTrades)
	t.Run("TeamToToTeamTrades", testTeamToManyToTeamTrades)
	t.Run("TeamToFromTeamWaivers", testTeamToManyFromTeamWaivers)
	t.Run("TeamToToTeamWaivers", testTeamToManyToTeamWaivers)
	t.Run("TradeToTradeTradePlayers", testTradeToManyTradeTradePlayers)
	t.Run("UserToUserTeams", testUserToManyUserTeams)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("ClubMatchupToRoundUsingRoundClubMatchups", testClubMatchupToOneSetOpRoundUsingRound)
	t.Run("ClubMatchupToClubUsingHomeClubClubMatchups", testClubMatchupToOneSetOpClubUsingHomeClubClub)
	t.Run("ClubMatchupToClubUsingAwayClubClubMatchups", testClubMatchupToOneSetOpClubUsingAwayClubClub)
	t.Run("PlayerRoundToPlayerUsingPlayerPlayerRounds", testPlayerRoundToOneSetOpPlayerUsingPlayer)
	t.Run("PlayerRoundToRoundUsingRoundPlayerRounds", testPlayerRoundToOneSetOpRoundUsingRound)
	t.Run("PlayerToClubUsingClubPlayers", testPlayerToOneSetOpClubUsingClub)
	t.Run("TeamMatchupToRoundUsingRoundTeamMatchups", testTeamMatchupToOneSetOpRoundUsingRound)
	t.Run("TeamMatchupToTeamUsingHomeTeamTeamMatchups", testTeamMatchupToOneSetOpTeamUsingHomeTeamTeam)
	t.Run("TeamMatchupToTeamUsingAwayTeamTeamMatchups", testTeamMatchupToOneSetOpTeamUsingAwayTeamTeam)
	t.Run("TeamMatchupToTeamUsingWinnerTeamMatchups", testTeamMatchupToOneSetOpTeamUsingWinnerTeam)
	t.Run("TeamPlayerToPlayerUsingPlayerTeamPlayers", testTeamPlayerToOneSetOpPlayerUsingPlayer)
	t.Run("TeamPlayerToTeamUsingTeamTeamPlayers", testTeamPlayerToOneSetOpTeamUsingTeam)
	t.Run("TeamRoundPlayerToTeamRoundUsingTeamRoundTeamRoundPlayers", testTeamRoundPlayerToOneSetOpTeamRoundUsingTeamRound)
	t.Run("TeamRoundPlayerToTeamPlayerUsingTeamPlayerTeamRoundPlayers", testTeamRoundPlayerToOneSetOpTeamPlayerUsingTeamPlayer)
	t.Run("TeamRoundToRoundUsingRoundTeamRounds", testTeamRoundToOneSetOpRoundUsingRound)
	t.Run("TeamRoundToTeamUsingTeamTeamRounds", testTeamRoundToOneSetOpTeamUsingTeam)
	t.Run("TeamToUserUsingUserTeams", testTeamToOneSetOpUserUsingUser)
	t.Run("TeamToLeagueUsingLeagueTeams", testTeamToOneSetOpLeagueUsingLeague)
	t.Run("TradePlayerToTeamPlayerUsingTeamPlayerTradePlayers", testTradePlayerToOneSetOpTeamPlayerUsingTeamPlayer)
	t.Run("TradePlayerToTradeUsingTradeTradePlayers", testTradePlayerToOneSetOpTradeUsingTrade)
	t.Run("TradePlayerToTeamUsingCurrentTeamTradePlayers", testTradePlayerToOneSetOpTeamUsingCurrentTeamTeam)
	t.Run("TradeToTeamUsingFromTeamTrades", testTradeToOneSetOpTeamUsingFromTeamTeam)
	t.Run("TradeToTeamUsingToTeamTrades", testTradeToOneSetOpTeamUsingToTeamTeam)
	t.Run("WaiverToTeamUsingFromTeamWaivers", testWaiverToOneSetOpTeamUsingFromTeamTeam)
	t.Run("WaiverToTeamUsingToTeamWaivers", testWaiverToOneSetOpTeamUsingToTeamTeam)
	t.Run("WaiverToPlayerUsingPlayerWaivers", testWaiverToOneSetOpPlayerUsingPlayer)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("ClubMatchupToRoundUsingRoundClubMatchups", testClubMatchupToOneRemoveOpRoundUsingRound)
	t.Run("ClubMatchupToClubUsingHomeClubClubMatchups", testClubMatchupToOneRemoveOpClubUsingHomeClubClub)
	t.Run("ClubMatchupToClubUsingAwayClubClubMatchups", testClubMatchupToOneRemoveOpClubUsingAwayClubClub)
	t.Run("PlayerRoundToPlayerUsingPlayerPlayerRounds", testPlayerRoundToOneRemoveOpPlayerUsingPlayer)
	t.Run("PlayerRoundToRoundUsingRoundPlayerRounds", testPlayerRoundToOneRemoveOpRoundUsingRound)
	t.Run("PlayerToClubUsingClubPlayers", testPlayerToOneRemoveOpClubUsingClub)
	t.Run("TeamMatchupToRoundUsingRoundTeamMatchups", testTeamMatchupToOneRemoveOpRoundUsingRound)
	t.Run("TeamMatchupToTeamUsingHomeTeamTeamMatchups", testTeamMatchupToOneRemoveOpTeamUsingHomeTeamTeam)
	t.Run("TeamMatchupToTeamUsingAwayTeamTeamMatchups", testTeamMatchupToOneRemoveOpTeamUsingAwayTeamTeam)
	t.Run("TeamMatchupToTeamUsingWinnerTeamMatchups", testTeamMatchupToOneRemoveOpTeamUsingWinnerTeam)
	t.Run("TeamRoundPlayerToTeamRoundUsingTeamRoundTeamRoundPlayers", testTeamRoundPlayerToOneRemoveOpTeamRoundUsingTeamRound)
	t.Run("TeamRoundPlayerToTeamPlayerUsingTeamPlayerTeamRoundPlayers", testTeamRoundPlayerToOneRemoveOpTeamPlayerUsingTeamPlayer)
	t.Run("TeamRoundToRoundUsingRoundTeamRounds", testTeamRoundToOneRemoveOpRoundUsingRound)
	t.Run("TeamRoundToTeamUsingTeamTeamRounds", testTeamRoundToOneRemoveOpTeamUsingTeam)
	t.Run("TradePlayerToTeamPlayerUsingTeamPlayerTradePlayers", testTradePlayerToOneRemoveOpTeamPlayerUsingTeamPlayer)
	t.Run("TradePlayerToTradeUsingTradeTradePlayers", testTradePlayerToOneRemoveOpTradeUsingTrade)
	t.Run("TradePlayerToTeamUsingCurrentTeamTradePlayers", testTradePlayerToOneRemoveOpTeamUsingCurrentTeamTeam)
	t.Run("TradeToTeamUsingFromTeamTrades", testTradeToOneRemoveOpTeamUsingFromTeamTeam)
	t.Run("TradeToTeamUsingToTeamTrades", testTradeToOneRemoveOpTeamUsingToTeamTeam)
	t.Run("WaiverToTeamUsingFromTeamWaivers", testWaiverToOneRemoveOpTeamUsingFromTeamTeam)
	t.Run("WaiverToTeamUsingToTeamWaivers", testWaiverToOneRemoveOpTeamUsingToTeamTeam)
	t.Run("WaiverToPlayerUsingPlayerWaivers", testWaiverToOneRemoveOpPlayerUsingPlayer)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("ClubToHomeClubClubMatchups", testClubToManyAddOpHomeClubClubMatchups)
	t.Run("ClubToAwayClubClubMatchups", testClubToManyAddOpAwayClubClubMatchups)
	t.Run("ClubToClubPlayers", testClubToManyAddOpClubPlayers)
	t.Run("LeagueToLeagueTeams", testLeagueToManyAddOpLeagueTeams)
	t.Run("PlayerToPlayerPlayerRounds", testPlayerToManyAddOpPlayerPlayerRounds)
	t.Run("PlayerToPlayerTeamPlayers", testPlayerToManyAddOpPlayerTeamPlayers)
	t.Run("PlayerToPlayerWaivers", testPlayerToManyAddOpPlayerWaivers)
	t.Run("RoundToRoundClubMatchups", testRoundToManyAddOpRoundClubMatchups)
	t.Run("RoundToRoundPlayerRounds", testRoundToManyAddOpRoundPlayerRounds)
	t.Run("RoundToRoundTeamMatchups", testRoundToManyAddOpRoundTeamMatchups)
	t.Run("RoundToRoundTeamRounds", testRoundToManyAddOpRoundTeamRounds)
	t.Run("TeamPlayerToTeamPlayerTeamRoundPlayers", testTeamPlayerToManyAddOpTeamPlayerTeamRoundPlayers)
	t.Run("TeamPlayerToTeamPlayerTradePlayers", testTeamPlayerToManyAddOpTeamPlayerTradePlayers)
	t.Run("TeamRoundToTeamRoundTeamRoundPlayers", testTeamRoundToManyAddOpTeamRoundTeamRoundPlayers)
	t.Run("TeamToHomeTeamTeamMatchups", testTeamToManyAddOpHomeTeamTeamMatchups)
	t.Run("TeamToAwayTeamTeamMatchups", testTeamToManyAddOpAwayTeamTeamMatchups)
	t.Run("TeamToWinnerTeamMatchups", testTeamToManyAddOpWinnerTeamMatchups)
	t.Run("TeamToTeamTeamPlayers", testTeamToManyAddOpTeamTeamPlayers)
	t.Run("TeamToTeamTeamRounds", testTeamToManyAddOpTeamTeamRounds)
	t.Run("TeamToCurrentTeamTradePlayers", testTeamToManyAddOpCurrentTeamTradePlayers)
	t.Run("TeamToFromTeamTrades", testTeamToManyAddOpFromTeamTrades)
	t.Run("TeamToToTeamTrades", testTeamToManyAddOpToTeamTrades)
	t.Run("TeamToFromTeamWaivers", testTeamToManyAddOpFromTeamWaivers)
	t.Run("TeamToToTeamWaivers", testTeamToManyAddOpToTeamWaivers)
	t.Run("TradeToTradeTradePlayers", testTradeToManyAddOpTradeTradePlayers)
	t.Run("UserToUserTeams", testUserToManyAddOpUserTeams)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("ClubToHomeClubClubMatchups", testClubToManySetOpHomeClubClubMatchups)
	t.Run("ClubToAwayClubClubMatchups", testClubToManySetOpAwayClubClubMatchups)
	t.Run("ClubToClubPlayers", testClubToManySetOpClubPlayers)
	t.Run("PlayerToPlayerPlayerRounds", testPlayerToManySetOpPlayerPlayerRounds)
	t.Run("PlayerToPlayerWaivers", testPlayerToManySetOpPlayerWaivers)
	t.Run("RoundToRoundClubMatchups", testRoundToManySetOpRoundClubMatchups)
	t.Run("RoundToRoundPlayerRounds", testRoundToManySetOpRoundPlayerRounds)
	t.Run("RoundToRoundTeamMatchups", testRoundToManySetOpRoundTeamMatchups)
	t.Run("RoundToRoundTeamRounds", testRoundToManySetOpRoundTeamRounds)
	t.Run("TeamPlayerToTeamPlayerTeamRoundPlayers", testTeamPlayerToManySetOpTeamPlayerTeamRoundPlayers)
	t.Run("TeamPlayerToTeamPlayerTradePlayers", testTeamPlayerToManySetOpTeamPlayerTradePlayers)
	t.Run("TeamRoundToTeamRoundTeamRoundPlayers", testTeamRoundToManySetOpTeamRoundTeamRoundPlayers)
	t.Run("TeamToHomeTeamTeamMatchups", testTeamToManySetOpHomeTeamTeamMatchups)
	t.Run("TeamToAwayTeamTeamMatchups", testTeamToManySetOpAwayTeamTeamMatchups)
	t.Run("TeamToWinnerTeamMatchups", testTeamToManySetOpWinnerTeamMatchups)
	t.Run("TeamToTeamTeamRounds", testTeamToManySetOpTeamTeamRounds)
	t.Run("TeamToCurrentTeamTradePlayers", testTeamToManySetOpCurrentTeamTradePlayers)
	t.Run("TeamToFromTeamTrades", testTeamToManySetOpFromTeamTrades)
	t.Run("TeamToToTeamTrades", testTeamToManySetOpToTeamTrades)
	t.Run("TeamToFromTeamWaivers", testTeamToManySetOpFromTeamWaivers)
	t.Run("TeamToToTeamWaivers", testTeamToManySetOpToTeamWaivers)
	t.Run("TradeToTradeTradePlayers", testTradeToManySetOpTradeTradePlayers)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("ClubToHomeClubClubMatchups", testClubToManyRemoveOpHomeClubClubMatchups)
	t.Run("ClubToAwayClubClubMatchups", testClubToManyRemoveOpAwayClubClubMatchups)
	t.Run("ClubToClubPlayers", testClubToManyRemoveOpClubPlayers)
	t.Run("PlayerToPlayerPlayerRounds", testPlayerToManyRemoveOpPlayerPlayerRounds)
	t.Run("PlayerToPlayerWaivers", testPlayerToManyRemoveOpPlayerWaivers)
	t.Run("RoundToRoundClubMatchups", testRoundToManyRemoveOpRoundClubMatchups)
	t.Run("RoundToRoundPlayerRounds", testRoundToManyRemoveOpRoundPlayerRounds)
	t.Run("RoundToRoundTeamMatchups", testRoundToManyRemoveOpRoundTeamMatchups)
	t.Run("RoundToRoundTeamRounds", testRoundToManyRemoveOpRoundTeamRounds)
	t.Run("TeamPlayerToTeamPlayerTeamRoundPlayers", testTeamPlayerToManyRemoveOpTeamPlayerTeamRoundPlayers)
	t.Run("TeamPlayerToTeamPlayerTradePlayers", testTeamPlayerToManyRemoveOpTeamPlayerTradePlayers)
	t.Run("TeamRoundToTeamRoundTeamRoundPlayers", testTeamRoundToManyRemoveOpTeamRoundTeamRoundPlayers)
	t.Run("TeamToHomeTeamTeamMatchups", testTeamToManyRemoveOpHomeTeamTeamMatchups)
	t.Run("TeamToAwayTeamTeamMatchups", testTeamToManyRemoveOpAwayTeamTeamMatchups)
	t.Run("TeamToWinnerTeamMatchups", testTeamToManyRemoveOpWinnerTeamMatchups)
	t.Run("TeamToTeamTeamRounds", testTeamToManyRemoveOpTeamTeamRounds)
	t.Run("TeamToCurrentTeamTradePlayers", testTeamToManyRemoveOpCurrentTeamTradePlayers)
	t.Run("TeamToFromTeamTrades", testTeamToManyRemoveOpFromTeamTrades)
	t.Run("TeamToToTeamTrades", testTeamToManyRemoveOpToTeamTrades)
	t.Run("TeamToFromTeamWaivers", testTeamToManyRemoveOpFromTeamWaivers)
	t.Run("TeamToToTeamWaivers", testTeamToManyRemoveOpToTeamWaivers)
	t.Run("TradeToTradeTradePlayers", testTradeToManyRemoveOpTradeTradePlayers)
}

func TestReload(t *testing.T) {
	t.Run("ClubMatchups", testClubMatchupsReload)
	t.Run("Clubs", testClubsReload)
	t.Run("Leagues", testLeaguesReload)
	t.Run("PlayerRounds", testPlayerRoundsReload)
	t.Run("Players", testPlayersReload)
	t.Run("Rounds", testRoundsReload)
	t.Run("TeamMatchups", testTeamMatchupsReload)
	t.Run("TeamPlayers", testTeamPlayersReload)
	t.Run("TeamRoundPlayers", testTeamRoundPlayersReload)
	t.Run("TeamRounds", testTeamRoundsReload)
	t.Run("Teams", testTeamsReload)
	t.Run("TradePlayers", testTradePlayersReload)
	t.Run("Trades", testTradesReload)
	t.Run("Users", testUsersReload)
	t.Run("Waivers", testWaiversReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("ClubMatchups", testClubMatchupsReloadAll)
	t.Run("Clubs", testClubsReloadAll)
	t.Run("Leagues", testLeaguesReloadAll)
	t.Run("PlayerRounds", testPlayerRoundsReloadAll)
	t.Run("Players", testPlayersReloadAll)
	t.Run("Rounds", testRoundsReloadAll)
	t.Run("TeamMatchups", testTeamMatchupsReloadAll)
	t.Run("TeamPlayers", testTeamPlayersReloadAll)
	t.Run("TeamRoundPlayers", testTeamRoundPlayersReloadAll)
	t.Run("TeamRounds", testTeamRoundsReloadAll)
	t.Run("Teams", testTeamsReloadAll)
	t.Run("TradePlayers", testTradePlayersReloadAll)
	t.Run("Trades", testTradesReloadAll)
	t.Run("Users", testUsersReloadAll)
	t.Run("Waivers", testWaiversReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("ClubMatchups", testClubMatchupsSelect)
	t.Run("Clubs", testClubsSelect)
	t.Run("Leagues", testLeaguesSelect)
	t.Run("PlayerRounds", testPlayerRoundsSelect)
	t.Run("Players", testPlayersSelect)
	t.Run("Rounds", testRoundsSelect)
	t.Run("TeamMatchups", testTeamMatchupsSelect)
	t.Run("TeamPlayers", testTeamPlayersSelect)
	t.Run("TeamRoundPlayers", testTeamRoundPlayersSelect)
	t.Run("TeamRounds", testTeamRoundsSelect)
	t.Run("Teams", testTeamsSelect)
	t.Run("TradePlayers", testTradePlayersSelect)
	t.Run("Trades", testTradesSelect)
	t.Run("Users", testUsersSelect)
	t.Run("Waivers", testWaiversSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("ClubMatchups", testClubMatchupsUpdate)
	t.Run("Clubs", testClubsUpdate)
	t.Run("Leagues", testLeaguesUpdate)
	t.Run("PlayerRounds", testPlayerRoundsUpdate)
	t.Run("Players", testPlayersUpdate)
	t.Run("Rounds", testRoundsUpdate)
	t.Run("TeamMatchups", testTeamMatchupsUpdate)
	t.Run("TeamPlayers", testTeamPlayersUpdate)
	t.Run("TeamRoundPlayers", testTeamRoundPlayersUpdate)
	t.Run("TeamRounds", testTeamRoundsUpdate)
	t.Run("Teams", testTeamsUpdate)
	t.Run("TradePlayers", testTradePlayersUpdate)
	t.Run("Trades", testTradesUpdate)
	t.Run("Users", testUsersUpdate)
	t.Run("Waivers", testWaiversUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("ClubMatchups", testClubMatchupsSliceUpdateAll)
	t.Run("Clubs", testClubsSliceUpdateAll)
	t.Run("Leagues", testLeaguesSliceUpdateAll)
	t.Run("PlayerRounds", testPlayerRoundsSliceUpdateAll)
	t.Run("Players", testPlayersSliceUpdateAll)
	t.Run("Rounds", testRoundsSliceUpdateAll)
	t.Run("TeamMatchups", testTeamMatchupsSliceUpdateAll)
	t.Run("TeamPlayers", testTeamPlayersSliceUpdateAll)
	t.Run("TeamRoundPlayers", testTeamRoundPlayersSliceUpdateAll)
	t.Run("TeamRounds", testTeamRoundsSliceUpdateAll)
	t.Run("Teams", testTeamsSliceUpdateAll)
	t.Run("TradePlayers", testTradePlayersSliceUpdateAll)
	t.Run("Trades", testTradesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
	t.Run("Waivers", testWaiversSliceUpdateAll)
}
