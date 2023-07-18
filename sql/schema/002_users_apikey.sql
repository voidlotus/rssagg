-- +goose Up
/*
generating some random bytes and casting it into a byte array
and we are using sha256 hash function to get fixed size output
and lastly encoded it in hexadecimal
*/
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT(
    encode(sha256(random()::text::bytea), 'hex')
);


-- +goose Down
ALTER TABLE users DROP COLUMN api_key;