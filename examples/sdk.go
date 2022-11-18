package main

import (
	"context"
	"fmt"
	"os"

	sailpoint "github.com/sailpoint-oss/golang-sdk/sdk-output"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World1!")

	auth := context.WithValue(context.Background(), sailpoint.ContextClientCredentials, sailpoint.ClientCredentials{ClientId: "", ClientSecret: ""})

	configuration := sailpoint.NewConfiguration("devrel")

	apiClient := sailpoint.NewAPIClient(configuration)

	//apiClient2 := sailpointsdk.NewAPIClient(sailpointsdk.NewConfiguration("test"))

	resp, r, err := apiClient.V3.TransformsApi.GetTransformsList(auth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransformsApi.GetTransformsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransformsList`: []Transform
	fmt.Fprintf(os.Stdout, "Response from `TransformsApi.GetTransformsList`: %v\n", resp[0].Name)

	fmt.Println(r.Body)

}
