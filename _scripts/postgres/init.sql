GRANT ALL PRIVILEGES ON DATABASE bank TO dev;

CREATE TABLE transfers (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    account_origin_id VARCHAR(36),
    account_destination_id VARCHAR(36) NOT NULL,
    idempotency_key VARCHAR(36) NOT NULL,
    amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE accounts (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    name VARCHAR NOT NULL,
    document VARCHAR(50) UNIQUE NOT NULL,
    balance BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL
);