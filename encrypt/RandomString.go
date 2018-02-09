package encrypt

import (
	"math/rand"
	"time"
)

//RadomString 生成指定长度的随机字符串，只包含26个小写字母
func RadomString(n int) (radom string) { //65-90 97-122
	for i := 0; i < n; i++ {
		base := time.Now().UnixNano()
		radom = string(append([]byte(radom), byte(base%26)+'a'))
	}
	return radom
}

func GetRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
