package session

import (
	//"fmt"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"io"
)

//go -> crypto
//aes
//cipher
//hmac

func GetMD5(baseString string) string {
	//md5
	h := md5.New()
	io.WriteString(h, baseString)
	ret := fmt.Sprintf("%x", h.Sum(nil))

	fmt.Printf("md5: %s\n", ret)

	return ret
}

func GetSHA1(baseString string) string {
	//sha1
	h := sha1.New()
	io.WriteString(h, baseString)
	ret := fmt.Sprintf("%x", h.Sum(nil))

	fmt.Printf("sha1: %s\n", ret)

	return ret
}

func GetSHA256(baseString string) string {
	//sha256
	h := sha256.New()
	io.WriteString(h, baseString)
	ret := fmt.Sprintf("%x", h.Sum(nil))

	fmt.Printf("sha256: %s\n", ret)

	return ret
}

func GetMD5Plus(baseString string, strPlus string) string {
	//ユーザが入力したパスワードに対してMD5で一度暗号化
	//得られたMD5の値の前後に管理者自身だけが知っているランダムな文字列を追加
	//再度MD5で暗号化
	h := md5.New()
	io.WriteString(h, baseString)
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	salt1 := "@#$%"
	salt2 := "^&*()"

	//salt1 + username + salt2+MD5を連結。
	io.WriteString(h, salt1)
	io.WriteString(h, strPlus)
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	ret := fmt.Sprintf("%x", h.Sum(nil))

	fmt.Printf("md5 Plus: %s\n", ret)

	return ret
}

func GetScrypt(baseString string) string {
	salt := "@#$%7G8r"
	//func Key(password, salt []byte, N, r, p, keyLen int) ([]byte, error) {
	dk, _ := scrypt.Key([]byte(baseString), []byte(salt), 16384, 8, 1, 32)

	//FIXME:文字化けする？？
	fmt.Printf("Scrypt: %s\n", dk)

	return string(dk)
}
