package main

import (
	"fmt"
)

/*
３つの参照型についてまとめる

スライス
マップ
チャネル

いずれもmakeによって生成される

*/

func main() {

	// Declare slices of type int
	var DeclareSliceData []int
	DeclareSliceData = make([]int, 10)
	fmt.Println(DeclareSliceData)

	// Generate a slice of type int whose element is 10
	GenerateSliceData := make([]int, 10)
	fmt.Println(GenerateSliceData)

	//

}
