package runTestCase

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"gorm.io/datatypes"
)

// StructToMap converts a struct to a map using "mapstructure" tags.
// If "mapstructure" tag is missing, it falls back to "json" tag, then field name.
func StructToMap(obj interface{}) map[string]interface{} {
	if obj == nil {
		return nil
	}
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	t := v.Type()

	out := make(map[string]interface{})

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		val := v.Field(i)

		// Handle anonymous/embedded structs (flatten them)
		if field.Anonymous {
			embeddedMap := StructToMap(val.Interface())
			for k, v := range embeddedMap {
				out[k] = v
			}
			continue
		}

		// Skip unexported fields
		if field.PkgPath != "" {
			continue
		}

		tag := field.Tag.Get("mapstructure")
		if tag == "" {
			tag = field.Tag.Get("json")
		}

		// Parse tag options (e.g. "name,omitempty")
		parts := strings.Split(tag, ",")
		name := parts[0]
		if name == "-" {
			continue
		}
		if name == "" {
			name = field.Name
		}

		// Handle zero values if omitempty is present
		omitempty := false
		if len(parts) > 1 {
			for _, p := range parts[1:] {
				if p == "omitempty" {
					omitempty = true
					break
				}
			}
		}

		if omitempty && isZero(val) {
			continue
		}

		// Special handling for datatypes.JSON
		if field.Type == reflect.TypeOf(datatypes.JSON{}) {
			if bytes, ok := val.Interface().(datatypes.JSON); ok {
				if len(bytes) > 0 {
					var dest interface{}
					// Try to unmarshal to interface{}
					if err := json.Unmarshal(bytes, &dest); err == nil {
						out[name] = dest
					}
				}
			}
			continue
		}

		// Generic Map handling (including datatypes.JSONMap)
		// We convert all maps with string keys to map[string]interface{}
		// This ensures compatibility with mapstructure and avoids custom type issues
		if val.Kind() == reflect.Map && val.Type().Key().Kind() == reflect.String {
			if !val.IsNil() {
				newMap := make(map[string]interface{})
				iter := val.MapRange()
				for iter.Next() {
					k := iter.Key().String()
					vVal := iter.Value()
					
					// Recursively handle nested structs or maps if necessary
					// For now, we just pass interface{}, relying on recursion in StructToMap if we were to call it,
					// but here we are dealing with values inside a map. 
					// If the value is a JSONMap or Struct, we might want to convert it too.
					// Simple approach: just unwrap interface. 
					// If strict deep conversion is needed, we'd need more logic.
					// But for JSONMap, it usually contains primitives or []interface{} or map[string]interface{}.
					
					// One special case: nested datatypes.JSONMap
					if vVal.Kind() == reflect.Map {
                         // We can't easily recurse StructToMap because it expects a struct? 
                         // Actually StructToMap handles Map logic now too (if we extract it).
                         // But let's just use the interface value, mapstructure usually handles nested maps fine
                         // AS LONG AS the type isn't obscure.
                         // Since we are peeling off the top-level custom type here, 
                         // we hope nested ones are standard or mapstructure can handle them.
                         // However, to be safe against nested JSONMap:
                         newMap[k] = convertMapValue(vVal.Interface())
					} else {
						newMap[k] = vVal.Interface()
					}
				}
				out[name] = newMap
			}
			continue
		}

		// Recurse for nested structs (excluding time.Time)
		if val.Kind() == reflect.Struct || (val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.Struct) {
			// Check if it's time.Time
			if isTime(val) {
				out[name] = val.Interface()
				continue
			}
			
			nested := StructToMap(val.Interface())
			// If nested is empty and omitempty is set, skip? 
			// But we already checked isZero. 
			// If nested is NOT nil, but map is empty (no exported fields?), we might still want to include it?
			// Usually StructToMap returns map. 
			if nested != nil {
				out[name] = nested
			} else if !omitempty {
                 // If conversion failed or nil, but we want it?
                 // If val was nil ptr, StructToMap returns nil.
            }
			continue
		}

		out[name] = val.Interface()
	}
	return out
}

// convertMapValue helper to recursively convert map[string]any-like types
func convertMapValue(v interface{}) interface{} {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Map && val.Type().Key().Kind() == reflect.String {
		if val.IsNil() {
			return nil
		}
		newMap := make(map[string]interface{})
		iter := val.MapRange()
		for iter.Next() {
			newMap[iter.Key().String()] = convertMapValue(iter.Value().Interface())
		}
		return newMap
	}
	return v
}
func isTime(v reflect.Value) bool {
	if v.Type() == reflect.TypeOf(time.Time{}) {
		return true
	}
	if v.Kind() == reflect.Ptr && v.Elem().Type() == reflect.TypeOf(time.Time{}) {
		return true
	}
	return false
}

func isZero(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Ptr, reflect.Interface:
        return v.IsNil()
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !isZero(v.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Struct:
        // Special case for time.Time
        if isTime(v) {
            return v.Interface().(time.Time).IsZero()
        }
        // For other structs, check if all fields are zero? 
        // Or simply check if it equals Zero value?
        z := reflect.Zero(v.Type())
		return v.Interface() == z.Interface()
	default:
		z := reflect.Zero(v.Type())
		return v.Interface() == z.Interface()
	}
}
