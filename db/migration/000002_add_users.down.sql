ALTER TABLE IF EXISTS "account" DROP CONSTRAINT IF EXISTS account_owner_fkey;

DROP INDEX IF EXISTS account_owner_currency_idx;

DROP TABLE IF EXISTS "users";