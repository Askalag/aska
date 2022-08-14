create table if not exists auth.users (
    id serial unique not null,
    login varchar(50) unique not null,
    first_name varchar(100),
    last_name varchar(100),
    password varchar(500) not null,
    email varchar(50) unique not null,
    active boolean default true,
    date_created timestamp default now(),
    last_modified timestamp default current_timestamp
)