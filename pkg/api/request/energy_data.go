package request

import "time"

type GetEnergyDataParams struct {
	StartTimestamp uint64 `json:"start_timestamp"`
	EndTimestamp   uint64 `json:"end_timestamp"`
	Interval       uint64 `json:"interval"`
}

func GetEnergyDataParamsHourly(start time.Time, end time.Time) GetEnergyDataParams {
	return GetEnergyDataParams{
		StartTimestamp: uint64(time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location()).Unix()),
		EndTimestamp:   uint64(time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 0, end.Location()).Unix()),
		Interval:       60,
	}
}

func GetEnergyDataParamsDaily(start time.Time) GetEnergyDataParams {
	ts := uint64(time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location()).Unix())
	return GetEnergyDataParams{
		StartTimestamp: ts,
		EndTimestamp:   ts,
		Interval:       1440,
	}
}

func GetEnergyDataParamsMonthly(start time.Time) GetEnergyDataParams {
	ts := uint64(time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location()).Unix())
	return GetEnergyDataParams{
		StartTimestamp: ts,
		EndTimestamp:   ts,
		Interval:       43200,
	}
}
