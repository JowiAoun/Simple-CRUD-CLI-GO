/*
main.go:

This file contains the main entry point for the University Database Command-Line Interface (CLI)
application. It orchestrates the interaction between the user interface functions, database
functions, and the main execution loop.

Variables:
- choice: Stores the user's menu choice.
- student_id: Stores the ID of a student.
- first_name: Stores the first name of a student.
- last_name: Stores the last name of a student.
- email: Stores the email address of a student.
- enrollment_date: Stores the enrollment date of a student.
- new_email: Stores the new email address for updating a student's email.
- students: Slice to hold student records.

Functions:
- main: The main entry-point function of the CLI application. It connects to the database,
displays the menu, and handles user input to perform various operations such as
retrieving, adding, updating, or deleting student records. It also closes the
database connection when the user chooses to exit.

Dependencies:
- DBConnect: Function to connect to the PostgreSQL database.
- PrintMenu: Function to display the menu and prompt the user for input.
- GetAllStudents: Function to retrieve all student records from the database.
- PrintStudents: Function to display student records.
- PrintOption2: Function to prompt the user for input to add a new student.
- AddStudent: Function to add a new student record to the database.
- PrintOption3: Function to prompt the user for input to update a student's email.
- UpdateStudentEmail: Function to update a student's email in the database.
- PrintOption4: Function to prompt the user for input to delete a student.
- DeleteStudent: Function to delete a student record from the database.
*/

package main

var (
	choice          = -1
	student_id      string
	first_name      string
	last_name       string
	email           string
	enrollment_date string
	new_email       string
	students        []Student
)

// main function is the entry point of the program where it connects to the database and presents a menu to the user.
func main() {
	DBConnect()

	for choice != 0 {

		PrintMenu(&choice)

		switch choice {
		case 0:
			DBClose()
			return
		case 1:
			GetAllStudents(&students)
			PrintStudents(&students)
		case 2:
			PrintOption2(&first_name, &last_name, &email, &enrollment_date)
			if AddStudent(first_name, last_name, email, enrollment_date) {
				PrintStr("Student has been added.")
			}
		case 3:
			PrintOption3(&student_id, &new_email)
			if UpdateStudentEmail(student_id, new_email) {
				PrintStr("Student email has been updated.")
			}
		case 4:
			PrintOption4(&student_id)
			if DeleteStudent(student_id) {
				PrintStr("Student has been deleted.")
			}
		}
	}
}
