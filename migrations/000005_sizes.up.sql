CREATE TABLE "go-coffeeshop".sizes (
	sizes_id serial4 NOT NULL,
	sizes_name varchar(100) NOT NULL,
	sizes_cost int8 NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT pk_sizes PRIMARY KEY (sizes_id)
);