package sailpoint

import (
	"net/http"
	"reflect"
)

func PaginateWithDefaults[T any](f interface{}) ([]T, *http.Response, error) {
	return Paginate[T](f, 0, 250, 10000)
}

func Paginate[T any](f interface{}, initialOffset int32, increment int32, limit int32) ([]T, *http.Response, error) {
	var offset int32 = initialOffset
	var returnObject []T
	var latestResponse *http.Response

	// Ensure we don't request more than the remaining limit
	effectiveIncrement := increment
	if limit > 0 && limit < increment {
		effectiveIncrement = limit
	}

	incrementResp := Invoke(f, "Limit", effectiveIncrement)

	for offset < limit || limit == 0 {
		offsetResp := Invoke(incrementResp[0].Interface(), "Offset", offset)

		resp := Invoke(offsetResp[0].Interface(), "Execute")

		actualValue := resp[0].Interface().([]T)
		latestResponse = resp[1].Interface().(*http.Response)
		err := resp[2].Interface()

		if err != nil {
			return returnObject, latestResponse, err.(error)
		}

		returnObject = append(returnObject, actualValue...)

		if int32(len(actualValue)) < effectiveIncrement {
			break
		}

		offset += effectiveIncrement

		if limit > 0 && offset+effectiveIncrement > limit {
			effectiveIncrement = limit - offset
			incrementResp = Invoke(f, "Limit", effectiveIncrement)
		}
	}

	return returnObject, latestResponse, nil
}

func Invoke(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, v := range args {
		inputs[i] = reflect.ValueOf(v)
	}
	return reflect.ValueOf(any).MethodByName(name).Call(inputs)
}
