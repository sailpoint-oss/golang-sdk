# golang-sdk

### Create your project

```bash
go mod init github.com/github-repo-name/projectname
```

### Create sdk.go file and copy [example](./examples/sdk.go) from this repository

* Add your ClientID and ClientSecret
* Update your tenant name in the configuration

```go
package main

import (
	"context"
	"fmt"
	"os"

	sailpoint "github.com/sailpoint-oss/golang-sdk/sdk-output"
)

func main() {

	auth := context.WithValue(context.Background(), sailpoint.ContextClientCredentials, sailpoint.ClientCredentials{ClientId: "", ClientSecret: ""})

	configuration := sailpoint.NewConfiguration("devrel")

	apiClient := sailpoint.NewAPIClient(configuration)

	resp, r, err := apiClient.V3.AccountsApi.ListAccounts(auth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccounts`: []Account
	fmt.Fprintf(os.Stdout, "First response from `AccountsApi.ListAccount`: %v\n", resp[0].Name)

}
```

### Install sdk

```bash
go mod tidy
```

### Run the example

```bash
go run sdk.go
```


### Handling Pagination

there is a built in pagination function that can be used to automatically call and collect responses from APIs that support pagination. Use the following syntax to call it:

```go
import (
	"context"
	"fmt"
	"os"

	sailpoint "github.com/sailpoint-oss/golang-sdk/sdk-output"
	// need to import the v3 library so we are aware of the sailpointsdk.Account struct
	sailpointsdk "github.com/sailpoint-oss/golang-sdk/sdk-output/v3"
)

func main() {

	auth := context.WithValue(context.Background(), sailpoint.ContextClientCredentials, sailpoint.ClientCredentials{ClientId: "", ClientSecret: ""})

	configuration := sailpoint.NewConfiguration("devrel")

	apiClient := sailpoint.NewAPIClient(configuration)

	// use the paginate function to get 1000 results instead of hitting the normal 250 limit
	resp, r, err := sailpoint.Paginate[sailpointsdk.Account](apiClient.V3.AccountsApi.ListAccounts(auth), 0, 1000)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccounts`: []Account
	fmt.Fprintf(os.Stdout, "First response from `AccountsApi.ListAccount`: %v\n", resp[0].Name)

}