CREATE TABLE art_files (
    id INT PRIMARY KEY, 
    art_id INT REFERENCES arts(id),
    file_id TEXT
);