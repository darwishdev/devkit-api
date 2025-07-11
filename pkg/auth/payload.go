package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	ID                uuid.UUID `json:"id"`
	TenantId          int32     `json:"tenant_id"`
	UserId            int32     `json:"user_id"`
	UserSecurityLevel int32     `json:"user_security_level"`
	Username          string    `json:"username"`
	TokenType         string    `json:"token_type"`
	IssuedAt          time.Time `json:"issued_at"`
	ExpiredAt         time.Time `json:"expired_at"`
}
type RefreshPayload struct {
	ID        uuid.UUID `json:"id"`
	UserId    int32     `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, userId int32, userSecurityLevel int32, tenantId int32, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:                tokenId,
		UserSecurityLevel: userSecurityLevel,
		TenantId:          tenantId,
		UserId:            userId,
		Username:          username,
		IssuedAt:          time.Now(),
		ExpiredAt:         time.Now().Add(duration),
	}

	return payload, nil
}
func NewRefreshPayload(userId int32, duration time.Duration) (*RefreshPayload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &RefreshPayload{
		ID:        tokenId,
		UserId:    userId,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func (payload *RefreshPayload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
