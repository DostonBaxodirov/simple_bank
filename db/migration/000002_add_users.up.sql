CREATE TABLE "users" (
 "username" varchar PRIMARY KEY,
 "full_name" varchar NOT NULL,
 "email" varchar UNIQUE NOT NULL,
 "hashed_password" varchar NOT NULL,
 "created_at" timestamptz NOT NULL DEFAULT 'now()',
 "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE UNIQUE INDEX ON "account" ("owner", "currency");
ALTER TABLE "account" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");