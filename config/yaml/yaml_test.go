//
// yaml_test.go
//
package yaml_test

import (
	"os"
	"testing"

	"yaml"
)


//
// Test Init function for single YAML file.
//
func TestInit_SingleYAMLFile(t *testing.T) {
	file, err := createTempYAMLFile(`
key1: value1
key2: value2
`)
    if err != nil {
        t.Fatalf("Failed to create first YAML file: %v", err)
    }
    defer os.Remove(file)

	os.Args = []string{"cmd", "-yaml", file}

	err = yaml.Init()
	if err != nil {
		t.Fatalf("Failed to execute Init. Error: %v\n", err)
	}

	if val, ok := yaml.Config["key1"].(string); !ok || val != "value1" {
		t.Errorf("Failed to read key1 value. Expected: value1, actual: %s\n", val)
	}
	if val, ok := yaml.Config["key2"].(string); !ok || val != "value2" {
		t.Errorf("Failed to read key1 value. Expected: value2, actual: %s\n", val)
	}
}


//
// Test Init function for two YAML files.
//
func TestInit_TwoYAMLFiles(t *testing.T) {
	file1, err := createTempYAMLFile(`
key1: old_value
key2: value2
`)
	if err != nil {
		t.Fatalf("Failed to create first YAML file: %v", err)
	}
	defer os.Remove(file1)

	file2, err := createTempYAMLFile(`
key1: new_value
key3: value3
`)
	if err != nil {
		t.Fatalf("Failed to create second YAML file: %v", err)
	}
	defer os.Remove(file2)

	os.Args = []string{"cmd", "-yaml", file1, "-yaml", file2}

	err = yaml.Init()
	if err != nil {
		t.Fatalf("Failed to execute Init. Error: %v\n", err)
	}

	if val, ok := yaml.Config["key1"].(string); !ok || val != "new_value" {
		t.Errorf("Failed to read key1 value. Expected: new_value, actual: %s\n", val)
	}
	if val, ok := yaml.Config["key2"].(string); !ok || val != "value2" {
		t.Errorf("Failed to read key3 value. Expected: value2, actual: %s\n", val)
	}
	if val, ok := yaml.Config["key3"].(string); !ok || val != "value3" {
		t.Errorf("Failed to read key3 value. Expected: value3, actual: %s\n", val)
	}
}


//
// create temporary YAML file.
//
func createTempYAMLFile(content string) (string, error) {
	tempFile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	if _, err := tempFile.WriteString(content); err != nil {
		return "", err
	}

	return tempFile.Name(), nil
}
