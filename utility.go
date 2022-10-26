package serr

import (
	"reflect"
	"strings"
)

// StandardFormat function
func StandardFormat() string {
	return "In %s[%s:%d] %s.%s"
}

// StandardColorFormat function
func StandardColorFormat() string {
	frmt := ""
	frmt += applyConsoleForeColor("In", consoleColorDarkGray) + " "
	frmt += applyConsoleForeColor("%s", consoleColorLightYellow)
	frmt += applyConsoleForeColor("[", consoleColorDarkGray)
	frmt += applyConsoleForeColor("%s:%d", consoleColorMagenta)
	frmt += applyConsoleForeColor("]", consoleColorDarkGray)
	frmt += " %s.%s"
	return frmt
}

// isExists returns true if the value exists in the array, otherwise returns false.
func isExists(value interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		sliceValue := reflect.ValueOf(array)
		for i := 0; i < sliceValue.Len(); i++ {
			if reflect.DeepEqual(value, sliceValue.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}

// resolvePath returns sort file path after trim by root paths.
func resolvePath(val string) string {
	for _, v := range rootPaths {
		if strings.HasPrefix(val, v) {
			val = val[len(v):]
			return val
		}
	}

	return val
}
