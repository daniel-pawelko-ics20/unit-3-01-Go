// Copyright (c) 2021 Daniel Pawelko All rights reserved
//
// Created by: Daniel Pawelko
// Created on: May 2021
// This file contains GO program that area of trapazoid

package main

import (
	"fmt"
)

// This main function will ask user for dimensions and output calculations
func main() {
	// Defining variables
	var abase int
	var bbase int
	var height int

	fmt.Println("Area of a Trapezoid, with JS\n")
	fmt.Println("A = [(a+b) / 2] x h\n")
	fmt.Println()

	// User Input
	fmt.Println("Please enter integers for dimensions:")
	fmt.Println("a base(mm): ")
	fmt.Scanln(&abase)
	fmt.Println()

	fmt.Println("b base(mm):")
	fmt.Scanln(&bbase)
	fmt.Println()

	fmt.Println("height(mm):")
	fmt.Scanln(&height)
	fmt.Println()

	// Print out area
	fmt.Println("Area is: ", ((abase+bbase)/2)*height, "mm²")
}
