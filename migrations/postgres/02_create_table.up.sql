create table if not exists urls (
    id uuid primary key default gen_random_uuid() not null,
    title varchar default "" not null,
    url varchar not null,
    short_url varchar(254) primary key not null,
    expires_at timestamp,
    expires_count bigint,
    cheng_count bigint,
    created_by uuid references users(id) not null,
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null,
);