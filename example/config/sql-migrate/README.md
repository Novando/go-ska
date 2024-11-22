Configuration file use for migrating database.

Run the migration with following bash syntax:
```shell
$ [...delaclare variable] sql-migrate [up | down] -config [relative path to sql-migrate-config.yml] -env [migrate-mysql | migrate-pg]
```

for example, to migrate into mysql:
```shell
$ DB_HOST=localhost DB_PORT=3306 DB_USER=root DB_PASS= DB_NAME=my_table sql-migrate up -config ./sql-migrate-config.yml -env migrate-mysql
```