package gencrypt

import (
	"testing"
)

func TestEncrypt(t *testing.T)  {
	t.Log("encrypt is " , Encrypt("helloworld~"))
	t.Log("encrypt is " , Encrypt(""))
}

func TestDecrypt(t *testing.T){
	t.Log("encrypt is " , Decrypt(Encrypt("helloworld~")))
	t.Log("encrypt is " , Decrypt(Encrypt("")))
}
