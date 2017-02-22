package drivers

type Driver interface {
	Columns(string) (map[string]string, error)
	Count(string) (int, error)
	ForeignKeys() (map[string]string, error)
	Host() string
	Indexes() (map[string]map[string]string, error)
	Tables() ([]string, error)
	ViewDefinition(string) (string, error)
	Views() ([]string, error)
}
