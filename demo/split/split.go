// split/split.go
package split

import (
	"strings"
	)

// split package with a single split function


// 将切片s分割成所有由sep分割的子字符串

// 返回这些分隔符之间的子字符串的一个片段.

func Split (s,sep string) (result []string) {

	result = make([]string,0,strings.Count(s,sep)+1)

	i := strings.Index(s,sep)

	for i > -1 {

		result = append(result,s[:i])

		s = s[i+len(sep):]

		i = strings.Index(s,sep)

	}

	result = append(result,s)

	return

}
