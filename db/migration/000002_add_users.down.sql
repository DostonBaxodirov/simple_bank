ALTER TABLE IF EXISTS "account" DROP IF EXISTS "owner_currency_key";

ALTER TABLE IF EXISTS "account" DROP IF EXISTS "account_owner_fkey";

DROP TABLE IF EXISTS "users";