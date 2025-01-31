package request

type MultipleRequestParams struct {
	jsonValue map[string]interface{}
}

func NewMultipleRequestParams(requests ...TapoRequest) MultipleRequestParams {
	return MultipleRequestParams{
		jsonValue: map[string]interface{}{
			"requests": requests,
		},
	}
}

func (m MultipleRequestParams) GetJsonValue() map[string]interface{} {
	return m.jsonValue
}
