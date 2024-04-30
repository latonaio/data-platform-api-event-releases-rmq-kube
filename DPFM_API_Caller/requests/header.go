package requests

type Header struct {
	Event		int		`json:"Event"`
	IsReleased	*bool	`json:"IsReleased"`
}
