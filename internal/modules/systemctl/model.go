package systemctl

type Service struct {
	ID            string
	ActiveState   string
	SubState      string
	UnitFileState string
	Description   string
}