package token

type Token struct {
	SignKey string
}

type TokenInterface interface {
	GenerateToken(accountExternalID string) (signedToken string, err error)
	ExtractAccountIDFromToken(token string) (accountExternalID string, err error)
}
