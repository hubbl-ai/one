package dm

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/golang/glog"
)

func serialize(obj interface{}) (
	buf []byte,
	err error,
) {
	m := make(map[string]interface{})

	defer func() {
		if value := recover(); value != nil {
			glog.Fatal(value)
		}

		if err != nil {
			msg := fmt.Sprintf("Error return from appendFields(%s, %v): %s",
				m,
				obj,
				err.Error())
			panic(msg)
		}
	}()

	m["@context"] = map[string]interface{}{
		"iata": "https://onerecord.iata.org",
	}

	vRecord := reflect.ValueOf(obj)

	if vRecord.Type().Kind() == reflect.Ptr {
		vRecord = vRecord.Elem()
	}

	m["@type"] = fmt.Sprintf("iata:%s", vRecord.Type().Name())

	n := vRecord.NumField()

	for ix := 0; ix < n; ix++ {
		vField := vRecord.Field(ix)

		if !vField.IsValid() {
			err = fmt.Errorf("invalid field %d", ix)
			glog.Error(err.Error())
			return nil, err
		}

		t := vField.Type()
		tag := vRecord.Type().Field(ix).Tag.Get("json")

		omitempty := false

		if tag != "" {
			parts := strings.Split(tag, ",")
			tag = parts[0]
			omitempty = len(parts) > 0 && parts[1] == "omitempty"
		} else {
			tag = t.Name()
		}

		switch t.Kind() {
		case reflect.String:
			value := vField.String()

			if value != "" || !omitempty {
				m[tag] = value
			}
		case reflect.Bool:
			value := vField.Bool()

			if value || !omitempty {
				m[tag] = value
			}
		case reflect.Float64:
			value := vField.Float()

			if value != 0 || !omitempty {
				m[tag] = value
			}
		case reflect.Int64:
			value := vField.Int()

			if value != 0 || !omitempty {
				m[tag] = value
			}
		case reflect.Pointer, reflect.Interface:
			value := vField.Interface()

			if value != nil || !omitempty {
				m[tag] = value
			}
		}
	}

	return json.Marshal(m)
}
