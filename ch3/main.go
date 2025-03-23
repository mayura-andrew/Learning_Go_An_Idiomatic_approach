package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// var x = [12]int{1, 5: 4, 10: 100, 15}
	// var y = [...]int{1, 2, 3, 4, 5}

	// Slices

	// s := []int{1, 2, 3, 4, 5}
	// sub := s[1:3]

	// fmt.Println(sub)
	// // Modifying a sub-slice modifies the original slice:
	// sub[0] = 99
	// fmt.Println(sub)
	// fmt.Println(s)
	// // fmt.Println(len(y))

	// // var x = []int{1, 5: 4, 6, 10: 100, 15}
	// // fmt.Println(x)

	// // nums := []int{1, 2, 3, 4}
	// // nums = append(nums, 6)
	// // nums = append(nums, 7, 8, 9)
	// // fmt.Println(nums)

	// fmt.Println("Length and capacity of a slice")

	// a := []string{"first", "second", "third"}
	// fmt.Println(a, len(a), cap(a))

	// /// Empting a slice
	// clear(a)
	// fmt.Println(a, len(a), cap(a))
	// a = append(a, "fourth")
	// fmt.Println(a, len(a), cap(a))

	// //  empty slice literal

	// var x = []int{}
	// //  zero length and zero capacity
	// fmt.Println(x, len(x), cap(x))

	//  array := make([]int, 5)
	//  array = append(array, 1, 2, 3, 4, 5)
	//  fmt.Println(array, len(array), cap(array))

	// Slicing

	// y := []string{"a", "b", "c", "d", "e"}
	// // z := y[:2]
	// j := y[1:]
	// k := y[1:3]
	// e := y[:]

	// fmt.Println(j, len(j), cap(j))

	// fmt.Println(k, len(k), cap(k))

	// fmt.Println(e, len(e), cap(e))

	// x := []string{"a", "b", "c", "d"}
	// y := x[:2]
	// z := x[1:]
	// x[1] = "y"
	// y[0] = "x"
	// z[1] = "z"
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)
	// fmt.Println("x:", x)

	// x := []int{1, 2, 3, 4}
	// fmt.Println(x[:3])
	// fmt.Println(x[1:])
	// num := copy(x[:3], x[1:])
	// fmt.Println(x, num)

	//   converting array to slice

	xArray := [4]int{5, 6, 7, 8}
	xSlice := xArray[:]
	fmt.Println(xSlice)
	fmt.Println(reflect.TypeOf(xSlice))
	fmt.Println(reflect.TypeOf(xArray))
	fmt.Printf("type of %T\n", xSlice)

	//  converting slice to array
	fmt.Println("-----------------")
	ySlice := []int{1, 2, 3, 4}
	yArray := [4]int(ySlice)
	smallArray := [2]int(ySlice)
	ySlice[0] = 10
	fmt.Println(ySlice)
	fmt.Println(yArray)
	fmt.Println(smallArray)

	// Safer approach using copy()
	fmt.Println("-----------------")

	zSlice := []int{1, 2, 3, 4}
	var safeArray [5]int

	copy(safeArray[:], zSlice) // copy elements from zSlice to safeArray
	fmt.Println("safeArray:", safeArray)

	// converting slice to array using make()
	fmt.Println("-----------------")
	slice := make([]int, 5)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	slice[4] = 5
	array := [5]int{}
	copy(array[:], slice) // copy elements from slice to array
	fmt.Println("array:", array)

	//  converting array to slice using make()
	fmt.Println("-----------------")

	array2 := [5]int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(array2))
	copy(slice2, array2[:]) // copy elements from array to slice
	fmt.Println("slice2:", slice2)

	//  converting array to slice using append()
	fmt.Println("-----------------")

	array3 := [5]int{1, 2, 3, 4, 5}
	slice3 := append([]int{}, array3[:]...) // append elements from array to slice
	fmt.Println("slice3:", slice3)

	//  converting array to slice using for loop

	fmt.Println("-----------------")
	array4 := [5]int{1, 2, 3, 4, 5}
	slice4 := make([]int, len(array4))
	for i := 0; i < len(array4); i++ {
		slice4[i] = array4[i]
	}
	fmt.Println("slice4:", slice4)

	//  converting array to slice using range

	fmt.Println("-----------------")

	array5 := [5]int{1, 2, 3, 4, 5}
	slice5 := make([]int, len(array5))
	for i, v := range array5 {
		slice5[i] = v
	}

	fmt.Println("slice5:", slice5)

	//  converting array to slice using reflect
	fmt.Println("-----------------")
	array6 := [5]int{1, 2, 3, 4, 5}
	slice6 := reflect.ValueOf(array6[:]).Interface().([]int) // convert array to slice using reflect
	fmt.Println("slice6:", slice6)

	//  converting array to slice using unsafe
	fmt.Println("-----------------")
	array7 := [5]int{1, 2, 3, 4, 5}
	slice7 := (*[5]int)(unsafe.Pointer(&array7))[:] // convert array to slice using unsafe
	fmt.Println("slice7:", slice7)
	fmt.Println(reflect.TypeOf(slice7))

	// converting slice into a pointer to an array

	fmt.Println("-----------------")
	slice8 := []int{1, 2, 3, 4, 5}
	arrayPointer := (*[5]int)(slice8) // convert slice to pointer to array

	fmt.Println("arrayPointer:", arrayPointer)

	slice8[0] = 10
	arrayPointer[1] = 20
	fmt.Println("slice8:", slice8)
	fmt.Println("arrayPointer:", arrayPointer)

	//  Strings

	str := "Hello, 🌍"
	byteSlice := []byte(str)

	fmt.Println("string:", str)
	fmt.Println("byte slice:", byteSlice)
	fmt.Println("byte slice length:", len(byteSlice))

	// iterating over a string using for loop

	// if you loop over a string normally (without range), you get individual bytes, not characters

	for i := 0; i < len(str); i++ {
		fmt.Printf("byte's in decimal number %d: %x\n", i, str[i])
		fmt.Printf("byte's in hexadecimal number %d: %d\n", i, str[i])
		fmt.Printf("byte's in hexadecimal number %d: %c\n", i, str[i])
	}

	fmt.Println("-----------------")

	// Maps

	totalWins := map[string]int{}
	totalWins["teamA"] = 10
	totalWins["teamB"] = 20
	totalWins["teamC"] = 30
	totalWins["teamD"] = 40

	fmt.Println(totalWins)
	totalWins["teamA"]++

	fmt.Println(totalWins)

	text := map[string]int{
		"hello":  1,
		"world":  2,
		"golang": 3,
	}

	fmt.Println(text)

	v, ok := text["hello"]
	fmt.Println(v, ok)

}
