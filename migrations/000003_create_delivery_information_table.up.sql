CREATE TABLE IF NOT EXISTS delivery_information (
  id uuid,
  user_id uuid NOT NULL,
  address VARCHAR(255) NOT NULL,
  city VARCHAR(32) NOT NULL,
  state VARCHAR(64) NOT NULL,
  country VARCHAR(64) NOT NULL,
  phone VARCHAR(32) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_user
    FOREIGN KEY(user_id) 
	    REFERENCES users(id)
);
