package chinese

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/yanyiwu/gojieba"
)

var (
	dictData = path.Join("./chinese", "dict.dat")
	ziData   = path.Join("./chinese", "zi.dat")
	dictArr  map[string]string
	ziArr    []string
	reg      *regexp.Regexp
	jieba    *gojieba.Jieba
)

var initials []string = strings.Split("b,p,m,f,d,t,n,l,g,k,h,j,q,x,r,zh,ch,sh,z,c,s", ",")
var sympolMap = map[string]string{
	"ā": "a1",
	"á": "a2",
	"ǎ": "a3",
	"à": "a4",
	"ē": "e1",
	"é": "e2",
	"ě": "e3",
	"è": "e4",
	"ō": "o1",
	"ó": "o2",
	"ǒ": "o3",
	"ò": "o4",
	"ī": "i1",
	"í": "i2",
	"ǐ": "i3",
	"ì": "i4",
	"ū": "u1",
	"ú": "u2",
	"ǔ": "u3",
	"ù": "u4",
	"ü": "v0",
	"ǘ": "v2",
	"ǚ": "v3",
	"ǜ": "v4",
	"ń": "n2",
	"ň": "n3",
	"": "m2",
}

func init() {
	reg = regexp.MustCompile("([" + getMapKeys() + "])")
	jieba = gojieba.NewJieba()
	initPhrases()
	initZi()
}

func getMapKeys() string {
	keys := ""
	for key, _ := range sympolMap {
		keys += key
	}
	return keys
}

func normalStr(str string) string {
	findRet := reg.FindString(str)
	if findRet == "" {
		return str
	}
	return strings.Replace(str, findRet, string([]byte(sympolMap[findRet])[0]), -1)
}

func firstLetter(str string) string {
	c := string(str[0])
	if str[0] >= 195 && str[0] <= 238 {
		c = str[:2]
		c2 := sympolMap[c]
		if c2 != "" {
			c = string(c2[0])
		}
	}
	return c
}

func initPhrases() {
	f, err := os.Open(dictData)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&dictArr); err != nil {
		log.Fatal(err)
	}
}

func initZi() {
	f, err := os.Open(ziData)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	ziArr = make([]string, 173466)
	rows := 0
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if rows == 173466 {
			break
		}
		ziArr[rows] = string(line)
		rows++
	}
}
