package main

import "fmt"

type Speaker struct{}

func (s Speaker) MakeSound() {
	fmt.Println("Boom-Boom! 🔊")
}

type Bell string

func (b Bell) MakeSound() {
	fmt.Println("Ting-Ting! 🔔")
}

type SoundMaker interface {
	MakeSound()
}

type EmployeeContract struct {
	EmpID         int
	Salary        float64
	TaxPercentage float64
}

type FreelancerContract struct {
	EmpID       int
	HourlyRate  float64
	HoursWorked int
}

// 'EmployeeContract' Salary is the base salary - calculated taxes
func (e EmployeeContract) CalculateSalary() float64 {
	return e.Salary - (e.TaxPercentage * e.Salary)
}

// 'FreelancerContract' salary is the hourly rate * hours worked -- they declare taxes themselves
func (f FreelancerContract) CalculateSalary() float64 {
	return f.HourlyRate * float64(f.HoursWorked)
}

type SalaryCalculator interface {
	CalculateSalary() float64
}

func main() {
	var s SoundMaker

	s = Bell("Acme Bell")
	s.MakeSound() // Ting-Ting! 🔔

	s = Speaker{}
	s.MakeSound() // Boom-Boom! 🔊
	homer := EmployeeContract{EmpID: 1, Salary: 479.60, TaxPercentage: 0.245}
	fmt.Printf("Homer's salary: $%.2f\n", homer.CalculateSalary()) // Homer's salary: $362.10

	deadpool := FreelancerContract{EmpID: 2, HourlyRate: 50_000.00, HoursWorked: 10}
	fmt.Printf("Deadpool's salary: $%.2f\n", deadpool.CalculateSalary()) // Deadpool's salary: $500000.00

	employees := []SalaryCalculator{homer, deadpool}
	fmt.Printf("Monthly sal. expense: $%.2f\n", salaryExpense(employees)) // Monthly sal. expense: "$500362.10

}

// salaryExpense() takes a slice of []SalaryCalculator and returns the sum of all salaries
func salaryExpense(salaries []SalaryCalculator) float64 {
	totalExpense := 0.0
	for _, v := range salaries {
		totalExpense += v.CalculateSalary()
	}
	return totalExpense
}
