package application

type AccountRequest struct {
	Document string `json:"document" validate:"required,numeric"`
	Name     string `json:"name" validate:"required,min=5"`
	Email    string `json:"email" validate:"required,email"`
	Currency string `json:"currency" validate:"required,len=3"`
	Age      int    `json:"age" validate:"required,min=16"`
}

type AccountBalanceResponse struct {
	Account    string  `json:"account"`
	LastUpdate string  `json:"last_update"`
	Balance    float64 `json:"balance"`
}
