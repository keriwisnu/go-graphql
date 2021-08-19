CREATE TABLE cart (
                      id serial PRIMARY KEY,
                      name VARCHAR (50) NOT NULL,
                      stock INT NOT NULL,
                      price INT NOT NULL,
                      codedisc VARCHAR (2) NOT NULL FOREIGN KEY
);