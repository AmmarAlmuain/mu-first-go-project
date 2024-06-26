package main

import (
	"fmt"
	"reflect"
	"slices"
)

var employees []*Employee = []*Employee{}

type Employee struct {
	id                   int
	OptionalEmployeeInfo map[string]any
}

type OptionalEmployeeInfo struct {
	firstName    string
	middleName   string
	lastName     string
	phone        int
	email        string
	salary       int
	taxes        int
	dateOfBirth  string
	role         string
	period       string
	departmentId int
	isManager    bool
	isCEO        bool
}

func GetEmployee(id int) *Employee {
	for _, v := range employees {
		if v.id == id {
			return v
		}
	}
	return nil
}

func NewEmployee(id int) *Employee {
	employees = append(employees, &Employee{id: id})
	return employees[len(employees)-1]
}

func (e *Employee) SetOptionalInfo(info OptionalEmployeeInfo) {
	typeOfInfo := reflect.TypeOf(info)
	e.OptionalEmployeeInfo = make(map[string]any)
	for i := 0; i < typeOfInfo.NumField(); i++ {
		fieldName := typeOfInfo.Field(i).Name
		fieldValue := reflect.ValueOf(info).Field(i)
		if !fieldValue.IsZero() {
			e.OptionalEmployeeInfo[fieldName] = fieldValue
		}
	}
}

func DeleteEmployee(id int) []*Employee {
	var newEmployeeList []*Employee = employees
	employees = slices.Delete(newEmployeeList, id, id+1)
	return employees
}

type Bill struct {
	id         int
	employeeId int
	amount     int
	isPaid     bool
	date       string
}

type Report struct {
	id          int
	employeeId  int
	description string
	date        string
}

type Task struct {
	id          int
	employeeId  int
	managerId   int
	description string
	state       string
	notes       []string
	startDate   string
	endDate     string
}

type Department struct {
	id   int
	name string
}

func main() {

	for {
		fmt.Println("hi im in loop still!")
		break
	}

	var employee *Employee = NewEmployee(1)
	var employee2 *Employee = NewEmployee(2)
	var employee3 *Employee = NewEmployee(3)
	info := OptionalEmployeeInfo{
		firstName: "John",
		lastName:  "Doe",
		email:     "john.doe@company.com",
	}

	fmt.Println(employee, employee2, employee3)

	fmt.Println(GetEmployee(2))

	employee.SetOptionalInfo(info)

	fmt.Println(employees)

	DeleteEmployee(2)

	fmt.Println(employees)

	// var firstName, lastName string = "Ammar", "Almuain"
	// employee.SetOptionalInfo(&OptionalEmployeeInfo{firstName: &firstName})
	// fmt.Print(employee.OptionalEmployeeInfo)
	// employee.SetOptionalInfo(&OptionalEmployeeInfo{lastName: &lastName})
	// fmt.Print(employee.OptionalEmployeeInfo)
}
