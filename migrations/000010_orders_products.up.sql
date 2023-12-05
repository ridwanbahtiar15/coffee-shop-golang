-- public.orders_products definition

-- Drop table

-- DROP TABLE public.orders_products;

CREATE TABLE public.orders_products (
	orders_products_id serial4 NOT NULL,
	orders_id int4 NOT NULL,
	products_id int4 NOT NULL,
	sizes_id int4 NOT NULL,
	orders_products_qty int8 NOT NULL,
	orders_products_subtotal int8 NOT NULL DEFAULT 0,
	hot_or_ice public."hot_or_ice_type" NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT pk_orders_products PRIMARY KEY (orders_products_id)
);


-- public.orders_products foreign keys

ALTER TABLE public.orders_products ADD CONSTRAINT fk_orders_products_orders FOREIGN KEY (orders_id) REFERENCES public.orders(orders_id);
ALTER TABLE public.orders_products ADD CONSTRAINT fk_orders_products_products FOREIGN KEY (products_id) REFERENCES public.products(products_id);
ALTER TABLE public.orders_products ADD CONSTRAINT fk_orders_products_sizes FOREIGN KEY (sizes_id) REFERENCES public.sizes(sizes_id);