package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Student struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	ClassId   int    `json:"classid"`
}

type Teacher struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	ClassId   int    `json:"classid"`
}

type Class struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Language     string    `json:"language"`
	StudentsList []Student `json:"studentslist"`
	TeachersList []Teacher `json:"teacherslist"`
}

type ClassList struct {
	Classes []Class `json:"classes"`
}

func GetClasses(db *sql.DB) ClassList {
	sqlClass := "SELECT * FROM classes"
	classRows, err := db.Query(sqlClass)

	if err != nil {
		fmt.Println("err")
		panic(err)
	}

	defer classRows.Close()

	result := ClassList{}

	for classRows.Next() {
		class := Class{}
		err := classRows.Scan(&class.Id, &class.Name, &class.Language)
		if err != nil {
			fmt.Println("err1")
			panic(err)
		}

		sqlteachers := "SELECT * FROM teachers WHERE classid = $1"
		teacherRows, err := db.Query(sqlteachers, &class.Id)
		if err != nil {
			fmt.Println("err2")
			panic(err)
		}

		defer teacherRows.Close()

		for teacherRows.Next() {
			teacher := Teacher{}
			err := teacherRows.Scan(&teacher.Id, &teacher.FirstName, &teacher.LastName, &teacher.ClassId)
			if err != nil {
				fmt.Println("err3")
				panic(err)
			}
			class.TeachersList = append(class.TeachersList, teacher)
		}

		sqlStudents := "SELECT * FROM Students WHERE classid = $1"

		studentRows, err := db.Query(sqlStudents, &class.Id)

		if err != nil {
			fmt.Println("err4")
			panic(err)
		}

		defer studentRows.Close()

		for studentRows.Next() {
			student := Student{}
			err := studentRows.Scan(&student.Id, &student.FirstName, &student.LastName, &student.ClassId)

			if err != nil {
				fmt.Println("err5")
				panic(err)
			}
			class.StudentsList = append(class.StudentsList, student)
		}

		result.Classes = append(result.Classes, class)
	}

	return result
}
