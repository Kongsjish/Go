package main

import "fmt"

type Book struct{
	title string
	author string
	bookid  int
}

func main(){
	var n [10]int
	var i int

	//结构体
	fmt.Println("结构体：")

	var book Book
	book.title="Go语言"
	book.author="Dont"
	book.bookid=1234

	fmt.Printf("title of book:%s\n",book.title)
	fmt.Printf("author of book:%s\n",book.author)
	fmt.Printf("bookid of book:%d\n",book.bookid)
	//Map集合
	//创建集合
	var kkmap map[string]string

	fmt.Println("集合：")
	kkmap=make(map[string]string)
	kkmap["js"]="JS"
	kkmap["css"]="Css"
	kkmap["html"]="Html"
    for kp:=range kkmap{
    	fmt.Println(kp,"是：",kkmap[kp])
	}

	//切片
    fmt.Println("切片：")
	var numbers = make([]int,3,5)
	PrintSlice(numbers)
	fmt.Println("数组：")
    //数组
	for i=0;i<10;i++{
		n[i]=i+1
	}
	for i=0;i<10;i++{
		fmt.Printf("Element[%d]=%d\n",i,n[i])
	}
	fmt.Println("HelloWorld")
}
func PrintSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
