package util

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func Sha1(content []byte) string {
	hash := sha1.New()
	hash.Write(content)
	slice := hash.Sum(nil)
	return fmt.Sprintf("%x", slice)
}

func Md5(content []byte) string {
	hash := md5.New()
	hash.Write(content)
	slice := hash.Sum(nil)
	return fmt.Sprintf("%x", slice)
}

func Equal(b1 []byte, b2 []byte) bool {
	return Sha1(b1) == Sha1(b2)
}
