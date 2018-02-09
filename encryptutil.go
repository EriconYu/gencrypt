package gencrypt

import "github.com/EriconYu/gencrypt/encrypt"

func Encrypt(content string)  string{
	enContent , ok := encrypt.CreateSecData(content , 20)
	if ok == false {
		return ""
	}
	return enContent
}

func Decrypt(content string)  string{
	deContent , ok := encrypt.ParseSecData(content , 20)
	if ok == false {
		return ""
	}
	return deContent
}
