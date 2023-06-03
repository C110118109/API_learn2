CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS departments
(
    d_id uuid default uuid_generate_v4() not null primary key,
    d_code text default add_department_code() not null,
    name text not null,
    created_time timestamp default now()  not null
);

