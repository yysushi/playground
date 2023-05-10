package main

import (
	"fmt"
	"sort"
	"testing"
)

func main() {
	// This reports a custom benchmark metric relevant to a
	// specific algorithm (in this case, sorting).
	br := testing.Benchmark(func(b *testing.B) {
		var compares int64
		for i := 0; i < b.N; i++ {
			s := []int{5, 4, 3, 2, 1}
			sort.Slice(s, func(i, j int) bool {
				compares++
				return s[i] < s[j]
			})
		}
		// This metric is per-operation, so divide by b.N and
		// report it as a "/op" unit.
		b.ReportMetric(float64(compares)/float64(b.N), "compares/op")
		// This metric is per-time, so divide by b.Elapsed and
		// report it as a "/ns" unit.
		b.ReportMetric(float64(compares)/float64(b.Elapsed().Nanoseconds()), "compares/ns")
	})
	fmt.Println(br)
}

// func myFunc(long bool) {
// 	_ = fmt.Sprintf("hello")
// 	if long {
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func BenchmarkHello(b *testing.B) {
// 	var i int
// 	log.Println(i, "started")
// 	for i = 0; i < b.N; i++ {
// 		myFunc(false)
// 	}
// 	log.Println(i, "finished")
// }

// func BenchmarkHelloLong(b *testing.B) {
// 	var i int
// 	log.Println(i, "started")
// 	for i = 0; i < b.N; i++ {
// 		myFunc(true)
// 	}
// 	log.Println(i, "finished")
// }

// TODO: RunParallel, testing.PB
