## About the project

This repository contains the complete build, with assets, for the SailPoint Golang SDK. 
The Golang SDK allows you to use Golang to interact with the SailPoint APIs and other resources to customize and extend your SailPoint platform. 

## Documentation 

To get started with the Golang SDK and learn about its various functionalities, refer to the [documentation](https://developer.sailpoint.com/idn/tools/sdk/go) on https://developer.sailpoint.com. 

## Calling any API with the generic client

The generic client calls any SailPoint versioned API by its path, including APIs
that do not have a resource-specific client or documentation yet. The path
includes the API version, for example `accounts/v1`. The client uses the same
configuration, authentication and retries as the other clients.

- Each path segment is URL-encoded. Pass the segments unencoded.
- Anything after a `?` in the path is sent as query parameters.
- A PATCH body is a JSON Patch list, sent as `application/json-patch+json`.

```go
accounts, _, err := apiClient.Generic.DefaultAPI.GenericGet(ctx, "accounts/v1").Limit(10).Execute()

patch := []map[string]interface{}{{"op": "replace", "path": "/attributes/input", "value": "lower"}}
_, _, err = apiClient.Generic.DefaultAPI.GenericPatch(ctx, "transforms/v1/2c91808a").RequestBody(patch).Execute()
```

The response is the decoded JSON (`interface{}`): usually `map[string]interface{}` or `[]interface{}`.

## Examples

In addition to the examples shown in the [documentation](https://developer.sailpoint.com/idn/tools/sdk/go), you can find more uses of the Golang SDK in the [examples](./examples/sdk.go).

## Contributing 

The SailPoint Developer Relations Team is always actively working to improve the SailPoint Golang SDK to make it as useful as possible to the SailPoint Developer Community. 
This means adding and maintaining new resources and functionalities for the SDK.
The most valuable resource of all, however, is the community itself! 
We greatly appreciate any input, feedback, and direct contributions you can provide to the Golang SDK, big or small. 

Refer to our [contribution guidelines](./CONTRIBUTING.md) to learn more about the different ways you can contribute to the SDK. 

Before you contribute, you [must sign our CLA](https://cla-assistant.io/sailpoint-oss/developer.sailpoint.com).

## License

Distributed under the MIT License. Refer to [the license](./LICENSE) for more information.

## Code of Conduct

We pledge to act and interact in ways that contribute to an open, welcoming, diverse, inclusive, and healthy community. Read our [code of conduct](./CODE_OF_CONDUCT.md) to learn more.

