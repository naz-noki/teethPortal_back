CREATE TABLE author_category (
    id SERIAL PRIMARY KEY, 
    category_id INT REFERENCES categories(id) ON DELETE CASCADE,
    author_id INT REFERENCES authors(id) ON DELETE CASCADE
);