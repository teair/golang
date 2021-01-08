
// split/split_test.go
package split

import  (
	"reflect"
	"testing"
	//"time"
	//. "fmt"
)

func TestSplit(t *testing.T) {		// 测试函数必须以Test开头,后缀名首字母必须大写,必须接收一个*.testing.T类型参数
	
	// 定义一个测试用例类型
	type test struct {
		input string
		sep string
		want []string
	}

	// 定义一个存储测试用例的切片
	tests := map[string]test{
		"simple" : {input:"a:b:c",sep:":", want : []string{"a","b","c"}},
		//"wrong" : {input:"abcdefg",sep:"m",want:[]string{"a","b","c","d","e","f","g"}},
		//"more sep" : {input:"abkdjgjlkgdjpoi",sep:"dj",want:[]string{"abk","gjlkg","poi"}},
		//"chinese split" : {input:"抽刀断水水更流",sep:"断",want:[]string{"抽刀","水水更流"}},
	}

	for name,tc := range tests{

		t.Run(name,func(t *testing.T){

		got := Split(tc.input,tc.sep)

		if !reflect.DeepEqual(got,tc.want){

			t.Errorf("name: %s , expect: %#v , got:%#v",name,tc.want,got)
	
		}
		
		})


	}

}


func BenchmarkSplit(b *testing.B) {

	//time.Sleep(5 * time.Second)		// 假设需要做一些需要耗时的无关操作

	//b.ResetTimer()				// 重置计时器

	for i:=0;i<b.N;i++ {

		Split("枯藤老树昏鸦","老")

	}

}




