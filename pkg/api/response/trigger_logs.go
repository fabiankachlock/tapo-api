package response

type TriggerLogs[T any] struct {
	StartId uint64 `json:"start_id"`
	Sum     uint64 `json:"sum"`
	Logs    []T    `json:"logs"`
}
