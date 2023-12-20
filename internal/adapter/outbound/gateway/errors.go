package gateway

import "fmt"

type ErrorResponse int

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("request failed: server returned %d status code", int(e))
}
