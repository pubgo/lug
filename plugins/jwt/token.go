package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/pubgo/lava/logz"
)

//DefaultManager can be replaced
var DefaultManager Manager = &jwtTokenManager{}

var (
	ErrInvalidExp = errors.New("expire time is illegal")
)

var logs = logz.New(Name)

//jwt claims RFC 7519
//https://tools.ietf.org/html/rfc7519#section-4.1.2
const (
	JWTClaimsExp = "exp"
	JWTClaimsSub = "sub"
)

// SecretFunc is a callback function to supply
// the key for verification.  The function receives the parsed,
// but unverified claims in Token.  This allows you to use properties in the
// claims of the token (such as `username`) to identify which key to use.
type SecretFunc func(claims interface{}, method SigningMethod) (interface{}, error)

//Sign gen token
func Sign(claims map[string]interface{}, secret interface{}, opts ...Option) (string, error) {
	return DefaultManager.Sign(claims, secret, opts...)
}

//Verify return claims
func Verify(tokenString string, f SecretFunc, opts ...Option) (map[string]interface{}, error) {
	return DefaultManager.Verify(tokenString, f, opts...)
}

//Manager manages token
type Manager interface {
	Sign(claims map[string]interface{}, secret interface{}, option ...Option) (string, error)
	Verify(tokenString string, f SecretFunc, opts ...Option) (map[string]interface{}, error)
}
type jwtTokenManager struct {
}

//Sign signature a token
func (j *jwtTokenManager) Sign(claims map[string]interface{}, secret interface{}, opts ...Option) (string, error) {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}
	c := jwt.MapClaims(claims)
	if o.Expire != "" {
		d, err := time.ParseDuration(o.Expire)
		if err != nil {
			return "", ErrInvalidExp
		}
		claims[JWTClaimsExp] = time.Now().Add(d).Unix()
	}
	var to *jwt.Token
	switch o.SigningMethod {
	case RS256:
		to = jwt.NewWithClaims(jwt.SigningMethodRS256, c)
		return to.SignedString(secret)
	case RS512:
		to = jwt.NewWithClaims(jwt.SigningMethodRS512, c)
		return to.SignedString(secret)
	case HS256:
		to = jwt.NewWithClaims(jwt.SigningMethodHS256, c)
		return to.SignedString(secret)
	default:
		to = jwt.NewWithClaims(jwt.SigningMethodHS256, c)
		return to.SignedString(secret)
	}

}

//Verify return claims
func (j *jwtTokenManager) Verify(tokenString string, f SecretFunc, opts ...Option) (map[string]interface{}, error) {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		sm := HS256
		if m, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			if m.Name == "HS256" {
				sm = HS256
			} else if m.Name == "RS512" {
				sm = RS512
			} else if m.Name == "RS256" {
				sm = RS256
			}
		}
		return f(token.Claims, sm)
	})
	if err != nil {
		return nil, err
	}
	var ve *jwt.ValidationError
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else if ok := errors.As(err, ve); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			logs.Error("not a valid jwt")
			return nil, err
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			logs.Error("token expired")
			return nil, err
		} else {
			logs.WithErr(err).Error("parse token err")
			return nil, err
		}
	} else {
		logs.WithErr(err).Error("parse token err")
		return nil, err
	}
}
