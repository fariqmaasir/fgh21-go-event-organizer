create table "events" (
    "id" serial primary key,
    "image" varchar(225),
    "title" varchar(80),
    "date" timestamptz,
    "descriptions" text,
    "location_id" int references locations (id),
    "created_by" int references users (id)
);