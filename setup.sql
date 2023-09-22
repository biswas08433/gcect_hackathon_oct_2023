DROP TABLE sessions;
DROP TABLE posts;
DROP TABLE threads;
DROP TABLE users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(64) NOT NULL UNIQUE,
    name VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    created_at timestamp NOT NULL
);

CREATE TABLE sessions(
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(64) NOT NULL UNIQUE,
    email VARCHAR(255),
    user_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE threads(
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(64) NOT NULL UNIQUE,
    topic TEXT,
    user_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP NOT NULL 
);

CREATE TABLE posts(
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(64) NOT NULL UNIQUE,
    body TEXT,
    user_id INTEGER REFERENCES users(id),
    thread_id INTEGER REFERENCES threads(id),
    created_at TIMESTAMP NOT NULL
);