package crypto

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func CryptoPassWithSalt(password string, salt ...interface{}) (string, error) {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		password = fmt.Sprintf(password+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(password))), nil
}
