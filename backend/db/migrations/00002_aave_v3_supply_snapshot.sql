-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS aave_v3_supply_snapshot (
	id BIGSERIAL PRIMARY KEY,
	network TEXT NOT NULL,
	apy TEXT NOT NULL,
	timestamp TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS aave_v3_supply_snapshot;
-- +goose StatementEnd
