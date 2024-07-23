CREATE TYPE STATUS AS ENUM ('processed','shipped','delivered');
CREATE TYPE CATEGORY AS ENUM ('default', 'interesting', 'testing');

CREATE TABLE IF NOT EXISTS users (
  user_id UUID PRIMARY KEY,
  username TEXT UNIQUE NOT NULL,
  email TEXT UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products (
  product_id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  price REAL NOT NULL,
  stock BIGINT NOT NULL,
  category_id CATEGORY NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS orders (
  order_id UUID PRIMARY KEY,
  user_id UUID NOT NULL,
  total_amount BIGINT NOT NULL,
  shipping_address TEXT NOT NULL,
  status STATUS NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY ( user_id ) REFERENCES users ( user_id )
);

CREATE TABLE IF NOT EXISTS order_items (
  order_id UUID,
  product_id UUID,
  quantity BIGINT,
  PRIMARY KEY (order_id, product_id),
  FOREIGN KEY (order_id) references orders (order_id),
  FOREIGN KEY (product_id) references products (product_id)
);
