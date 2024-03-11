# Simple University CRUD DBMS - GO CLI

## Description

Simple server-database application allowing the user to perform CRUD operations (Create, Read, Update, Delete)
on a PostgreSQL database table. This application is written in Go, and features database initialization in
efforts to reduce overhead before first usage. The next best step would be to containerize the application with Docker.

## Features

- **Get All Students**: Display a list of all students with their attributes of ID, first name, last name, email, and enrollment date.
- **Add Student**: Insert a new student record into the database.
- **Delete Student**: Remove a student record from the database, based on the student's ID.
- **Update Student Email**: Modify the email address of an existing student record.

## Requirements

- [Go](https://go.dev/doc/install) 1.12
- [PostgreSQL](https://www.postgresql.org/download/) server

## Setup
1. Clone this repository using `git clone https://github.com/JowiAoun/Simple-University-CRUD-DBMS-GO-CLI.git`
2. Make sure that PostgreSQL is installed and running
3. In `/src/db.go`, replace the `DB_PASS` constant with the PostgreSQL server you would like to use, and the other constants if not default
4. To compile & run, execute the following in the terminal: `go run src/main.go src/db.go src/view.go`
5. Ensure that there are no errors to proceed. Existence of an error would most likely be about incorrect database connection details


## Usage

After correctly setting up the application, make use of the command line interface to perform operations on the database provided.

## Demonstration Video:
 - A demonstration video about the setup & usage can be found [here]().
