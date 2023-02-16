DROP TABLE IF EXISTS product;
CREATE TABLE IF NOT EXISTS product(
    id bigserial PRIMARY KEY,
    name varchar(225) NOT NULL ,
    price DECIMAL(10,5) NOT NULL,
    description varchar(225) NOT NULL ,
    created_at timestamptz NOT NULL DEFAULT NOW()
);


CREATE INDEX idx_product_time ON product (created_at);
CREATE INDEX idx_product_name ON product (name);
CREATE INDEX idx_product_price ON product (price);


INSERT INTO product ( name, price, description) values ('Iphone 11', 10.00, 'test 1' );
INSERT INTO product ( name, price, description) values ('Iphone 12', 100.00, 'test 2' );


