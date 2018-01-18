create table users(
    id serial primary key,
    first_name varchar(50) not null,
    last_name varchar(50) not null,
    email varchar(250) not null,
    password char(64) not null,
    salt char(32) not null
);
