package main

var (
	choice          int = -1
	student_id      int
	first_name      string
	last_name       string
	email           string
	enrollment_date string
	new_email       string
	students        []Student
)

type Student struct {
	Student_id int
	First_name string
	Last_name  string
	Email      string
	Enrolment  string
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
			PrintOption1()
			GetAllStudents(&students)
			PrintStudents(&students)
		case 2:
			PrintOption2(&first_name, &last_name, &email, &enrollment_date)
			AddStudent(first_name, last_name, email, enrollment_date)
		case 3:
			PrintOption3(&student_id, &new_email)
			UpdateStudentEmail(student_id, new_email)
		case 4:
			PrintOption4(&student_id)
			DeleteStudent(student_id)
		}
	}
}
