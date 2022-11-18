package event

import (
	"context"
	"encoding/json"
	"time"

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

// GetEvents returns a list of events
func GetEvents(duration time.Duration) (string, error) {
	ctx := datadog.NewDefaultContext(context.Background())
	api := newDatadogEventsClient()
	resp, _, err := api.ListEvents(
		ctx,
		time.Now().Add(time.Minute*-duration).Unix(),
		time.Now().Unix(),
		*datadogV1.NewListEventsOptionalParameters())

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	return string(responseContent), err
}
