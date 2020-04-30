package main

import "fmt"

// Only define a set of behavior... there is nothing that is
// "Concrete", no state should ever be stored, no values ever
// set in an interface
// Behavior!
type reader interface {
	read(b []byte) (int, error)
}

type file struct {
	name string
}

type pipe struct {
	name string
}

func (file) read(b []byte) (int, error) {
	s := "Going Go Programming"
	copy(b, s)
	return len(s), nil
}

func (pipe) read(b []byte) (int, error) {
	s := `{name: "Bill"}`
	copy(b, s)
	return len(s), nil
}

func main() {
	f := file{"data.json"}
	p := pipe{"pipe.json"}

	// this is passed by value
	retrieve(f)
	retrieve(p)
}

func retrieve(r reader) error {
	data := make([]byte, 100)
	len, err := r.read(data)
	if err != nil {
		return err
	}
	fmt.Println(string(data[:len]))
	return nil
}
