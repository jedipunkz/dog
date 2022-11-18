// Get a list of events returns "OK" response

package main

import (
	"fmt"
	"log"

	"github.com/jedipunkz/dog/internal/pkg/datadog/event"
)

func main() {
	var datadogEvent event.DatadogEvent

	events, err := datadogEvent.GetEvents(30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(events)
}
