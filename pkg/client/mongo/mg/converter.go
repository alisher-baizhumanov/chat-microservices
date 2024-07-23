package mg

import (
	"go.mongodb.org/mongo-driver/bson"
)

func toBsonM(data map[string]any) bson.M {
	result := bson.M{}
	for key, value := range data {
		switch v := value.(type) {
		case map[string]any:
			result[key] = toBsonM(v) // Recursively convert nested maps
		case []any:
			result[key] = toBsonSlice(v) // Convert slices
		default:
			result[key] = v
		}
	}
	return result
}

func toBsonSlice(data []any) []any {
	result := make([]any, len(data))
	for i, value := range data {
		switch v := value.(type) {
		case map[string]any:
			result[i] = toBsonM(v) // Recursively convert nested maps
		case []any:
			result[i] = toBsonSlice(v) // Recursively convert nested slices
		default:
			result[i] = v
		}
	}
	return result
}
