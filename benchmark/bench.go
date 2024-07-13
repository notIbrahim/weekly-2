package benchmark

import (
	"fmt"
	"testing"
	"weekly-2/data"
)

// How Exactly they run for
func ChannelingTest(Bench *testing.B) {
	var NonChannelBench, ChannelBench []any
	fmt.Println("Testing Benchmark")
	NonChannelBench = data.Work_NonChannel()
	ChannelBench = NonChannelBench
	Bench.Run("All_Total_Summary", func(Bench *testing.B) {
		fmt.Println("Total Salary Sums with non channel")
		for i := 0; i < Bench.N; i++ {
			data.All_Total_Summary(false, NonChannelBench)
		}
		fmt.Printf("Time Lapsed %v\n", Bench.Elapsed().Nanoseconds())
	})
	Bench.Run("All_Total_Summary", func(Bench *testing.B) {
		fmt.Println("Total Salary Sums with channel")
		for i := 0; i < Bench.N; i++ {
			data.All_Total_Summary(true, ChannelBench)
		}
		fmt.Printf("Time Lapsed %v\n", Bench.Elapsed().Nanoseconds())
	})
}
