package limiter

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// TestExample Limiter is new instance
func TestTime(t *testing.T) {
	go func() {
		// params
		nbJobs := 6
		nbConc := 5
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
		//	limit.Wait()
		time.Sleep(20 * time.Second)
		fmt.Printf("no wait,goroutine:%v,chan count:%v", runtime.NumGoroutine(), len(limit))
		time.Sleep(10 * time.Minute)
	}()
	go func() {
		// params
		nbJobs := 6
		nbConc := 5
		jobTime := time.Millisecond * 2000

		limit := New(nbConc)
		m := make([]int, 0)
		for i := 0; i < nbJobs; i++ {
			i := i
			limit.Execute(func() {
				fmt.Println(i, runtime.NumGoroutine(), "|||||||||||")
				m = append(m, i)
				time.Sleep(jobTime)
			})
		}
		limit.Wait()
		fmt.Printf("wait end,goroutine:%v,chan count:%v", runtime.NumGoroutine(), len(limit))

	}()
	time.Sleep(10 * time.Minute)
}
