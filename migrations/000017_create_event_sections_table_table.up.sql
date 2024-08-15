create table "event_sections"(
"id"      serial primary key,
"name"    varchar(255),
"price"   int,
"quantity" varchar(100),
"events_id" int references "events"("id")
);
