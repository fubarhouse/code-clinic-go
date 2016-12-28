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
			results = append(results, DataSet{ result[0], result[1], result[2], result[3], result[4], result[5], result[6], result[7], result[8]})
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
	var filename = "./Exercise Files/problems/Problem01/data.txt"
	var data = loadData(1, filename)

	// Process and store required data
	var aTMd = median("Air_Temp", data)
	var aTMn = mean("Air_Temp", data)
	var wSMd = median("Wind_Speed", data)
	var wSMn = mean("Wind_Speed", data)
	var bPMd = median("Barometric_Press", data)
	var bPMn = mean("Barometric_Press", data)

	// Print the data above required to complete code clinic task #1
	fmt.Printf("Median of %v is %.2f\n", "Air_Temp", aTMd)
	fmt.Printf("Mean of %v is %.2f\n", "Air_Temp", aTMn)

	fmt.Printf("Median of %v is %.2f\n", "Wind_Speed", wSMd)
	fmt.Printf("Mean of %v is %.2f\n", "Wind_Speed", wSMn)

	fmt.Printf("Median of %v is %.2f\n", "Barometric_Press", bPMd)
	fmt.Printf("Mean of %v is %.2f\n", "Barometric_Press", bPMn)
}