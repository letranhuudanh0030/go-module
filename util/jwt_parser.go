package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
	"todo/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type TokenData struct {
	Username  string
	Useragent string
	IPAdress  string
	Createdat int64
	Expires   int64
	IsAdmin   bool
	Avatar    string
	Birthday  string
	Gender    string
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("x-csv-token")
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 && onlyToken[0] == "Bearer" {
		return onlyToken[1]
	}

	return ""
}

func VerifyToken(c *fiber.Ctx, authType string) (*jwt.Token, error) {
	tokenString := extractToken(c)

	if len(tokenString) == 0 {
		msg := config.TOKEN_INCORRECT
		return nil, errors.New(msg)
	}

	if authType != "sso" {
		token, err := jwt.Parse(tokenString, jwtKeyFunc)
		if err != nil {
			return nil, err
		}
		return token, nil
	} else {
		ssoToken, err := VerifyTokenPublicKey(c, tokenString)
		if err != nil {
			return nil, err
		}
		return ssoToken, nil
	}

}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(config.Get("JWT_SECRET_KEY")), nil
}

func ExtractTokenData(c *fiber.Ctx) (*TokenData, error) {
	token, err := VerifyToken(c, "internal") // set temp demo
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return &TokenData{
			Username:  fmt.Sprint(claims["username"]),
			Createdat: int64((claims["iat"]).(float64)),
			Expires:   int64((claims["exp"]).(float64)),
		}, nil
	}

	return nil, err
}

// func ExtractTokenData(c *fiber.Ctx) (*TokenData, error) {
// 	authType := c.Get("X-Auth-Type")

// 	token, err := VerifyToken(c, authType)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Setting and checking token and credentials.
// 	if token.Valid {
// 		if authType != "sso" {
// 			claims := token.Claims.(jwt.MapClaims)
// 			//--- INTERNAL TOKEN
// 			return &TokenData{
// 				Username:  fmt.Sprint(claims["username"]),
// 				Useragent: fmt.Sprint(claims["useragent"]),
// 				IPAdress:  fmt.Sprint(claims["ipaddress"]),
// 				Createdat: int64((claims["createdat"]).(float64)),
// 				Expires:   int64((claims["exp"]).(float64)),
// 				IsAdmin:   bool(claims["isAdmin"].(bool)),
// 			}, nil
// 		} else {
// 			claims := token.Claims.(*SSOClaims)
// 			//check is admin
// 			c.Locals("user", fmt.Sprint(claims.UserName))
// 			var userSSO *userModel.UserSSO
// 			userSSO, errSSO := CheckPermission(c)
// 			if errSSO != nil {
// 				return nil, errSSO
// 			}
// 			isAdmin := userSSO.IsAdmin
// 			//--- SSO TOKEN
// 			return &TokenData{
// 				Username:  fmt.Sprint(claims.UserName),
// 				Createdat: claims.StandardClaims.IssuedAt,
// 				Expires:   claims.StandardClaims.ExpiresAt,
// 				IsAdmin:   isAdmin,
// 				Avatar:    userSSO.UserFile,
// 				Birthday:  userSSO.Birthday,
// 				Gender:    userSSO.Gender,
// 			}, nil
// 		}
// 	}

// 	return nil, err
// }

type multiString string

type SSOClaims struct {
	Audience multiString `json:"aud,omitempty"`
	jwt.StandardClaims
	GiveName   string `json:"given_name,omitempty"`
	UserName   string `json:"username,omitempty"`
	EmployeeId string `json:"employeeID,omitempty"`
}

func (c *SSOClaims) Valid() error {
	c.StandardClaims.IssuedAt = time.Now().Unix()
	// c.StandardClaims.ExpiresAt = time.Now().Unix()
	valid := c.StandardClaims.Valid()
	return valid
}

func (ms *multiString) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		switch data[0] {
		case '"':
			var s string
			if err := json.Unmarshal(data, &s); err != nil {
				return err
			}
			*ms = multiString(s)
		case '[':
			var s []string
			if err := json.Unmarshal(data, &s); err != nil {
				return err
			}
			*ms = multiString(strings.Join(s, ","))
		}
	}
	return nil
}

func VerifyTokenPublicKey(c *fiber.Ctx, tokenString string) (*jwt.Token, error) {
	PublicKey := "-----BEGIN CERTIFICATE-----\n" + config.Get("SSO_PUBLIC_KEY") + "\n-----END CERTIFICATE-----"
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(PublicKey))
	if err != nil {
		fmt.Println(err)
		msg := config.UNAUTHORIZED
		return nil, errors.New(msg)
	}

	token, err := jwt.ParseWithClaims(tokenString, &SSOClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil || !token.Valid {
		msg := config.UNAUTHORIZED
		// msg := err.Error()
		return nil, errors.New(msg)
	}

	return token, nil
}

/** Mobile */

type TokenDataMobile struct {
	Username   string
	Useragent  string
	IPAdress   string
	Permission int
	Createdat  int64
	Expires    int64
}

func ExtractTokenDataMobile(c *fiber.Ctx) (*TokenDataMobile, error) {
	authType := c.Get("X-Auth-Type")

	token, err := VerifyToken(c, authType)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return &TokenDataMobile{
			Username:   fmt.Sprint(claims["username"]),
			Useragent:  fmt.Sprint(claims["useragent"]),
			IPAdress:   fmt.Sprint(claims["ipaddress"]),
			Permission: int(claims["permission"].(float64)),
			Createdat:  int64((claims["createdat"]).(float64)),
			Expires:    int64((claims["exp"]).(float64)),
		}, nil
	}

	return nil, err
}

/**
EncodedData
*/

type Data struct {
	EmployeeId    string `json:"employee_id"`
	DateInSeconds int64  `json:"date_in_seconds"`
	Coordinates   string `json:"coordinates"`
	ShiftId       int64  `json:"shift_id"`
}

func DecodeData(encodedData string) (*Data, error) {
	if len(encodedData) == 0 {
		return nil, errors.New("data is incorrect")
	}

	data, err := jwt.Parse(encodedData, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Get("JWT_DATA_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := data.Claims.(jwt.MapClaims)
	if ok && data.Valid {
		return &Data{
			EmployeeId:    fmt.Sprint(claims["employee_id"]),
			DateInSeconds: int64((claims["date_in_seconds"]).(float64)),
			Coordinates:   fmt.Sprint(claims["coordinates"]),
			ShiftId:       int64((claims["shift_id"]).(float64)),
		}, nil
	}

	return nil, err
}
