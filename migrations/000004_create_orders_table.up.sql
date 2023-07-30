CREATE TABLE IF NOT EXISTS orders (
  id uuid,
  user_id uuid NOT NULL,
  delivery_information_id uuid NOT NULL,
  delivery_date VARCHAR(32) NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_user
    FOREIGN KEY(user_id) 
	    REFERENCES users(id),
  CONSTRAINT fk_delivery_information
    FOREIGN KEY(delivery_information_id) 
      REFERENCES delivery_information(id)
);
