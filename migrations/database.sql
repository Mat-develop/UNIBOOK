CREATE DATABASE IF NOT EXISTS unibook;
USE unibook;

CREATE TABLE IF NOT EXISTS users(
    id int auto_increment primary key unique,
    name varchar(50) not null, 
    nick varchar(50) not null unique, 
    email varchar(50) not null unique,
    image_url varchar(255),
    password varchar(100) not null, 
    created_at timestamp default current_timestamp()
) ENGINE=INNODB;