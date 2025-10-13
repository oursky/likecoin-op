#!/bin/bash

set -e

migrations=$(ls migrations | sort)

mkdir -p migration-results

for migration in $migrations; do
    echo Migrating $migration
    psql -v ON_ERROR_STOP=1 $DB_CONNECTION_STR < \
        migrations/$migration > \
        migration-results/$migration
done
