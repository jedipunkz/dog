package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

// DatadogEvent is structure for datadog evnet
type DatadogEvent struct {
	api *datadogV1.EventsApi
}

func (d *DatadogEvent) newDatadogEventsClient() {
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewEventsApi(apiClient)
	d.api = api
}

// GetEvents returns a list of events
func (d *DatadogEvent) GetEvents(duration time.Duration) (string, error) {
	ctx := datadog.NewDefaultContext(context.Background())
	d.newDatadogEventsClient()
	resp, _, err := d.api.ListEvents(
		ctx,
		time.Now().Add(time.Minute*-duration).Unix(),
		time.Now().Unix(),
		*datadogV1.NewListEventsOptionalParameters())

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	return string(responseContent), err
}
