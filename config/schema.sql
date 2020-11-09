create table if not exists app_tracking
(
    id          serial primary key,
    bundle      varchar(255) not null,
    category    varchar(128) not null,
    developerId varchar(128),
    developer   varchar(500),
    Geo         varchar(10)  not null,
    startAt     timestamp,
    period      int
);

