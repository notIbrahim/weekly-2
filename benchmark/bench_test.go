package benchmark

import (
	"testing"
)

func TestBenchmark(t *testing.T) {
	testing.Benchmark(ChannelingTest)
}
