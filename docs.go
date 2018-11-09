/*
Package gokv contains a simple key-value store abstraction in the form of a Go interface.
Implementations of the gokv.Store interface can be found in the sub-packages.

Usage

Example code for using Redis:

	package main

	import (
		"fmt"

		"github.com/philippgille/gokv/redis"
	)

	type foo struct {
		Bar string
	}

	func main() {
		options := redis.DefaultOptions // Address: "localhost:6379", Password: "", DB: 0

		// Create client
		client := redis.NewClient(options)

		// Store value
		val := foo{
			Bar: "baz",
		}
		err := client.Set("foo123", val)
		if err != nil {
			panic(err)
		}

		// Retrieve value
		retrievedVal := new(foo)
		found, err := client.Get("foo123", retrievedVal)
		if err != nil {
			panic(err)
		}
		if !found {
			panic("Value not found")
		}

		fmt.Printf("foo: %+v", *retrievedVal) // Prints `foo: {Bar:baz}`

		// Delete value
		err = client.Delete("foo123")
		if err != nil {
			panic(err)
		}
	}

More details can be found on https://github.com/philippgille/gokv.
*/
package gokv
