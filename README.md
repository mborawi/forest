# heat

$> createdb heat


$> psql -d heat -f misc/schema.sql
$> psql -d heat -c "\copy employees FROM 'misc/employees.csv' WITH (FORMAT csv);"
$> psql -d heat -c "\copy leaves FROM 'misc/leaves.csv' WITH (FORMAT csv);"

Or 
 $> cd misc
 $> ./setup.sh