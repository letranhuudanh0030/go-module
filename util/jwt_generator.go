package util

import (
	"strconv"
	"time"
	"todo/config"

	"github.com/golang-jwt/jwt"
)

func GenerateAccessToken(username, userAgent, ipAddress string, isAdmin bool) (string, error) {
	secret := config.Get("JWT_SECRET_KEY")
	timeExpire := config.Get("JWT_EXPIRED_TIME")

	minutesCount, _ := strconv.Atoi(timeExpire)

	claims := jwt.MapClaims{}

	claims["username"] = username
	claims["useragent"] = userAgent
	claims["ipaddress"] = ipAddress
	claims["createdat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	claims["isAdmin"] = isAdmin

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func GenerateAccessTokenMobile(username, userAgent, ipAddress string, permission int) (string, error) {
	secret := config.Get("JWT_SECRET_KEY")
	timeExpire := config.Get("JWT_EXPIRED_TIME")

	minutesCount, _ := strconv.Atoi(timeExpire)

	claims := jwt.MapClaims{}

	claims["username"] = username
	claims["useragent"] = userAgent
	claims["ipaddress"] = ipAddress
	claims["permission"] = permission
	claims["createdat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func EncodeDataTokenMobile(employeeId string, dateInSeconds int64, coordinates string, shiftId int) (string, error) {
	secret := config.Get("JWT_DATA_SECRET_KEY")
	timeExpire := config.Get("JWT_DATA_EXPIRED_TIME")

	minutesCount, _ := strconv.Atoi(timeExpire)

	claims := jwt.MapClaims{}

	claims["employee_id"] = employeeId
	claims["date_in_seconds"] = dateInSeconds
	claims["coordinates"] = coordinates
	claims["shift_id"] = shiftId
	claims["createdat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
