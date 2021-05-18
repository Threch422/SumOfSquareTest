package main

import "fmt"

// HENNGE Challenge 2021
// Calculating Sum of Square
// no for statement in this code --> using Recursion as alternative method


/*	sumOfSquareCal
 *	paras: source []int, slice of the input integers in testcase
 *		   index int, index in the slice of input integer in testcase
 *	return: result int, the sum of square of testcase
 */
func sumOfSquareCal(source []int, index int) int {
	result := 0
	// Stoping condition
	if index < 0 {
		return result
	}

	// Calculate only if number is >=0, else skip it
	if source[index] >= 0 {
		result = source[index] * source[index]
		return result + sumOfSquareCal(source, index - 1)
	} else {
		return result + sumOfSquareCal(source, index - 1)
	}
}

/*	gettingInput
 *	paras: arr []int, slice used to store the result of input integer
 *		   err error, error in fmt.Scanf, used to check whether reaching the end of input of testcase
 *	return: arr []int, slice used to store the result of input integer
 */
func gettingInput(arr []int, err error, length int) []int {
	// use to check whether inputting is end or not
	if err != nil {
		return arr
	}

	// Declare input variable
	var input int

	// Start to appending the input into single slice
	_, err = fmt.Scanf("%d", &input)
	
	// Validate the input is within the range and ignore the exceeded input in each testcase
	if input >= -100 && input <= 100 && len(arr) < length {
		arr = append(arr, input)
	}
	return gettingInput(arr, err, length)
}

/*	sumOfSquare
 *	paras: result []int, slice used to store the result of sum of square of each testcase
 *		   testCaseIndex *int, used to record how many testcase we have to go with later
 *	return: result []int, slice used to store the result of sum of square of each testcase
 */
func sumOfSquare(result []int, testCaseIndex *int) []int {
	var (
		numOfElement int
		currSumOfSquare int
	)
	_, err := fmt.Scanf("%d\n", &numOfElement)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error in sumOfSquare()")
	}

	// Validate the number of integers in each testcase is within the range
	if numOfElement < 1 || numOfElement > 100 {
		fmt.Println("Number of integer in this testcase is invalid")
		return result
	}

	currSlice := make([]int, 0)
	
	if *testCaseIndex == 1 {
		currSlice = gettingInput(currSlice, nil, numOfElement)
		currSumOfSquare = sumOfSquareCal(currSlice, len(currSlice) - 1)
		result = append(result, currSumOfSquare)
		return result
	} else {
		currSlice = gettingInput(currSlice, nil, numOfElement)
		currSumOfSquare = sumOfSquareCal(currSlice, len(currSlice) - 1)
		result = append(result, currSumOfSquare)
		*testCaseIndex -= 1
		return sumOfSquare(result, testCaseIndex)
	}

}

/*	printSliceOnebyOne
 *	paras: source []int, the source slicee we needed to print
 *		   index int, the index number of the slice to access the element currently
 *			length int, the length of the slice
 *	return: void, we just print the slice out by recursion
 */
func printSliceOnebyOne(source []int, index int, length int) {
	if index < length {
		fmt.Println(source[index])
		printSliceOnebyOne(source, index + 1, length)
	}	
}


func main() {
	// Getting testcase number first
	var testCaseNum int

	// Error message printing
	_, err := fmt.Scanf("%d\n", &testCaseNum)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error happens in main()")
		return
	}

	if testCaseNum < 1 || testCaseNum > 100 {
		fmt.Println("Number of testcase is invalid")
		return
	}

	result := make([]int, 0)
	result = sumOfSquare(result, &testCaseNum)
	printSliceOnebyOne(result, 0, len(result))

}
