-- +migrate Up

CREATE TABLE accounts (
    id TEXT PRIMARY KEY,
    document_num INT NOT NULL
);

-- +migrate Down

DROP TABLE accounts;
