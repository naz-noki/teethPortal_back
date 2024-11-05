CREATE TABLE auction (
    id INT PRIMARY KEY, 
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    price INT NOT NULL, 
    lot INT REFERENCES arts(id),
    seller INT REFERENCES authors(id),
    buyer INT REFERENCES users(id),
    paid BOOLEAN DEFAULT FALSE, 
    sent BOOLEAN DEFAULT FALSE
);