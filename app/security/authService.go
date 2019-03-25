package security

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"intelliq/app/common"
	utility "intelliq/app/common"
	"net/http"
	"strings"
	"time"

	"github.com/gbrlsnchs/jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var skipURIS = []string{"login", "forgot"}

//EnableSecurity enables app security
func EnableSecurity(router *gin.Engine) {
	router.Use(authenticateRequest())
	router.Use(enableCors())
}

func enableCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"PUT", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:           []string{"Origin", "Content-Type", "X-Requested-With", "Accept"},
		AllowCredentials:       false,
		MaxAge:                 12 * time.Hour,
		AllowBrowserExtensions: true,
	})
}

func authenticateRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestURL := ctx.Request.RequestURI
		for _, uri := range skipURIS {
			if strings.Contains(requestURL, uri) {
				ctx.Next()
				return
			}
		}
		sessionToken, err := ctx.Cookie(common.COOKIE_SESSION_KEY)
		if err != nil {
			fmt.Println("COOKIE FETCH ERROR: ", err)
			ctx.AbortWithStatusJSON(http.StatusForbidden,
				common.GetErrorResponse(common.MSG_USER_SESSION_ERROR))
			return
		}
		if len(sessionToken) == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden,
				common.GetErrorResponse(common.MSG_USER_SESSION_ERROR))
			return
		}
		isSessionOK, status := VerifyToken(sessionToken)
		if !isSessionOK {
			ctx.AbortWithStatusJSON(http.StatusForbidden,
				common.GetErrorResponse(common.MSG_USER_AUTH_ERROR+"\n"+status))
			return
		}
		ctx.Next()
	}
}

//GenerateToken generates JWT token
func GenerateToken(subject, val string, expiry int) string {
	privKey := getPrivateKey()
	if privKey != nil {
		now := time.Now()
		ecdsa512 := jwt.NewECDSA(jwt.SHA256, privKey, &privKey.PublicKey)
		jot := &jwt.JWT{
			Claims: &jwt.Claims{
				Issuer:         common.APP_NAME,
				Subject:        subject,
				ExpirationTime: now.Add(time.Duration(expiry) * time.Minute).Unix(),
				ID:             val,
				IssuedAt:       now.Unix(),
			},
		}
		token, err := jwt.Sign(jot, ecdsa512)
		if err != nil {
			fmt.Println("TOKEN GEN ERR = ", err)
		}
		return string(token)
	}
	return ""
}

func getPrivateKey() *ecdsa.PrivateKey {
	privKeyString := utility.ReadFile(common.PRIVATE_KEY_FILEPATH)
	if privKeyString == nil {
		fmt.Println("Could not fetch private key from file")
		return nil
	}
	privKey := decodePriv(privKeyString)
	if privKey == nil {
		fmt.Println("Could not generate private key")
		return nil
	}
	return privKey
}

func decodePriv(pemEncoded []byte) *ecdsa.PrivateKey {
	block, _ := pem.Decode(pemEncoded)
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)
	return privateKey
}

//VerifyToken verifies token against key,expiry time
func VerifyToken(token string) (bool, string) {
	privKey := getPrivateKey()
	if privKey != nil {
		now := time.Now()
		var ecdsa512 *jwt.ECDSA
		ecdsa512 = jwt.NewECDSA(jwt.SHA256, privKey, &privKey.PublicKey)
		raw, err := jwt.Verify([]byte(token), ecdsa512)
		if err != nil {
			return false, "Invalid Token"
		}
		var jot jwt.JWT
		if err = raw.Decode(&jot); err != nil {
			fmt.Println(err)
			return false, "Corrupt Token"
		}
		expValidator := jwt.ExpirationTimeValidator(now)
		issuerValidator := jwt.IssuerValidator(common.APP_NAME)
		issuedAtValidator := jwt.IssuedAtValidator(now)
		if err := jot.Validate(expValidator, issuerValidator, issuedAtValidator); err != nil {
			var status string
			switch err {
			case jwt.ErrIssValidation:
				status = "Unidentified Issuer"
			case jwt.ErrExpValidation:
				status = "Token Expired"
			case jwt.ErrIatValidation:
				status = "Pre-Issued token"
			}
			return false, status
		}
		return true, ""

	}
	return false, ""
}

//SetCookie sets cookie attribute
func SetCookie(ctx *gin.Context, body string, expiry int) {
	ctx.SetCookie(common.COOKIE_SESSION_KEY, body,
		expiry*60, "", "localhost",
		true, true)
}

//RemoveCookie removes cookie attribute
func RemoveCookie(ctx *gin.Context) {
	ctx.SetCookie(common.COOKIE_SESSION_KEY, "",
		-1, "", "localhost",
		true, true)
}
