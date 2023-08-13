package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8-slaves/config"
	"github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8-slaves/models"
	"github.com/go-playground/validator/v10"
)

// PingHandler handles ping requests
func Ping(w http.ResponseWriter, r *http.Request) {
	// Send response to client
	_, err := w.Write([]byte("ping successful"))
	if err != nil {
		log.Printf("Error sending response: %v", err)
	}
}

// RegisterUserHandler handles user registration requests
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	// Unmarshal request
	var req models.RegisterUserDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
	// validate request
	if err := validator.New().Struct(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.(*validator.InvalidValidationError).Error()))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
	// Make connection to database
	db := config.SetupDatabaseConnection()
	// Close connection database
	defer config.CloseDatabaseConnection(db)
	// Save user to database
	user := models.User{
		Username: req.Username,
		Password: req.Password,
		Salt:     req.Salt,
	}
	if err := db.Create(&user).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// LoginPrecheckHandler handles login precheck requests and get the salt of the user from the database using the username from the request URL
func LoginPrecheckHandler(w http.ResponseWriter, r *http.Request) {
	// Get username from body
	var req models.LoginPrecheckDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
	// validate request
	if err := validator.New().Struct(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.(*validator.InvalidValidationError).Error()))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
	// Make connection to database
	db := config.SetupDatabaseConnection()
	// Close connection database
	defer config.CloseDatabaseConnection(db)
	// Using the username, find the user in the database
	var user models.User
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
	resp := models.LoginPrecheckResponseDTO{
		Username: user.Username,
		Salt:     user.Salt,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
	}
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	// Unmarshal request
	var req models.LoginUserDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
	// validate request
	if err := validator.New().Struct(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.(*validator.InvalidValidationError).Error()))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
	// Make connection to database
	db := config.SetupDatabaseConnection()
	// Close connection database
	defer config.CloseDatabaseConnection(db)
	// Using the username, find the user in the database
	var user models.User
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
	// Compare the password with the password in the database
	if user.Password != req.SaltedHashedPassword {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Invalid credentials"))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
	// Get JWT_SECRET from the Layer8 Master Server
	port := os.Getenv("LAYER8_MASTER_PORT")
	respSecret, err := http.Get("http://localhost:" + port + "//api/v1/jwt-secret")
	if err != nil {
		log.Printf("failed to get picture: %v", err)
		return
	}
	defer respSecret.Body.Close()
	// Convert the response body to a string
	RespBodyByte, err := ioutil.ReadAll(respSecret.Body)
	if err != nil {
		log.Printf("failed to read response body: %v", err)
		return
	}
	// Convert RespBodyByte to string
	JWT_SECRET := []byte(string(RespBodyByte))

	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &models.Claims{
		UserName: user.Username,
		UserID:   user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "GlobeAndCitizen",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWT_SECRET)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
	}
	resp := models.LoginUserResponseDTO{
		Token: tokenString,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
	}
}

func GetContentHandler(w http.ResponseWriter, r *http.Request) {
	// Unmarshal request
	var req models.ContentReqDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
	// validate request
	if err := validator.New().Struct(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.(*validator.InvalidValidationError).Error()))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
	// Validate the JWT token
	masterPort := os.Getenv("LAYER8_MASTER_PORT")
	respSecret, err := http.Get("http://localhost:" + masterPort + "//api/v1/jwt-secret")
	if err != nil {
		log.Printf("failed to connect to master server: %v", err)
		return
	}
	defer respSecret.Body.Close()
	// Convert the response body to a string
	RespBodyByte, err := ioutil.ReadAll(respSecret.Body)
	if err != nil {
		log.Printf("failed to read response body: %v", err)
		return
	}
	// Convert RespBodyByte to string
	JWT_SECRET := []byte(string(RespBodyByte))
	// Parse the token
	token, err := jwt.ParseWithClaims(req.Token, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWT_SECRET, nil
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
	// Check if the token is valid
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Invalid token"))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
	// Get the user id from the token
	claims, ok := token.Claims.(*models.Claims)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Error getting user id"))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
	fmt.Println("User id: ", claims.UserID)
	// expiryDate := time.Unix(claims.ExpiresAt, 0)
	fmt.Println("Token expires at: ", claims.ExpiresAt)
	port := os.Getenv("CONTENT_SERVER_PORT")
	// Make request to the content server
	resp, err := http.Get("http://localhost:" + port + "/image" + "?id=" + req.Choice)
	if err != nil {
		log.Printf("failed to get picture: %v", err)
		return
	}
	defer resp.Body.Close()

	// Convert the response body to a string
	RespBodyByteImg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response body: %v", err)
		return
	}

	// Convert RespBodyByte to string
	RespBodyString := string(RespBodyByteImg)

	// Send the response back to the WASM module
	_, err = w.Write([]byte(RespBodyString))
	if err != nil {
		log.Printf("failed to send response: %v", err)
		return
	}
}