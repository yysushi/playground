package fernet_test

import (
	"encoding/base64"
	"fmt"
	"testing"
	"time"

	"github.com/fernet/fernet-go"
	"github.com/stretchr/testify/assert"
)

func TestFernetGenerate(t *testing.T) {
	var k fernet.Key
	k.Generate()
	kb := [32]byte(k)
	fmt.Println(k.Encode())
	assert.Equal(t, k.Encode(), base64.URLEncoding.EncodeToString(kb[:]))
	s, _ := fernet.DecodeKey(k.Encode())
	fmt.Println(*s)
	fmt.Println(fmt.Sprintf("%s", *s))
	assert.Equal(t, fmt.Sprintf("%s", *s), fmt.Sprintf("%s", kb))
}

func TestFernetEncryption(t *testing.T) {
	var k fernet.Key
	k.Generate()
	tok, err := fernet.EncryptAndSign([]byte("hello"), &k)
	if err != nil {
		t.Error(err)
	}
	msg := fernet.VerifyAndDecrypt(tok, 60*time.Second, []*fernet.Key{&k})
	assert.Equal(t, "hello", string(msg))
}
