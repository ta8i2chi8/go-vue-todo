USE todo-db;

CREATE TABLE IF NOT EXISTS todos(
	id        INT(20) AUTO_INCREMENT PRIMARY KEY,
	title     VARCHAR(20) NOT NULL,
	context   TEXT NOT NULL,
    limitDate TIMESTAMP NOT NULL,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);