package main

import (
	"fmt"
)

var (
  bmr float32
  weight float32
  height float32
  age float32
  bmr_const float32
  result float32
  sedentary float32
  lightly float32
  moderately float32
  active float32
  sup_active float32
)

func main() {
  fmt.Print("Enter age: ")
  fmt.Scanln(&age)
  fmt.Print("Input height in cm: ")
  fmt.Scanln(&height)
  fmt.Print("Enter weight: ")
  fmt.Scanln(&weight)
  result = get_bmr(age, weight, height)
  fmt.Printf("\nYour basal metabolic rate is: %v\n\n",result)
  daily_calorie(result)
}

func get_bmr(age float32, weight float32, height float32) float32 {
  bmr_const = 66.47
  bmr = bmr_const + (13.75 * weight) + (5.003 * height) - (6.755 * age)
  return bmr
}


// Gotta make a map instead of a slice to return these values
// Or maybe figure out something else idk
func daily_calorie(result float32) {
  //var calories = [5]float32{sedentary,lightly,moderately,active,sup_active}
  sedentary = result * 1.2
  lightly = result * 1.375
  moderately = result * 1.55
  active = result * 1.725
  sup_active = result * 1.9
  fmt.Print("\nAmount of calories required to maintain your weight depending on your level of activity:\n\n")
  fmt.Println("Sedentary:", sedentary)
  fmt.Println("Lightly active:", lightly)
  fmt.Println("Moderately active:", moderately)
  fmt.Println("Very active:", active)
  fmt.Println("Super active:", sup_active)
}
