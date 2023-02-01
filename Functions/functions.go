package functions

import (
	"fmt"
	"sort"
)

func Fibonacci(n int) []int {
	f1 := 0
	f2 := 1
	var f3 int

	newArray := make([]int, n+1)
	newArray[0] = f1
	newArray[1] = f2

	fmt.Print("\n Fibonacci serie is: ", f1, f2)

	for i := 3; i <= n; i++ {
		f3 = f1 + f2
		newArray[i-1] = f3
		f1 = f2
		f2 = f3

		fmt.Print(" ", f3)

	}

	return newArray[:n]
}

func SortArrayWithSortPackage(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	// fmt.Println("The sorted array with SelectionSort is: ", arr)
	return arr
}

func MaxElement(arr []int) int {
	maxim := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxim {
			maxim = arr[i]
		}
	}
	return maxim
}

func MinElement(arr []int) int {
	minim := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < minim {
			minim = arr[i]
		}
	}
	return minim
}

func FrequencyArray(arr []int, helper []int) {
	for i := 0; i < len(arr); i++ {
		helper[arr[i]]++
	}
}

func FindDublicateElements(arr []int) map[int]int {

	maximArray := MaxElement(arr)

	//helper array is like frequency array in c++
	helper := make([]int, maximArray+1)

	FrequencyArray(arr, helper)

	var infoElem = make(map[int]int)

	for i := 0; i < len(helper); i++ {
		if helper[i] >= 2 {
			//fmt.Printf("The element %v appears %v times.\n", i, helper[i])
			infoElem[i] = helper[i]
		}
	}
	return infoElem
}

func DeletionWithNewArray(arr []int, helper []int) []int {
	nr := 0
	for j := 0; j < len(helper); j++ {
		if helper[j] > 0 {
			nr++
		}
	}
	newArray := make([]int, nr)
	k := 0
	for i := 0; i < len(arr); i++ {
		if helper[arr[i]] > 0 {
			newArray[k] = arr[i]
			helper[arr[i]] = 0
			k++
		}
	}
	// fmt.Println("The new array after the deletions: ", newArray)
	return newArray
}

// func DeletionFromArray(arr []int, helper []int) []int {
// 	//O(n^2)
// 	n := len(arr)
// 	for i := len(arr) - 1; i >= 0; i-- {
// 		if helper[arr[i]] >= 2 {
// 			helper[arr[i]]--
// 			for j := i; j < len(arr)-1; j++ {
// 				arr[j] = arr[j+1]
// 			}
// 			n--
// 		}
// 	}
// 	var newArray []int
// 	newArray = append(newArray, arr[:n]...)
// 	return newArray
// 	// fmt.Println("The new array after the deletions: ", newArray)
// }

func DeleteDublicateElements(arr []int) []int {
	maximArray := MaxElement(arr)

	//helper array is like frequency array in c++
	helper := make([]int, maximArray+1)

	FrequencyArray(arr, helper)

	//The variant with adding the elements that interest me in a new array to avoid O(n^2)
	newArray := DeletionWithNewArray(arr, helper)

	return newArray
}

func FindLongestPalindromicSubstring(str string) string {
	//a palindrome can have even length or odd length. For a certain fixed position i, it is checked if there is a
	//palindrome of odd length centered in position i and then a palindrome of even length centered in positions i and i+1
	b := []byte(str)

	lgmax := 1

	var newString []byte

	var copie_p int
	var copie_q int

	for i := 1; i < len(b); i++ {
		//search for palindrome of odd length  centered in position i
		p := i - 1
		q := i + 1

		for p >= 0 && q < len(b) && b[p] == b[q] {
			p--
			q++
		}

		p++
		q--

		if lgmax < q-p+1 {
			lgmax = q - p + 1
			copie_p = p
			copie_q = q
		}

		//search for palindrome of even length centered in position i and i+1
		p = i
		q = i + 1
		for p >= 0 && q < len(b) && b[p] == b[q] {
			p--
			q++
		}

		q--

		if lgmax < q-p+1 {
			lgmax = q - p + 1
			copie_p = p
			copie_q = q
		}
	}

	newString = append(newString, b[copie_p:copie_q+1]...)
	return string(newString[:])

}

func FindElementMissing(arr []int) int {
	maximArray := MaxElement(arr)
	minimArray := MinElement(arr)
	fmt.Printf("\nThe array has elements in range [ %v , %v ].", minimArray, maximArray)

	helper := make([]int, maximArray+1)
	FrequencyArray(arr, helper)

	fmt.Print("\nThe element that are missing is: ")
	for i := minimArray; i < len(helper); i++ {
		if helper[i] == 0 {
			return i
		}
	}

	return 0
}
