package main

import "log"

// basic types (integer, string, boolean)
// integers:
var myInt int
var myInt16 int16
var myInt32 int32
var myInt64 int64
var myUInt uint // unsigned integer
// decimal
var myFloat float32
var myFloat64 float64

// string -> strings are immutable!
var myString = "Mohsen"

// boolean
var myBool = false

func main() {
	myInt = 10
	myUInt = 20
	myFloat = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUInt, myFloat, myFloat64, myString, myBool)
}
