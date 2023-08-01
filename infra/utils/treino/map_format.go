package utils

import "net/url"

func UrlValuesToMap(values url.Values) (mapValues map[string]string) {
	for key, value := range values {
		mapValues = map[string]string{key: value[0]}
	}
	return mapValues
}

func MapValueToArray(mapValues map[string]string) (values []interface{}) {
	for _, value := range mapValues {
		values = append(values, value)
	}
	return values
}

func MapKeyToArray(mapValues map[string]string) (keys []interface{}) {
	for key := range mapValues {
		keys = append(keys, key)
	}
	return keys
}

func MapToGormConditions(mapValues map[string]string) (conditions string) {
	for key := range mapValues {
		if len(conditions) != 0 {
			conditions += " AND "
		}
		conditions += key + "= ?"
	}
	return conditions
}
