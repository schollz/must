package must

import (
	"reflect"
)

// Must takes a function and optional parameters and
// evaluates the function, returning the first result no matter what.
func Must(fn interface{}, params ...interface{}) interface{} {
	// recover
	defer func() {
		if r := recover(); r != nil {
			// do nothing
		}
	}()
	fnVal := reflect.ValueOf(fn)
	fnType := fnVal.Type()

	if fnType.Kind() != reflect.Func {
		// oops no function, oh well
		return nil
	}

	in := make([]reflect.Value, len(params))
	for i, param := range params {
		in[i] = reflect.ValueOf(param)
	}

	results := fnVal.Call(in)

	// don't bother checking errors even
	// lastResult := results[len(results)-1]
	// if !lastResult.IsNil() {
	// 	// do nothing
	// }

	return results[0].Interface()
}
