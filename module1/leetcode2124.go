package module1

// difference bwn  'a'  and "a"  and comparision bwn s[i] and 'a'

func CheckString(s string) bool {
	var (
		bindex int
		aindex int
	)
	aflag := true
	bflag := true
	for i, j := 0, len(s)-1; i < len(s)-1 && j >= 0; i, j = i+1, j-1 {

		if s[i] == 'b' && bflag {
			bindex = i
			bflag = false
		}
		if s[j] == 'a' && aflag {
			aindex = j
			aflag = false
		}
	}

	if aflag || bflag {
		return true
	}
	if aindex < bindex {
		return true
	}
	return false

}

func CheckString2(s string) bool {
	for i := 0; i <= len(s)-2; i++ {
		if s[i] == 'b' && s[i+1] == 'a' {
			return false
		}
	}
	return true
}
