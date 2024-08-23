package main

import (
	"fmt"
	"stargen/starforged/delves"
)

func main() {
	t := delves.NewTheme()

	fmt.Printf("Theme: %s\n", t.Name)
	fmt.Printf("- Description: %s\n", t.Description)
	for k := 0; k < 3; k++ {
		fmt.Printf("- Feature: %s\n", t.GetFeatureString())
		fmt.Printf("  - Danger: %s\n", t.GetDangerString())
		fmt.Printf("  - Danger: %s\n", t.GetDangerString())
	}
}
