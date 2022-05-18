// Sample logging-quickstart writes a log entry to Cloud Logging.
package sayhi

import (
	"fmt"
)

// "context"
// "log"

// "cloud.google.com/go/logging"

func Test() {
	// ctx := context.Background()

	// // Sets your Google Cloud Platform project ID.
	// projectID := "projects/spry-framework-350102" // "projects/spry-framework-350102"

	// // Creates a client.
	// client, err := logging.NewClient(ctx, projectID)
	// if err != nil {
	// 	log.Fatalf("Failed to create client: %v", err)
	// }
	// defer client.Close()

	// // Sets the name of the log to write to.
	// logName := "my-log"

	// // Selects the log to write to.
	// logger := client.Logger(logName)

	// // Adds an entry to the log buffer.
	// logger.Log(logging.Entry{
	// 	Payload: struct {
	// 		Anything string
	// 	}{
	// 		Anything: "this is a message to confirm logging",
	// 	},
	// 	Severity: logging.Debug,
	// })

	fmt.Println("hi")
}
