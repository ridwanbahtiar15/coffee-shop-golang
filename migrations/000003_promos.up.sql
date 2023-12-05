CREATE TABLE "go-coffeeshop".promos (
	promos_id serial4 NOT NULL,
	promos_name varchar(100) NOT NULL,
	promos_start varchar(25) NOT NULL,
	promos_end varchar(25) NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT pk_promos PRIMARY KEY (promos_id),
	CONSTRAINT promos_promos_name_key UNIQUE (promos_name)
);