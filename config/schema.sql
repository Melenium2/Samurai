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
create type developerContacts as
(
    email    text,
    contacts text
);
create table if not exists meta_tracking
(
    id               bigserial primary key not null,
    bundleId         int references app_tracking,
    title            varchar(300),
    price            varchar(50),
    picture          varchar(192),
    screenshots      varchar(192)[],
    rating           varchar(50),
    reviewCount      varchar(50),
    ratingHistogram  varchar(50)[],
    description      text,
    shortDescription text,
    recentChanges    text,
    releaseDate      varchar(50),
    lastUpdateDate   varchar(50),
    appSize          varchar(50),
    installs         varchar(50),
    version          varchar(100),
    androidVersion   varchar(100),
    contentRating    varchar(100),
    devContacts      developerContacts,
    privacyPolicy    text,
    date             timestamp
);
