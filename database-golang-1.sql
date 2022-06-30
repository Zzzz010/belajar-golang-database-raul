use golang_mysql_raul;
use belajar_golang_database;

CREATE TABLE customer (
    id VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
)  ENGINE=INNODB;

desc customer;
show create table customer;
select * from customer;

delete from customer;

alter table customer
add column email varchar(100),
add column balance integer default 0,
add column rating double default 0.0,
add column created_at timestamp default current_timestamp,
add column birth_date date,
add column married boolean default false;

insert into customer (id, name, email, balance, rating, birth_date, married)
values ('A001', 'Raung', 'raung@gmail.com', 100000, 4.8, '2000-11-05', false),
('A002', 'Jayan', 'jayan@gmail.com', 1000000, 5.0, '2000-05-11', true);

insert into customer (id, name, email, balance, rating, birth_date, married)
values ('A004', 'Zzzz', null, 5000, 2.5, '2004-12-02', true);

create table user
(
username varchar(100) not null,
password varchar(100) not null,
primary key (username)
) engine = innodb;

insert into user(username, password) values('admin', 'admin');

create table comments
(
id int not null auto_increment,
email varchar(100) not null,
comment text,
primary key (id)
) engine = innoDB;
