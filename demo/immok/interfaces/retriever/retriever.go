package retriever

import (
	"demo/immok/interfaces/retriever/mock"
	"demo/immok/interfaces/retriever/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

// PostRetriever 组合接口
type PostRetriever interface {
	Retriever
	Poster
}

const url = "https://www.imooc.com"

// 接口的使用着
func downloader(r Retriever) string {
	return r.Get(url)
}

func session(s PostRetriever) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func post(post Poster) {
	post.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

func UseMain() {
	var r Retriever
	retriever := &mock.Retriever{
		Contents: "this is a fake immoc.com",
	}
	r = retriever
	inspect(r)

	// Type assertion	断言
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever)
	} else {
		fmt.Println("not a mockRetriever!")
	}

	fmt.Println()

	r = &real.Retriever{ // 如果 real.Retriever 结构体特别大值的拷贝非常耗内存可以切换为指针调用
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Second,
	}
	inspect(r)

	// Type assertion	断言
	if realRetriever, ok := r.(*real.Retriever); ok {
		// 获取真实的 retriever 与 switch 判断类型相同
		fmt.Println(realRetriever)
	} else {
		fmt.Println("not a realRetriever!")
	}

	fmt.Println("Try a session")

	fmt.Println(session(retriever))

}

func inspect(r Retriever) {
	fmt.Printf("%T , %v\n", r, r) // 看 r 是什么类型、值是什么
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("real retriever:", v.UserAgent)
	}
}
