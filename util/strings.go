package util

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	accents    = []rune("ÀÁÂÃÈÉÊÌÍÒÓÔÕÙÚÝàáâãèéêìíòóôõùúýĂăĐđĨĩŨũƠơƯưẠạẢảẤấẦầẨẩẪẫẬậẮắẰằẲẳẴẵẶặẸẹẺẻẼẽẾếỀềỂểỄễỆệỈỉỊịỌọỎỏỐốỒồỔổỖỗỘộỚớỜờỞởỠỡỢợỤụỦủỨứỪừỬửỮữỰựỲỳỴỵỶỷỸỹ")
	nonAccents = []rune("AAAAEEEIIOOOOUUYaaaaeeeiioooouuyAaDdIiUuOoUuAaAaAaAaAaAaAaAaAaAaAaAaEeEeEeEeEeEeEeEeIiIiOoOoOoOoOoOoOoOoOoOoOoOoUuUuUuUuUuUuUuYyYyYyYy")
)

func Find(s []string, x string) int {
	for i, n := range s {
		if x == n {
			return i
		}
	}
	return -1
}

func FindRune(s []rune, x rune) int {
	for i, n := range s {
		if x == n {
			return i
		}
	}
	return -1
}

func Contains(s []string, x string) bool {
	for _, n := range s {
		if x == n {
			return true
		}
	}
	return false
}

func Include(s []string, x []string) bool {
	strs := Same(s, x)
	return len(strs) == len(x)
}

func Remove(s []string, x string) (ret []string) {
	ret = s
	idx := Find(s, x)
	if idx < 0 {
		return
	}
	for ok := true; ok; ok = (idx > 0) {
		ret = append(ret[:idx], ret[idx+1:]...)
		idx = Find(ret, x)
	}
	return
}

// Set Difference: A - B
func Difference(a, b []string) (diff []string) {
	m := make(map[string]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}

// Set Same: A | B
func Same(a, b []string) (union []string) {
	m := make(map[string]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; ok {
			union = append(union, item)
		}
	}
	return
}

func Compare(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	s := Same(a, b)
	return len(s) == len(a)
}

func Identical(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, item := range a {
		if item != b[i] {
			return false
		}
	}
	return true
}

func Unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func SplitAndTrim(str string, sep string) (array []string) {
	strs := strings.Split(str, sep)
	for i := range strs {
		s := strings.TrimSpace(strs[i])
		if len(s) > 0 {
			array = append(array, strings.TrimSpace(strs[i]))
		}
	}
	return
}

func LowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

func getAccentCharacter(r rune) string {
	index := FindRune(accents, r)
	if index >= 0 {
		return string(nonAccents[index : index+1])
	}
	return string(r)
}

func GetNonAccent(s string) string {
	var arr []string
	for _, r := range s {
		arr = append(arr, getAccentCharacter(r))
	}
	return strings.Join(arr, "")
}

func GetSqlLikeString(s string) string {
	txt := strings.ReplaceAll(s, `_`, `\_`)
	txt = "%" + strings.ToLower(txt) + "%"
	return txt
}

func ToUppers(ss []string) (arr []string) {
	for _, s := range ss {
		arr = append(arr, strings.ToUpper(s))
	}
	return
}

func ToLowers(ss []string) (arr []string) {
	for _, s := range ss {
		arr = append(arr, strings.ToLower(s))
	}
	return
}

func ArrayToString(strArray []string, split string) (str string) {
	str = strings.Join(strArray, split)
	return
}
