// Created By: Lamees Hemdan
// Created On: March 2023
//
// This file contains the main function for the go_app application.

package main

import "fmt"

func main() {
	var aBase, bBase, height float64

	// This program calculates the area of a trapezoid.
	fmt.Println("Enter the base1 of the trapezoid: ")
	fmt.Scan(&aBase)
	fmt.Println("Enter the base2 of the trapezoid: ")
	fmt.Scan(&bBase)
	fmt.Println("Enter the height of the trapezoid: ")
	fmt.Scan(&height)
	area := (aBase + bBase) * height / 2
	fmt.Println("The area of the trapezoid is:", area, "cm²")

	fmt.Println("\nDone.")
}
