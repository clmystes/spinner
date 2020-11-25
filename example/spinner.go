package main

import (
	"fmt"
	"github.com/clmystes/spinner"
	"time"
)

func main() {
	s := spinner.New("Processing staff...")
	s.Start()
	time.Sleep(1000 * time.Millisecond)
	s.Stop()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Done.")
}
