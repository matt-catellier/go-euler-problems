package main
// https://projecteuler.net/problem=1

import (
  "fmt"
  "math"
)

func Loop(limit int) int {
  var sum int = 0
  for i:= 0; i < limit; i++ {
    if i % 3 == 0 || i % 5 == 0 {
      sum += i
    }
  }
  return sum;
}

func sumN(multiple int, limit int, sum int, num int) int {
  if multiple * num < limit {
    return sumN(multiple, limit, sum + multiple * num, num + 1)
  }
  return sum
}

func Recursive(limit int) int {
  return sumN(3, limit, 0, 0) + sumN(5, limit, 0, 0) - sumN(3*5, limit, 0, 0)
}

func SumN(limit float64, multiple float64) int {
  d := math.Floor(limit / multiple)
  return int(math.Floor(multiple * (d * (d+1)) / 2))
}

func Euler1(limit float64) int {
  return SumN(limit-1, 3) + SumN(limit-1, 5) - SumN(limit-1, 15)
}


func main() {
  // find multiples of 3 and 5 below 1000
  fmt.Println(Loop(1000))
  fmt.Println(Recursive(1000))
  fmt.Println(Euler1(1000))
}
