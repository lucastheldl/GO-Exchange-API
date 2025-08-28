DROP INDEX IF EXISTS idx_products_user_id;

ALTER TABLE products 
DROP CONSTRAINT IF EXISTS fk_products_user_id;