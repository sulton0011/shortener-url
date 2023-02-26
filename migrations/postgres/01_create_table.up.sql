create table if not exists users (
    id uuid primary key default gen_random_uuid() not null,
    name varchar not null,
    surname varchar not null,
    middle_name varchar not null,
    email varchar not null,
    login varchar unique not null,
    password varchar not null,
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null
);
