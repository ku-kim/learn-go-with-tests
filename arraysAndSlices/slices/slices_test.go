package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSlicesAssignment(t *testing.T) {
	var s1 []byte // slice 할당 (기본값, cap 설정 X)

	// 빈 Slice 할당 후 Slice 할당
	var s2 []byte
	s2 = make([]byte, 5, 5) // 크기 5, cap 5
	// 와 동일 s2 == []byte{0, 0, 0, 0, 0}

	// make + Simple 버전
	s3 := make([]byte, 5)
	//s3 := make([]byte, 5, 5)

	// make 사용하지 않고 기본값 있는 slice 할당
	s4 := []int{1, 2, 3} // slice 할당 (기본값, cap 3 설정)

	b := []string{"g", "o", "l", "a", "n", "g"}
	sliceB := b[:2] // slice 로 넣으면 포인터로 접근하기에 b와 SliceB 연결되어있음
	sliceB[0] = "AAA"

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(b)
	fmt.Println(sliceB)

}

func TestSlicesBasic(t *testing.T) {
	mySlice := []int{1, 2, 3, 4, 5}

	fmt.Println(reflect.DeepEqual([]int{2, 3}, (mySlice[1:3])))
	fmt.Println(reflect.DeepEqual([]int{1, 2, 3}, (mySlice[:3])))
	fmt.Println(reflect.DeepEqual([]int{1, 2, 3, 4, 5}, (mySlice[:])))
}

func TestSlicesBasic2(t *testing.T) {
	x := [3]string{"A", "B", "C"}

	// poitn 복사이기에 x,y 똑같이 바라봄
	y := x[:] // slice "y" points to the underlying array "x"

	// 아예 새롭게 복사
	z := make([]string, len(x))
	copy(z, x[:]) // slice "z" is a copy of the slice created from array "x"

	y[1] = "BBBBB" // the value at index 1 is now "Belka" for both "y" and "x"

	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)
	/*
		[3]string [A BBBBB C]
		[]string [A BBBBB C]
		[]string [A B C]
	*/
}

func TestSlicesAppend(t *testing.T) {
	var mySlice []int

	mySlice = append(mySlice, 3)

	fmt.Println(mySlice)

}
