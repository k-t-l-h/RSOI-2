#!/bin/bash
set -e

export PGPASSWORD=test
psql -U program -d services <<-EOSQL

  CREATE TABLE warranty
  (
      id            SERIAL CONSTRAINT warranty_pkey PRIMARY KEY,
      comment       VARCHAR(1024),
      item_uid      UUID         NOT NULL CONSTRAINT idx_warranty_item_uid UNIQUE,
      status        VARCHAR(255) NOT NULL,
      warranty_date TIMESTAMP    NOT NULL
  );

  CREATE TABLE items
  (
      id              SERIAL CONSTRAINT items_pkey PRIMARY KEY,
      available_count INTEGER      NOT NULL,
      model           VARCHAR(255) NOT NULL,
      size            VARCHAR(255) NOT NULL
  );

  CREATE TABLE order_item
  (
      id             SERIAL CONSTRAINT order_item_pkey PRIMARY KEY,
      canceled       BOOLEAN,
      order_item_uid UUID NOT NULL CONSTRAINT idx_order_item_order_item_uid UNIQUE,
      order_uid      UUID NOT NULL,
      item_id        INTEGER CONSTRAINT fk_order_item_item_id REFERENCES items
  );

  CREATE TABLE orders
  (
      id         SERIAL CONSTRAINT orders_pkey PRIMARY KEY,
      item_uid   UUID         NOT NULL,
      order_date TIMESTAMP    NOT NULL,
      order_uid  UUID         NOT NULL CONSTRAINT idx_orders_order_uid UNIQUE,
      status     VARCHAR(255) NOT NULL,
      user_uid   UUID         NOT NULL
  );

  CREATE TABLE users
  (
      id       SERIAL CONSTRAINT users_pkey PRIMARY KEY,
      name     VARCHAR(255) NOT NULL CONSTRAINT idx_user_name UNIQUE,
      user_uid UUID         NOT NULL CONSTRAINT idx_user_user_uid UNIQUE
  );

  INSERT INTO items(id,available_count,	model,	size)
  VALUES(1, 1000, 'Lego 8070', 'M');
  INSERT INTO items(id,available_count,	model,	size)
  VALUES(3, 1000, 'Lego 42070', 'L');
  INSERT INTO items(id,available_count,	model,	size)
  VALUES(2, 1000, 'Lego 8880', 'L');

  INSERT INTO users(id,name,	user_uid)
  VALUES(1, 'Alex', '6d2cb5a0-943c-4b96-9aa6-89eac7bdfd2b');
EOSQL
