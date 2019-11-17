package helper

import "strings"

func SubStrByLen(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

func SubStrByEnd(str string, start, end int) string {
	rs := []rune(str)
	rl := len(rs)

	if start < 0 {
		start = rl - 1 + start
	}

	if end < 0 {
		end = rl + end
	}

	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

//第一个字母大写
func FirstToUper(str string) string {
	s := []byte(str)
	return strings.ToUpper(string(s[0])) + string(s[1:])
}

//首字母小写
func FirstLower(str string) string {
	s := []byte(str)
	return strings.ToLower(string(s[0])) + string(s[1:])
}

//驼峰转下划线
func HumpToUnder(str string) string {
	var back string
	strArr := []byte(str)
	for i := 0; i < len(strArr); i++ {
		if strArr[i] >= 65 && strArr[i] <= 90 {
			if i > 0 {
				back += "_"
			}
			back += strings.ToLower(string(strArr[i]))
		} else {
			back += string(strArr[i])
		}
	}
	return back
}
