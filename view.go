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

func PrintOption2(first_name *string, last_name *string, email *string, enrollment_date *string) {
	fmt.Printf("Enter student's first name: ")
	*first_name, _ = r.ReadString('\n')

	fmt.Printf("Enter last name: ")
	*last_name, _ = r.ReadString('\n')

	fmt.Printf("Enter email: ")
	*email, _ = r.ReadString('\n')

	fmt.Printf("Enter enrolment date (i.e.: 2024-09-03): ")
	*enrollment_date, _ = r.ReadString('\n')
}

func PrintOption3(student_id *string, new_email *string) {
	fmt.Printf("Enter student's ID: ")
	*student_id, _ = r.ReadString('\n')
	fmt.Printf("Enter new email: ")
	*new_email, _ = r.ReadString('\n')
}

func PrintOption4(student_id *string) {
	fmt.Printf("Enter student's ID: ")
	*student_id, _ = r.ReadString('\n')
}

func PrintStudents(students *[]Student) {
	fmt.Println("Students in the university")

	fmt.Printf("%-9s %-16s %-16s %-32s %-10s\n", "ID", "FIRST NAME", "LAST NAME", "EMAIL", "ENROLLMENT DATE")
	fmt.Println(strings.Repeat("-", 92))

	for _, student := range *students {
		fmt.Printf(
			"%-9s %-16s %-16s %-32s %-10s\n",
			student.Student_id, student.First_name, student.Last_name, student.Email, student.Enrollment,
		)
	}
}

func PrintStr(str string, args ...interface{}) {
	fmt.Println(StrFormat(str, args))
}

func StrFormat(str string, args ...interface{}) string {
	return fmt.Sprintf(str, args...)
}
