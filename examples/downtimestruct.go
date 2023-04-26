package examples

import (
	"encoding/json"
	"github/erastus/reflect/types"
	"io/ioutil"
	"log"
	"reflect"
)

func ReflandDowntimeStruct(l *log.Logger) {
	var out types.DowntimeStruct
	jsonFile, err := ioutil.ReadFile("downtimes.json")
	if err != nil {
		log.Printf("Unable to read Json: %+v", err)
	}
	err = json.Unmarshal(jsonFile, &out)
	if err != nil {
		l.Printf("error %+v", err)
	}

	refltype := reflect.TypeOf(out)
	reflval := reflect.ValueOf(out)
	for i := 0; i < refltype.NumField(); i++ {
		l.Println(refltype.Field(i).Type.Kind())
		//reffield := reflval.Field(i)
		switch refltype.Field(i).Type.Kind() {
		case reflect.Int:
			l.Printf("Kind -> %v Field Name -> %s, Field Value ->  %d", refltype.Field(i).Type.Kind(), refltype.Field(i).Name, reflval.Field(i).Int())
		case reflect.String:
			l.Printf("Kind -> %v Field Name -> %s, Field Value ->  %s", refltype.Field(i).Type.Kind(), refltype.Field(i).Name, reflval.Field(i).String())
		case reflect.Slice:
			for a := 0; a < reflval.Field(i).Len(); a++ {
				l.Printf("Kind -> %v Field Name -> %s, Field Value ->  %v", refltype.Field(i).Type.Kind(), refltype.Field(i).Name, reflval.Field(i).Index(a))
			}
		case reflect.Bool:
			l.Printf("Kind -> %v Field Name -> %s, Field Value ->  %v", refltype.Field(i).Type.Kind(), refltype.Field(i).Name, reflval.Field(i).Bool())
		case reflect.Struct:
			for z := 0; z < reflval.Field(i).NumField(); z++ {
				l.Printf("Kind -> %v Field Name -> %s, Field Value ->  %v", refltype.Field(i).Type.Kind(), refltype.Field(i).Name, reflval.Field(i).Field(z).Interface())
			}
		default:
			l.Printf("Default case    Kind -> %v, Type -> %v", reflval.Field(i).Kind(), reflval.Field(i).Type())

		}

	}

}
