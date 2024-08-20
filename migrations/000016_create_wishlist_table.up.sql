create table "wishlist" (
    id serial primary key,
    event_id int REFERENCES events (id),
    user_id int references users (id)
);