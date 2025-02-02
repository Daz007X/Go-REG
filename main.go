package main

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
  "fmt"
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
	app.Post("/students", createNewStudent)

	app.Listen(":3000")
}

func getStudents(c *fiber.Ctx) error {
	return c.JSON(students)
}

func createNewStudent(c *fiber.Ctx) error {
	student := new(Student)
	if err := c.BodyParser(student); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Generate a new UUID for the student
	student.ID = uuid.New()
	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()

	students = append(students, *student)
	return c.JSON(student)
}