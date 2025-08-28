ALTER TABLE products 
ADD CONSTRAINT fk_products_user_id 
FOREIGN KEY (user_id) REFERENCES users(id) 
ON DELETE CASCADE 
ON UPDATE CASCADE;

CREATE INDEX idx_products_user_id ON products(user_id);