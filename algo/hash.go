package algo

import (
	"log"
	"reflect"
	"strconv"
)

func CharCode(char string) int {
	if len(char) == 1 {
		return int(byte(char[0]))
	}
	return -1
}

func HashCode(values ...interface{}) uint64 {
	var h uint64 = 0 
	var pow uint64 = 1
	var m uint64 = 1e9 + 9
	if len(values) == 1 {
		input := generic(values[0])
		for _, b := range input {
			char := uint64(CharCode(string(b)))
			h = (h + char*pow)%m
			pow = ((pow << 8) - pow )%m
			log.Println(pow)
		}
		// h =
		return h
	}
	index := make(map[string]int, len(values))
	max := 0
	for _, value := range values {
		input := generic(value)
		if len(input)-1 > max {
			max = len(input) - 1
		}
		index[input] = 0
	}

	for max >= 0 {
		h = (h << 8)
		for input, i := range index {
			if i < len(input) {
				char := uint64(CharCode(string(input[i])))
				h = (h + char*pow)
				index[input]++
			}
		}
		h = h % m
		pow = ((pow << 8) - pow)%m
		log.Println(pow)
		max--
	}
	return h
}

func generic(value interface{}) string {
	v := reflect.TypeOf(value)

	if v.Kind() == reflect.Int {
		r := strconv.Itoa(value.(int))
		return r
	}

	if v.Kind() == reflect.String {
		r := value.(string)
		return r
	}
	return ""
}
