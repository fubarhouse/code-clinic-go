// Code Clinic Problem #1 Solution
// Features much more than the task requested, in an effort to learn Go.
//
// Built with Gogland EAP 1.0
// * Data structure for all data
// * Data stored in slice of struct
// * Structs package converts struct to slice where needed.
// * API Call which returns one complete value
// * API Call which returns one whole string field from any entry
// * API Call which returns one not so whole float64 field from any entry
// * Error handling function which is called when the requested file isn't found
// * Function dedicated to finding the mean of dataset with the usage of (string) keys.
// * Function dedicated to finding the median of dataset with the usage of (string) keys.
// * (Added later) Prevent false readings from Median function by sorting returned values prior to return.
// * (Added later) API function to create and return a DataSet object, which allows scalable and programmable adding of items into the data.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"structs" // https://github.com/fatih/structs
	"sort"
)

// DataSet struct to house all data
type DataSet struct {
	date			string
	time			string
	Air_Temp		string
	Barometric_Press	string
	Dew_Point		string
	Relative_Humidity	string
	Wind_Dir		string
	Wind_Gust		string
	Wind_Speed		string
}

// Generic error handling function which can be scaled per error object
func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

// This function returns a single newly created DataSet struct object.
func makeDataObject(data []string) DataSet {
	return DataSet{ data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7], data[8]}
}

// The main data function which generates structured data from the inputted file path
func loadData(offset int, filename string) (slice []DataSet) {
	var file, err = os.Open(filename)
	defer file.Close()
	errorHandler(err)
	var count = 0
	scanner := bufio.NewScanner(file)
	var results = make([]DataSet, 0, len(scanner.Text()))
	for scanner.Scan() {
		if count > offset {
			var result = strings.Fields(scanner.Text())
			results = append(results, makeDataObject(result))
		}
		count += 1
	}
	return results
}

// Go and get a single data entry and return as structured data
func getEntry(i int, data []DataSet) (map[string]interface{}) {
	return structs.Map(data[i])
}

// Return a single field from a single entry as a string value
func getFieldString(i int, s string, data []DataSet) (string) {
	var components = structs.Map(data[i])
	var stringComponent = ""
	stringComponent += components[s].(string)
	return stringComponent
}

// Return a single field from a single entry as a float64 value
func getFieldFloat(i int, s string, data []DataSet) (float64) {
	var components = structs.Map(data[i])
	var stringComponent = ""
	stringComponent += components[s].(string)
	var floatComponent, _ = strconv.ParseFloat(stringComponent, 64)
	return floatComponent
}

// Returns the median of multiple data sets based upon input data
func median(value string, data []DataSet) (float64) {
	var numberList[]float64
	for item := range data {
		var currentNumber = getFieldFloat(item,value,data)
		numberList = append(numberList, currentNumber)
	}
	sort.Float64s(numberList)
	var returnValue = numberList[len(numberList) / 2]
	return returnValue
}

// Returns the mean of multiple data sets based upon input data
func mean(value string, data []DataSet) (float64) {
	var currentTotal float64 = 0.00;
	for i := 0; i < len(data); i++  {
		currentTotal += getFieldFloat(i, value, data)
	}
	return currentTotal / float64(len(data))
}

// Our main function
func main() {
	// Specify file, Generate data
	var filename= "./Exercise Files/problems/Problem01/data.txt"
	var data= loadData(1, filename)

	// Process required data and output problem requirements

	fmt.Printf("Median of %v is %.2f\n", median("Air_Temp", data))
	fmt.Printf("Median of %v is %.2f\n", median("Wind_Speed", data))
	fmt.Printf("Median of %v is %.2f\n", median("Barometric_Press", data))

	fmt.Printf("Mean of %v is %.2f\n", mean("Air_Temp", data))
	fmt.Printf("Mean of %v is %.2f\n", mean("Wind_Speed", data))
	fmt.Printf("Mean of %v is %.2f\n", mean("Barometric_Press", data))
}