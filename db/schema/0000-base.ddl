CREATE TABLE `Users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `full_name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `hashed_password` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `Teams` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `user_id` int NOT NULL,
  `league_id` int NOT NULL,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `Clubs` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `Leagues` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `Players` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `club_id` int,
  `name` varchar(255) NOT NULL,
  `age` timestamp NOT NULL,
  `position` int NOT NULL,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `TeamPlayers` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `player_id` int NOT NULL,
  `team_id` int NOT NULL,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `TradePlayers` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `team_player_id` int,
  `trade_id` int,
  `current_team` int,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `Trades` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `from_team` int,
  `to_team` int,
  `status` int,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `Waiver` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `from_team` int,
  `to_team` int,
  `player_id` int,
  `status` int,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `Round` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `round_name`  varchar(255),
  `round_number` int,
  `start_date` timestamp,
  `end_date` timestamp,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `TeamMatchups` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `round_id` int,
  `home_team` int,
  `away_team` int,
  `home_team_score` int,
  `away_team_score` int,
  `winner` int,
  `status` int,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `TeamRounds` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `round_id` int,
  `team_id` int,
  `score` int,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `TeamRoundPlayers` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `team_round_id` int,
  `team_player_id` int,
  `status` int,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `PlayerRounds` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `player_id` int,
  `round_id` int,
  `status` int,
  `score` int,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

CREATE TABLE `ClubMatchups` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `round_id` int,
  `home_club` int,
  `away_club` int,
  `game_time` timestamp,
  `home_club_score` int,
  `away_club_score` int,
  `status` int,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp
);

ALTER TABLE `Teams` ADD FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`);

ALTER TABLE `Teams` ADD FOREIGN KEY (`league_id`) REFERENCES `Leagues` (`id`);

ALTER TABLE `Players` ADD FOREIGN KEY (`club_id`) REFERENCES `Clubs` (`id`);

ALTER TABLE `TeamPlayers` ADD FOREIGN KEY (`player_id`) REFERENCES `Players` (`id`);

ALTER TABLE `TeamPlayers` ADD FOREIGN KEY (`team_id`) REFERENCES `Teams` (`id`);

ALTER TABLE `TradePlayers` ADD FOREIGN KEY (`team_player_id`) REFERENCES `TeamPlayers` (`id`);

ALTER TABLE `TradePlayers` ADD FOREIGN KEY (`trade_id`) REFERENCES `Trades` (`id`);

ALTER TABLE `TradePlayers` ADD FOREIGN KEY (`current_team`) REFERENCES `Teams` (`id`);

ALTER TABLE `Trades` ADD FOREIGN KEY (`from_team`) REFERENCES `Teams` (`id`);

ALTER TABLE `Trades` ADD FOREIGN KEY (`to_team`) REFERENCES `Teams` (`id`);

ALTER TABLE `Waiver` ADD FOREIGN KEY (`from_team`) REFERENCES `Teams` (`id`);

ALTER TABLE `Waiver` ADD FOREIGN KEY (`to_team`) REFERENCES `Teams` (`id`);

ALTER TABLE `Waiver` ADD FOREIGN KEY (`player_id`) REFERENCES `Players` (`id`);

ALTER TABLE `TeamMatchups` ADD FOREIGN KEY (`round_id`) REFERENCES `Round` (`id`);

ALTER TABLE `TeamMatchups` ADD FOREIGN KEY (`home_team`) REFERENCES `Teams` (`id`);

ALTER TABLE `TeamMatchups` ADD FOREIGN KEY (`away_team`) REFERENCES `Teams` (`id`);

ALTER TABLE `TeamMatchups` ADD FOREIGN KEY (`winner`) REFERENCES `Teams` (`id`);

ALTER TABLE `TeamRounds` ADD FOREIGN KEY (`round_id`) REFERENCES `Round` (`id`);

ALTER TABLE `TeamRounds` ADD FOREIGN KEY (`team_id`) REFERENCES `Teams` (`id`);

ALTER TABLE `TeamRoundPlayers` ADD FOREIGN KEY (`team_round_id`) REFERENCES `TeamRounds` (`id`);

ALTER TABLE `TeamRoundPlayers` ADD FOREIGN KEY (`team_player_id`) REFERENCES `TeamPlayers` (`id`);

ALTER TABLE `PlayerRounds` ADD FOREIGN KEY (`player_id`) REFERENCES `Players` (`id`);

ALTER TABLE `PlayerRounds` ADD FOREIGN KEY (`round_id`) REFERENCES `Round` (`id`);

ALTER TABLE `ClubMatchups` ADD FOREIGN KEY (`round_id`) REFERENCES `Round` (`id`);

ALTER TABLE `ClubMatchups` ADD FOREIGN KEY (`home_club`) REFERENCES `Clubs` (`id`);

ALTER TABLE `ClubMatchups` ADD FOREIGN KEY (`away_club`) REFERENCES `Clubs` (`id`);
