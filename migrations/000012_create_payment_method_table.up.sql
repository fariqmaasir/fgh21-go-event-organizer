create table "payment_method" (
    id serial primary key,
    name varchar(80)
);

insert into
    "payment_method" (name)
values ('Card'),
    ('Bank Transfer'),
    ('Retail'),
    ('E-Money');