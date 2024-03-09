package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DB_USER  = "abcd"
	DB_PASS  = "abcd"
	DB_NAME  = "abcd"
	DB_PORT  = 5432
	DB_TABLE = "abcd"
)

var (
	conn *sql.DB
)

func DBConnect() {
	connStr := fmt.Sprintf("postgresql://%s:%s@localhost:%d/%s?sslmode=disable", DB_USER, DB_PASS, DB_PORT, DB_NAME)
	conn, _ = sql.Open("postgres", connStr)
}

func dbSetup() {
	//TODO: First, check if db exists. If not, run this function
}

func DBClose() {
	err := conn.Close()
	if err != nil {
		return
	}
}

func GetAllStudents(students *[]Student) bool {
	rows, err := conn.QueryContext(context.Background(), fmt.Sprintf("SELECT * FROM %s", DB_TABLE))
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
