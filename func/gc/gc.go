package main
import (
	"runtime"
	"fmt"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)
}