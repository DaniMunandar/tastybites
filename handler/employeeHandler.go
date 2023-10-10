package handler

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"release5/entity"
)

func GetEmployeeInput() (entity.Employee, error) {
	scanner := bufio.NewScanner(os.Stdin)
	employee := entity.Employee{}

	fmt.Print("Enter first name: ")
	scanner.Scan()
	if len(scanner.Text()) < 1 {
		return entity.Employee{}, errors.New("empty input")
	}
	employee.First_name = scanner.Text()

	fmt.Print("Enter last name: ")
	scanner.Scan()
	if len(scanner.Text()) < 1 {
		return entity.Employee{}, errors.New("empty input")
	}
	employee.Last_name = scanner.Text()

	fmt.Print("Enter position: ")
	scanner.Scan()
	if len(scanner.Text()) < 1 {
		return entity.Employee{}, errors.New("empty input")
	}
	employee.Position = scanner.Text()

	return employee, nil
}

func AddEmployee(db *sql.DB, employee entity.Employee) error {
	_, err := db.Exec("INSERT INTO Employees (first_name, last_name, position) VALUES (?,?,?)", employee.First_name, employee.Last_name, employee.Position)
	return err
}
