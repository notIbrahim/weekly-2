package models

import "errors"

const (
	PERMANENT StatusInformation = "PERMANENT"
	CONTRACT  StatusInformation = "CONTRACT"
	TRAINEE   StatusInformation = "TRAINEE"
)

type Employee struct {
	EmpID        uint64
	Fullname     string
	Salary       uint64
	Total_Salary uint64
	Status       StatusInformation
}

type PermanentEmployee struct {
	Employee
	Insurance uint64
}

type ContractEmployee struct {
	Employee
	Overtime uint64
}

type TraineeEmployee struct {
	Employee
	Allowance uint64
}

type StatusInformation string

type CheckingStatus func(t_Status string) StatusInformation

func Money_on_Status(ListStatus StatusInformation) uint64 {
	switch ListStatus {
	case PERMANENT:
		return uint64(500000)
	case CONTRACT:
		return uint64(55000)
	case TRAINEE:
		return uint64(100000)
	default:
		return uint64(0)
	}
}

func ProccededStatus(StatusCheck string) StatusInformation {
	switch StatusCheck {
	case string(PERMANENT):
		return PERMANENT
	case string(CONTRACT):
		return CONTRACT
	case string(TRAINEE):
		return TRAINEE
	default:
		return StatusInformation("NONE")
	}
}

func CheckStatus(TempStatus string, StatusFunction CheckingStatus) StatusInformation {
	return StatusFunction(TempStatus)
}

type Tunjangan interface {
	InsuranceEmployee()
	OvertimeEmployee()
	AllowanceEmployee()
}

// func (NewData *Employee) CreateEmployee(id uint64, fullname string, salary uint64, statusEmployee string) *Employee {
// 	statusConvert := CheckStatus(statusEmployee, ProccededStatus)
// 	return &Employee{
// 		EmpID:        uint64(id),
// 		Fullname:     string(fullname),
// 		Status:       statusConvert,
// 		Total_Salary: uint64(salary),
// 	}
// }

func (TempEmployee *PermanentEmployee) InsuranceEmployee(status string) (uint64, error) {
	PreCheckStatus := CheckStatus(status, ProccededStatus)

	switch PreCheckStatus {
	case PERMANENT:
		Money := Money_on_Status(PreCheckStatus)
		return Money, nil
	default:
		return 0, errors.New("Unknown Status Please Check Again")
	}
	// 	Type of Employee must checking
}

func (TempEmployee *ContractEmployee) OvertimeEmployee(status string) (uint64, error) {
	PreCheckStatus := CheckStatus(status, ProccededStatus)
	switch PreCheckStatus {
	case CONTRACT:
		return Money_on_Status(PreCheckStatus), nil
	default:
		return 0, errors.New("Unknown Status Please Check Again")
	}

	// 	Type of Employee must checking
}

// We already know these wasnt scallable then
func (TempEmployee *TraineeEmployee) AllowanceEmployee(status string) (uint64, error) {
	// 	Type of Employee must checking
	PreCheckStatus := CheckStatus(status, ProccededStatus)

	switch PreCheckStatus {
	case TRAINEE:
		return Money_on_Status(PreCheckStatus), nil
	default:
		return 0, errors.New("Unknown Status Please Check Again")
	}
}
