create table "users"(
    "id" serial primary key,
    "email" varchar(80) unique,
    "password" varchar(255),
    "username" varchar(80)
);