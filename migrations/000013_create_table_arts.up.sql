CREATE TABLE arts(
    id SERIAL PRIMARY KEY, 
    name TEXT NOT NULL,
    description TEXT,
    content TEXT,
    author_id INT REFERENCES authors(id),
    type art_type,
    files_idx TEXT
);