CREATE TABLE "go-coffeeshop".roles (
	roles_id serial4 NOT NULL,
	roles_name varchar(100) NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT pk_roles PRIMARY KEY (roles_id)
);