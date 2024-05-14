package jwtauth

import (
	"context"
	"errors"
	"math"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/neabparinya11/Golang-Project/pkg/utils"
	"google.golang.org/grpc/metadata"
)

type (
	AuthFactory interface {
		SignToken() string
	}

	Claims struct {
		PlayerId string `json:"player_id"`
		RoleCode int    `json:"role_code"`
	}

	AuthMapClaim struct {
		*Claims
		jwt.RegisteredClaims
	}

	AuthConcrete struct {
		Secret []byte
		Claims *AuthMapClaim `json:"claims"`
	}

	AccessToken struct {
		*AuthConcrete
	}

	RefreshToken struct {
		*AuthConcrete
	}

	ApiKey struct {
		*AuthConcrete
	}
)

func now() time.Time {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return time.Now().In(loc)
}

func (a *AuthConcrete) SignToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.Claims)
	ss, _ := token.SignedString(a.Secret)
	return ss
}

// t is a second unit
func JwtTimeDurationCal(t int64) *jwt.NumericDate {
	return jwt.NewNumericDate(utils.LocalTime().Add(time.Duration(t * int64(math.Pow10(9)))))
}
func JwtTimeRepeaterAdapter(t int64) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Unix(t, 0))
}

func NewAccessToken(secret string, expiredAt int64, claims *Claims) AuthFactory {
	return &AccessToken{
		AuthConcrete: &AuthConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaim{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "neab",
					Subject:   "access-token",
					Audience:  []string{"neabparinya.com"},
					ExpiresAt: JwtTimeDurationCal(expiredAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
}

func NewRefreshToken(secret string, expiredAt int64, claims *Claims) AuthFactory {
	return &RefreshToken{
		AuthConcrete: &AuthConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaim{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "neab",
					Subject:   "refresh-token",
					Audience:  []string{"neabparinya.com"},
					ExpiresAt: JwtTimeDurationCal(expiredAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
}

func ReloadToken(secret string, expiredAt int64, claims *Claims) string {
	obj := &RefreshToken{
		AuthConcrete: &AuthConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaim{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "neab",
					Subject:   "refresh-token",
					Audience:  []string{"neabparinya.com"},
					ExpiresAt: JwtTimeRepeaterAdapter(expiredAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}

	return obj.SignToken()
}

func NewApiKey(secret string) AuthFactory {
	return &ApiKey{
		AuthConcrete: &AuthConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaim{
				Claims: &Claims{},
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "neab",
					Subject:   "api-key",
					Audience:  []string{"neabparinya.com"},
					ExpiresAt: JwtTimeDurationCal(31560000),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
}

func ParseToken(secret string, tokenString string) (*AuthMapClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthMapClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error: Unexpected sign method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("error: Invalid format")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("error: Token is expired")
		} else {
			return nil, errors.New("error: Token invalid")
		}
	}

	if claim, ok := token.Claims.(*AuthMapClaim); ok {
		return claim, nil
	} else {
		return nil, errors.New("error: Claim type invalid")
	}
}

var apiKeyInstant string
var once sync.Once

func SetApiKey(secret string) {
	once.Do(func() {
		apiKeyInstant = NewApiKey(secret).SignToken()
	})
}
func GetApiKey() string {
	return apiKeyInstant
}

func SetApiKeyInContext(pctx *context.Context) {
	*pctx = metadata.NewOutgoingContext(*pctx, metadata.Pairs("auth", apiKeyInstant))
}
