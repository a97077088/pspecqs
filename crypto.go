package pspecqs

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/d1str0/pkcs7"
)
func EncodePass(pwd string,authcode string)string{
	p1:=smd5(pwd)
	p2:=smd5(fmt.Sprintf("%s%s",p1,authcode))
	return p2
}
func smd5(s string)string{
	m5:=md5.New()
	m5.Write([]byte(s))
	return hex.EncodeToString(m5.Sum(nil))
}
func EncodeUserName(username string)string{
	key:="ihaierForTodoKey"
	iv:="ihaierForTodo_Iv"
	cb,err:=aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	blk:=cipher.NewCBCEncrypter(cb, []byte(iv))
	src,err:=pkcs7.Pad([]byte(username),blk.BlockSize())
	dst:=make([]byte,len(src))
	blk.CryptBlocks(dst, src)
	return base64.StdEncoding.EncodeToString(dst)
}