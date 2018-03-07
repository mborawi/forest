DROP TABLE IF EXISTS leaves CASCADE;
 create table leaves (
	id SERIAL PRIMARY KEY,
	employeeId Integer,
	LeaveDate date,
	LeaveId integer  
);
DROP TABLE IF EXISTS employees CASCADE;
 create table employees (
	id SERIAL PRIMARY KEY,
	first_name VARCHAR(50),
	last_name VARCHAR(50),
	email VARCHAR(80),
	phone VARCHAR(25),
	job_title VARCHAR(100),
	branchId integer,
	managerId integer 
);
