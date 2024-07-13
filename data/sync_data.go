package data

import (
	"fmt"
	"math/rand"
	"sync"
	"weekly-2/models"
)

var DataEmployee []any
var commonNames = []string{
	"John Doe", "Jane Smith", "Michael Johnson", "Emily Brown", "David Wilson",
	"Sarah Martinez", "James Taylor", "Linda Anderson", "Robert Thomas", "Mary Garcia",
}

func randomFullName() string {
	return commonNames[rand.Intn(len(commonNames))]
}

func randRange(min int, max int) int {
	var summation int
	for {
		summation = rand.Intn(max) + min
		if summation >= max {
			continue
		} else {
			break
		}
	}
	return summation
}

func randStatus() models.StatusInformation {
	statuses := []models.StatusInformation{
		models.PERMANENT,
		models.CONTRACT,
		models.TRAINEE,
	}

	size := len(statuses)
	if size == 0 {
		return models.StatusInformation("NONE")
	}

	index := rand.Intn(size)
	return statuses[index]
}

func All_Total_Summary(feature_flag bool, data []any) any {
	if feature_flag {
		var WorkGroups sync.WaitGroup
		var mutex sync.Mutex
		// 	var AssignValue uint64
		All_Channels := make(chan uint64)
		for _, v := range DataEmployee {
			switch PerPerson := v.(type) {
			case models.PermanentEmployee:
				WorkGroups.Add(1)
				go func() {
					defer WorkGroups.Done()
					mutex.Lock()
					PerPerson.Total_Salary = PerPerson.Salary + PerPerson.Insurance
					All_Channels <- PerPerson.Total_Salary
					mutex.Unlock()
				}()
			case models.ContractEmployee:
				WorkGroups.Add(1)
				go func() {
					defer WorkGroups.Done()
					mutex.Lock()
					PerPerson.Total_Salary = PerPerson.Salary + PerPerson.Overtime
					// All_Total_Salary = All_Total_Salary + PerPerson.Total_Salary
					// AssignValue += PerPerson.Total_Salary
					All_Channels <- PerPerson.Total_Salary
					mutex.Unlock()
				}()
			case models.TraineeEmployee:
				WorkGroups.Add(1)
				go func() {
					defer WorkGroups.Done()
					mutex.Lock()
					PerPerson.Total_Salary = PerPerson.Salary + PerPerson.Allowance
					All_Channels <- PerPerson.Total_Salary
					// All_Total_Salary = All_Total_Salary + PerPerson.Total_Salary
					mutex.Unlock()
				}()
			}
		}

		go func() {
			WorkGroups.Wait()
			close(All_Channels)
		}()

		TotalValues := uint64(0)
		for value := range All_Channels {
			TotalValues += value
		}

		return TotalValues
	} else {
		var All_Total_Salary uint64
		for _, v := range data {
			switch PerPerson := v.(type) {
			case models.PermanentEmployee:
				PerPerson.Total_Salary = PerPerson.Salary + PerPerson.Insurance
				All_Total_Salary = All_Total_Salary + PerPerson.Total_Salary
			case models.ContractEmployee:
				PerPerson.Total_Salary = PerPerson.Salary + PerPerson.Overtime
				All_Total_Salary = All_Total_Salary + PerPerson.Total_Salary
			case models.TraineeEmployee:
				PerPerson.Total_Salary = PerPerson.Salary + PerPerson.Allowance
				All_Total_Salary = All_Total_Salary + PerPerson.Total_Salary
			}
		}

		return All_Total_Salary
	}
}

func Work_NonChannel() []any {
	var FinalData []any
	var WorkGroups sync.WaitGroup
	var Mutex sync.Mutex
	for i := 0; i < 2000; i++ {
		WorkGroups.Add(1)
		go func(index int) {
			defer WorkGroups.Done()
			Employees := models.Employee{
				EmpID:        uint64(index + 1),
				Fullname:     randomFullName(),
				Salary:       uint64(randRange(5000, 15000)),
				Total_Salary: 0,
				Status:       randStatus(),
			}
			switch Employees.Status {
			case models.PERMANENT:
				temp_Permanent := models.PermanentEmployee{
					Employee: Employees,
				}
				temp_insurance, err := temp_Permanent.InsuranceEmployee(string(Employees.Status))
				if err != nil {
					fmt.Printf("Error assigning insurance for PermanentEmployee %d: %v\n", Employees.EmpID, err)
					break
				}
				temp_Permanent.Insurance = temp_insurance
				// Employees.Total_Salary = temp_Permanent.Insurance + temp_Permanent.Salary
				Mutex.Lock()
				DataEmployee = append(DataEmployee, temp_Permanent)
				Mutex.Unlock()
			case models.CONTRACT:
				temp_Contract := models.ContractEmployee{
					Employee: Employees,
				}
				temp_overtime, err := temp_Contract.OvertimeEmployee(string(Employees.Status))
				if err != nil {
					fmt.Printf("Error assigning insurance for PermanentEmployee %d: %v\n", Employees.EmpID, err)
					break
				}
				temp_Contract.Overtime = temp_overtime
				// Employees.Total_Salary = temp_Contract.Overtime + temp_Contract.Salary
				Mutex.Lock()
				DataEmployee = append(DataEmployee, temp_Contract)
				Mutex.Unlock()
			case models.TRAINEE:
				temp_Trainee := models.TraineeEmployee{
					Employee: Employees,
				}
				temp_allowance, err := temp_Trainee.AllowanceEmployee(string(Employees.Status))
				if err != nil {
					fmt.Printf("Error assigning insurance for PermanentEmployee %d: %v\n", Employees.EmpID, err)
					break
				}
				temp_Trainee.Allowance = temp_allowance
				// Employees.Total_Salary = temp_Trainee.Allowance + temp_Trainee.Salary
				Mutex.Lock()
				DataEmployee = append(DataEmployee, temp_Trainee)
				Mutex.Unlock()
			default:
				break
			}
		}(i)
	}
	WorkGroups.Wait()
	FinalData = append(FinalData, DataEmployee...)
	return FinalData
}
