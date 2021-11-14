CREATE DATABASE 'order';

USE 'order';

create table instruments
(
    id            varchar(36)                         not null,
    currency_code varchar(36)                         not null,
    code          varchar(16)                         not null,
    name          varchar(32)                         not null,
    description   varchar(256)                        null,
    lot           int                                 not null comment 'the minimum quantity that can be sold at a time',
    type          text                                not null comment 'the type of financial instrument, for example - a currency pair, stocks, abligations, bonds (for simplicity, I did not put it into the directory)',
    created_at    timestamp default CURRENT_TIMESTAMP not null,
    updated_at    timestamp                           null,
    deleted_at    timestamp                           null,
    constraint instrument_code_uindex
        unique (code),
    constraint instrument_id_uindex
        unique (id)
);

create index instruments_code_index
    on instruments (code);

alter table instruments
    add primary key (id);


create table orders
(
    id            varchar(36)                         not null,
    account_id    varchar(36)                         not null,
    instrument_id varchar(36)                         not null,
    type          varchar(64)                         not null comment 'type of order, for example: buy, sell (for simplicity, I did not put it in the directory)',
    price         decimal                             not null comment 'the price at which the user wants to purchase the instrument',
    volume        int                                 not null comment 'number of lots',
    status        varchar(64)                         null comment 'order status, for example: completed, canceled (for simplicity, I did not put it into the directory)',
    created_at    timestamp default CURRENT_TIMESTAMP not null,
    updated_at    timestamp                           null,
    constraint orders_order_id_uindex
        unique (id),
    constraint orders_instruments_id_fk
        foreign key (instrument_id) references instruments (id)
)
    comment 'User order table';

create index orders_account_id_instrument_id_type_index
    on orders (account_id, instrument_id, type);

create index orders_instrument_id_type_status_created_at_index
    on orders (instrument_id, type, status, created_at);

alter table orders
    add primary key (id);