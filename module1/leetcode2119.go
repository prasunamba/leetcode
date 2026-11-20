package module1

func IsSameAfterReversals(num int) bool {
	if num%10 == 0 && num != 0 {
		return false
	}
	return true
}
