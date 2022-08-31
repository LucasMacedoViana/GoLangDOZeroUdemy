CREATE DATABASE IF NOT EXISTS devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id integer serial primary key,
    nome varchar (255) not null,
    nick varchar (255) not null unique,
    email varchar(255) not null unique,
    senha varchar (255) not null,
    criadoEm timestamp default current_timestamp
) ENGINE=INNODB
