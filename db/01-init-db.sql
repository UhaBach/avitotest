CREATE DATABASE avitotest;

CREATE TABLE users (
    id bigserial not null primary key,
    name varchar(255) not null check(name !='')
);

Create TABLE segments(
    id serial not null primary key,
    name varchar(255) not null unique check(name !='')
);

CREATE TABLE users_segments(
    user_id integer references users(id),
    segment_id integer references segments(id),
    constraint users_segments_pk PRIMARY KEY(user_id, segment_id)
);