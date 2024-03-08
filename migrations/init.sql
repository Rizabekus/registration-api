-- CREATE DATABASE IF NOT EXISTS mobydev;



CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NULL,
    email VARCHAR(255) NOT NULL,
    mobile_number VARCHAR(255) NULL,
    date_of_birth DATE  NULL,
    hashed_password VARCHAR(255) NOT NULL,
    usertype VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS session_cookies(
    user_id INT REFERENCES users(id),
    session VARCHAR(255) NOT NULL
);