package storage

type Storage interface {
	Initialize() error
	GetWorkDays(year string, month string) ([]map[string]interface{}, error)
	MarkWorking(year string, month string) error
}
