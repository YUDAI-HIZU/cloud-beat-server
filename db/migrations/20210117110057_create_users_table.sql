
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table users (
    id                 bigint(20) not null auto_increment primary key,
    uid                varchar(255) unique not null,
    display_name       varchar(255) default null,
    web_url            varchar(255) default null,
    introduction       varchar(255) default null,
    icon_name    varchar(255) default null,
    cover_name   varchar(255) default null,
    created_at         datetime not null,
    updated_at         datetime not null
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table users;