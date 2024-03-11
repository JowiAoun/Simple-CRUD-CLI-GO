/*
db.go:

This file contains functions for interacting with a PostgreSQL database in the context
of a University Database Command-Line Interface (CLI). It includes functions for connecting
to the database, setting up the necessary schema and tables, retrieving, adding, updating,
and deleting student records.

Constants:
- DB_PASS: Password for the database user. (Replace with actual password)
- DB_USER: Database username.
- DB_PORT: Port number for the PostgreSQL database.
- DB_SCHEMA: Database schema name.
- DB_NAME: Database name.
- DB_TABLE: Table name for storing student records.

Variables:
- conn: Global variable for the database connection.

Structs:
- Student: Represents a student entity with fields such as student ID, first name, last name, email, and enrollment date.

Functions:
- DBConnect: Connects to the PostgreSQL database using the provided credentials and sets up the necessary schema and tables if they don't exist.
- dbSetup: Sets up the database schema and table. If they already exist, it connects to the existing database.
- DBClose: Closes the database connection.
- GetAllStudents: Retrieves all student records from the database and populates the provided slice with Student structs.
- AddStudent: Adds a new student record to the database.
- UpdateStudentEmail: Updates the email address of a student in the database.
- DeleteStudent: Deletes a student record from the database based on the student ID.

Dependencies:
- context: Provides support for context management.
- database/sql: Package for working with SQL databases.
- fmt: Provides formatted I/O functions.
- github.com/lib/pq: PostgreSQL driver for Go.
- os: Provides a platform-independent interface to operating system functionality.
- time: Package for handling time-related operations.
*/

package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"time"
)

const (
	DB_PASS   = "Wy5w0UY5l55G1Pf" // REPLACE HERE
	DB_USER   = "postgres"
	DB_PORT   = 5432
	DB_SCHEMA = "q1"
	DB_NAME   = "comp3005a3"
	DB_TABLE  = "students"
)

var (
	conn *sql.DB
)

// Student represents a student entity with fields such as student ID, first name, last name, email, and enrollment date.
type Student struct {
	Student_id string
	First_name string
	Last_name  string
	Email      string
	Enrollment string
}

// DBConnect establishes a connection to the PostgreSQL database using the provided credentials and sets up the necessary schema and tables.
func DBConnect() {
	// Connection string for the PostgreSQL database
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@localhost:%d/?sslmode=disable",
		DB_USER, DB_PASS, DB_PORT)
	_conn, err := sql.Open("postgres", connStr)

	// Check if connection is successful
	if err = _conn.Ping(); err != nil {
		PrintStr("Error: could not connect to database.\nPlease ensure that the constant values are valid.")
		os.Exit(1)
	}

	conn = _conn

	dbSetup()
}

// dbSetup sets up the database schema and table. If they already exist, it connects to the existing database.
func dbSetup() {
	// Defer setting search path to the specified schema
	defer func() {
		_, errSetSchema := conn.Exec(StrFormat(`
		SET search_path TO %s`,
			DB_SCHEMA))
		if errSetSchema != nil {
			PrintStrF("\nError: could not set schema search path:\n", errSetSchema)
		}
	}()

	// Create database if it does not exist
	_, errDb := conn.Exec(StrFormat(`CREATE DATABASE %s`, DB_NAME))
	if errDb != nil {
		connStr := fmt.Sprintf(
			"postgresql://%s:%s@localhost:%d/%s?sslmode=disable",
			DB_USER, DB_PASS, DB_PORT, DB_NAME)
		conn, _ = sql.Open("postgres", connStr)

		return
	}

	// Close current connection
	err := conn.Close()
	if err != nil {
		os.Exit(1)
	}

	// Reconnect to the newly created database
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@localhost:%d/%s?sslmode=disable",
		DB_USER, DB_PASS, DB_PORT, DB_NAME)
	conn, _ = sql.Open("postgres", connStr)

	// Create schema if it does not exist
	_, errSchema := conn.Exec(StrFormat(`
		CREATE SCHEMA %s`,
		DB_SCHEMA))
	if errSchema != nil {
		PrintStrF("\nError: could not create schema:\n", errSchema)
	}

	// Create table if it does not exist
	_, errTable := conn.Exec(StrFormat(`
		CREATE TABLE %s.%s (
		student_id SERIAL PRIMARY KEY,
		first_name VARCHAR(16) NOT NULL,
		last_name VARCHAR(16) NOT NULL,
		email VARCHAR(32) UNIQUE NOT NULL,
		enrollment_date DATE)`,
		DB_SCHEMA, DB_TABLE))
	if errTable != nil {
		PrintStrF("\nError: could not create table:\n", errTable)
	}

	// Insert sample data into the table
	_, errInsert := conn.Exec(StrFormat(`
		INSERT INTO %s.%s
		(first_name, last_name, email, enrollment_date) VALUES
		('John', 'Doe', 'john.doe@example.com', '2023-09-01'),
		('Jane', 'Smith', 'jane.smith@example.com', '2023-09-01'),
		('Jim', 'Beam', 'jim.beam@example.com', '2023-09-02');`,
		DB_SCHEMA, DB_TABLE))
	if errInsert != nil {
		PrintStrF("\nError: could not insert values into table:\n", errInsert)
	}
}

