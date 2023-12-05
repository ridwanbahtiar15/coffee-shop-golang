CREATE TABLE public.users (
	users_id serial4 NOT NULL,
	users_fullname varchar(100) NOT NULL,
	users_email varchar(100) NOT NULL,
	users_password varchar(120) NOT NULL,
	users_phone varchar(15) NULL DEFAULT 0,
	users_address text NULL DEFAULT 0,
	users_image varchar NOT NULL DEFAULT 'profile.jpg'::character varying,
	roles_id int8 NOT NULL DEFAULT '2'::bigint,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	deleted_at timestamp NULL,
	is_active int4 NOT NULL DEFAULT 0,
	CONSTRAINT pk_users PRIMARY KEY (users_id),
	CONSTRAINT users_users_email_key UNIQUE (users_email),
	CONSTRAINT users_users_phone_key UNIQUE (users_phone)
);


-- public.users foreign keys

ALTER TABLE public.users ADD CONSTRAINT fk_users_roles FOREIGN KEY (roles_id) REFERENCES public.roles(roles_id);