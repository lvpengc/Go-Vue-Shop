package hash

/* 对字符串取hash值 */

// Md5String 获取字符串md5值
func Md5String(s string, salt string) string {
	return Md5Byte([]byte(s), []byte(salt))
}

// Sha1String 获取字符串sha1值
func Sha1String(s string, salt string) string {
	return Sha1Byte([]byte(s), []byte(salt))
}

// Sha256String 获取字符串sha256值
func Sha256String(s string, salt string) string {
	return Sha256Byte([]byte(s), []byte(salt))
}

// Sha512String 获取字符串sha512值
func Sha512String(s string, salt string) string {
	return Sha512Byte([]byte(s), []byte(salt))
}
