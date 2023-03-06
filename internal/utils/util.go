package utils

import (
	"encoding/csv"
	"io"
	"time"

	"fiber-base-go/internal/model"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func ParseCSV(r io.Reader, batchSize int) ([][]*model.Student, error) {
	reader := csv.NewReader(r)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	header := rows[0]
	dataRows := rows[1:]

	// Map the header names to their column indexes
	var nameIndex, classIndex, birthdayIndex int
	for i, col := range header {
		switch col {
		case "name":
			nameIndex = i
		case "class":
			classIndex = i
		case "birthday":
			birthdayIndex = i
		}
	}

	students := make([][]*model.Student, 0, len(dataRows)/batchSize+1)
	batch := make([]*model.Student, 0, batchSize)

	for _, row := range dataRows {
		// Parse the birthday field into a time.Time value
		birthday, err := time.Parse("2006-01-02", row[birthdayIndex])
		if err != nil {
			return nil, err
		}

		student := &model.Student{
			Name:     row[nameIndex],
			Class:    row[classIndex],
			Birthday: birthday,
		}

		batch = append(batch, student)
		if len(batch) >= batchSize {
			students = append(students, batch)
			batch = make([]*model.Student, 0, batchSize)
		}
	}

	if len(batch) > 0 {
		students = append(students, batch)
	}

	return students, nil
}

func GenerateToken(email string, secretKey string, expirationTime time.Duration) (string, error) {
	// Create the claims for the JWT token
	claims := &model.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Generate the JWT token with the claims and secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(tokenString string, secretKey string) (*model.Claims, error) {
	// Parse the JWT token with the claims and secret key
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Check if the token is valid and contains the expected claims
	claims, ok := token.Claims.(*model.Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}
