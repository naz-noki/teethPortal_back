CREATE TABLE refresh_tokens(
    id SERIAL PRIMARY KEY,
    token TEXT NOT NULL UNIQUE,
    user_id INT REFERENCES users(id)
);