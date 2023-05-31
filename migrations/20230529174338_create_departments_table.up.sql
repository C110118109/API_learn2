CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS departments
(
    department_id uuid default uuid_generate_v4() not null primary key,
    department_code text default add_department_code() not null,
    name text not null,
    created_by uuid null ,
    created_time timestamp default now()  not null
);

