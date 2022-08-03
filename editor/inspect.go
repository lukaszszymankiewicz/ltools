package editor

import (
	"image"
	"reflect"
)

func InspectRect(st interface{}, names []string) int {
	val := reflect.ValueOf(st)

	if len(names) == 0 {
		return int(val.Int())
	}

	for i := 0; i < val.NumField(); i++ {
		if val.Type().Field(i).Name == names[0] {
			f := val.Field(i)
			return InspectRect(f.Interface(), names[1:])
		}
	}
	return 0
}

func ReadRect(st interface{}, names []string) image.Rectangle {
	val := reflect.ValueOf(st)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {

		if len(names) == 0 {
			f := val.Field(i)

			if f.CanInterface() {
				MinX := InspectRect(f.Interface(), []string{"Rect", "Min", "X"})
				MinY := InspectRect(f.Interface(), []string{"Rect", "Min", "Y"})
				MaxX := InspectRect(f.Interface(), []string{"Rect", "Max", "X"})
				MaxY := InspectRect(f.Interface(), []string{"Rect", "Max", "Y"})

				return image.Rect(MinX, MinY, MaxX, MaxY)
			}

		} else if val.Type().Field(i).Name == names[0] {
			f := val.Field(i)
			return ReadRect(f.Interface(), names[1:])
		}
	}

	// this damn recursion validation in golang!
	return image.Rect(0, 0, 0, 0)
}
