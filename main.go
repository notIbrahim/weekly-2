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
}
