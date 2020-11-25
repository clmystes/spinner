package main

import (
	"fmt"
	"github.com/clmystes/spinner"
	"time"
)

func main() {
	s := spinner.New("Processing stuff...")
	s.Start()
	time.Sleep(1000 * time.Millisecond)
	s.Stop()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Done.")
}
