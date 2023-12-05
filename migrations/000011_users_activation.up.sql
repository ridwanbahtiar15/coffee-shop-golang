-- public.users_activation definition

-- Drop table

-- DROP TABLE public.users_activation;

CREATE TABLE public.users_activation (
	users_activation_id serial4 NOT NULL,
	users_id int4 NOT NULL,
	"token" varchar(20) NOT NULL,
	CONSTRAINT pk_users_activation PRIMARY KEY (users_activation_id)
);


-- public.users_activation foreign keys

ALTER TABLE public.users_activation ADD CONSTRAINT fk_users_activation_users FOREIGN KEY (users_id) REFERENCES public.users(users_id);