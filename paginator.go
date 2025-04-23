package sailpoint

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"strings"

	v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
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

func PaginateSearchApi(ctx context.Context, apiClient *APIClient, search v3.Search, initialOffset int32, increment int32, limit int32) ([]map[string]interface{}, *http.Response, error) {
	var offset int32 = initialOffset
	var returnObject []map[string]interface{}
	var r *http.Response

	if len(search.Sort) != 1 {
		return nil, nil, errors.New("search must include exactly one sort parameter to paginate properly")
	}

	for offset < limit {
		if len(returnObject) > 0 {
			search.SearchAfter = []string{returnObject[len(returnObject)-1][strings.Trim(search.Sort[0], "-")].(string)}
		}
		// convert the expected return values to their respective types
		actualValue, latestResponse, err := apiClient.V3.SearchAPI.SearchPost(ctx).Limit(increment).Search(search).Execute()

		if err != nil {
			return returnObject, latestResponse, err.(error)
		}

		// append the results to the main return object
		returnObject = append(returnObject, actualValue...)
		r = latestResponse

		// check if this is the last set in the response. This could be enhanced by inspecting the header for the max results
		if int32(len(actualValue)) < increment {
			break
		}

		offset += increment
	}
	return returnObject, r, nil
}

func Invoke(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, v := range args {
		inputs[i] = reflect.ValueOf(v)
	}
	return reflect.ValueOf(any).MethodByName(name).Call(inputs)
}
