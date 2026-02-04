-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_deposits (
    id UUID PRIMARY KEY,
    address TEXT NOT NULL,
    value TEXT NOT NULL,
    dest_domain INT NOT NULL,
    dest_gateway_mint TEXT NOT NULL,
    attestation TEXT,
    signature TEXT,
    deposit_tx_hash TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    unified_at TIMESTAMPTZ,
    deposited_at TIMESTAMPTZ,
    PRIMARY KEY (id, owner)
);

CREATE TABLE user_unified_permits (
    deposit_id UUID REFERENCES user_deposits(id),
    owner TEXT NOT NULL,
    token TEXT NOT NULL,
    value TEXT NOT NULL,
    deadline BIGINT NOT NULL,
    signature TEXT,
    domain INT NOT NULL,
    gateway_wallet TEXT NOT NULL,
    tx_hash TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deposited_at TIMESTAMPTZ,
    PRIMARY KEY (deposit_id, owner, domain)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_deposits;
DROP TABLE user_unified_permits;
-- +goose StatementEnd
