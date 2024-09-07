CREATE TABLE images(
    id SERIAL PRIMARY KEY, 
    path TEXT NOT NULL UNIQUE,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    user_id INT REFERENCES users(id)
);