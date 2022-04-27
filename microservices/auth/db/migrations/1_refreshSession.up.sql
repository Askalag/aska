create table if not exists refresh_session (
    id serial unique not null,
    user_id int not null references users(id) on delete cascade,
    refresh_token uuid unique not null,
    ip varchar(15) not null,
    expires_in timestamp not null,
    created_at timestamp not null default now()
)