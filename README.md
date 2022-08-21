# MintRaffle

Setup steps:

DB:
- setup a mysql db with user name "root" & password "Root1234" with localhost port 3306, db name: "MintRaffle"
- load the schema to the DB
```
CREATE TABLE `projects` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text,
  `offical_link` varchar(255) DEFAULT NULL,
  `max_winner` int NOT NULL,
  `due_time` datetime NOT NULL,
  `status` varchar(45) NOT NULL DEFAULT 'opening' COMMENT 'opening, pending_for_raffle, closed',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `submissions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `project_id` int DEFAULT NULL,
  `wallet_address` varchar(255) DEFAULT NULL,
  `winner` tinyint(1) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_project_id_idx` (`project_id`),
  CONSTRAINT `fk_project_id` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
```

Go:
- install go
- clone the project then run go mod tidy
- to start up the RESTful API server:
  `go run cmd/mint-raffle/mint-raffle.go`
- to run the raffle script:
  `go run cmd/raffle/raffle.go`
  
