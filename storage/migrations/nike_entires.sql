
create table if not exists mrporter.nike_entries(
    id serial not null unique ,
    launch varchar,
    username varchar,
    password varchar,
    entry_time timestamp with time zone,
    status varchar,
    entered int,
    style_id varchar,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    unique (username, launch)
)