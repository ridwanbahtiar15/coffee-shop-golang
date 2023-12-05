CREATE TABLE "go-coffeeshop".categories (
	categories_id int4 NOT NULL DEFAULT nextval('categorys_categorys_id_seq'::regclass),
	categories_name varchar(100) NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT pk_categories PRIMARY KEY (categories_id)
);