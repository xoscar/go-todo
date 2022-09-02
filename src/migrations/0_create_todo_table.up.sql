CREATE TABLE IF NOT EXISTS todos (
	id SERIAL PRIMARY KEY,
	title varchar NOT NULL,
	status varchar NOT NULL
);