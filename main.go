package main

import (
	"fmt"
	"log"
	"path"
	"runtime"

	//make sure you add this

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/golang-migrate/migrate/v4"
)

func main() {

	//databse URL

	databaseURL := "postgres://postgres:55120@127.0.0.1:5432/google?sslmode=disable"

	//migration path

	_, b, _, _ := runtime.Caller(0)

	log.Println(b)

	migrationPath := fmt.Sprintf("file://%s/postgres/migration", path.Dir(b))

	log.Println(migrationPath)

	//creating a migration instanace "m"

	m, err := migrate.New(migrationPath, databaseURL)
	if err != nil {
		log.Println("error in creating a migration instanace")
		log.Fatal(err)
	}

	//Performing migration Up() operation ...
	// err = m.Up()

	// if err != nil {
	// 	log.Println("error in performing migration Up")
	// 	log.Fatal(err)
	// }

	// log.Println("Migration Up done!")

	//Performing migration Down() operation ...

	err = m.Down()

	if err != nil {
		log.Println("error in performing migration Down")
		log.Fatal(err)
	}

	log.Println("Migration Down done!")

}
