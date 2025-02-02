package main

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
  	"fmt"
	"reflect"
)

type Student struct {
	ID           uuid.UUID `json:"id"`
	FirstName    string    `json:"first_name"`
	MiddleName   string    `json:"middle_name"`
	LastName     string    `json:"last_name"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	Sex          string    `json:"sex"`
	Address      string    `json:"address"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	Zip          string    `json:"zip"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	Status       string    `json:"status"`
	FacultyID    uuid.UUID `json:"faculty_id"`
	DepartmentID uuid.UUID `json:"department_id"`
	ProgramID    uuid.UUID `json:"program_id"`
	Level        int       `json:"level"`
}

var students []Student

func main() {
  newUUID := uuid.New()
	fmt.Println(newUUID) 
	app := fiber.New()

	app.Get("/students", getStudents)
	app.Get("/student/:id", getStudentById)
	app.Post("/student", createNewStudent)
	app.Put("/student/:id", updateStudent)
	app.Delete("/student/:id", deletedStudent)

	app.Listen(":3000")
}

func getStudents(c *fiber.Ctx) error {
	return c.JSON(students)
}

func getStudentById(c *fiber.Ctx) error {
	studentId := c.Params("id") 
	for _, student := range students {
		if student.ID.String() == studentId {
			return c.JSON(student)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func createNewStudent(c *fiber.Ctx) error {
	student := new(Student)
	if err := c.BodyParser(student); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	student.ID = uuid.New()
	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()
	students = append(students, *student)
	return c.JSON(student)
}

func updateStudent(c *fiber.Ctx) error {
	studentId := c.Params("id")
	student := new(Student)
	if err := c.BodyParser(student); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, s := range students {
		if s.ID.String() == studentId {
			val := reflect.ValueOf(student).Elem()
			currentVal := reflect.ValueOf(&students[i]).Elem()
			for j := 0; j < val.NumField(); j++ {
				field := val.Field(j)
				if !reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
					currentVal.Field(j).Set(field)
				}
			}
			students[i].UpdatedAt = time.Now()
			return c.JSON(students[i])
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}
func deletedStudent(c *fiber.Ctx) error {
	studentId := c.Params("id")
	for i, student := range students {
		if student.ID.String() == studentId {
			students = append(students[:i], students[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}