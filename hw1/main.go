package main

import (
	"fmt"
	"hw1/employee"
	"hw1/employee/director"
)

func main() {
	var directorE employee.Employee = &director.Director{}

	fmt.Println(directorE.Salary())

	directorE.SetSalary(1000)

	fmt.Println(directorE)

	fmt.Println(directorE.Salary())
}
