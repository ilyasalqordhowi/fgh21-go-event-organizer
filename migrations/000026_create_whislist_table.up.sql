create table "whislist"(
"id" serial primary key,
"user_id" int references "users"("id"),
"event_id" int references "events"("id")
);