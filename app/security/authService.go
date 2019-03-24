package security

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"intelliq/app/common"
	"time"

	"github.com/gbrlsnchs/jwt"
	"github.com/gin-gonic/gin"
)

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
	privKeyString := common.ReadFile(common.PRIVATE_KEY_FILEPATH)
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
		false, true)
}

//RemoveCookie removes cookie attribute
func RemoveCookie(ctx *gin.Context) {
	ctx.SetCookie(common.COOKIE_SESSION_KEY, "",
		-1, "", "localhost",
		true, true)
}
