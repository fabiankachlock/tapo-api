package request

import "time"

type GetEnergyDataParams struct {
	StartTimestamp uint64 `json:"start_timestamp"`
	EndTimestamp   uint64 `json:"end_timestamp"`
	Interval       uint64 `json:"interval"`
}

// GetEnergyDataParamsHourly creates [GetEnergyDataParams] for hourly data.
func GetEnergyDataParamsHourly(start time.Time, end time.Time) GetEnergyDataParams {
	return GetEnergyDataParams{
		StartTimestamp: uint64(time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location()).Unix()),
		EndTimestamp:   uint64(time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 0, end.Location()).Unix()),
		Interval:       60,
	}
}

// GetEnergyDataParamsDaily creates [GetEnergyDataParams] for daily data.
func GetEnergyDataParamsDaily(start time.Time) GetEnergyDataParams {
	ts := uint64(time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location()).Unix())
	return GetEnergyDataParams{
		StartTimestamp: ts,
		EndTimestamp:   ts,
		Interval:       1440,
	}
}

// GetEnergyDataParamsMonthly creates [GetEnergyDataParams] for monthly data.
func GetEnergyDataParamsMonthly(start time.Time) GetEnergyDataParams {
	ts := uint64(time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location()).Unix())
	return GetEnergyDataParams{
		StartTimestamp: ts,
		EndTimestamp:   ts,
		Interval:       43200,
	}
}
