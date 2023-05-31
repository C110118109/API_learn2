CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS requests
(
    request_id  uuid default uuid_generate_v4() not null primary key,
    employee_id uuid not null,
    request_itemlist_id uuid not null,
    reason text not null,
    request_date timestamp not null,
    created_by uuid null ,
    created_time timestamp default now() not null
);