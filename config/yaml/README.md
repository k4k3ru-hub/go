# YAML Utility for Go

This is a lightweight and efficient YAML configuration utility for Go.
It simplifies loading and accessing YAML configuration files with support for multiple files and nested structures.


## Features
- Load YAML configuration files with a simple interface
- Support for multiple YAML files (overriding keys in order)
- Convenient functions to fetch values as different types (`string`, `bool`, `int64`, `float64`, `array`)
- Support for nested keys using dot notation (e.g., `config.GetString("key1.subkey")`)


## Installation

```console
go get github.com/k4k3ru-hub/go/config/yaml
```


## Usage

### Basic

1. Create a YAML file.

```yaml
key1: value1
key2: true
key3: 10
key4: 3.14
key5:
- a
- b
- c
key6: 
- 1
- 2
- 3
key7:
  key8: value8
```

2. Import the module

```go
import "github.com/k4k3ru-hub/go/config/yaml"
```

3. Initialize the module

```go
if err := yaml.Init(); err != nil {
    // Error Handling
}
```

4. Read the values

```go
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
```

5. Run Go

```consle
go run main.go -yaml config.yaml
```


### Multiple YAML Files (Merging Configurations)

1. Create one more YAML (`config2.yaml`) which you want to merge

```config2.yaml
key1: new-value1
```

2. Run Go

```consle
go run main.go -yaml config.yaml -yaml config2.yaml
```


## Support me
I am a Japanese developer, and your support is a great encouragement for my work!
In addition to support, feel free to reach out with comments, feature requests, or development inquiries!

Thank you for your support😊

[![Support on Ko-fi](https://img.shields.io/badge/Ko--fi-Support%20Me-blue?style=flat-square&logo=ko-fi)](https://ko-fi.com/k4k3ru)


## License
This repository is open-source and distributed under the MIT License.
