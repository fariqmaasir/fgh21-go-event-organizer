create table "event_sections" (
    "id" serial primary key,
    "name" varchar(100),
    "price" int,
    "quantity" int,
    "event_id" int references "events" ("id")
);

insert into
    "event_sections" (
        name,
        price,
        quantity,
        event_id
    )
values ('REG', 500000, 10, 1),
    ('VIP', 700000, 10, 1),
    ('VVIP', 1000000, 10, 1),
    ('REG', 500000, 10, 2),
    ('VIP', 700000, 10, 2),
    ('VVIP', 1000000, 10, 2),
    ('REG', 500000, 10, 3),
    ('VIP', 700000, 10, 3),
    ('VVIP', 1000000, 10, 3),
    ('REG', 500000, 10, 4),
    ('VIP', 700000, 10, 4),
    ('VVIP', 1000000, 10, 4),
    ('REG', 500000, 10, 5),
    ('VIP', 700000, 10, 5),
    ('VVIP', 1000000, 10, 5);