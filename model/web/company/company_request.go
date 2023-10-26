package company

type Request struct {
	Name string `validate:"required,max=100" json:"name"`
}
