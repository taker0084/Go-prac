package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main(){
	// connect to a database
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=test_connect user=sekitakeru password=")
	if err != nil {
		log.Fatalf("Unable to connect :%v\n", err)
	}
	defer conn.Close()

	log.Println("Connected to database")

	//test my connection
	err = conn.Ping()
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	log.Println("Ping database success!")
	//get rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}
	//insert a row
	query := `insert into users (first_name, last_name) values ($1, $2)`
	_, err = conn.Exec(query, "Jack", "Brown")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Insert a row success!")
	//get rows from database again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}
	//update a row
	stmt:= `update users set first_name=$1 where id =$2`
	_,err = conn.Exec(stmt, "Jackie", 4)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Update a row success!")
	//get rows from database again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}
	//get one row by id
	query = `select id, first_name, last_name from users where id=$1`

	var firstName, lastName string
	var id int

	row := conn.QueryRow(query, 1)
	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("QueryRow returns", id, firstName, lastName)
	//delete a row
	query = `delete from users where id=$1`
	_, err = conn.Exec(query, 6)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Delete a row success!")
	//get rows again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllRows(conn *sql.DB)error{
	rows, err := conn.Query("SELECT id,first_name, last_name FROM users")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	var firstName, lastName string
	var id int
	for rows.Next(){
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println("Record is", id, firstName, lastName)
	}
	if err = rows.Err(); err!=nil{
		log.Fatal("Error scanning rows", err)
	}
	fmt.Println("-----------------------------")
	return nil
}