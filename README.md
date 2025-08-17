## Blowfish Encryption in Go (From Scratch)

* Implements core Blowfish encryption and decryption without external libraries.
* Accepts 8-character (64-bit) plaintext input.
* Converts input to binary; splits into two 32-bit halves (L and R).
* Uses a hardcoded 448-bit key string as the P-array substitute.
* Performs 14 Feistel rounds using custom function `F(L)`.
* Feistel function splits input into four 8-bit blocks and applies: `(S1 + S2) XOR S3 + S4`.
* XOR operations and swaps mimic Blowfish round logic.
* Final L and R are combined and converted back to characters.
* Decryption follows reverse order of round keys (P14 → P1).
* No use of actual S-boxes or full key schedule.

[Description of a New Variable-Length Key, 64-Bit Block Cipher (Blowfish)](https://www.schneier.com/academic/archives/1994/09/description_of_a_new.html)
