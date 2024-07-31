package common

import "encoding/json"

func LogStruct(v any) string {
	marshal, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func GetPointer[T any](value T) *T {
	return &value
}

func GetValue[T any](pointer *T) T {
	if pointer == nil {
		var zeroValue T
		return zeroValue
	}
	return *pointer
}
