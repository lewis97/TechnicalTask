-- +migrate Up

CREATE TABLE accounts (
    id TEXT PRIMARY KEY,
    document_num INT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL
);

CREATE TYPE operation_type AS ENUM ('INVALID', 'CASH PURCHASE', 'INSTALLMENT PURCHASE', 'WITHDRAWAL', 'PAYMENT');

CREATE TABLE transactions (
    id TEXT PRIMARY KEY,
    account_id TEXT NOT NULL,
    operation_id OPERATION_TYPE NOT NULL,
    amount INT NOT NULL,
    event_time TIMESTAMP NOT NULL,
    FOREIGN KEY (account_id) REFERENCES accounts(id)
        ON UPDATE CASCADE
);

-- +migrate Down

DROP TABLE accounts;
