package model

type ValidatorResult struct {
	IsPass  bool
	Error   string `json:"error"`
	Message string `json:"message"`
}
