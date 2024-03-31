#!/bin/bash
export PGPASSWORD=$POSTGRES_PASSWORD

declare -a sql_files
sql_files=`ls ./schema`

echo "sql_files: ${sql_files}"

for i in ${sql_files[*]}
do
 if [[ $i == *".down.sql"* ]]; then
        psql -h "$POSTGRES_HOST" -p $POSTGRES_PORT -U "$POSTGRES_USER" -d $POSTGRES_DB -f "./schema/"$i
    fi
done
