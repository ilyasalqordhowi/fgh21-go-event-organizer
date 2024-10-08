create table "events"(
"id" serial primary key,
"image"    varchar(255),
"title"  varchar(100),
"date"   varchar(50),
"descriptions" text,
"location_id" int references "location"("id"),
"created_by" int references "users"("id")
);