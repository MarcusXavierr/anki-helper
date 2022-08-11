// Package  check provides ...
package check

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
