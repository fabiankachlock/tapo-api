package childdevices

import "time"

type DeviceInfoT31X struct {
	DeviceInfoGenericChildDevice

	// 0 when within comfort zone, otherwise the difference
	CurrentHumidityException int8  `json:"current_humidity_exception"`
	CurrentHumidity          uint8 `json:"current_humidity"`
	// 0 when within comfort zone, otherwise the difference
	CurrentTemperatureException float32         `json:"current_temp_exception"`
	CurrentTemperature          float32         `json:"current_temp"`
	TemperatureUnit             TemperatureUnit `json:"temp_unit"`
}

type TemperatureUnit string

const (
	Celsius    TemperatureUnit = "celsius"
	Fahrenheit TemperatureUnit = "fahrenheit"
)

type TemperaturHumidityRecordsRaw struct {
	LocalTime                   int64           `json:"local_time"`
	Past24hHumidityException    []int16         `json:"past24h_humidity_exception"`
	Past24hHumidity             []int16         `json:"past24h_humidity"`
	Past24hTemperatureException []int16         `json:"past24h_temp_exception"`
	Past24hTemperature          []int16         `json:"past24h_temp"`
	TemperatureUnit             TemperatureUnit `json:"temp_unit"`
}

type TemperaturHumidityRecord struct {
	Datetime             time.Time `json:"datetime"`
	HumidityException    int8      `json:"humidity_exception"`
	Humidity             uint8     `json:"humidity"`
	TemperatureException float32   `json:"temp_exception"`
	Temperature          float32   `json:"temp"`
}

type TemperaturHumidityRecords struct {
	Datetime        time.Time                  `json:"datetime"`
	Records         []TemperaturHumidityRecord `json:"records"`
	TemperatureUnit TemperatureUnit            `json:"temp_unit"`
}

func (r TemperaturHumidityRecordsRaw) ToRecords() TemperaturHumidityRecords {
	datetime := time.Unix(r.LocalTime, 0)
	intervalMinute := 0
	if datetime.Minute() >= 45 {
		intervalMinute = 45
	} else if datetime.Minute() >= 30 {
		intervalMinute = 30
	} else if datetime.Minute() >= 15 {
		intervalMinute = 15
	}

	intervalTime := time.Date(datetime.Year(), datetime.Month(), datetime.Day(), datetime.Hour(), intervalMinute, 0, 0, datetime.Location())
	records := make([]TemperaturHumidityRecord, 0, len(r.Past24hTemperature))

	for i := len(r.Past24hTemperature) - 1; i >= 0; i-- {
		humidityException := r.Past24hHumidityException[i]
		humidity := r.Past24hHumidity[i]
		temperatureException := r.Past24hTemperatureException[i]
		temperature := r.Past24hTemperature[i]

		if temperature != -1000 && temperatureException != -1000 && humidity != -1000 && humidityException != -1000 {
			record := TemperaturHumidityRecord{
				Datetime:             intervalTime,
				HumidityException:    int8(humidityException),
				Humidity:             uint8(humidity),
				TemperatureException: float32(temperatureException) / 10.0,
				Temperature:          float32(temperature) / 10.0,
			}
			records = append(records, record)
		}

		intervalTime = intervalTime.Add(-15 * time.Minute)
	}

	// Reverse the records slice
	for i, j := 0, len(records)-1; i < j; i, j = i+1, j-1 {
		records[i], records[j] = records[j], records[i]
	}

	return TemperaturHumidityRecords{
		Datetime:        datetime,
		TemperatureUnit: r.TemperatureUnit,
		Records:         records,
	}
}
