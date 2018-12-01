package limiter

import (
	"fmt"
	"log"
	"runtime"
	"testing"
	"time"
)

// TestExample should take a bit more than 400ms (less than 440ms too)
func TestTime(t *testing.T) {
	// params
	nbJobs := 1600
	nbConc := 100
	jobTime := time.Millisecond * 2000

	start := time.Now()
	limit := New(nbConc)
	m := make(chan int, 1600)
	for i := 0; i < nbJobs; i++ {
		i := i
		limit.Execute(func() {
			fmt.Println(i, runtime.NumGoroutine(), "========")
			m <- i
			time.Sleep(jobTime)
		})
	}
	limit.Wait()
	fmt.Println("done", cap(m))
	close(m)
	for index := 0; index < 160; index++ {
		<-m
	}
	duration := time.Now().Sub(start)
	expected := time.Duration(jobTime) * time.Duration(nbJobs) / time.Duration(nbConc)
	expectedTenPercent := time.Duration(float64(expected) * 1.1)
	log.Println(expectedTenPercent, " > ", duration, " > ", expected)
	// if duration > expected && duration < expectedTenPercent {
	// 	log.Println(expectedTenPercent, " > ", duration, " > ", expected)
	// } else {
	// 	t.FailNow()
	// }
}
