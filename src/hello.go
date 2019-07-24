package main

import ("fmt"
        "time"
        "math"
      "math/rand")


func add(x int, y int) int {
	return x + y
}

func main() {
  fmt.Println("The time is", time.Now())
  fmt.Println("This is my first go language program! \n")
  fmt.Println("My favorite number is", rand.Intn(19)," \n")
  fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
  fmt.Println(math.Pi)
  fmt.Println(add(42, 13))
}
