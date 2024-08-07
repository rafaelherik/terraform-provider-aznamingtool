# Error Handling 

The utils package provides utilities for handling errors in a structured and consistent way. It defines custom error types and functions to work with these errors effectively.

## Constants

e.g, `ErrorClientNotInitialized`
```go 
var ErrorClientNotInitialized = constError("client not initialized")
```

A predefined constant error for scenarios where a client is not initialized.

## Types

### `constError` 
*A custom error type that implements the error interface.*

#### Methods

* `Error() string` Returns the string representation of the constError.

```go
func (err constError) Error() string
```

* `Is(target error) bool`

Checks if the target error matches the constError. It also supports matching if the target error message has a prefix of the constError message.

```go
func (err constError) Is(target error) bool
```
* `wrap(inner error) error`

Creates a new wrapError which wraps the given inner error with the constError message.

```go
func (err constError) wrap(inner error) error
```


### `wrapError` 

*A custom error type that includes an inner error and a message.*

#### Methods

* `Error() string`

Returns the combined error message of wrapError, including the inner error if present.

```go
func (err wrapError) Error() string
```

* `Unwrap() error`

Returns the inner error of wrapError.

```go
func (err wrapError) Unwrap() error
```

* `Is(target error) bool`

Checks if the target error matches the wrapError's message using constError's Is method.

```go
func (err wrapError) Is(target error) bool
```

## Usage

* Example

Here is an example of how to use the utils package to handle errors:

```go
package main

import (
	"errors"
	"fmt"
	"yourmodule/utils"
)

func initializeClient() error {
	return utils.ErrorClientNotInitialized
}

func main() {
	err := initializeClient()
	if errors.Is(err, utils.ErrorClientNotInitialized) {
		fmt.Println("Client is not initialized:", err)
	}

	// Wrapping an error
	innerErr := errors.New("network timeout")
	wrappedErr := utils.ErrorClientNotInitialized.wrap(innerErr)
	fmt.Println("Wrapped error:", wrappedErr)

	// Checking wrapped error
	if errors.Is(wrappedErr, utils.ErrorClientNotInitialized) {
		fmt.Println("Matched wrapped error:", wrappedErr)
	}
}

```

### In this example:

* The `initializeClient` function returns an `ErrorClientNotInitialized` error.
* The main function checks if the error matches `ErrorClientNotInitialized`.
* An inner error (network timeout) is wrapped with `ErrorClientNotInitialized`.
* The wrapped error is printed and checked if it matches `ErrorClientNotInitialized`.
* This approach provides a consistent and structured way to handle errors in your Go application.
