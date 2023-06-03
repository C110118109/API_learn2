CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table employees
(
    e_id uuid default uuid_generate_v4() not null primary key,
    name text not null,
    department_id uuid not null,
    role_id uuid not null,
    created_time timestamp default now() not null
);

