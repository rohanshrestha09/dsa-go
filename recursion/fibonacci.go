package recursion

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return n - 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
