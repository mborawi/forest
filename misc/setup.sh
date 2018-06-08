#! /bin/bash

go run generator.go && 
psql -d heat -f schema.sql &&
psql -d heat -c "\copy employees (id,first_name,last_name,email,phone,job_title,branch_id,manager_id,start_date) FROM 'employees.csv' WITH (FORMAT csv);" && 
psql -d heat -c "\copy leave_ranges (id,employee_id,start,finish,hours,name,category,type) FROM 'LeaveRanges.csv' WITH (FORMAT csv);" && 
psql -d heat -c "\copy leave_names (id,name) FROM 'lookupData/LeaveNames.csv' WITH (FORMAT csv);" &&
psql -d heat -c "\copy leave_categories (id,name) FROM 'lookupData/LeaveCategories.csv' WITH (FORMAT csv);" &&
psql -d heat -c "\copy leave_types (id,name) FROM  'lookupData/LeaveTypes.csv' WITH (FORMAT csv);"&&
psql -d heat -c "INSERT INTO leaves(employee_id, leave_date,leave_name_id, leave_type_id, leave_category_id) SELECT employee_id, generate_series(start,finish,'1 day') as leave_date, GetLeaveNameId(name) as leave_name_id ,GetLeaveTypeId(type) as leave_type_id, GetLeaveCategoryId(category) AS leave_catgeory_id FROM leave_ranges;"&&
psql -d heat -c "DELETE FROM leaves WHERE EXTRACT(dow FROM leave_date) IN (0, 6);"&&
psql -d heat -c "UPDATE employees SET full_name = first_name || ' ' || last_name;"&&
psql -d heat -c "CREATE EXTENSION btree_gist;"&&
psql -d heat -c "CREATE EXTENSION pg_trgm CASCADE;"&&
psql -d heat -c "CREATE INDEX full_name_idx ON employees USING GIST (full_name gist_trgm_ops);"


