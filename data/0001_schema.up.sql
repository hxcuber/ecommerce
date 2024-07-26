CREATE TYPE STATUS AS ENUM ('processed','shipped','delivered');

CREATE TABLE IF NOT EXISTS users (
  user_id UUID PRIMARY KEY,
  username TEXT UNIQUE NOT NULL,
  email TEXT UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS categories (
  category_id BIGSERIAL PRIMARY KEY,
  description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
  product_id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  price REAL NOT NULL,
  stock BIGINT NOT NULL,
  category_id BIGINT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  FOREIGN KEY (category_id) REFERENCES categories (category_id)
);

CREATE TABLE IF NOT EXISTS orders (
  order_id UUID PRIMARY KEY,
  user_id UUID NOT NULL,
  total_amount BIGINT NOT NULL,
  shipping_address TEXT NOT NULL,
  status STATUS NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE IF NOT EXISTS order_items (
  order_id UUID,
  product_id UUID,
  quantity BIGINT NOT NULL,
  PRIMARY KEY (order_id, product_id),
  FOREIGN KEY (order_id) references orders (order_id),
  FOREIGN KEY (product_id) references products (product_id)
);

INSERT INTO categories (description) VALUES ('cat1');
INSERT INTO categories (description) VALUES ('cat2');
INSERT INTO categories (description) VALUES ('cat3');
INSERT INTO categories (description) VALUES ('cat4');
INSERT INTO categories (description) VALUES ('cat5');
INSERT INTO categories (description) VALUES ('cat6');
INSERT INTO categories (description) VALUES ('cat7');
INSERT INTO categories (description) VALUES ('cat8');
INSERT INTO categories (description) VALUES ('cat9');
INSERT INTO categories (description) VALUES ('cat10');
