package main

import (
  "fmt"
)

func main() {
  fmt.Println(canPlaceFlowers([]int{1,0,0,0,0,1}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
    count := 0
    for i := range flowerbed {
      adj := flowerbed[i] == 1
      if i > 0 && flowerbed[i-1] == 1 {
        adj = true
      }
      if i < len(flowerbed)-1 && flowerbed[i+1] == 1 {
        adj = true
      }
      if !adj {
        count++
        flowerbed[i] = 1
      }
    }
    return count >= n
}
