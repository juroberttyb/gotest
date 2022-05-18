// Sample logging-quickstart writes a log entry to Cloud Logging.
package main

import (
	"fmt"
	hi "github.com/juroberttyb/gotest/sayhi"
)

func main() {
	fmt.Println("hello")
	hi.Test()
}
