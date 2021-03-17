


create table if not exists nike_accounts (
     id serial not null unique,
     login varchar,
     password varchar,
     priority integer,
     created_at timestamp with time zone,
     updated_at timestamp with time zone,
     deleted_at timestamp with time zone,
     unique (login)
)