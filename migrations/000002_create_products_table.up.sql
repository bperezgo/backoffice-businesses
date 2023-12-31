CREATE TABLE IF NOT EXISTS products (
  id uuid,
  sku VARCHAR(16) NOT NULL,
  name VARCHAR(127) NOT NULL,
  description VARCHAR(511) NOT NULL,
  price NUMERIC NOT NULL,
  stock INTEGER NOT NULL,
  available_media JSON NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  PRIMARY KEY (id)
);