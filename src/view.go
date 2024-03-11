/*
view.go:

This file contains the user interface functions for interacting with a University Database
Command-Line Interface (CLI). It includes functions for displaying a menu, getting user
input for various operations such as adding, updating, or deleting students, and printing
out student information.

Functions:
- PrintMenu: Displays a menu with options for interacting with the university database CLI and prompts the user to make a selection.
- PrintOption2: Prompts the user to input details of a new student (first name, last name, email, enrollment date).
- PrintOption3: Prompts the user to input a student ID and a new email for updating a student's email.
- PrintOption4: Prompts the user to input a student ID for deleting a student.
- PrintStudents: Displays a list of students with their IDs, first names, last names, emails, and enrollment dates.
- PrintStrF: Formats and prints a string using Printf-style formatting.
- PrintStr: Prints a string.
- StrFormat: Formats a string using Sprintf.

Dependencies:
- bufio: Buffered I/O for reading user input.
- fmt: Formatting and printing functions.
- os: Operating system functions for interacting with the standard input/output.
- strings: String manipulation functions.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	r = bufio.NewReader(os.Stdin)
)

// PrintMenu prints the main menu of the University Database Command-Line Interface (CLI),
// prompts the user for a selection, and validates the input.
func PrintMenu(choice *int) {
	fmt.Println("\nUniversity DB CLI")
	fmt.Println("-----------------")
	fmt.Println("(1) Get all students")
	fmt.Println("(2) Add student")
	fmt.Println("(3) Update student email")
	fmt.Println("(4) Delete student")
	fmt.Println("(0) Exit")

	fmt.Printf("Enter your selection: ")
	_, err := fmt.Scanln(choice)
	if err != nil {
		fmt.Println("\nError: please enter a valid integer for a choice.")
		*choice = -1
		PrintMenu(choice)
	} else if *choice < 0 || *choice > 4 {
		fmt.Println("\nError: please enter one of the valid options [0, 4]")
		*choice = -1
		PrintMenu(choice)
	}
	fmt.Println()
}

// PrintOption2 prompts the user to input details of a new student such as first name, last name, email, and enrollment date.
func PrintOption2(first_name *string, last_name *string, email *string, enrollment_date *string) {
	fmt.Printf("Enter student's first name: ")
	*first_name, _ = r.ReadString('\n')
	*first_name = strings.TrimSpace(*first_name)

	fmt.Printf("Enter last name: ")
	*last_name, _ = r.ReadString('\n')
	*last_name = strings.TrimSpace(*last_name)

	fmt.Printf("Enter email: ")
	*email, _ = r.ReadString('\n')
	*email = strings.TrimSpace(*email)

	fmt.Printf("Enter enrolment date (i.e.: 2024-09-03): ")
	*enrollment_date, _ = r.ReadString('\n')
	*enrollment_date = strings.TrimSpace(*enrollment_date)
}

// PrintOption3 prompts the user to input a student ID and a new email for updating a student's email.
func PrintOption3(student_id *string, new_email *string) {
	fmt.Printf("Enter student's ID: ")
	*student_id, _ = r.ReadString('\n')
	*student_id = strings.TrimSpace(*student_id)

	fmt.Printf("Enter new email: ")
	*new_email, _ = r.ReadString('\n')
	*new_email = strings.TrimSpace(*new_email)
}

// PrintOption4 prompts the user to input a student ID for deleting a student.
func PrintOption4(student_id *string) {
	fmt.Printf("Enter student's ID: ")
	*student_id, _ = r.ReadString('\n')
	*student_id = strings.TrimSpace(*student_id)
}

// PrintStudents prints the list of students with their IDs, first names, last names, emails, and enrollment dates.
func PrintStudents(students *[]Student) {
	fmt.Println("Students in the university")

	fmt.Printf("%-9s %-16s %-16s %-32s %-10s\n", "ID", "FIRST NAME", "LAST NAME", "EMAIL", "ENROLLMENT")
	fmt.Println(strings.Repeat("-", 87))

	for _, student := range *students {
		fmt.Printf(
			"%-9s %-16s %-16s %-32s %-10s\n",
			student.Student_id, student.First_name, student.Last_name, student.Email, student.Enrollment,
		)
	}
}

// PrintStrF prints a formatted string with the provided format and arguments.
func PrintStrF(str string, args ...interface{}) {
	fmt.Println(StrFormat(str, args[0]))
}

// PrintStr prints a string.
func PrintStr(str string) {
	fmt.Println(str)
}

// StrFormat formats a string using Sprintf.
func StrFormat(str string, args ...interface{}) string {
	return fmt.Sprintf(str, args...)
}
