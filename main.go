package main

import (
	"fmt"
	"stargen/starforged/delves"
)

func main() {
	t := delves.NewTheme()

	fmt.Printf("Theme: %s\n", t.Name)
	fmt.Printf("- Description: %s\n", t.Description)
	fmt.Printf("- Feature: %s\n", t.GetFeatureString())
	fmt.Printf("  - Danger: %s\n", t.GetDangerString())
	fmt.Println()
}
