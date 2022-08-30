CREATE DATABASE IF NOT EXISTS devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id integer serial primary key,
    nome varchar (255) not null,
    nick varchar (255) not null unique,
    email varchar(255) not null unique,
    senha varchar (20) not null unique,
    criadoEm timestamp default current_timestamp
) ENGINE=INNODB