create table "role"(
    "id" serial primary key,
    "name" varchar(80)
);
create table "users"(
    "id" serial primary key,
    "email" varchar(80) unique,
    "password" varchar(255),
    "username" varchar(80),
    "role_id" int REFERENCES"role"("id")
);