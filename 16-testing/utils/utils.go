package utils

func IsPrime(no int32) bool {
	if no <= 0 {
		return false
	}
	if no <= 3 {
		return true
	}
	for i := int32(2); i <= no-1; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
