package utilService

import (
	// "os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func createJWTToken(payload map[string]interface{}, secretKey string, tokenExpiration int64) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))
    token.Claims.(jwt.MapClaims)["exp"] = tokenExpiration
    signedToken, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", err
    }
    return signedToken, nil
}

func GetToken(userId string, role string ) (string, error){
	payload := map[string]interface{}{
		"sub": "12345",        
		"iat": time.Now().Unix(),      
		"exp": time.Now().Add(time.Hour * 48).Unix(),  
		"https://hasura.io/jwt/claims": map[string]interface{}{
			"x-hasura-allowed-roles": []string{"user", "user-admin"}, 
			"x-hasura-default-role":  "user",           
			"x-hasura-user-id":       userId,           
			"x-hasura-role":          role,              
		},	
		"metadata": map[string]interface{}{
			"userId": userId,
		},
	}
	tokenExpiration := time.Now().Add(time.Hour * 48).Unix()
	token, err := createJWTToken(payload, "minab_secret_keyminab_secret_keyminab_secret_keyminab_secret_key",tokenExpiration )
	if err != nil {
		return "", err
	}
	return token, nil
}
