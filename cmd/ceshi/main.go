package main

import (
	"fmt"
)

func handleCommentPush(hasConversation bool) {
	if hasConversation {

	}
}


type shuzu struct {
	a int
}
func ceshishuzu(ids []shuzu) {
	a := shuzu{1}
	ids = append(ids, a)
}

func ceshipointer(ids *[]shuzu) {
	a := shuzu{1}
	ids = append(ids, a)
}

func main() {
	//handleCommentPush(true)
	//a := 1 == 1
	//fmt.Println(a)
	//var s interface{} = a
	//
	//fmt.Println(s.(bool))
	ids := []shuzu{{5}, {6}}
	ceshipointer(&ids)
	fmt.Println(ids)
}
