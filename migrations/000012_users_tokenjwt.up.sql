-- public.users_tokenjwt definition

-- Drop table

-- DROP TABLE public.users_tokenjwt;

CREATE TABLE public.users_tokenjwt (
	users_tokenjwt_id serial4 NOT NULL,
	users_id int4 NOT NULL,
	token_jwt varchar(300) NOT NULL,
	CONSTRAINT pk_users_tokenjwt PRIMARY KEY (users_tokenjwt_id)
);


-- public.users_tokenjwt foreign keys

ALTER TABLE public.users_tokenjwt ADD CONSTRAINT fk_users_tokenjwt_users FOREIGN KEY (users_id) REFERENCES public.users(users_id);