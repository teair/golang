package regexp

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func IsIP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

	} else {

		curtime := time.Now().Unix()

		h := md5.New()

		// fmt.Fprintln(w, strconv.FormatInt(curtime, 10))

		io.WriteString(h, strconv.FormatInt(curtime, 10))

		token := fmt.Sprintf("%x", h.Sum(nil))

		t, err := template.ParseFiles("regexp/view/regexp.html")

		if err != nil {
			log.Fatalln("failed to parse template")
		}

		t.Execute(w, token)
	}
}

func RegexMain() {

	http.HandleFunc("/regexp", IsIP)

	// http.HandleFunc("/test", Test)

	err := http.ListenAndServe("127.0.0.1:9091", nil)

	if err != nil {
		fmt.Println("Listen and Server failed!", err)
	}
}

func Spider() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("failed to get!")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("failed to readAll!")
		return
	}

	src := string(body)

	// 将html标签全转换成小写
	re, _ := regexp.Compile(`<[\S\s]+?>`)
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	// 去除 style 标签
	re, _ = regexp.Compile(`<style[\S\s]+?</style>`)
	src = re.ReplaceAllString(src, "")
	// 去除HTMLUnscape的style
	re, _ = regexp.Compile(`&lt;style&gt;[\S\s]+?&lt;/style&gt;`)
	src = re.ReplaceAllString(src, "")

	// 去除script标签
	re, _ = regexp.Compile(`<script[\S\s]+?</script>`)
	src = re.ReplaceAllString(src, "")
	re, _ = regexp.Compile(`&lt;script&gt;[\S\s]+?&lt;script&gt;`)
	src = re.ReplaceAllString(src, "")

	// 去除所有尖括号内的html代码并换成换行符
	re, _ = regexp.Compile(`<[\S\s]+?>`)
	src = re.ReplaceAllString(src, "\n")

	// 去除连续的换行符
	re, _ = regexp.Compile(`\s{2,}`)
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))
}

func Test() {
	// 获取最长不重复子串
	s := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"

	// freq := make([]int, 128)
	// var res = 0
	// start, end := 0, -1

	// for start < len(s) {
	// 	if end+1 < len(s) && freq[s[end+1]] == 0 {
	// 		end++
	// 		freq[s[end]]++
	// 	} else {
	// 		freq[s[start]]--
	// 		start++
	// 	}
	// 	res = max(res, end-start+1)
	// }
	// fmt.Println(res, freq)

	// lastOccurred := make(map[byte]int)
	// start := 0
	// maxLength := 0
	// for i, ch := range []byte(s) {
	// 	if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
	// 		start = lastI + 1
	// 	}
	// 	if i-start+1 > maxLength {
	// 		maxLength = i - start + 1
	// 	}
	// 	lastOccurred[ch] = i
	// }
	// fmt.Fprintln(w, maxLength)

	// 查找最长不连续子串
	// tmp := ""
	// max := 0
	// for i := 0; i < len(s); i++ {
	// 	ind := strings.Index(tmp, string(s[i]))
	// 	if ind != -1 {
	// 		// 有重复
	// 		tmp += string(s[i])
	// 		tmp = tmp[ind+1:]
	// 	} else {
	// 		// 无重复
	// 		tmp += string(s[i])
	// 	}

	// 	if len(tmp) > max {
	// 		max = len(tmp)
	// 	}
	// }
	// fmt.Fprintln(w, max)

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			// 当前视窗口无重复则判断窗口左侧是否超过右侧
			if i+1 > end {
				end = i + 1
			}
		} else {
			// 有重复则移动窗口
			start += index + 1
			end += index + 1
		}
	}
	fmt.Println(end - start)

}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
