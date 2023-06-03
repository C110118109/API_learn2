CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS request_itemlists
(
    ri_id uuid default uuid_generate_v4() not null primary key,
    request_id uuid not null,
    name text not null,
    application text not null,
    quanity integer not null,
    unit text not null,
    price integer not null,
    total integer not null,
    created_by uuid null ,
    created_time timestamp default now() not null
);
