// Package domain contains custom error types and definitions.
package domain

type InvalidRefreshToken struct{}

func (p *InvalidRefreshToken) Error() string {
	return "Invalid refresh token"
}
