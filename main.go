package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	Functions "project2_api/demo/Functions"

	"github.com/gorilla/mux"
)

// object with array
type Arrays struct {
	ID          string `json:"ID"`
	Array       []int  `json:"Array"`
	Description string `json:"Description"`
}

type allArrays []Arrays

var arrays = allArrays{}

// object with words
type Words struct {
	Word        string `json:"Word"`
	Description string `json:"Description"`
}

type allStrings []Words

var words = allStrings{}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! This is 2nd project with REST API.")
	fmt.Fprintln(w)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Fibonacci: \":port-number/fibonacci/n\"		where \"n\" is the number of elements")
	fmt.Fprintln(w, "Sort Array: \":port-number/sort\"")
	fmt.Fprintln(w, "Find duplicates: \":port-number/duplicates\"")
	fmt.Fprintln(w, "Delete duplicates: \":port-number/deleteduplicates/\"")
	fmt.Fprintln(w, "Find the longest palindromic string: \":port-number/longestpalindromic\"")
	fmt.Fprintln(w, "Missing element: \":port-number/missing\"")
}

// parameter:
//
//	w http.ResponseWriter is used to write the response to the client
//	r *http.Request which represents the HTTP request received from the client
func FibonacciEndpoint(w http.ResponseWriter, r *http.Request) {
	//next line uses the mux.Vars function to extract any variables from the URL
	//that was used to make the request. In this case, we expect to find a variable named "n".
	vars := mux.Vars(r)
	//next line uses the strconv.Atoi function to convert the "n" variable from a string to an integer
	//the _ is a blank identifier, which is used to ignore the second return value of strconv.Atoi, which is an error value.
	n, _ := strconv.Atoi(vars["n"])

	fiboArray := Functions.Fibonacci(n) //this line calls the Fibonacci function

	//creating a new object Arrays

	newArray := Arrays{
		ID:          "1",
		Array:       fiboArray,
		Description: "This is a Fibonacci array.",
	}

	arrays = append(arrays, newArray)

	fmt.Fprintf(w, "Fibonacci serie is: ")

	// w.WriteHeader(http.StatusOK) //indicates that the request was successful
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fiboArray)
}

// in CreatingArray we get all the information from postman
// newArray *Arrays because I want to change in SortArrayUpdateEndpoint
func CreatingArray(w http.ResponseWriter, r *http.Request, newArray *Arrays) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, newArray) //not &newArray because is already pointer

	arrays = append(arrays, *newArray)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(*newArray)

}

func SortArrayUpdateEndpoint(w http.ResponseWriter, r *http.Request) {

	var newArray Arrays
	CreatingArray(w, r, &newArray)
	fmt.Println(newArray)
	// fmt.Fprintf(w, "\nUnsorted array:\n")
	// json.NewEncoder(w).Encode(newArray)

	newArray.Array = Functions.SelectionSort(newArray.Array)
	newArray.Description = "It is sorted."

	for i, array := range arrays {
		if array.ID == newArray.ID {
			arrays[i].Array = newArray.Array
			break
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newArray)

	// // vars := mux.Vars(r)
	// unsortedArray := []int{5, 10, 3, 4, 2, 7}
	// fmt.Fprintf(w, "\nUnsorted array:")
	// json.NewEncoder(w).Encode(unsortedArray)
	// var sortedArray []int
	// switch vars["method"] {
	// case "package":
	// 	sortedArray = Functions.SortArrayWithSortPackage(unsortedArray)
	// 	fmt.Fprintf(w, "\nThe sorted array with package sort is: ")
	// case "selectionsort":
	// 	sortedArray = Functions.SelectionSort(unsortedArray)
	// 	fmt.Fprintf(w, "\nThe sorted array with SelectionSort is: ")
	// default:
	// 	fmt.Fprintf(w, "\nPlease enter the url: \"localhost:port-number/sort/method\". \n The method can be: \"package\" or \"selectionsort\"")
	// }
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(sortedArray)

}

func DuplicatesEndpoint(w http.ResponseWriter, r *http.Request) {

	var newArray Arrays
	CreatingArray(w, r, &newArray)
	fmt.Println(newArray)

	infoElemMap := Functions.FindDublicateElements(newArray.Array)

	// for key, value := range infoElemMap {
	// 	fmt.Fprintln(w, "The element: ", key, "appears: ", value, " times.")
	// }
	fmt.Fprintf(w, "\nThe map is: ")
	json.NewEncoder(w).Encode(infoElemMap)

	for key, value := range infoElemMap {
		fmt.Fprintf(w, "\nThe value %v appears for %v times.", key, value)
	}

	w.Header().Set("Content-Type", "application/json")

}

func DeleteDuplicatesEndpoint(w http.ResponseWriter, r *http.Request) {

	var newArray Arrays
	CreatingArray(w, r, &newArray)
	fmt.Println(newArray)

	newArray.Array = Functions.DeleteDublicateElements(newArray.Array)
	newArray.Description = "Duplicates have been deleted."

	for i, array := range arrays {
		if array.ID == newArray.ID {
			arrays[i].Array = newArray.Array
			break
		}
	}
	json.NewEncoder(w).Encode(newArray)
	w.Header().Set("Content-Type", "application/json")
}

func CreatingString(w http.ResponseWriter, r *http.Request, newString *Words) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, newString) //not &newArray because is already pointer

	words = append(words, *newString)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(*newString)

}
func LongestPalindromicEndpoint(w http.ResponseWriter, r *http.Request) {
	// str := "capacrotitormagnific"
	// fmt.Fprintf(w, "The string is: ")
	// json.NewEncoder(w).Encode(str)

	var newString Words
	CreatingString(w, r, &newString)

	longPalindr := Functions.FindLongestPalindromicSubstring(newString.Word)
	fmt.Fprintf(w, "The longest palindromic substring is: ")
	json.NewEncoder(w).Encode(longPalindr)

	w.Header().Set("Content-Type", "application/json")
}

func MissingElementEndpoint(w http.ResponseWriter, r *http.Request) {
	var newArray Arrays
	CreatingArray(w, r, &newArray)
	fmt.Println(newArray)

	minimArray := Functions.MinElement(newArray.Array)
	maximArray := Functions.MaxElement(newArray.Array)
	fmt.Fprintf(w, "\nThe array has elements in range [ %v , %v ].\n", minimArray, maximArray)

	miss := Functions.FindElementMissing(newArray.Array)
	fmt.Fprintf(w, "The element that are missing is: ")
	json.NewEncoder(w).Encode(miss)

	w.Header().Set("Content-Type", "application/json")

}

func main() {
	// creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	router.HandleFunc("/", homeLink)

	router.HandleFunc("/fibonacci/{n}", FibonacciEndpoint).Methods("GET")
	router.HandleFunc("/sort/", SortArrayUpdateEndpoint).Methods("POST")
	router.HandleFunc("/duplicates/", DuplicatesEndpoint).Methods("GET")
	router.HandleFunc("/deleteduplicates/", DeleteDuplicatesEndpoint).Methods("POST")
	router.HandleFunc("/longestpalindromic/", LongestPalindromicEndpoint).Methods("GET")
	router.HandleFunc("/missing/", MissingElementEndpoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}
