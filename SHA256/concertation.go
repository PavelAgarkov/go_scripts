package SHA256

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"os"
)

func Convertation() ([]byte, error) {

	var Sha384 = flag.Bool("Sha384", false, "Вызов функции sha384")
	var Sha512 = flag.Bool("Sha512", false, "Вызов функции sha512")

	flag.Parse()

	var hash256 [32]byte
	var hash384 [48]byte
	var hash512 [64]byte

	var hash []byte

	args := os.Args[1:]

	if len(args) == 1 {
		hash256 = SHA256(args[0])
		// [:] позволяет сделать срез []byte из среза любой размерности например [48]byte равный по значениям и длине
		hash := hash256[:]
		return hash, nil

	} else {
		if len(args) == 2 {

			if *Sha384 {
				hash384 = SHA384(args[1])
				hash := hash384[:]
				return hash, nil
			}

			if *Sha512 {
				hash512 = SHA512(args[1])
				hash := hash512[:]
				return hash, nil
			}
		}
	}

	return hash, nil
}

func SHA256(arg string) [32]byte {
	return sha256.Sum256([]byte(arg))
}

func SHA384(arg string) [48]byte {
	return sha512.Sum384([]byte(arg))
}

func SHA512(arg string) [64]byte {
	return sha512.Sum512([]byte(arg))
}
