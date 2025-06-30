package types

type Student struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name" validate:"required,min=3,max=15"`
	Age   int    `json:"age" validate:"required,gte=18,lte=100"`
	Email string `json:"email" validate:"required,email"`
}
