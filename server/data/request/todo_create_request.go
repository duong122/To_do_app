package request

type TodoCreateRequest struct {
	Name string `validate:"require min=1, max=100" json:"name"`
}
