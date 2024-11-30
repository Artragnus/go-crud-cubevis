CREATE TABLE users(
    id UUID PRIMARY KEY UNIQUE NOT NULL,
    name VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE address (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    address VARCHAR(255) NOT NULL,
    number VARCHAR(255),
    zip_code VARCHAR(255),
    city VARCHAR(255) NOT NULL,
    state VARCHAR(255) NOT NULL
);

