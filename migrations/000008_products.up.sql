CREATE TABLE public.products (
	products_id serial4 NOT NULL,
	products_name varchar(100) NOT NULL,
	products_price int4 NOT NULL,
	products_desc text NULL,
	products_stock int4 NOT NULL,
	products_image varchar NULL,
	categories_id int4 NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT pk_prducts PRIMARY KEY (products_id)
);


-- public.products foreign keys

ALTER TABLE public.products ADD CONSTRAINT fk_products_categorys FOREIGN KEY (categories_id) REFERENCES public.categories(categories_id);