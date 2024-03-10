package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

const (
	DB_USER   = "postgres"
	DB_PASS   = "Wy5w0UY5l55G1Pf"
	DB_PORT   = 5432
	DB_SCHEMA = "q1"
	DB_NAME   = "comp3005a3"
	DB_TABLE  = "students"
)

var (
	conn *sql.DB
)

func DBConnect() {
	connStr := fmt.Sprintf("postgresql://%s:%s@localhost:%d/?sslmode=disable", DB_USER, DB_PASS, DB_PORT)
	_conn, err := sql.Open("postgres", connStr)

	if err = _conn.Ping(); err != nil {
		PrintStr("Error: could not connect to database", err)
		os.Exit(1)
	}

	conn = _conn

	dbSetup()
}

func dbSetup() {
	_, errDb := conn.Exec(StrFormat(`CREATE DATABASE %s`, DB_NAME))
	if errDb != nil {
		return
	}

	conn.Close()

	connStr := fmt.Sprintf("postgresql://%s:%s@localhost:%d/%s?sslmode=disable", DB_USER, DB_PASS, DB_PORT, DB_NAME)
	conn, _ = sql.Open("postgres", connStr)

	_, errSchema := conn.Exec(StrFormat(`CREATE SCHEMA %s`, DB_SCHEMA))
	if errSchema != nil {
		PrintStr("\nError: could not create schema:\n", errSchema)
	}

	_, errTable := conn.Exec(StrFormat(`CREATE TABLE %s.%s (
            student_id SERIAL PRIMARY KEY,
            first_name VARCHAR(32) NOT NULL,
            last_name VARCHAR(32) NOT NULL,
            email VARCHAR(64) UNIQUE NOT NULL,
            enrollment_date TIMESTAMP
        		)`, DB_SCHEMA, DB_TABLE))
	if errTable != nil {
		PrintStr("\nError: could not create table:\n", errTable)
	}

	_, errInsert := conn.Exec(StrFormat(`INSERT INTO %s.%s
            (first_name, last_name, email, enrollment_date) VALUES
						('John', 'Doe', 'john.doe@example.com', '2023-09-01'),
						('Jane', 'Smith', 'jane.smith@example.com', '2023-09-01'),
						('Jim', 'Beam', 'jim.beam@example.com', '2023-09-02');`, DB_SCHEMA, DB_TABLE))
	if errInsert != nil {
		PrintStr("\nError: could not insert values into table:\n", errInsert)
	}
}

func DBClose() {
	err := conn.Close()
	if err != nil {
		return
	}
}

func GetAllStudents(students *[]Student) bool {
	rows, err := conn.QueryContext(context.Background(), StrFormat(`SELECT * FROM %s.%s`, DB_SCHEMA, DB_TABLE))
	if err != nil {
		PrintStr("Error: could not execute query.")
		return false
	}

	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.Student_id, &student.First_name, &student.Last_name, &student.Email, &student.Enrolment); err != nil {
			return true
		}
		*students = append(*students, student)
	}

	return true
}

func AddStudent(first_name string, last_name string, email string, enrollment_date string) bool {
	return false
}

func UpdateStudentEmail(student_id int, new_email string) bool {
	return false
}

func DeleteStudent(student_id int) bool {
	return false
}
