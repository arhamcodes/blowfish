package main

import "math/big"

func initializePArray() [18]uint32 {
	piHex := "243f6a8885a308d313198a2e03707344a4093822299f31d0082efa98ec4e6c89" +
		"452821e638d01377be5466cf34e90c6cc0ac29b7c97c50dd3f84d5b5b5470917" +
		"921db3140175f162faf17831e7a0b16730e56f5a43b8eae68013c4a2e2224562" +
		"b3b4c3fe3f10b5a0a868327946a3b0bf57f9089c2a8e8"
	piBigInt, _ := new(big.Int).SetString(piHex, 16)
	var PArray [18]uint32
	for i := 0; i < len(PArray); i++ {
		if piBigInt.Sign() == 0 {
			break
		}
		PArray[i] = uint32(piBigInt.Uint64())
		piBigInt.Rsh(piBigInt, 32)
	}
	return PArray
}

func initializeSBoxes(pArray [18]uint32) [4][256]uint32 {
	sboxes := [4][256]uint32{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 256; j++ {
			sboxes[i][j] = pArray[i*256+j]
		}
	}
	return sboxes
}
