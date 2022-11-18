package event

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

// func newDatadogClient() *datadog.APIClient {
func newDatadogEventsClient() *datadogV1.EventsApi {
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewEventsApi(apiClient)
	return api
}

func getEvents() (string, error) {
	ctx := datadog.NewDefaultContext(context.Background())
	api := newDatadogEventsClient()
	resp, _, err := api.ListEvents(ctx, 1668784677, 1668785249, *datadogV1.NewListEventsOptionalParameters())

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListEvents`:\n%s\n", responseContent)
	return string(responseContent), err
}
