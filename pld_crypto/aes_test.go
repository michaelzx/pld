package pld_crypto

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	x := []byte("世界上最邪恶最专制的现代奴隶制国家--朝鲜")
	key := []byte("hgfedcba87654321")
	b := AesBase64Kit.EncryptBase64(x, key)
	fmt.Println(b)

	x2, err := AesBase64Kit.DecryptBase64(b, []byte("hgfedcba87654321"))
	if err != nil {
		fmt.Println("解码失败")
	}
	fmt.Println(string(x2))
}
