CREATE TABLE authors(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT,
    user_id INT REFERENCES users(id)
);