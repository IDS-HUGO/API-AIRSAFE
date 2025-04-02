package domain

type Admin struct {
	ID         int    `json:"id"`
	Usuario    string `json:"usuario"`
	Contrasena string `json:"contrasena"`
}
