package request

type TodoUpdateRequest struct {
	Id   int
	Name string `validate:"require min=1, max=100" json:"name"`
}
