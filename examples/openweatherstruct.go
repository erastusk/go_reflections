package examples

import (
	"encoding/json"
	"github/erastus/reflect/types"
	"io/ioutil"
	"log"
	"reflect"
)

func ReflandOpenWeatherStruct(l *log.Logger) {

	var out types.OpenWeatherStruct
	jsonFile, err := ioutil.ReadFile("tampa.json")
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
		switch refltype.Field(i).Type.Kind() {
		case reflect.Struct:
			l.Printf("Kind -> %v, Type -> %v, Number of Fields -> %v, Struct Type Name - > %v", reflval.Field(i).Kind(), reflval.Field(i).Type(), reflval.Field(i).NumField(), reflval.Field(i).Type().Name())
			for e := 0; e < reflval.Field(i).NumField(); e++ {
				l.Printf("Struct Type Name- > %s, %v:%v", reflval.Field(i).Type().Name(), reflval.Field(i).Type().Field(e).Name, reflval.Field(i).Field(e).Interface())
			}
		case reflect.Slice:
			l.Printf("Kind -> %v, Type -> %v, Slice length - > %v", reflval.Field(i).Kind(), reflval.Field(i).Type(), reflval.Field(i).Len())
			for c := 0; c < reflval.Field(i).Len(); c++ {
				//l.Println(reflval.Field(i).Type().Name(), reflval.Field(i).Index(c), reflval.Field(i).Index(c).Type().Kind())
				for z := 0; z < reflval.Field(i).Index(c).NumField(); z++ {
					l.Printf("Struct Type Name- > %s, Value -> %v", reflval.Field(i).Type().Name(), reflval.Field(i).Index(c).Field(z).Interface())
				}
			}
		case reflect.String:
			l.Printf("Field Name -> %s, Field Value -> %v", refltype.Field(i).Name, reflval.Field(i).String())

		case reflect.Int:
			l.Printf("Field Name -> %s, Field Value ->  %d", refltype.Field(i).Name, reflval.Field(i).Int())
		case reflect.Int32:
			l.Printf("Field Name -> %s, Field Value ->  %d", refltype.Field(i).Name, reflval.Field(i).Int())
		default:
			l.Printf("Default case    Kind -> %v, Type -> %v", reflval.Field(i).Kind(), reflval.Field(i).Type())
		}
	}

}
