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
(6, 'approve', '0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913', '095ea7b3', '{"dcced49bcc2022f18d087b37325862e1d8ddf620b4b2f07db1061b85900ef3b0"}'),
(6, 'supply', '0xA238Dd80C259a72e81d7e4664a9801593F98d1c5', '617ba037', '{"220f2c6e9fc386a9244e977253160f39168afa8282a15c86cd4ffa9586f67d52", "8b5f85f2e3e9cc39005c9de4107342626375e40d13bca8c27e8fc9ed8709186e"}'),
(6, 'withdraw', '0xA238Dd80C259a72e81d7e4664a9801593F98d1c5', '69328dec', '{"3decc60727aad9ec67796601fd262350e023d527bb46e21ba4fc987dc91684ad", "8b5f85f2e3e9cc39005c9de4107342626375e40d13bca8c27e8fc9ed8709186e"}');

INSERT INTO strategies (domain, name, target_address, selector, proof) VALUES 
(3, 'approve', '0xaf88d065e77c8cC2239327C5EDb3A432268e5831', '095ea7b3', '{"08b758f49aa00413bb4baf65678a9c87f261c3e1a07337927ed23be610ac93aa", "e684453d432b92ee35ff296de2b9c1d7c99b5d8dd4fe4b1ff9872e822c346eed"}'),
(3, 'supply', '0x794a61358D6845594F94dc1DB02A252b5b4814aD', '617ba037', '{"3ba87b9193cd227d49d88e1944b9fc11598446411ac53f66fd603f024ff9ab91", "e684453d432b92ee35ff296de2b9c1d7c99b5d8dd4fe4b1ff9872e822c346eed"}'),
(3, 'withdraw', '0x794a61358D6845594F94dc1DB02A252b5b4814aD', '69328dec', '{"45ac8193aa47ae9131cdadb3abc4ce3a936875d88ff1df730dbdc3b9f4b97ae0"}');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE strategies;
-- +goose StatementEnd
