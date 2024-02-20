# Migrate package in Go

A database migration, also known as a schema migration, is a set of changes to be made to a structure of objects within a relational database.

It is a way to manage and implement incremental changes to the structure of data in a controlled, programmatic manner. These changes are often reversible, meaning they can be undone or rolled back if required.

The process of migration helps to change the database schema from its current state to a new desired state, whether it involves adding tables and columns, removing elements, splitting fields, or changing types and constraints.

By managing these changes in a programmatic way, it becomes easier to maintain consistency and accuracy in the database, as well as keep track of the history of modifications made to it.


# Migration format

Migrations can be written in different formats: as a ```.go``` file or a ```.json``` file. But the most popular and universal format is ```.sql```, which is raw SQL. Such migrations are easily readable and easier to work with.


# How to setup Migration

to use migration package follow following steps

## Root module intialization

first make sure you have intialized a root package file using ```go mod init``` command like this ```go mod init github.com/shubhammishra-1```

## Install Go Migrate Package

Install go ```migrate``` package using this command

```go
go get github.com/golang-migrate/migrate/v4
```

## To create sql files for migration

to create a specific ```.sql``` file replace ```name_of_sql_file``` placeholder with your database operation name use below command and excute this command. 


# To Create Psql Migration

``` go
migrate create -ext sql -dir postgres/migration name_of_sql_file

```

Consider ```postgres://postgres:password@127.0.0.1:5432/database_name?sslmode=disable``` database URL

it will create two migration files one will be ending with ```up.sql``` and another one will be ```down.sql``` in ```postgres/migration``` directory

Now you can put your sql logic which you want to execute on database in ```up.sql``` and ```down.sql``` files

generallly creation, updation , reading operation related logic put into ```up.sql``` file. and deletion,rollback etc.. into ```down.sql```


# To Create MySQL Migration

```bash
migrate create -ext sql -dir mysql -seq name_of_sql_file

```

it will create create a directory that contains `up.sql` and `down.sql` files in MySQL directory just like Psql above.


# Execution of migration files

There are two ways to excute written migration files

## Way 1 using Command line for Psql


to execute all ```up.sql``` ending files

```bash
migrate -source file://postgres/migration -database postgres://postgres:password@127.0.0.1:5432/database_name?sslmode=disable up 
```

to execute all ```down.sql``` ending files

```bash
migrate -source file://postgres/migration -database postgres://postgres:password@127.0.0.1:5432/database_name?sslmode=disable down 
```


## For Mysql CLI

will update...



# if any error occurs

```Note``` if you migration failed makesure visit this docmentation 

```https://stackoverflow.com/questions/59616263/dirty-database-version-error-when-using-golang-migrate```

Or just apply this command on your database

```bash 
select * from schema_migrations;
update schema_migrations set dirty =false
```


## Way 2 using go program for Psql

first install migration library ```go get github.com/golang-migrate/migrate/v4```

make sure you have imported these 3 libraries at least in your main file

```"github.com/golang-migrate/migrate/v4"```

```_ "github.com/golang-migrate/migrate/v4/source/file"```

```_ "github.com/golang-migrate/migrate/v4/database/postgres"```


Now create a migration instanace "m" which will be used perform Up(), Down() etc.. operations
New returns a new Migrate instance from a source URL and a database URL. The URL scheme is defined by each driver. 

it require path of ```.sql``` migration files and databae URL

```go 
_, b, _, _ := runtime.Caller(0)
migrationPath := fmt.Sprintf("file://%s/postgres/migration", path.Dir(b))
```

Creating migration instance

```go 
m, err := migrate.New(migrationPath, databaseURL)
```


To Perform migration ```Up()``` operation.. just use your created migration instance
Up() looks at the currently active migration version and will migrate all the way up (applying all up migrations). 

```go 
err = m.Up()
```

To perfoem migration ```Down()``` operation..
Down looks at the currently active migration version and will migrate all the way down (applying all down migrations).

```go 
err = m.Down()
```

# For MySQL

see `mysql.go file` almost everything is same



# Reference

```https://pkg.go.dev/github.com/golang-migrate/migrate```