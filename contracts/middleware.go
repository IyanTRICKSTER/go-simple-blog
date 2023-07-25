package contracts

const AuthContex = "auth-ctx"

type IAuthenticatedRequest interface {
	GetUserID() uint
	GetUsername() string
	Validate() (bool, error)
}
