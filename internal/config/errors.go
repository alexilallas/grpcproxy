package config

import "fmt"

type MissingEnv string

func (e MissingEnv) Error() string {
	return fmt.Sprintf("missing env %s", string(e))
}
