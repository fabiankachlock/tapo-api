package request

import "time"

type EnergyDataParams struct {
	// jsonValue is the map that will be marshaled into the JSON body of the request.
	// A map must be used explicitly, because otherwise there is no way of differentiating
	// between an empty value and a value that was not set.
	jsonValue map[string]interface{}

	StartTimestamp uint64 `json:"start_timestamp"`
	EndTimestamp   uint64 `json:"end_timestamp"`
	Interval       uint64 `json:"interval"`
}

func NewEnergyDataParams() EnergyDataParams {
	return EnergyDataParams{
		jsonValue: map[string]interface{}{},
	}
}

func (c EnergyDataParams) GetJsonValue() map[string]interface{} {
	return c.jsonValue
}

func (c EnergyDataParams) setInterval(interval uint64) EnergyDataParams {
	c.jsonValue["interval"] = interval
	return c
}

func (c EnergyDataParams) setStartTimestamp(startTimestamp uint64) EnergyDataParams {
	c.jsonValue["start_timestamp"] = startTimestamp
	return c
}

func (c EnergyDataParams) setEndTimestamp(endTimestamp uint64) EnergyDataParams {
	c.jsonValue["end_timestamp"] = endTimestamp
	return c
}

// Hourly configures the request to return data ina hourly interval.
// start and end are an inclusive interval that ,ust not be greater than 8 days.
func (c EnergyDataParams) Hourly(start time.Time, end time.Time) EnergyDataParams {
	timezone := time.Now().Location()
	startTs := uint64(time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, timezone).Unix())
	endTs := uint64(time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 0, timezone).Unix())
	return c.setInterval(60).setStartTimestamp(startTs).setEndTimestamp(endTs)
}

// Daily configures the request to return data in a daily interval.
// start must be the first day of a quarter.
func (c EnergyDataParams) Daily(start time.Time) EnergyDataParams {
	timezone := time.Now().Location()
	ts := uint64(time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, timezone).Unix())
	return c.setInterval(1440).setStartTimestamp(ts).setEndTimestamp(ts)
}

// Monthly configures the request to return data in a monthly interval.
// start must be the first day of a year.
func (c EnergyDataParams) Monthly(start time.Time) EnergyDataParams {
	ts := uint64(time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location()).Unix())
	return c.setInterval(43200).setStartTimestamp(ts).setEndTimestamp(ts)
}
