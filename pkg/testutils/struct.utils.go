package testutils

import (
	"math/rand"
	"reflect"
	"time"
)

// ------------------------------------------------ PACKAGE VARIABLES --------------------------------------------------

// All characters that are valid in a URL
const randomChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!*'();:@&=+$,/?#[]_.~-"

// Seeded random numbers
var seed = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// ------------------------------------------------- PUBLIC FUNCTIONS --------------------------------------------------

// Populates the given (pointer) to a structure with random data in every field.
func PopulateTestData(obj interface{}) {

	// Get the concrete type
	//println("Address of object: ", obj)
	//println("Obj type : ", reflect.TypeOf(obj).String())
	//println("Obj type : ", reflect.TypeOf(obj).Kind().String())
	//println("Obj value : ", reflect.ValueOf(obj).Kind().String())

	typ := reflect.TypeOf(obj).Elem()
	val := reflect.ValueOf(obj).Elem()

	// Loop over each property
	for i := 0; i < typ.NumField(); i++ {

		ft := typ.Field(i)
		fv := val.Field(i)

		switch ft.Type.Kind() {
		case reflect.String:
			fv.SetString(RandomString(20))
		case reflect.Int:
			fv.SetInt(int64(RandomInt(0, 9999999)))
		case reflect.Struct:
			// Create a new instance of the struct
			newStruct := reflect.New(ft.Type).Elem()
			ptr := newStruct.Addr().Interface()

			// Recursively populate it
			PopulateTestData(ptr)

			// Set to the newly created struct
			fv.Set(reflect.ValueOf(ptr).Elem())
		default:
			println("Unknown type for PopulateTestData: ", ft.Type.Kind().String())
		}
	}

}

// Generate a random string of the given length
func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = randomChars[seed.Intn(len(randomChars))]
	}
	return string(b)
}

// Generate a random integer between the given values (inclusive)
func RandomInt(min int, max int) int {
	return seed.Intn(max-min) + min
}
