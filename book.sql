use book;

create table books
(
    id          int auto_increment primary key,
    book        varchar(100)  null,
    author      varchar(100)  null,
    translator  varchar(100)  null,
    publisher   varchar(100)  null,
    cover       varchar(100)  null,
    tag         json          null comment '[]',
    reading     json          null comment '[{"start_time":"","end_time":""}]',
    reading_cnt int default 1 null,
    favour      json          null comment '[{"page":"","time":"","content":"","comment":""}]',
    favour_cnt  int default 0 null
);


create table users
(
    id          int auto_increment
        primary key,
    user_name   varchar(100)     null,
    user_pwd    varchar(200)     null,
    nick_name   varchar(100)     null,
    email       varchar(120)     null,
    permissions bigint default 0 null,
    constraint users_user_name_uindex
        unique (user_name)
);