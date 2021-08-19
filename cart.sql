CREATE TABLE cart (
                       id serial PRIMARY KEY,
                       name VARCHAR (50) NOT NULL,
                       quantity INT NOT NULL,
                       price INT NOT NULL,
                       totprice INT NOT NULL
);