package main

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil
}

func init() {
	students = append(students, &Student{Id: "S001", Name: "Jung", Grade: 2})
	students = append(students, &Student{Id: "S002", Name: "Setia", Grade: 2})
	students = append(students, &Student{Id: "S003", Name: "Agung", Grade: 3})
}
