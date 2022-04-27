package encryption

// CheckPassword 验证密码是否正确
func CheckPassword(password, salt, inputPassword string) bool {
	if password == EncodeMD5(inputPassword+salt) {
		return true
	}
	return false
}