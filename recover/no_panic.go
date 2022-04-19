package recover

import "fmt"

// iPanic: reference https://go.dev/blog/defer-panic-and-recover
func iPanic(i int) {
	if i > 3 {
		fmt.Println("panicking!")
		panic(i)
	}
	defer fmt.Printf("deferring %d\n", i)
	fmt.Printf("printing %d\n", i)
	iPanic(i + 1)
}

func iHandlePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovering %d\n", r)
		}
	}()

	iPanic(0)
}
