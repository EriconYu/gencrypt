package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

//MD5 ...
type MD5 struct {
}

//CreateMD52Hex ....
func (m *MD5) CreateMD52Hex(param string) (md5str string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(param))
	cipherStr := md5Ctx.Sum(nil)
	md5str = hex.EncodeToString(cipherStr)
	return md5str
}
