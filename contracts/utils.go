package contracts

type IHashUtils interface {
	Hash(payload string) string
	HashCheck(hashed string, payload string) (bool, error)
}

type IJWTUtils interface {
	GenerateToken(userID uint, lifeSpan int, secretKey string) (string, error)
	ValidateToken(token string, secretKey string) bool
	ExtractPayloadFromToken(token string, secretKey string) (map[string]interface{}, error)
}
