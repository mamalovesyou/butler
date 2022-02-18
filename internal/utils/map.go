package utils

// MergeMaps merges maps of type map[string]interface{}
// Note: It doesn't handle duplicate keys, it will keep the latest map override
func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
