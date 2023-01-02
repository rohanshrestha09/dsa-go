package recursion

import "fmt"

func Run() {
	fmt.Println("Factorial :", factorial(5))

	fmt.Print("Fibonacci series : ")
	for i := 1; i <= 8; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}

	fmt.Println("\nTower of Hanoi")
	TOH(4, 'S', 'A', 'D')
}
