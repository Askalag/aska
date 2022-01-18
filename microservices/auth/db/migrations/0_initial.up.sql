create table if not exists users (
    id serial unique not null,
    login varchar(50) unique not null,
    name varchar(100),
    password varchar(500) not null,
    email varchar(50) unique not null,
    active boolean default false,
    date_created timestamp default now(),
    date_modified timestamp default current_timestamp
)