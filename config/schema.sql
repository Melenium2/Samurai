create table if not exists app_tracking
(
    id          serial primary key not null,
    bundle      varchar(255)       not null,
    category    varchar(128)       not null,
    developerId varchar(128),
    developer   varchar(500),
    Geo         varchar(10)        not null,
    startAt     timestamp,
    period      int
);
create table if not exists category_tracking
(
    id       serial primary key not null,
    bundleId int references app_tracking,
    type     varchar(128)       not null,
    place    int                not null,
    date     timestamp          not null
);
create table if not exists keyword_tracking
(
    id       serial primary key not null,
    bundleId int references app_tracking,
    type     varchar(128)       not null,
    place    int                not null,
    date     timestamp          not null
);
