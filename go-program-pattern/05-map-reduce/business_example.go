package main

import (
	"reflect"
	"fmt"
)

type Employee struct {
	Name string
	Age int
	Vacation int
	Salary int
}

var list = []Employee{
	{"Hao", 44, 0, 8000},
	{"Bob", 34, 10, 5000},
	{"Alice", 23, 5, 9000},
	{"Jack", 26, 0, 4000},
	{"Tom", 48, 9, 7500},
	{"Marry", 29, 0, 6000},
	{"Mike", 32, 8, 4000},
}

func EmployeeCountIf(lst []Employee, fn func(e *Employee) bool) int {
	count := 0
	for i, _ := range lst {
		if fn(&lst[i]) {
			count++
		}
	}
	return count
}

func EmployeeFilterIn(list []Employee, fn func(e *Employee) bool) []Employee {
	var res []Employee
	for i, _ := range list {
		if fn(&list[i]) {
			res = append(res, list[i])
		}
	}
	return res
}

func EmployeeSumIf(lst []Employee, fn func(e *Employee)int) int {
	sum := 0
	for i, _ := range lst {
		sum += fn(&lst[i])
	}
	return sum
}

func main() {
	//getOldPeople()
	old := EmployeeCountIf(list, func (e *Employee) bool {
		return e.Age > 40
	})
	fmt.Printf("old people numbers:%v\n", old)


	//getNoVacationEmployee()
	noVacationEmployee := EmployeeFilterIn(list, func (e *Employee) bool {
		return e.Vacation == 0
	})
	fmt.Printf("no vacation employee is:%+v\n", noVacationEmployee)


	totalPay := EmployeeSumIf(list, func(e *Employee) int {
		return e.Salary
	})
	fmt.Printf("Total Salary: %d\n", totalPay)//Total Salary: 43500

	type A int
	a := A(2)
	fmt.Printf("type:%+v, value:%+v", reflect.TypeOf(a), a)
}
