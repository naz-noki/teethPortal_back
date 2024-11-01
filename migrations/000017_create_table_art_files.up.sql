CREATE TABLE art_files (
    id SERIAL PRIMARY KEY, 
    art_id INT REFERENCES arts(id),
    file_id TEXT
);