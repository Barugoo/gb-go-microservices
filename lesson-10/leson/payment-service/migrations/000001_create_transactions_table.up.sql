CREATE TABLE IF NOT EXISTS transactions (
    id integer serial primary key,
    external_id text null,
    amount integer not null,
    status varchar(255) not null,
    created_at timestamp not null default NOW(),
    updated_at timestamp not null default NOW(),
)


CREATE TABLE IF NOT EXISTS movie_ownership (
    id integer serial primary key,
    movie_id integer not null,
    user_id integer not null,
    created_at timestamp not null default NOW(),
    updated_at timestamp not null default NOW(),
)

