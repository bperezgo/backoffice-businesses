CREATE TABLE IF NOT EXISTS orders_products (
  order_id uuid NOT NULL,
  product_id uuid NOT NULL,
  PRIMARY KEY (order_id, product_id),
  CONSTRAINT fk_order
    FOREIGN KEY(order_id) 
	    REFERENCES orders(id),
  CONSTRAINT fk_product
    FOREIGN KEY(product_id) 
      REFERENCES products(id)
);
