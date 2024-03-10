package main

import (
	"fmt"
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
}

func PrintOption1() {

}

func PrintOption2(first_name *string, last_name *string, email *string, enrollment_date *string) {
	fmt.Println("Enter student's first name: ")
	_, _ = fmt.Scanln(&first_name)
	fmt.Println("Enter last name: ")
	_, _ = fmt.Scanln(&last_name)
	fmt.Println("Enter email: ")
	_, _ = fmt.Scanln(&email)
	fmt.Println("Enter first name: ")
	_, _ = fmt.Scanln(&enrollment_date)
}

func PrintOption3(student_id *int, new_email *string) {
	fmt.Println("Enter student's ID: ")
	_, _ = fmt.Scanln(&student_id)
	fmt.Println("Enter new email: ")
	_, _ = fmt.Scanln(&new_email)
}

func PrintOption4(student_id *int) {
	fmt.Println("Enter student's ID: ")
	_, _ = fmt.Scanln(&student_id)
}

func PrintStudents(students *[]Student) {
	fmt.Println("All students")
	fmt.Println("------------")

	for _, student := range *students {
		fmt.Printf("%9d %16s %16s %32s %10s\n", student.Student_id, student.First_name, student.Last_name, student.Email, student.Enrolment)
	}
}

func PrintStr(str string, args ...interface{}) {
	fmt.Printf(str, args...)
}

func StrFormat(str string, args ...interface{}) string {
	return fmt.Sprintf(str, args...)
}
