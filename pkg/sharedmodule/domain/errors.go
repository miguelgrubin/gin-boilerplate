package domain

type InvalidRefreshToken struct{}

func (p *InvalidRefreshToken) Error() string {
	return "Invalid refresh token"
}
