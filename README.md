# CSRF Protection Middleware for Go

This library provides a middleware function for handling CSRF protection in Go HTTP servers. It generates a CSRF token and sets it as a cookie in the response, and checks for a valid CSRF token in the request body or X-CSRF-Token header.
Installation

To install the library, use the go get command:
```go
go get github.com/mail4metablocks/csrf-protection-middleware-go
```

## Usage

To use the CSRF middleware, wrap your application's routes in the CSRFMiddleware function:


```go

package main

import (
	"net/http"
	"github.com/mail4metablocks/csrf-protection-middleware-go"
)

func main() {
	http.Handle("/", csrf.CSRFMiddleware(http.HandlerFunc(handleIndex)))
	http.Handle("/login", csrf.CSRFMiddleware(http.HandlerFunc(handleLogin)))
	http.ListenAndServe(":8080", nil)
}
```

To include the CSRF token in your forms, use a hidden field with the name csrf_token:

```go
<form method="post">
  <input type="hidden" name="csrf_token" value="{{csrf_token}}">
  <!-- form fields -->
</form>

```
