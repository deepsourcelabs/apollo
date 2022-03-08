package dependency

type RegisterRequest struct {
	ID   uint   `json:"id"` // TODO: probably should use a UUID
	Name string `json:"name"`
	URI  string `json:"uri"`
}
