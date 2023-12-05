-- public.orders definition

-- Drop table

-- DROP TABLE public.orders;

CREATE TABLE public.orders (
	orders_id serial4 NOT NULL,
	users_id int4 NOT NULL,
	deliveries_id int4 NOT NULL,
	promos_id int4 NOT NULL DEFAULT 0,
	payment_methods_id int4 NOT NULL,
	orders_status public."status_type" NOT NULL DEFAULT 'Pending'::status_type,
	orders_total int8 NOT NULL DEFAULT 0,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT pk_orders PRIMARY KEY (orders_id)
);


-- public.orders foreign keys

ALTER TABLE public.orders ADD CONSTRAINT fk_orders_deliveries FOREIGN KEY (deliveries_id) REFERENCES public.deliveries(deliveries_id);
ALTER TABLE public.orders ADD CONSTRAINT fk_orders_payment_methods FOREIGN KEY (payment_methods_id) REFERENCES public.payment_methods(payment_methods_id);
ALTER TABLE public.orders ADD CONSTRAINT fk_orders_promos FOREIGN KEY (promos_id) REFERENCES public.promos(promos_id);
ALTER TABLE public.orders ADD CONSTRAINT fk_orders_users FOREIGN KEY (users_id) REFERENCES public.users(users_id);