// DBClose closes the database connection.
func DBClose() {
	err := conn.Close()
	if err != nil {
		return
	}
}

// GetAllStudents retrieves all student records from the database and populates the provided slice with Student structs.
func GetAllStudents(students *[]Student) bool {
	*students = (*students)[:0] // Clear array
	rows, err := conn.QueryContext(context.Background(), StrFormat(`
        SELECT * FROM %s`,
		DB_TABLE))
	if err != nil {
		PrintStrF("Error: could not execute query:\n", err)
		return false
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	// Iterate over the rows and scan them into Student structs
	for rows.Next() {
		var student Student
		// Scan the row into the student struct
		if err := rows.Scan(&student.Student_id, &student.First_name, &student.Last_name, &student.Email, &student.Enrollment); err != nil {
			PrintStrF("Error: could not scan row:\n", err)
			return false
		}
		var enrollmentDate time.Time
		// Format the enrollment date without timezone information
		if err := rows.Scan(&student.Student_id, &student.First_name, &student.Last_name, &student.Email, &enrollmentDate); err != nil {
			PrintStrF("Error: could not scan row:\n", err)
			return false
		}
		student.Enrollment = enrollmentDate.Format("2006-01-02")
		// Append the scanned student to the students slice
		*students = append(*students, student)
	}
	// Check if any error occurred during row iteration
	if err := rows.Err(); err != nil {
		PrintStrF("Error: rows iteration failed:\n", err)
		return false
	}

	return true
}

// AddStudent adds a new student record to the database.
func AddStudent(first_name string, last_name string, email string, enrollment_date string) bool {
	tx, err := conn.Begin()
	if err != nil {
		PrintStrF("Error: could not begin transaction:", err)
		return false
	}
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				PrintStrF("Error: could not rollback transaction:", rollbackErr)
			}
			return
		}
		if commitErr := tx.Commit(); commitErr != nil {
			PrintStrF("Error: could not commit transaction:", commitErr)
		}
	}()

	// Execute the SQL query to insert a new student record
	_, err = tx.Exec(StrFormat(`
        INSERT INTO %s.%s
        (first_name, last_name, email, enrollment_date) VALUES
        ('%s', '%s', '%s', '%s');`,
		DB_SCHEMA, DB_TABLE, first_name, last_name, email, enrollment_date))
	if err != nil {
		PrintStrF("Error: could not execute query:", err)
		return false
	}

	return true
}

// UpdateStudentEmail updates the email of a student record in the database.
func UpdateStudentEmail(student_id string, new_email string) bool {
	tx, err := conn.Begin()
	if err != nil {
		PrintStrF("Error: could not begin transaction:", err)
		return false
	}
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				PrintStrF("Error: could not rollback transaction:", rollbackErr)
			}
			return
		}
		if commitErr := tx.Commit(); commitErr != nil {
			PrintStrF("Error: could not commit transaction:", commitErr)
		}
	}()

	// Execute the SQL query to update the email of a student record
	_, err = tx.Exec(StrFormat(`
        UPDATE %s.%s
        SET email = '%s'
        WHERE student_id = %s;`,
		DB_SCHEMA, DB_TABLE, new_email, student_id))
	if err != nil {
		PrintStrF("Error: could not execute query:", err)
		return false
	}

	return true
}

// DeleteStudent deletes a student record from the database.
func DeleteStudent(student_id string) bool {
	tx, err := conn.Begin()
	if err != nil {
		PrintStrF("Error: could not begin transaction:", err)
		return false
	}
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				PrintStrF("Error: could not rollback transaction:", rollbackErr)
			}
			return
		}
		if commitErr := tx.Commit(); commitErr != nil {
			PrintStrF("Error: could not commit transaction:", commitErr)
		}
	}()

	// Execute the SQL query to delete a student record
	_, err = tx.Exec(StrFormat(`
        DELETE FROM %s.%s
        WHERE student_id = %s;`, DB_SCHEMA, DB_TABLE, student_id))
	if err != nil {
		PrintStrF("Error: could not execute query:", err)
		return false
	}

	return true
}
