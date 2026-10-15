-- +goose Up
CREATE TYPE product_conditions AS ENUM ('новый', 'б/у');
CREATE TYPE listing_status AS ENUM('активно', 'отложено','закрыто','снято модератором');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE listing (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL CHECK(btrim(title) != ''),
    price INT NOT NULL CHECK (price >= 0),
    phone_number TEXT NOT NULL,
    whatsapp_number TEXT,
    category_id INT NOT NULL REFERENCES category(id),
    product_condition product_conditions NOT NULL,
    description TEXT NOT NULL CHECK(btrim(description) != ''),
    owner_id INT NOT NULL REFERENCES users(id),
    status listing_status NOT NULL DEFAULT 'активно',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL ,
    last_raised_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE photos(
    id SERIAL PRIMARY KEY,
    storage_key TEXT NOT NULL,
    listing_id INT REFERENCES listing(id) NOT NULL,
    position INT NOT NULL CHECK (position > 0),

    UNIQUE(listing_id, position)
);


-- +goose Down
DROP TABLE photos;
DROP TABLE listing;
DROP TABLE category;
DROP TABLE users;
DROP TYPE listing_status;
DROP TYPE product_conditions;
