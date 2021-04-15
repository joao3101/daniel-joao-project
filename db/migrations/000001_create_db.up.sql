CREATE TABLE "Users" (
  "id" SERIAL PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "created_at" int NOT NULL,
  "deleted_at" int
);

CREATE TABLE "Teams" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "user_id" int NOT NULL,
  "league_id" int NOT NULL,
  "created_at" int NOT NULL,
  "deleted_at" int
);

CREATE TABLE "Clubs" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" int NOT NULL,
  "deleted_at" int
);

CREATE TABLE "Leagues" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" int NOT NULL,
  "deleted_at" int
);

CREATE TABLE "Players" (
  "id" SERIAL PRIMARY KEY,
  "club_id" int,
  "name" varchar NOT NULL,
  "age" int NOT NULL,
  "position" int NOT NULL,
  "created_at" int NOT NULL,
  "deleted_at" int
);

CREATE TABLE "TeamPlayers" (
  "id" SERIAL PRIMARY KEY,
  "player_id" int NOT NULL,
  "team_id" int NOT NULL,
  "created_at" int NOT NULL,
  "deleted_at" int
);

CREATE TABLE "TradePlayers" (
  "id" SERIAL PRIMARY KEY,
  "team_player_id" int,
  "trade_id" int,
  "current_team" int,
  "created_at" int NOT NULL,
  "deleted_at" int
);

CREATE TABLE "Trades" (
  "id" SERIAL PRIMARY KEY,
  "from_team" int,
  "to_team" int,
  "status" int,
  "created_at" int NOT NULL,
  "deleted_at" int
);

CREATE TABLE "Waiver" (
  "id" SERIAL PRIMARY KEY,
  "from_team" int,
  "to_team" int,
  "player_id" int,
  "status" int,
  "created_at" int NOT NULL,
  "deleted_at" int
);

CREATE TABLE "Round" (
  "id" SERIAL PRIMARY KEY,
  "round_name" varchar,
  "round_number" int,
  "start_date" int,
  "end_date" int,
  "created_at" int NOT NULL,
  "deleted_at" int
);

CREATE TABLE "TeamMatchups" (
  "id" SERIAL PRIMARY KEY,
  "round_id" int,
  "home_team" int,
  "away_team" int,
  "home_team_score" int,
  "away_team_score" int,
  "winner" int,
  "status" int,
  "created_at" int NOT NULL,
  "deleted_at" int
);

CREATE TABLE "TeamRounds" (
  "id" SERIAL PRIMARY KEY,
  "round_id" int,
  "team_id" int,
  "score" int,
  "created_at" int NOT NULL,
  "deleted_at" int
);

CREATE TABLE "TeamRoundPlayers" (
  "id" SERIAL PRIMARY KEY,
  "team_round_id" int,
  "team_player_id" int,
  "status" int,
  "created_at" int NOT NULL,
  "deleted_at" int
);

CREATE TABLE "PlayerRounds" (
  "id" SERIAL PRIMARY KEY,
  "player_id" int,
  "round_id" int,
  "status" int,
  "score" int,
  "created_at" int NOT NULL,
  "deleted_at" int
);

CREATE TABLE "ClubMatchups" (
  "id" SERIAL PRIMARY KEY,
  "round_id" int,
  "home_club" int,
  "away_club" int,
  "game_time" int,
  "home_club_score" int,
  "away_club_score" int,
  "status" int,
  "created_at" int NOT NULL,
  "deleted_at" int
);

ALTER TABLE "Teams" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "Teams" ADD FOREIGN KEY ("league_id") REFERENCES "Leagues" ("id");

ALTER TABLE "Players" ADD FOREIGN KEY ("club_id") REFERENCES "Clubs" ("id");

ALTER TABLE "TeamPlayers" ADD FOREIGN KEY ("player_id") REFERENCES "Players" ("id");

ALTER TABLE "TeamPlayers" ADD FOREIGN KEY ("team_id") REFERENCES "Teams" ("id");

ALTER TABLE "TradePlayers" ADD FOREIGN KEY ("team_player_id") REFERENCES "TeamPlayers" ("id");

ALTER TABLE "TradePlayers" ADD FOREIGN KEY ("trade_id") REFERENCES "Trades" ("id");

ALTER TABLE "TradePlayers" ADD FOREIGN KEY ("current_team") REFERENCES "Teams" ("id");

ALTER TABLE "Trades" ADD FOREIGN KEY ("from_team") REFERENCES "Teams" ("id");

ALTER TABLE "Trades" ADD FOREIGN KEY ("to_team") REFERENCES "Teams" ("id");

ALTER TABLE "Waiver" ADD FOREIGN KEY ("from_team") REFERENCES "Teams" ("id");

ALTER TABLE "Waiver" ADD FOREIGN KEY ("to_team") REFERENCES "Teams" ("id");

ALTER TABLE "Waiver" ADD FOREIGN KEY ("player_id") REFERENCES "Players" ("id");

ALTER TABLE "TeamMatchups" ADD FOREIGN KEY ("round_id") REFERENCES "Round" ("id");

ALTER TABLE "TeamMatchups" ADD FOREIGN KEY ("home_team") REFERENCES "Teams" ("id");

ALTER TABLE "TeamMatchups" ADD FOREIGN KEY ("away_team") REFERENCES "Teams" ("id");

ALTER TABLE "TeamMatchups" ADD FOREIGN KEY ("winner") REFERENCES "Teams" ("id");

ALTER TABLE "TeamRounds" ADD FOREIGN KEY ("round_id") REFERENCES "Round" ("id");

ALTER TABLE "TeamRounds" ADD FOREIGN KEY ("team_id") REFERENCES "Teams" ("id");

ALTER TABLE "TeamRoundPlayers" ADD FOREIGN KEY ("team_round_id") REFERENCES "TeamRounds" ("id");

ALTER TABLE "TeamRoundPlayers" ADD FOREIGN KEY ("team_player_id") REFERENCES "TeamPlayers" ("id");

ALTER TABLE "PlayerRounds" ADD FOREIGN KEY ("player_id") REFERENCES "Players" ("id");

ALTER TABLE "PlayerRounds" ADD FOREIGN KEY ("round_id") REFERENCES "Round" ("id");

ALTER TABLE "ClubMatchups" ADD FOREIGN KEY ("round_id") REFERENCES "Round" ("id");

ALTER TABLE "ClubMatchups" ADD FOREIGN KEY ("home_club") REFERENCES "Clubs" ("id");

ALTER TABLE "ClubMatchups" ADD FOREIGN KEY ("away_club") REFERENCES "Clubs" ("id");
