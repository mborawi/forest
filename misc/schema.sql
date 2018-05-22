DROP TABLE IF EXISTS leave_ranges CASCADE;
 CREATE TABLE leave_ranges (
	id SERIAL PRIMARY KEY,
	employee_id Integer,
	start timestamp,
	finish timestamp,
	name VARCHAR(50),
	category VARCHAR(50),
	type VARCHAR(50),
	created_at timestamp DEFAULT CURRENT_TIMESTAMP
);


DROP TABLE IF EXISTS leaves CASCADE;
 CREATE TABLE leaves (
	id SERIAL PRIMARY KEY,
	employee_id Integer,
	leave_date timestamp,
	leave_duration Integer, -- zero for whole day,minutes should be only used for less than day leaves (max=7.22hrsx60)
	leave_id integer, 
	created_at timestamp DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX date_idx ON leaves (leave_date);
CREATE INDEX employee_idx ON leaves (employee_id);


DROP TABLE IF EXISTS employees CASCADE;
 create table employees (
	id SERIAL PRIMARY KEY,
	first_name VARCHAR(50),
	last_name VARCHAR(50),
	email VARCHAR(80),
	phone VARCHAR(25),
	job_title VARCHAR(100),
	branch_id integer,
	manager_id integer,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS  leave_names CASCADE;
 create table leave_names(
 	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
 	created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS leave_categories CASCADE;
 create table leave_categories(
 	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
 	created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS leave_types CASCADE;
 create table leave_types(
 	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
 	created_at timestamp DEFAULT CURRENT_TIMESTAMP
);
