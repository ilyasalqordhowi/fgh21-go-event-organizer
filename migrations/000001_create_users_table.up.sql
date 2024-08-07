create table "users"(
    "id" serial primary key,
    "email" varchar(80) unique,
    "password" varchar(80),
    "username" varchar(80)
);