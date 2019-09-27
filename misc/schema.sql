DROP TABLE IF EXISTS leave_ranges CASCADE;
 CREATE TABLE leave_ranges (
	id SERIAL PRIMARY KEY,
	employee_id Integer,
	start timestamp,
	finish timestamp,
	hours decimal,
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
	leave_name_id integer, 
	leave_type_id integer, 
	leave_category_id integer, 
	created_at timestamp DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX date_idx ON leaves (leave_date);
CREATE INDEX employee_idx ON leaves (employee_id);


DROP TABLE IF EXISTS employees CASCADE;
 create table employees (
	id SERIAL PRIMARY KEY,
	first_name VARCHAR(50),
	last_name VARCHAR(50),
	full_name VARCHAR(100),
	email VARCHAR(80),
	phone VARCHAR(25),
	job_title VARCHAR(100),
	branch_id integer,
	manager_id integer,
	start_date timestamp,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS  leave_names CASCADE;
 create table leave_names(
 	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
 	created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS branches CASCADE;
 create table branches(
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

CREATE OR REPLACE FUNCTION GetLeaveNameId(input text) 
RETURNS INTEGER AS $leave_name_id$
DECLARE
	leave_name_id INTEGER;
BEGIN
	SELECT id INTO leave_name_id 
	FROM leave_names 
	WHERE name ILIKE input;
	RETURN leave_name_id;
END;
$leave_name_id$ LANGUAGE plpgsql VOLATILE;

CREATE OR REPLACE FUNCTION GetLeaveTypeId(input text) 
RETURNS INTEGER AS $leave_type_id$
DECLARE
	leave_type_id INTEGER;
BEGIN
	SELECT id INTO leave_type_id 
	FROM leave_types 
	WHERE name ILIKE input;
	RETURN leave_type_id;
END;
$leave_type_id$ LANGUAGE plpgsql VOLATILE;

CREATE OR REPLACE FUNCTION GetLeaveCategoryId(input text) 
RETURNS INTEGER AS $leave_category_id$
DECLARE
	leave_category_id INTEGER;
BEGIN
	SELECT id INTO leave_category_id 
	FROM leave_categories 
	WHERE name ILIKE input;
	RETURN leave_category_id;
END;
$leave_category_id$ LANGUAGE plpgsql VOLATILE;


CREATE OR REPLACE FUNCTION team_leaves(empId integer, nyrs integer, incMan boolean) 
	RETURNS TABLE (dom text, pcount bigint, ucount bigint)
    AS $$
    BEGIN

    IF incMan THEN
    	RETURN QUERY SELECT to_char(leave_date, 'dd-mm'),
			COUNT(leave_name_id=1 OR NULL) AS PCOUNT,
			COUNT(leave_name_id=2 OR NULL) AS UCOUNT 
			FROM employees 
			INNER JOIN leaves 
			ON employees.id = leaves.employee_id 
			WHERE  employees.manager_id = empId 
			OR employees.id = empId 
			AND leave_date >= now()-make_interval(years => nyrs)  
			GROUP BY to_char(leave_date, 'dd-mm')  
			ORDER BY COUNT(leave_name_id=1 OR NULL) +  COUNT(leave_name_id=2 OR NULL) DESC;
    ELSE
    	RETURN QUERY SELECT to_char(leave_date, 'dd-mm'),
			COUNT(leave_name_id=1 OR NULL) AS PCOUNT,
			COUNT(leave_name_id=2 OR NULL) AS UCOUNT 
			FROM employees 
			INNER JOIN leaves 
			ON employees.id = leaves.employee_id 
			WHERE  employees.manager_id=empId  
			AND leave_date >= now()-make_interval(years => nyrs)  
			GROUP BY to_char(leave_date, 'dd-mm')  
			ORDER BY COUNT(leave_name_id=1 OR NULL) +  COUNT(leave_name_id=2 OR NULL) DESC;
    END IF;
    END;
    $$ LANGUAGE plpgsql
    VOLATILE;



