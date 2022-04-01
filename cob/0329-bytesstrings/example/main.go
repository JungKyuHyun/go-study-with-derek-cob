package main

// 스펠링이 틀려서.. 여튼 새로운 방법들을 알게됨.. 이렇게 패키지 이름을 변경해서 사용할 수 있음
import bytestrings "github.com/jungkyuhyun/go-study-with-derek-cob/cob/0329-bytesstrings"


func main() {
	if err := bytestrings.WorkWithBuffer(); err != nil {
		panic(err)
	} 
	// 아래 내용을 모두 표준 출력(stdout)에 출력한다
	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}