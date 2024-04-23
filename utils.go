package main

import (
	"fmt"
	"strconv"
)

func addArrays(a, b []int64) []int64 {
	result := make([]int64, len(a))
	for i := range a {
		result[i] = a[i] + b[i]
	}

	return result
}

func xorArrays(a, b []int64) []int64 {
	result := make([]int64, len(a))
	for i := range a {
		result[i] = a[i] ^ b[i]
	}

	return result
}

//func xorArray_with_Pvalue(L []int64, pValue int64) []int64 {
//	for i := 0; i < len(L); i++ {
//		L[i] = L[i] ^ pValue
//	}
//	return L
//}

func swap(L, R *[]int64) {
	*L, *R = *R, *L
}

func making_bit_array(input_as_a_bit_stream string) []int64 {
	var bits_array = []string{}
	for _, r := range input_as_a_bit_stream {
		bits_array = append(bits_array, string(r))
	}
	new_int_array_of_bits := make([]int64, len(bits_array))
	for i := 0; i < len(bits_array); i++ {
		bit, err := strconv.ParseInt(bits_array[i], 2, 32)
		if err != nil {
			fmt.Println("Error parsing bit:", err)
			continue
		}
		new_int_array_of_bits[i] = bit
	}
	return new_int_array_of_bits
}
