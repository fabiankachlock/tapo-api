package request

type TriggerLogsParams struct {
	// jsonValue is the map that will be marshaled into the JSON body of the request.
	// A map must be used explicitly, because otherwise there is no way of differentiating
	// between an empty value and a value that was not set.
	jsonValue map[string]interface{}
}

func NewTriggerLogsParams(pageSize uint64, startId uint64) TriggerLogsParams {
	return TriggerLogsParams{
		jsonValue: map[string]interface{}{
			"page_size": pageSize,
			"start_id":  startId,
		},
	}
}

func (c TriggerLogsParams) GetJsonValue() map[string]interface{} {
	return c.jsonValue
}
