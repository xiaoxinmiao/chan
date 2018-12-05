package limiter

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// TestExample keep 10 goroutine always running
func TestTime(t *testing.T) {
	// params
	nbJobs := 160
	nbConc := 10
	jobTime := time.Millisecond * 2000

	limit := New(nbConc)
	m := make([]int, 0)
	for i := 0; i < nbJobs; i++ {
		i := i
		limit.Execute(func() {
			fmt.Println(i, runtime.NumGoroutine(), "========")
			m = append(m, i)
			time.Sleep(jobTime)
		})
	}
	time.Sleep(10 * time.Second)
	fmt.Println("done", len(m))
	for i := 0; i < nbJobs; i++ {
		i := i
		limit.Execute(func() {
			fmt.Println(i, runtime.NumGoroutine(), "========")
			m = append(m, i)
			time.Sleep(jobTime)
		})
	}
	time.Sleep(10 * time.Minute)
}
