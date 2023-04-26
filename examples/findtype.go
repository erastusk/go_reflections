package examples

import (
	"log"
	"reflect"
)

func TypeExamples(l *log.Logger) {
	t := 5
	u := "string"
	d := map[string]int{}
	e := make(map[string]interface{})
	type Foo struct {
		k string
		p int
	}
	f := Foo{"me", 1}
	//Type ----> Foo

	l.Println(reflect.TypeOf(t), reflect.TypeOf(u), reflect.TypeOf(d), reflect.TypeOf(e), reflect.TypeOf(f))
}

func KindExamples(l *log.Logger) {
	t := 5
	u := "string"
	d := map[string]int{}
	e := make(map[string]interface{})
	type Foo struct {
		k string
		p int
	}
	f := Foo{"me", 1}
	//Type ----> Struct, Map, String, Int
	l.Println(reflect.TypeOf(t).Kind(), reflect.TypeOf(u).Kind(), reflect.TypeOf(d).Kind(), reflect.TypeOf(e).Kind(), reflect.TypeOf(f).Kind())

}

func ElemExamples(l *log.Logger) {
	// Valid when Kind is slice, pointer, map, channel, array

	d := map[string]int{"me": 1}
	e := []string{"how", "are", "you", "today"}
	ctx := make(chan string)
	//Type ----> Struct, Map, String, Int
	l.Println(reflect.TypeOf(d).Elem().Name(), reflect.TypeOf(e).Elem().Name(), reflect.TypeOf(ctx).Elem().Name())
	l.Println(reflect.TypeOf(d).Elem().Kind(), reflect.TypeOf(e).Elem().Kind(), reflect.TypeOf(ctx).Elem().Kind())

}
func ReflandSTructs(l *log.Logger) {

	type Foo struct {
		k string
		p int
		s string
	}
	f := Foo{"me", 1, "you"}
	l.Println(reflect.TypeOf(f).NumField())
	for i := 0; i < reflect.TypeOf(f).NumField(); i++ {
		l.Println(reflect.TypeOf(f).Field(i), reflect.TypeOf(f).Field(i).Name, reflect.TypeOf(f).Field(i).Type.Name(), reflect.ValueOf(f).Field(i))
	}
}
