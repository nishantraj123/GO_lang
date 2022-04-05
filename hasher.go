package hasher

import(
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func Sha256Hasher(userString string) string{
	//generate hash
	hash_:=sha256.Sum256([]byte(userString))
	// encode to string (from byte)
	hashedString:=hex.EncodeToString(hash_[:])
	// make return
	return hashedString
}

func Md5Hasher(userString string) string{
	//generate hash
	hash_:=md5.Sum([]byte(userString))
	// encode to string (from byte)
	hashedString:=hex.EncodeToString(hash_[:])
	// make return
	return hashedString
}