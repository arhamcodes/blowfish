package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	const sixtyfour_bit_cypher_text string = "arhamtma"
	var sixtyfour_bit_cypher_string string
	for _, c := range sixtyfour_bit_cypher_text {
		sixtyfour_bit_cypher_string += fmt.Sprintf("%08b", c)
	}

	L := sixtyfour_bit_cypher_string[:len(sixtyfour_bit_cypher_string)/2]
	R := sixtyfour_bit_cypher_string[len(sixtyfour_bit_cypher_string)/2:]

	var big_L, big_R big.Int
	L_64_bit_cypher, ok := big_L.SetString(L, 10)
	if !ok {
		fmt.Println("error in setting L_64_bit_cypher to big.L")
	}
	R_64_bit_cypher, ok := big_R.SetString(R, 10)
	if !ok {
		fmt.Println("error in setting R_64_bit_cypher to big.R")
	}
	const fourfourtyeight_bit_cypher_text string = "AO;Joi;HIUASVBBDOI;;oindOIDN;ODnouiasdnoiNAOICHOIAhCOIPC"
	var fourfourtyeight_bit_cypher_string string
	for _, c := range fourfourtyeight_bit_cypher_text {
		fourfourtyeight_bit_cypher_string += fmt.Sprintf("%08b", c)
	}

	var temp big.Int
	//fourfourtyeight_bit_cypher, ok := temp.SetString(fourfourtyeight_bit_cypher_string, 10)
	//if !ok {
	//	fmt.Println("error in setting fourfourtyeight_bit_cypher to big.Int")
	//}

	P1, _ := temp.SetString(fourfourtyeight_bit_cypher_string[:32], 10)
	P2, _ := temp.SetString(fourfourtyeight_bit_cypher_string[32:64], 10)
	P3, _ := temp.SetString(fourfourtyeight_bit_cypher_string[64:96], 10)
	P4, _ := temp.SetString(fourfourtyeight_bit_cypher_string[96:128], 10)
	P5, _ := temp.SetString(fourfourtyeight_bit_cypher_string[128:160], 10)
	P6, _ := temp.SetString(fourfourtyeight_bit_cypher_string[160:192], 10)
	P7, _ := temp.SetString(fourfourtyeight_bit_cypher_string[192:224], 10)
	P8, _ := temp.SetString(fourfourtyeight_bit_cypher_string[224:256], 10)
	P9, _ := temp.SetString(fourfourtyeight_bit_cypher_string[256:288], 10)
	P10, _ := temp.SetString(fourfourtyeight_bit_cypher_string[288:320], 10)
	P11, _ := temp.SetString(fourfourtyeight_bit_cypher_string[320:352], 10)
	P12, _ := temp.SetString(fourfourtyeight_bit_cypher_string[352:384], 10)
	P13, _ := temp.SetString(fourfourtyeight_bit_cypher_string[384:416], 10)
	P14, _ := temp.SetString(fourfourtyeight_bit_cypher_string[416:], 10)

	L_64_bit_cypher.Xor(L_64_bit_cypher, P1)
	R_64_bit_cypher.Xor(R_64_bit_cypher, F(L_64_bit_cypher.String()))
	swap(L_64_bit_cypher, R_64_bit_cypher)
	L_64_bit_cypher.Xor(L_64_bit_cypher, P2)
	R_64_bit_cypher.Xor(R_64_bit_cypher, F(L_64_bit_cypher.String()))
	swap(L_64_bit_cypher, R_64_bit_cypher)
	L_64_bit_cypher.Xor(L_64_bit_cypher, P3)
	R_64_bit_cypher.Xor(R_64_bit_cypher, F(L_64_bit_cypher.String()))
	swap(L_64_bit_cypher, R_64_bit_cypher)
	L_64_bit_cypher.Xor(L_64_bit_cypher, P4)
	R_64_bit_cypher.Xor(R_64_bit_cypher, F(L_64_bit_cypher.String()))
	swap(L_64_bit_cypher, R_64_bit_cypher)
	L_64_bit_cypher.Xor(L_64_bit_cypher, P5)
	R_64_bit_cypher.Xor(R_64_bit_cypher, F(L_64_bit_cypher.String()))
	swap(L_64_bit_cypher, R_64_bit_cypher)
	L_64_bit_cypher.Xor(L_64_bit_cypher, P6)
	R_64_bit_cypher.Xor(R_64_bit_cypher, F(L_64_bit_cypher.String()))
	swap(L_64_bit_cypher, R_64_bit_cypher)
	L_64_bit_cypher.Xor(L_64_bit_cypher, P7)
	R_64_bit_cypher.Xor(R_64_bit_cypher, F(L_64_bit_cypher.String()))
	swap(L_64_bit_cypher, R_64_bit_cypher)
	L_64_bit_cypher.Xor(L_64_bit_cypher, P8)
	R_64_bit_cypher.Xor(R_64_bit_cypher, F(L_64_bit_cypher.String()))
	swap(L_64_bit_cypher, R_64_bit_cypher)
	L_64_bit_cypher.Xor(L_64_bit_cypher, P9)
	R_64_bit_cypher.Xor(R_64_bit_cypher, F(L_64_bit_cypher.String()))
	swap(L_64_bit_cypher, R_64_bit_cypher)
	L_64_bit_cypher.Xor(L_64_bit_cypher, P10)
	R_64_bit_cypher.Xor(R_64_bit_cypher, F(L_64_bit_cypher.String()))
	swap(L_64_bit_cypher, R_64_bit_cypher)
	L_64_bit_cypher.Xor(L_64_bit_cypher, P11)
	R_64_bit_cypher.Xor(R_64_bit_cypher, F(L_64_bit_cypher.String()))
	swap(L_64_bit_cypher, R_64_bit_cypher)
	L_64_bit_cypher.Xor(L_64_bit_cypher, P12)
	R_64_bit_cypher.Xor(R_64_bit_cypher, F(L_64_bit_cypher.String()))
	swap(L_64_bit_cypher, R_64_bit_cypher)
	L_64_bit_cypher.Xor(L_64_bit_cypher, P13)
	R_64_bit_cypher.Xor(R_64_bit_cypher, F(L_64_bit_cypher.String()))
	swap(L_64_bit_cypher, R_64_bit_cypher)
	L_64_bit_cypher.Xor(L_64_bit_cypher, P14)
	R_64_bit_cypher.Xor(R_64_bit_cypher, F(L_64_bit_cypher.String()))

	crypt := L_64_bit_cypher.String() + R_64_bit_cypher.String() + "0" + "1"
	var text string
	for i := 0; i < len(crypt); i += 8 {
		b, _ := strconv.ParseInt(crypt[i:i+8], 2, 64)
		text += string(b)
	}
	println(text)
}

func F(bits_stream string) *big.Int {
	S1 := bits_stream[:8]
	S2 := bits_stream[8:16]
	S3 := bits_stream[16:24]
	S4 := bits_stream[24:]

	var temp big.Int
	S1_big, _ := temp.SetString(S1, 10)
	S2_big, _ := temp.SetString(S2, 10)
	S3_big, _ := temp.SetString(S3, 10)
	S4_big, _ := temp.SetString(S4, 10)

	S1_big.Add(S1_big, S2_big)
	S1_big.Xor(S1_big, S3_big)
	S1_big.Add(S1_big, S4_big)

	return S1_big
}

func swap(L, R *big.Int) {
	L, R = R, L
}
