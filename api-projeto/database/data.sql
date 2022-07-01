create table team(
	id_team serial primary key,
	name_team varchar(50)
)

create table project(
	id_projeto serial primary key,
	name varchar(50),
	description text
)

create table task(
	id_task serial primary key,
	name varchar(50),
	id_projeto int references project(id_projeto)
)

create table member(
	id serial primary key,
	name varchar(50),
	id_time int references team(id_team),
	id_task int references task(id_task)
)