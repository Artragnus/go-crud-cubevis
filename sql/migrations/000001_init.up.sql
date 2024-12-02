CREATE TABLE  users(
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE addresses (
   id UUID PRIMARY KEY,
   user_id UUID NOT NULL REFERENCES users(id),
   address VARCHAR(255) NOT NULL,
   number VARCHAR(255) NOT NULL,
   zip_code VARCHAR(255) NOT NULL,
   city VARCHAR(255) NOT NULL,
   state VARCHAR(255) NOT NULL
);

CREATE TABLE  products (
   id SERIAL PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   value INTEGER NOT NULL
);

CREATE TABLE orders (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    product_id INTEGER NOT NULL REFERENCES products(id),
    quantity INTEGER NOT NULL,
    total_value INTEGER NOT NULL,
    address_id UUID NOT NULL REFERENCES addresses(id)
);


