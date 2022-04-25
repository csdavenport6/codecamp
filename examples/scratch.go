package main

func main() {
	addNums(1, 2)
}

func addNums(x int, y int) (int, int) {
	return x + y, x + y + 1
}
