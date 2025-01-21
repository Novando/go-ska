Configuration file use for migrating database.

Run the migration with following bash syntax:
```shell
$ [...delaclare variable] sql-migrate [up | down] -config [relative path to sql-migrate-config.yml] -env [migrate-mysql | migrate-pg]
```

Available variable:
`DB_HOST`
**REQUIRED**, hostname (usually IP address or domain) of the database

`DB_PORT`
**REQUIRED**, the port to access database

`DB_USER`
**REQUIRED**, username to use to access the database

`DB_PASS`
**REQUIRED**, password of the user used to access the database

`DB_NAME`
**REQUIRED**, the name of the database

`APP_NAME`
Optional, the name of your app

for example, to migrate into mysql:
```shell
$ APP_NAME=go_ska DB_HOST=localhost DB_PORT=3306 DB_USER=root DB_PASS= DB_NAME=my_table sql-migrate up -config ./sql-migrate-config.yml -env migrate-mysql
```