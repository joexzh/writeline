package main

import (
	"fmt"
	"github.com/joexzh/writeline"
	"log"
	"sync"
	"time"
)

func main() {
	names := []string{"cat", "dog"}
	var wg sync.WaitGroup

	lw, err := writeline.New(len(names) * 2)
	if err != nil {
		log.Fatal(err)
	}
	lw.SetFlushDuration(2 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {

			lw.WriteLine(0, fmt.Sprintf("%v, %vth loop", names[0], i))
			lw.WriteLine(0+1, fmt.Sprintf("%v...%v/%v", names[0], i+1, 100))
			lw.Flush()
			time.Sleep(100 * time.Millisecond)
		}
		lw.WriteLine(0, fmt.Sprintf("%v done", names[0]))
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {

			lw.WriteLine(2, writeline.Style(writeline.Green+writeline.Bold, fmt.Sprintf("%v, %vth loop", names[1], i))+"test")
			lw.WriteLine(2+1, fmt.Sprintf("%v...%v/%v", names[1], i+1, 100))
			time.Sleep(50 * time.Millisecond)
		}
		lw.WriteLine(2, fmt.Sprintf("%v done", names[1]))
		wg.Done()
	}()

	wg.Wait()
	lw.WriteLastLine("dog 100%")
	lw.WriteNewLine("all done")
	lw.Close()
	fmt.Print(lw.Lines())
}
