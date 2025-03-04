//
// main.go
//
package main

import (
	"fmt"

	"github.com/k4k3ru-hub/go/config/yaml"
)


//
// Main.
//
func main() {
	// Initialize YAML config.
	if err := yaml.Init(); err != nil {
		fmt.Printf("Error: %s\v", err)
		return
	}

	// Output.
	fmt.Printf("key1: %s\n", yaml.GetString("key1"))
	fmt.Printf("key2: %t\n", yaml.GetBool("key2"))
	fmt.Printf("key3: %d\n", yaml.GetInt64("key3"))
	fmt.Printf("key4: %f\n", yaml.GetFloat64("key4"))
	for i, v := range yaml.GetArray("key5") {
		fmt.Printf("key5 (%d): %v\n", i, v)
	}
	for i, v := range yaml.GetArrayInt("key6") {
		fmt.Printf("key6 (%d): %v\n", i, v)
	}
	fmt.Printf("key7 > key8: %s\n", yaml.GetString("key7.key8"))
}
