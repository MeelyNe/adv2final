CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(50) UNIQUE,
    password   VARCHAR(255),
    email      VARCHAR(100) UNIQUE,
    full_name  VARCHAR(100),
    created_at TIMESTAMP
);

CREATE TABLE accounts
(
    id             SERIAL PRIMARY KEY,
    user_id        INTEGER REFERENCES users (id),
    account_number VARCHAR(20) UNIQUE,
    balance        DECIMAL(10, 2),
    created_at     TIMESTAMP
);

CREATE TABLE transactions
(
    id         SERIAL PRIMARY KEY,
    account_id INTEGER REFERENCES accounts (id),
    type       VARCHAR(20),
    amount     DECIMAL(10, 2),
    timestamp  TIMESTAMP
);
