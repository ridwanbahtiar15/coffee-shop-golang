CREATE TABLE "go-coffeeshop".payment_methods (
	payment_methods_id serial4 NOT NULL,
	payment_methods_name varchar(100) NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT pk_payment_methods PRIMARY KEY (payment_methods_id)
);