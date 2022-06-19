package config

type CreateUserRequest struct {
	Username 		string `json:"username"`
	Email 			string `json:"email"`
}

type CreateNoteRequest struct {
	Data 			string `json:"data"`
}