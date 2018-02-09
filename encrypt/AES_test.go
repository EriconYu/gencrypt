package encrypt

import (
	"encoding/base64"
	"testing"
)

//TestEncrypt ...
func TestEncrypt(t *testing.T) {
	var a AesEncrypt
	//a.PrivateCode = RadomString(20)
	a.PrivateCode = "1234567890123456"
	aesB, err := a.Encrypt("helloworld")
	if err != nil {
		t.Error("挂了！！！", err)
	} else {
		aes := base64.StdEncoding.EncodeToString(aesB)
		t.Log("aesEnc is", aes)
	}
}

//TestDecrypt ...
func TestDectypt(t *testing.T) {
	var a AesEncrypt
	a.PrivateCode = "1234567890123456"
	aesB, err := a.Encrypt("helloworld")
	if err != nil {
		t.Error("1挂了！！！", err)
	} else {
		aes := base64.StdEncoding.EncodeToString(aesB)
		t.Log("aesEnc is", aes)

		aesB, err = base64.StdEncoding.DecodeString(aes)
		if err != nil {
			t.Error("2挂了！！！", err)
			return
		}
		decsrc, err1 := a.Decrypt(aesB)
		if err1 != nil {
			t.Error("3挂了！！！", err1)
			return
		}
		t.Log("decsrc is ", decsrc)

	}
}
