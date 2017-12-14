package chinese

import (
	"strings"
	"unicode/utf8"
)

type py struct {
	style   int
	segment bool
}

func (this *py) perStr(pinyinStrs string) string {
	switch this.style {
	case 1:
		ret := strings.Split(pinyinStrs, ",")
		return firstLetter(ret[0])
	case 2:
		for i := 0; i < len(initials); i++ {
			if strings.Index(pinyinStrs, initials[i]) == 0 {
				return initials[i]
			}
		}
		ret := strings.Split(pinyinStrs, ",")
		return firstLetter(ret[0])
	case 3:
		ret := strings.Split(pinyinStrs, ",")
		return normalStr(ret[0])
	case 4:
		ret := strings.Split(pinyinStrs, ",")
		return ret[0]
	}
	return ""
}

func (this *py) doConvert(strs string) []string {
	bytes := []byte(strs)
	pinyinArr := make([]string, 0)
	nohans := ""
	var tempStr string
	var single string
	for len(bytes) > 0 {
		r, w := utf8.DecodeRune(bytes)
		bytes = bytes[w:]
		single = ziArr[int(r)]
		tempStr = string(r)
		if len(single) == 0 {
			nohans += tempStr
		} else {
			if len(nohans) > 0 {
				pinyinArr = append(pinyinArr, nohans)
				nohans = ""
			}
			pinyinArr = append(pinyinArr, this.perStr(single))
		}
	}
	if len(nohans) > 0 {
		pinyinArr = append(pinyinArr, nohans)
	}
	return pinyinArr
}

func (this *py) convert(strs string) []string {
	retArr := make([]string, 0)
	if this.segment {
		jiebaed := jieba.Cut(strs, true)
		for _, item := range jiebaed {
			mapValuesStr, exist := dictArr[item]
			mapValuesArr := strings.Split(mapValuesStr, ",")
			if exist {
				for _, v := range mapValuesArr {
					retArr = append(retArr, this.perStr(v))
				}
			} else {
				converted := this.doConvert(item)
				for _, v := range converted {
					retArr = append(retArr, v)
				}
			}
		}
	} else {
		retArr = this.doConvert(strs)
	}
	return retArr
}
