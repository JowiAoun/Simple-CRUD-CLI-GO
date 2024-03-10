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

type Student struct {
	Student_id string
	First_name string
	Last_name  string
	Email      string
	Enrollment string
}

func DBConnect() {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@localhost:%d/?sslmode=disable",
		DB_USER, DB_PASS, DB_PORT)
	_conn, err := sql.Open("postgres", connStr)

	if err = _conn.Ping(); err != nil {
		PrintStr("Error: could not connect to database.\nPlease ensure that the constant values are valid.\n", err.Error())
		os.Exit(1)
	}

	conn = _conn

	dbSetup()
}

func dbSetup() {
	defer func() {
		_, errSetSchema := conn.Exec(StrFormat(`
		SET search_path TO %s`,
			DB_SCHEMA))
		if errSetSchema != nil {
			PrintStr("\nError: could not set schema search path:\n", errSetSchema)
		}
	}()

	_, errDb := conn.Exec(StrFormat(`CREATE DATABASE %s`, DB_NAME))
	if errDb != nil {
		connStr := fmt.Sprintf(
			"postgresql://%s:%s@db:%d/%s?sslmode=disable",
			DB_USER, DB_PASS, DB_PORT, DB_NAME)
		conn, _ = sql.Open("postgres", connStr)

		return
	}

	err := conn.Close()
	if err != nil {
		os.Exit(1)
	}

	connStr := fmt.Sprintf(
		"postgresql://%s:%s@db:%d/%s?sslmode=disable",
		DB_USER, DB_PASS, DB_PORT, DB_NAME)
	conn, _ = sql.Open("postgres", connStr)

	_, errSchema := conn.Exec(StrFormat(`
		CREATE SCHEMA %s`,
		DB_SCHEMA))
	if errSchema != nil {
		PrintStr("\nError: could not create schema:\n", errSchema)
	}

	_, errTable := conn.Exec(StrFormat(`
		CREATE TABLE %s.%s (
		student_id SERIAL PRIMARY KEY,
		first_name VARCHAR(16) NOT NULL,
		last_name VARCHAR(16) NOT NULL,
		email VARCHAR(32) UNIQUE NOT NULL,
		enrollment_date DATE)`,
		DB_SCHEMA, DB_TABLE))
	if errTable != nil {
		PrintStr("\nError: could not create table:\n", errTable)
	}

	_, errInsert := conn.Exec(StrFormat(`
		INSERT INTO %s.%s
		(first_name, last_name, email, enrollment_date) VALUES
		('John', 'Doe', 'john.doe@example.com', '2023-09-01'),
		('Jane', 'Smith', 'jane.smith@example.com', '2023-09-01'),
		('Jim', 'Beam', 'jim.beam@example.com', '2023-09-02');`,
		DB_SCHEMA, DB_TABLE))
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
	*students = (*students)[:0] // clear array
	rows, err := conn.QueryContext(context.Background(), StrFormat(`
        SELECT * FROM %s`,
		DB_TABLE))
	if err != nil {
		PrintStr("Error: could not execute query:\n", err)
		return false
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.Student_id, &student.First_name, &student.Last_name, &student.Email, &student.Enrollment); err != nil {
			PrintStr("Error: could not scan row:\n", err)
			return false
		}
		var enrollmentDate time.Time
		if err := rows.Scan(&student.Student_id, &student.First_name, &student.Last_name, &student.Email, &enrollmentDate); err != nil {
			PrintStr("Error: could not scan row:\n", err)
			return false
		}
		student.Enrollment = enrollmentDate.Format("2006-01-02") // Format date without timezone information
		*students = append(*students, student)
	}
	if err := rows.Err(); err != nil {
		PrintStr("Error: rows iteration failed:\n", err)
		return false
	}

	return true
}

func AddStudent(first_name string, last_name string, email string, enrollment_date string) bool {
	tx, err := conn.Begin()
	if err != nil {
		PrintStr("Error: could not begin transaction:", err)
		return false
	}
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				PrintStr("Error: could not rollback transaction:", rollbackErr)
			}
			return
		}
		if commitErr := tx.Commit(); commitErr != nil {
			PrintStr("Error: could not commit transaction:", commitErr)
		}
	}()

	_, err = tx.Exec(StrFormat(`
        INSERT INTO %s.%s
        (first_name, last_name, email, enrollment_date) VALUES
        ('%s', '%s', '%s', '%s');`,
		DB_SCHEMA, DB_TABLE, first_name, last_name, email, enrollment_date))
	if err != nil {
		PrintStr("Error: could not execute query:", err)
		return false
	}

	return true
}

func UpdateStudentEmail(student_id string, new_email string) bool {
	tx, err := conn.Begin()
	if err != nil {
		PrintStr("Error: could not begin transaction:", err)
		return false
	}
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				PrintStr("Error: could not rollback transaction:", rollbackErr)
			}
			return
		}
		if commitErr := tx.Commit(); commitErr != nil {
			PrintStr("Error: could not commit transaction:", commitErr)
		}
	}()

	_, err = tx.Exec(StrFormat(`
        UPDATE %s.%s
        SET email = '%s'
        WHERE student_id = %s;`,
		DB_SCHEMA, DB_TABLE, new_email, student_id))
	if err != nil {
		PrintStr("Error: could not execute query:", err)
		return false
	}

	return true
}

func DeleteStudent(student_id string) bool {
	tx, err := conn.Begin()
	if err != nil {
		PrintStr("Error: could not begin transaction:", err)
		return false
	}
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				PrintStr("Error: could not rollback transaction:", rollbackErr)
			}
			return
		}
		if commitErr := tx.Commit(); commitErr != nil {
			PrintStr("Error: could not commit transaction:", commitErr)
		}
	}()

	_, err = tx.Exec(StrFormat(`
        DELETE FROM %s.%s
        WHERE student_id = %s;`, DB_SCHEMA, DB_TABLE, student_id))
	if err != nil {
		PrintStr("Error: could not execute query:", err)
		return false
	}

	return true
}
