package main

import "fmt"


func sort (data []int)[]int{
	for i:=1;i<len(data);i++{
		key:=data[i]
		j:=i-1
		for j>=0 &&data[j]>key{
			data[j+1]=data[j]
			j--
		}
		data[j+1]=key
	}
	return data
}



func main(){
// data:=[]int{2,3,1,5,6,4,7,8,6,3,1,5,8,9,5}
// value:=sort(data)
// fmt.Println(value)

fmt.Println(sort([]int{1,4,3,6,7,8,9,4,3,5,6,4,4,3,5,8,9,0,2}))
}