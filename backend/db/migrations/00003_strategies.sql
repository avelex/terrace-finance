-- +goose Up
-- +goose StatementBegin
CREATE TABLE strategies (
    id SERIAL PRIMARY KEY,
    domain INT NOT NULL,
    name TEXT NOT NULL,
    target_address TEXT NOT NULL,
    selector TEXT NOT NULL,
    proof TEXT[] NOT NULL
);

INSERT INTO strategies (domain, name, target_address, selector, proof) VALUES 
(0, 'approve', '0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238', '095ea7b3', '{"84b601305c2f80bfffbe275104c5cd778188f3499f00b168e25a8d490b82b9f7", "afeccea8d162dec84dd57182707e7faaaf384d00bd1b2b67c81a0d94978ed99d"}'),
(0, 'supply', '0x6Ae43d3271ff6888e7Fc43Fd7321a503ff738951', '617ba037', '{"4a6e631487d4db53ef5ec312de001ed3f2fdcc27f2a6e945b86b087f2d8bc320", "afeccea8d162dec84dd57182707e7faaaf384d00bd1b2b67c81a0d94978ed99d"}'),
(0, 'withdraw', '0x6Ae43d3271ff6888e7Fc43Fd7321a503ff738951', '69328dec', '{"0x01367063546e523ec46cd5c696b5569fb9516373c2874a8d83d561283142c581"}');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE strategies;
-- +goose StatementEnd
