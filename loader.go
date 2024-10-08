// package confhandler

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// 	"reflect"
// 	"strconv"

// 	"gopkg.in/yaml.v3"
// )

// // LoadConfigToStruct loads a YAML file and maps its values to a provided struct.
// // It also dynamically resolves environment variables and parses values to the correct types.
// func LoadConfigToStruct(filePath string, out interface{}) error {
// 	// Load YAML file
// 	data, err := os.ReadFile(filePath)
// 	if err != nil {
// 		return err
// 	}

// 	// Unmarshal the YAML into a generic map
// 	var configData map[string]interface{}
// 	err = yaml.Unmarshal(data, &configData)
// 	if err != nil {
// 		return err
// 	}

// 	// Map the values to the struct using reflection
// 	return mapToStruct(configData, out)
// }

// // mapToStruct maps the values from a map to the fields of a struct using reflection.
// func mapToStruct(data map[string]interface{}, out interface{}) error {
// 	v := reflect.ValueOf(out)
// 	if v.Kind() != reflect.Ptr || v.IsNil() {
// 		return errors.New("output must be a non-nil pointer to a struct")
// 	}

// 	v = v.Elem() // Dereference the pointer to get the struct

// 	// Iterate through the struct fields
// 	for i := 0; i < v.NumField(); i++ {
// 		field := v.Type().Field(i)
// 		fieldValue := v.Field(i)

// 		// Get the field tag (e.g., yaml:"field_name")
// 		tag := field.Tag.Get("yaml")
// 		if tag == "" {
// 			tag = field.Name // Fallback to field name if no tag
// 		}

// 		// Get the corresponding value from the config data
// 		rawValue, ok := data[tag]
// 		if !ok {
// 			continue // If the key is not found in the map, skip it
// 		}

// 		// Resolve environment variables and parse value based on field type
// 		resolvedValue := ResolveEnvVars(fmt.Sprintf("%v", rawValue))

// 		switch fieldValue.Kind() {
// 		case reflect.String:
// 			fieldValue.SetString(resolvedValue)
// 		case reflect.Int, reflect.Int64:
// 			parsedValue, err := strconv.Atoi(resolvedValue)
// 			if err != nil {
// 				return err
// 			}
// 			fieldValue.SetInt(int64(parsedValue))
// 		case reflect.Bool:
// 			parsedValue, err := strconv.ParseBool(resolvedValue)
// 			if err != nil {
// 				return err
// 			}
// 			fieldValue.SetBool(parsedValue)
// 		case reflect.Float64:
// 			parsedValue, err := strconv.ParseFloat(resolvedValue, 64)
// 			if err != nil {
// 				return err
// 			}
// 			fieldValue.SetFloat(parsedValue)
// 		default:
// 			return fmt.Errorf("unsupported field type: %s", field.Type)
// 		}
// 	}

// 	return nil
// }

// package confhandler

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// 	"reflect"
// 	"strconv"

// 	"gopkg.in/yaml.v3"
// )

// // LoadConfigToStruct loads a YAML file and maps its values to a provided struct.
// // It also dynamically resolves environment variables and parses values to the correct types.
// func LoadConfigToStruct(filePath string, out interface{}) error {
// 	// Load YAML file
// 	data, err := os.ReadFile(filePath)
// 	if err != nil {
// 		return err
// 	}

// 	// Unmarshal the YAML into a generic map
// 	var configData map[string]interface{}
// 	err = yaml.Unmarshal(data, &configData)
// 	if err != nil {
// 		return err
// 	}

// 	// Map the values to the struct using reflection
// 	return mapToStruct(configData, out)
// }

// // mapToStruct maps the values from a map to the fields of a struct using reflection.
// func mapToStruct(data map[string]interface{}, out interface{}) error {
// 	v := reflect.ValueOf(out)
// 	if v.Kind() != reflect.Ptr || v.IsNil() {
// 		return errors.New("output must be a non-nil pointer to a struct")
// 	}

// 	v = v.Elem() // Dereference the pointer to get the struct

// 	// Iterate through the struct fields
// 	for i := 0; i < v.NumField(); i++ {
// 		field := v.Type().Field(i)
// 		fieldValue := v.Field(i)

// 		// Get the field tag (e.g., yaml:"field_name")
// 		tag := field.Tag.Get("yaml")
// 		if tag == "" {
// 			tag = field.Name // Fallback to field name if no tag
// 		}

// 		// Get the corresponding value from the config data
// 		rawValue, ok := data[tag]
// 		if !ok {
// 			continue // If the key is not found in the map, skip it
// 		}

