


create table if not exists support_atc.users (
         id serial not null unique ,
         member_id varchar unique,
         token varchar ,
         private_channel varchar ,
         created_at timestamp with time zone,
         updated_at timestamp with time zone,
         deleted_at timestamp with time zone
)