package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

func main() {
	PArray := initializePArray()
	const key = "blowfish is a symmetric-key block cipher algorithm used in cryptography."

	keyBytes := []byte(key)
	keyBytes_in_uint32 := []uint32{}
	for i := 0; i < len(keyBytes); i += 4 {
		keyBytes_in_uint32 = append(keyBytes_in_uint32, binary.BigEndian.Uint32(keyBytes[i:i+4]))
	}
	for i := 0; i < 18; i++ {
		PArray[i] = PArray[i] ^ keyBytes_in_uint32[i%len(keyBytes_in_uint32)]
	}
	PArray_in_bytes := []byte{byte(PArray[0]), byte(PArray[1]), byte(PArray[2]), byte(PArray[3]), byte(PArray[4]),
		byte(PArray[5]), byte(PArray[6]), byte(PArray[7]), byte(PArray[8]), byte(PArray[9]), byte(PArray[10]),
		byte(PArray[11]), byte(PArray[12]), byte(PArray[13]), byte(PArray[14]), byte(PArray[15]), byte(PArray[16]),
		byte(PArray[17])}
	var binary_string string
	for i := 0; i < len(PArray_in_bytes); i++ {
		binary_string = binary_string + strconv.FormatInt(int64(PArray[i]), 2)
	}
	PArray_in_bits := making_bit_array(binary_string)
	P1 := PArray_in_bits[:32]
	P2 := PArray_in_bits[32:64]
	P3 := PArray_in_bits[64:96]
	P4 := PArray_in_bits[96:128]
	P5 := PArray_in_bits[128:160]
	P6 := PArray_in_bits[160:192]
	P7 := PArray_in_bits[192:224]
	P8 := PArray_in_bits[224:256]
	P9 := PArray_in_bits[256:288]
	P10 := PArray_in_bits[288:320]
	P11 := PArray_in_bits[320:352]
	P12 := PArray_in_bits[352:384]
	P13 := PArray_in_bits[384:416]
	P14 := PArray_in_bits[416:448]

	PArray_in_bits_indexed := [][]int64{P1, P2, P3, P4, P5, P6, P7, P8, P9, P10, P11, P12, P13, P14}
	fmt.Println(len(PArray_in_bits_indexed))

	//sboxes := initializeSBoxes(PArray)
	var entered_string string
	for {
		fmt.Print("Enter a string of 9 characters: ")
		fmt.Scan(&entered_string)
		if len(entered_string) == 9 && len(entered_string) != 0 {
			break
		}
		fmt.Println("Invalid input. Please try again.")
	}
	fmt.Println("Entered string: ", entered_string)

	bits_as_a_byte_array := []byte(entered_string)
	for i := 0; i < len(bits_as_a_byte_array); i++ {
		binary_string = binary_string + strconv.FormatInt(int64(bits_as_a_byte_array[i]), 2)
	}

	bits_array_int := making_bit_array(binary_string)
	bits_array_int_left := bits_array_int[:32]
	bits_array_int_right := bits_array_int[32:]
	for i := 0; i < 16; i++ {
		bits_array_int_left = xorArrays(bits_array_int_left, PArray_in_bits_indexed[i])
		bits_array_int_right = xorArrays(F(bits_array_int_left), bits_array_int_right)
		swap(&bits_array_int_left, &bits_array_int_right)
	}
	swap(&bits_array_int_left, &bits_array_int_right)
	bits_array_int_right = xorArrays(bits_array_int_right, PArray_in_bits_indexed[16])
	bits_array_int_left = xorArrays(bits_array_int_left, PArray_in_bits_indexed[17])
	bits_array_int = append(bits_array_int_left, bits_array_int_right...)
	fmt.Println(bits_array_int)
}

func F(L []int64) []int64 {
	S1 := L[:8]
	S2 := L[8:16]
	S3 := L[16:24]
	S4 := L[24:]
	result := addArrays(xorArrays(addArrays(S1, S2), S3), S4)
	return result
}