// 		// Handle nested structs by recursively calling mapToStruct
// 		if fieldValue.Kind() == reflect.Struct {
// 			nestedData, ok := rawValue.(map[string]interface{})
// 			if !ok {
// 				return fmt.Errorf("expected map for nested struct, got: %T", rawValue)
// 			}
// 			err := mapToStruct(nestedData, fieldValue.Addr().Interface())
// 			if err != nil {
// 				return err
// 			}
// 			continue
// 		}

// 		// Resolve environment variables and parse value based on field type
// 		resolvedValue := ResolveEnvVars(fmt.Sprintf("%v", rawValue))

// 		switch fieldValue.Kind() {
// 		case reflect.String:
// 			fieldValue.SetString(resolvedValue)
// 		case reflect.Int, reflect.Int64:
// 			parsedValue, err := strconv.Atoi(resolvedValue)
// 			if err != nil {
// 				return err
// 			}
// 			fieldValue.SetInt(int64(parsedValue))
// 		case reflect.Bool:
// 			parsedValue, err := strconv.ParseBool(resolvedValue)
// 			if err != nil {
// 				return err
// 			}
// 			fieldValue.SetBool(parsedValue)
// 		case reflect.Float64:
// 			parsedValue, err := strconv.ParseFloat(resolvedValue, 64)
// 			if err != nil {
// 				return err
// 			}
// 			fieldValue.SetFloat(parsedValue)
// 		default:
// 			return fmt.Errorf("unsupported field type: %s", field.Type)
// 		}
// 	}

// 	return nil
// }

package confhandler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"

	"gopkg.in/yaml.v3"
)

// LoadConfigToStruct loads a YAML file and maps its values to a provided struct.
// It also dynamically resolves environment variables and parses values to the correct types.
func LoadConfigToStruct(filePath string, out interface{}) error {
	// Load YAML file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Unmarshal the YAML into a generic map
	var configData map[string]interface{}
	err = yaml.Unmarshal(data, &configData)
	if err != nil {
		return err
	}

	// Map the values to the struct using reflection
	return mapToStruct(configData, out)
}

// mapToStruct maps the values from a map to the fields of a struct using reflection.
func mapToStruct(data map[string]interface{}, out interface{}) error {
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("output must be a non-nil pointer to a struct")
	}

	v = v.Elem() // Dereference the pointer to get the struct

	// Iterate through the struct fields
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fieldValue := v.Field(i)

		// Get the field tag (e.g., yaml:"field_name")
		tag := field.Tag.Get("yaml")
		if tag == "" {
			tag = field.Name // Fallback to field name if no tag
		}

		// Get the corresponding value from the config data
		rawValue, ok := data[tag]
		if !ok {
			continue // If the key is not found in the map, skip it
		}

		// Handle nested structs by recursively calling mapToStruct
		if fieldValue.Kind() == reflect.Struct {
			nestedData, ok := rawValue.(map[string]interface{})
			if !ok {
				return fmt.Errorf("expected map for nested struct, got: %T", rawValue)
			}
			err := mapToStruct(nestedData, fieldValue.Addr().Interface())
			if err != nil {
				return err
			}
			continue
		}

		// Handle slices (for example, the Migration field)
		if fieldValue.Kind() == reflect.Slice {
			sliceData, ok := rawValue.([]interface{})
			if !ok {
				return fmt.Errorf("expected slice for field %s, got: %T", field.Name, rawValue)
			}
			stringSlice := make([]string, len(sliceData))
			for i, item := range sliceData {
				stringSlice[i] = ResolveEnvVars(fmt.Sprintf("%v", item))
			}
			fieldValue.Set(reflect.ValueOf(stringSlice))
			continue
		}

		// Resolve environment variables and parse value based on field type
		resolvedValue := ResolveEnvVars(fmt.Sprintf("%v", rawValue))

		switch fieldValue.Kind() {
		case reflect.String:
			fieldValue.SetString(resolvedValue)
		case reflect.Int, reflect.Int64:
			parsedValue, err := strconv.Atoi(resolvedValue)
			if err != nil {
				return err
			}
			fieldValue.SetInt(int64(parsedValue))
		case reflect.Bool:
			parsedValue, err := strconv.ParseBool(resolvedValue)
			if err != nil {
				return err
			}
			fieldValue.SetBool(parsedValue)
		case reflect.Float64:
			parsedValue, err := strconv.ParseFloat(resolvedValue, 64)
			if err != nil {
				return err
			}
			fieldValue.SetFloat(parsedValue)
		default:
			return fmt.Errorf("unsupported field type: %s", field.Type)
		}
	}

	return nil
}
