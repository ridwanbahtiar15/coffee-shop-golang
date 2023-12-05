CREATE TABLE "go-coffeeshop".deliveries (
	deliveries_id int4 NOT NULL DEFAULT nextval('deliverys_deliverys_id_seq'::regclass),
	deliveries_name varchar(100) NOT NULL,
	deliveries_cost int4 NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT pk_deliveries PRIMARY KEY (deliveries_id)
);