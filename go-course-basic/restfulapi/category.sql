create database restful_api_golang;

create table category (
	id integer primary key auto_increment,
	name varchar(200) not null
) engine = InnoDB;

select * from category;