CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "price" bigint NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "product_id" bigint NOT NULL,
  "total_price" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "orders" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

CREATE INDEX ON "products" ("id");

CREATE INDEX ON "products" ("name");

CREATE INDEX ON "orders" ("code");

CREATE INDEX ON "orders" ("product_id");
