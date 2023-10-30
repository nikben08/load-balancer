package contracts

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
}

type GoogleOAuthRequest struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Locale  string `json:"locale"`
}

type AuthResponse struct {
	Code    int
	Message string
	Token   string
}

type AuthErrorResponse struct {
	Code    int
	Message string
}
