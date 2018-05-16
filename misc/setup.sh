#! /bin/bash

go run generator.go && 
psql -d heat -f schema.sql &&
psql -d heat -c "\copy employees FROM 'employees.csv' WITH (FORMAT csv);" && 
psql -d heat -c "\copy lnames FROM 'LeaveNames.csv' WITH (FORMAT csv);"
psql -d heat -c "\copy lcategories FROM 'LeaveCategories.csv' WITH (FORMAT csv);"
psql -d heat -c "\copy ltypes FROM 'LeaveTypes.csv' WITH (FORMAT csv);"