package test

import (
	"time"
)

// Entity information
type Entity struct {
	A    string  //string
	B    *string // pointer
	c    string  // not exported
	d    *string
	E    int
	f    *int
	G    uint `structs:"y"`
	H    uint `json:"h"`
	I    bool
	j    bool `xml:"j"` // not exported, with tag
	K    Bar
	L    *Bar
	Buzz `cool`
	*Bar "beans"
	M    []string
	n    *[]string
	O    map[string]interface{}
	p    *map[string]interface{}
	q    interface{}
	r    *interface{}
	s    []interface{}
	t    *[]interface{}
	u    time.Time
	v    *time.Time
	Fizz struct {
		Legs int `json:"legs"`
		Buzz
	}
	w             chan int
	length, width int
	ToString      func(int) string
}

// Buzz test struct
type Buzz struct {
	A *string
	B int
}

// Bar test struct
type Bar struct {
	E string
	F int
	g []string
}
