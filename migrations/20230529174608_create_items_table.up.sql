CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS items
(
    i_id uuid default uuid_generate_v4() not null primary key,
    name text not null,
    price integer not null,
    unit text not null,
    created_time timestamp default now() not null
);
