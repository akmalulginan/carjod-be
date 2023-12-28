package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/akmalulginan/carjod-be/domain"
	"github.com/akmalulginan/carjod-be/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthCustomClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

// jwt service
type JWTServiceInterface interface {
	GenerateToken(domain.User) string
	ValidateToken(token string) (*jwt.Token, error)
	Claims(ctx *gin.Context) AuthCustomClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

// auth-jwt
func NewJWTAuthService() JWTServiceInterface {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "break-social-media",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (s *jwtService) GenerateToken(user domain.User) string {
	claims := &AuthCustomClaims{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			Issuer:   s.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}

func (s *jwtService) Claims(ctx *gin.Context) (claims AuthCustomClaims) {
	const BEARER_SCHEMA = "Bearer"
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA)+1:]

	token, _ := s.ValidateToken(tokenString)

	byt, _ := json.Marshal(token.Claims)
	json.Unmarshal(byt, &claims)

	return claims
}

func Authorize(s JWTServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Cannot Get resources", "error": "Need Token to Get Resources"})
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA)+1:]
		token, err := s.ValidateToken(tokenString)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Cannot Get resources", "error": "Unauthorized"})
			return
		}

		if !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Cannot Get resources", "error": "Unauthorized"})
			return
		}

		claim := s.Claims(ctx)

		ctx.Set(string(utils.KeyUserId), claim.UserId)
	}
}
