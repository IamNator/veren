create table if not exists user
(
    id              int unsigned auto_increment,
    first_name      varchar(50)                         not null,
    middle_name     varchar(50)                         null,
    last_name       varchar(50)                         not null,
    location        text                                null,
    dob             datetime                            null,
    modified_at     timestamp default CURRENT_TIMESTAMP not null,
    hashed_password text                                null,
    constraint user_id_uindex
        unique (id)
)
    comment 'holds user information';

alter table user
    add primary key (id);

create table if not exists device
(
    id          int unsigned auto_increment,
    user_id     int unsigned                        not null,
    name        varchar(50)                         not null,
    description text                                null,
    modified_at timestamp default CURRENT_TIMESTAMP not null,
    constraint device_id_uindex
        unique (id),
    constraint device_user_id_fk
        foreign key (user_id) references user (id)
            on update cascade on delete cascade
)
    comment 'saves user gadgets';

alter table device
    add primary key (id);

create table if not exists log
(
    id          int unsigned auto_increment,
    device_id   int unsigned                        not null,
    state       tinyint(1)                          null comment 'saves the state of the device
',
    modified_at timestamp default CURRENT_TIMESTAMP not null,
    constraint log_id_uindex
        unique (id),
    constraint log_device_id_fk
        foreign key (device_id) references device (id)
            on update cascade on delete cascade
)
    comment 'saves the state of users device';

alter table log
    add primary key (id);


