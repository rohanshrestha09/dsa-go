package recursion

import "fmt"

// TOH src = Source aux = Auxiliary dest = Destination n = Number of disks
func TOH(n int, src, aux, dest byte) {
	if n == 1 {
		fmt.Printf("%c -> %c\n", src, dest)
	} else {
		TOH(n-1, src, dest, aux)
		TOH(1, src, aux, dest)
		TOH(n-1, aux, src, dest)
	}
}
