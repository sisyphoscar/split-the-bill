CREATE TABLE members (
	id serial4 NOT NULL,
	name varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	deleted_at timestamp NULL
);