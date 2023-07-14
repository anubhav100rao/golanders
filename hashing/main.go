package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// base mod str
func calculateHash(base int64, mod int64, str string) int64 {

	var hash int64 = 0
	var f int64 = 1
	for i := 0; i < len(str); i++ {
		pos := int64(str[i] - 'A' + 1)
		hash = (hash + pos*f) % mod
		f = f * base % mod
	}
	return hash
}

func cryptoHash(str string) string {
	md5Hash := md5.New()
	sha256Hash := sha256.New()
	sha512Hash := sha512.New()

	md5Hash.Write([]byte("JetBrains Academy"))
	sha256Hash.Write([]byte("JetBrains Academy"))
	sha512Hash.Write([]byte("JetBrains Academy"))

	fmt.Printf("%x\n", md5Hash.Sum(nil))
	fmt.Printf("%x\n", sha256Hash.Sum(nil))
	fmt.Printf("%x\n", sha512Hash.Sum(nil))

	return string(sha512Hash.Sum(nil))
}

func bcryptHash(str string) string {
	password := "TrustNo1"
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

func main() {
	base := int64(3)
	mod := int64(13)

	// getting collisions after mod changes
	println(calculateHash(base, mod, "ABCD"))
	println(calculateHash(base, mod, "ABCQ"))

	println(calculateHash(base, mod, "CCDB"))
	println(calculateHash(base, mod, "ACCD"))

	cryptoHash("JetBrains Academy")

}
