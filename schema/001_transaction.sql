-- +migrate Up

--role
CREATE TABLE IF NOT EXISTS "transaction"(
   id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
   store_id BIGINT,
   terminal_id BIGINT,
   amount  FLOAT,
   account FLOAT,
  
    created_at TIMESTAMP DEFAULT (NOW()),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);
-- +migrate Down
--table
DROP TABLE IF EXISTS "transaction";

