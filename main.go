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

type Student struct {
	Student_id string
	First_name string
	Last_name  string
	Email      string
	Enrollment string
}

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
