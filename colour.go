package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to simulate the rotation of colors
func rotateColors(colors []string, rotations int) []string {
	n := len(colors)
	// Simulate the rotation
	for i := 0; i < rotations; i++ {
		// Rotate by moving the last element to the front
		last := colors[n-1]
		for j := n - 1; j > 0; j-- {
			colors[j] = colors[j-1]
		}
		colors[0] = last
	}
	return colors
}

// Function to calculate the probability of a specific color in a given position
func calculateColorProbability(colors []string, rotations, trials int, targetColor string) float64 {

	count := 0

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Simulate the process multiple times to estimate the probability
	for i := 0; i < trials; i++ {
		// Copy the colors array to avoid modifying the original
		rotatedColors := append([]string(nil), colors...)
		rotatedColors = rotateColors(rotatedColors, rotations)

		// Check if the target color is at the first position (you can adjust the position)
		if rotatedColors[0] == targetColor {
			count++
		}
	}

	// Calculate the probability
	probability := float64(count) / float64(trials)
	return probability
}

func main() {
	// List of available colors
	colors := []string{"Red", "Green", "Blue", "Yellow", "Orange", "Purple"}

	// Number of rotations
	rotations := 3

	// Number of trials for simulation
	trials := 10000

	// Target color to check probability for
	targetColor := "Red"

	// Calculate the probability of the target color being in the first position
	probability := calculateColorProbability(colors, rotations, trials, targetColor)

	// Output the result
	fmt.Printf("Probability of %s being in the first position after %d rotations: %.4f\n", targetColor, rotations, probability)
}
