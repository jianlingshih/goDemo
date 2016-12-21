package main

import ()

//前面小节介绍了如何存储密码 但是有时候 我们想把一些敏感数据加密后存储起来 在将来的某个时候 随需将它们解密出来
//此时我们应该在选用对称加密算法来满足我们的需求

//base64加解密
//如果web应用足够简单 数据的安全性没有那么严格的要求 那么可以采用一种比较简单的加解密方法是base64 这种方式实现起来比较简单
//Go语言的base64包已经很好的支持这个了  请看下面的例子
package main

import (
	"encoding/base64"
	"fmt"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))

}
func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}
func main() {
	//encode
	hello := "你好 世界!Hello world"
	debyte := base64Encode([]byte(hello))
	fmt.Println(debyte)
	//decode
	enbyte, err := base64Decode(debyte)
	if err != nil {
		fmt.Println(err.Error())
	}
	if hello != string(enbyte) {
		fmt.Println("hello is not equal to enbyte")
	}
	fmt.Println(string(enbyte))

}
//高级加解密
//Go语言的crypto里面支持对称加密的高级加解密包有
//crypto/aes 包 AES  是美国联邦政府采用的一种区块加密标准
//crypto/des 包 DEA (Data Encryption Algorithm) 是一种对称加密算法 是目前使用最广泛的秘钥系统 特别是在保护金融数据的安全中

//因为这两种算法使用方法类似 所以在此 我们仅用aes包 为例 来讲解它们的使用 请看下面的例子

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func main() {
	//需要去加密的字符串
	plaintext := []byte("My name is QQQ")
	//如果传入加密串的话，plaint就是传入的字符串
	if len(os.Args) > 1 {
		plaintext = []byte(os.Args[1])
	}

	//aes的加密字符串
	key_text := "gqyqwetyhguhqqqq"
	if len(os.Args) > 2 {
		key_text = os.Args[2]
	}

	fmt.Println(len(key_text))

	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}

	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s=>%x\n", plaintext, ciphertext)

	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)
}

//上面通过调用函数 aes.NewCipher(参数key必须是 16 24 或者32位的[]byte 分别对应 AES-128 AES-192 AES-256算法)
//返回了一个cipher.Block接口 这个接口实现了三个功能
type Block interface{
 // BlockSize returns the cipher's block size.
    BlockSize() int

    // Encrypt encrypts the first block in src into dst.
    // Dst and src may point at the same memory.
    Encrypt(dst, src []byte)

    // Decrypt decrypts the first block in src into dst.
    // Dst and src may point at the same memory.
    Decrypt(dst, src []byte)

}
//这三个函数实现了 加解密操作 详细的操作请看上面的例子

//总结
//这小节介绍了几种加解密的算法 在开发web应用的时候 可以根据需求采用不同的方式进行加解密 一般的应用可以采用base64算法
//需要高级的话可以采用aes或者des算法












