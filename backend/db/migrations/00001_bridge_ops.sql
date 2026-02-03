-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS bridge_ops (
	id TEXT PRIMARY KEY,
	from_terrace TEXT NOT NULL,
	from_domain INT NOT NULL,
    from_chain_id TEXT NOT NULL,
	to_terrace TEXT NOT NULL,
	to_domain INT NOT NULL,
    to_chain_id TEXT,
	send_amount TEXT NOT NULL,
    received_amount TEXT,
    sent_tx_hash TEXT NOT NULL,
    received_tx_hash TEXT,
	cctp_message TEXT,
	cctp_attestation TEXT,
	sent_at TIMESTAMPTZ NOT NULL,
    received_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS bridge_ops;
-- +goose StatementEnd
