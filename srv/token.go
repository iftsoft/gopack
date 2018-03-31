package srv

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
)


type TokenCoder struct {
	secret	[]byte
}

func (this *TokenCoder)InitCoder(tokenKey string) {
	if len(tokenKey) >= 32 {
		tokenKey = tokenKey[:32]
	} else 	if len(tokenKey) < 32 && len(tokenKey) >= 24 {
		tokenKey = tokenKey[:24]
	} else if len(tokenKey) < 24 && len(tokenKey) >= 16 {
		tokenKey = tokenKey[:16]
	} else if len(tokenKey) < 16 {
		tokenKey = "0123456789ABCDEFFEDCBA9876543210"
	}
	this.secret	= []byte(tokenKey)
}

func (this *TokenCoder)EncodeToken(token interface{}) (dump string, err error) {
	js, err := json.Marshal(&token)
	if err != nil {
		srvLog.Error("Token marshal error: ", err.Error())
		return dump, err
	}
	c, err := aes.NewCipher(this.secret)
	if err != nil {
		srvLog.Error("Encrypt token error: %s.", err.Error())
		return dump, err
	}
	iv := this.secret[:aes.BlockSize]
	encrypted := make([]byte, len(js))
	encrypter := cipher.NewCFBEncrypter(c, iv)
	encrypter.XORKeyStream(encrypted, js)

	dump = base64.StdEncoding.EncodeToString(encrypted)
	srvLog.Trace("Encrypt %s to %s.", string(js), dump)

	return dump, nil
}

func (this *TokenCoder)DecodeToken(key string, token interface{}) (err error) {
	res, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		srvLog.Error("Decode token error: %s.", err.Error())
		return err
	}
	c, err := aes.NewCipher(this.secret)
	if err != nil {
		srvLog.Error("Decrypt token error: %s.", err.Error())
		return err
	}
	iv := this.secret[:aes.BlockSize]
	decrypter := cipher.NewCFBDecrypter(c, iv)
	decrypted := make([]byte, len(res))
	decrypter.XORKeyStream(decrypted, res)
	srvLog.Trace("Decrypt %s to %s.", key, string(decrypted))

	if err = json.Unmarshal(decrypted, token); err != nil {
		srvLog.Error("Token unmarshal error: ", err.Error())
		return err
	}
	return nil
}


