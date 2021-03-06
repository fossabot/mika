package postgres

const schema = `
create table torrent
(
    info_hash bytea check (octet_length(info_hash) = 20) not null primary key,
    release_name varchar(255) not null,
    total_uploaded int default 0 not null,
    total_downloaded int default 0 not null,
    total_completed smallint default 0 not null,
    is_deleted bool default 'f' not null,
    is_enabled bool default 't' not null,
    reason varchar(255) default '' not null,
    multi_up decimal(5,2) default 1.00 not null,
    multi_dn decimal(5,2) default 1.00 not null,
    constraint uq_release_name
        unique (release_name)
);

create table users
(
    user_id SERIAL
        primary key,
    passkey varchar(20) not null,
    download_enabled bool default 't' not null,
    is_deleted bool default 'f' not null,
    constraint user_passkey_uindex
        unique (passkey)
);

create table peers
(
    peer_id bytea  check (octet_length(peer_id) = 20) not null,
    info_hash bytea  check (octet_length(info_hash) = 20) not null,
    user_id int not null,
    addr_ip int not null,
    addr_port smallint not null,
    total_downloaded int default 0 not null,
    total_uploaded int default 0 not null,
    total_left int default 0 not null,
    total_time int default 0 not null,
    total_announces int default 0 not null,
    speed_up int default 0 not null,
    speed_dn int default 0 not null,
    speed_up_max int default 0 not null,
    speed_dn_max int default 0 not null,
    location point not null,
    created_on timestamp not null,
    updated_on timestamp not null,
    primary key (info_hash, peer_id),
    constraint peers_torrent_fk
        foreign key (info_hash) references torrent (info_hash)
            on update cascade on delete cascade,
    constraint peers_user_id_fk
        foreign key (user_id) references users (user_id)
            on update cascade on delete cascade
);

create table whitelist
(
    client_prefix varchar(10) not null
        primary key,
    client_name varchar(20) not null
);
`
