package main

import (
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// connect to cluster on localhost and create demo keyspace for practice
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "demo"
	session, _ := cluster.CreateSession()
	// closing session after main finishes
	defer session.Close()

	if err := session.Query("INSERT INTO user (lastname, age, city, email, firstname) VALUES ('Jackson', 34, 'Huntsville', 'foo@bar.baz', 'Michael')").Exec(); err != nil {
		log.Fatal(err)
	}

}

// tested this on windows
