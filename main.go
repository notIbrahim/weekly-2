package main

import (
	"fmt"
	"weekly-2/data"
)

func main() {
	fmt.Println("Week 2 Test")
	var Non_Channels, Channels []any
	Non_Channels = data.Work_NonChannel()
	Channels = Non_Channels
	Total_Salary_Non_Channel := data.All_Total_Summary(false, Non_Channels)
	fmt.Print("Using Non Channel Total Salary = ", Total_Salary_Non_Channel)
	Total_Salary_All_Channel := data.All_Total_Summary(true, Channels)
	fmt.Print("Using Channel Total Salary = ", Total_Salary_All_Channel)
	// go func() {
	// 	Total_Salary_All := All_Total_Summary()
	// 	fmt.Print(Total_Salary_All)
	// }()
	// for i := 0; i < 2000; i++ {
	// 	Employees := models.Employee{
	// 		EmpID:        uint64(i + 1),
	// 		Fullname:     randomFullName(),
	// 		Salary:       uint64(randRange(5000, 15000)),
	// 		Total_Salary: 0,
	// 		Status:       randStatus(),
	// 	}

	// 	switch Employees.Status {
	// 	case models.PERMANENT:
	// 		temp_Permanent := models.PermanentEmployee{
	// 			Employee: Employees,
	// 		}
	// 		temp_insurance, err := temp_Permanent.InsuranceEmployee(string(Employees.Status))
	// 		if err != nil {
	// 			fmt.Printf("Error assigning insurance for PermanentEmployee %d: %v\n", Employees.EmpID, err)
	// 			continue
	// 		}
	// 		temp_Permanent.Insurance = temp_insurance
	// 		Employees.Total_Salary = temp_Permanent.Insurance + temp_Permanent.Salary
	// 		DataEmployee = append(DataEmployee, Employees)
	// 	case models.CONTRACT:
	// 		temp_Contract := models.ContractEmployee{
	// 			Employee: Employees,
	// 		}
	// 		temp_overtime, err := temp_Contract.OvertimeEmployee(string(Employees.Status))
	// 		if err != nil {
	// 			fmt.Printf("Error assigning insurance for PermanentEmployee %d: %v\n", Employees.EmpID, err)
	// 			continue
	// 		}
	// 		temp_Contract.Overtime = temp_overtime
	// 		Employees.Total_Salary = temp_Contract.Overtime + temp_Contract.Salary
	// 		DataEmployee = append(DataEmployee, Employees)
	// 	case models.TRAINEE:
	// 		temp_Trainee := models.TraineeEmployee{
	// 			Employee: Employees,
	// 		}
	// 		temp_allowance, err := temp_Trainee.AllowanceEmployee(string(Employees.Status))
	// 		if err != nil {
	// 			fmt.Printf("Error assigning insurance for PermanentEmployee %d: %v\n", Employees.EmpID, err)
	// 			continue
	// 		}
	// 		temp_Trainee.Allowance = temp_allowance
	// 		Employees.Total_Salary = temp_Trainee.Allowance + temp_Trainee.Salary
	// 		DataEmployee = append(DataEmployee, Employees)
	// 	default:
	// 		break
	// 	}
	// }

	// fmt.Print(DataEmployee...)

	// var AllSalary uint64
	// for _, v := range DataEmployee {
	// 	switch PerPerson := v.(type) {
	// 	case *models.PermanentEmployee:
	// 		{
	// 			PerPerson.Total_Salary = PerPerson.Salary + PerPerson.Insurance
	// 		}
	// 	case *models.ContractEmployee:
	// 		{
	// 			PerPerson.Total_Salary = PerPerson.Salary + PerPerson.Overtime
	// 		}
	// 	case *models.TraineeEmployee:
	// 		{
	// 			PerPerson.Total_Salary = PerPerson.Salary + PerPerson.Allowance
	//
	// 		}
	// 	}
	// }

}
