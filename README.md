# Go Code Clinic

## Problem #1

This attempt features much more than the task requested, in an effort to learn Go.

I've attempted to make the data searchable by index and field name, as a result it's stored in a sliced structured data.

* Built with Gogland EAP 1.0
* Data structure for all data
* Data stored in slice of struct
* Structs package converts struct to slice where needed.
* API Call which returns one complete value
* API Call which returns one whole string field from any entry
* API Call which returns one not so whole float64 field from any entry
* Error handling function which is called when the requested file isn't found
* Function dedicated to finding the mean of dataset with the usage of (string) keys.
* Function dedicated to finding the median of dataset with the usage of (string) keys.
* (Added later) Prevent false readings from Median function by sorting returned values prior to return.
* (Added later) API function to create and return a DataSet object, which allows scalable and programmable adding of items into the data. 