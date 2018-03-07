#! /bin/bash

go run generator.go && psql -d heat -f schema.sql && psql -d heat -c "\copy employees FROM 'employees.csv' WITH (FORMAT csv);" && psql -d heat -c "\copy leaves FROM 'leaves.csv' WITH (FORMAT csv);"