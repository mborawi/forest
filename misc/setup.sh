#! /bin/bash

go run generator.go && 
psql -d heat -f schema.sql &&
psql -d heat -c "\copy employees (id,first_name,last_name,email,phone,job_title,branch_id,manager_id) FROM 'employees.csv' WITH (FORMAT csv);" && 
psql -d heat -c "\copy leave_ranges (id,employee_id,start,finish,name,category,type) FROM 'LeaveRanges.csv' WITH (FORMAT csv);" && 
psql -d heat -c "\copy leave_names (id,name) FROM 'lookupData/LeaveNames.csv' WITH (FORMAT csv);" &&
psql -d heat -c "\copy leave_categories (id,name) FROM 'lookupData/LeaveCategories.csv' WITH (FORMAT csv);" &&
psql -d heat -c "\copy leave_types (id,name) FROM  'lookupData/LeaveTypes.csv' WITH (FORMAT csv);"