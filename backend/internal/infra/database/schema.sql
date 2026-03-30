CREATE TABLE users (
	id TEXT PRIMARY KEY, -- uuid
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL
);
