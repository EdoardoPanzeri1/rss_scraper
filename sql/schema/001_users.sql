-- +goose UP
create table users (
    id int primary key,
    created_at text not null,
    updated_at text not null,
    name text not null
);

-- +goose Down
drop table users;