package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/config"
	"github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/models"
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
}

// LoginPrecheckHandler handles login precheck requests and get the salt of the user from the database using the username from the request URL
func loginPrecheckHandler(w http.ResponseWriter, r *http.Request) {
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

func loginUserHandler(w http.ResponseWriter, r *http.Request) {
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
	// hashedPassword, err := HashPassword(req.Password)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	_, err := w.Write([]byte(err.Error()))
	// 	if err != nil {
	// 		log.Printf("Error sending response: %v", err)
	// 	}
	// 	return
	// }
	// Compare the password with the password in the database
	if user.Password != req.Password {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Invalid credentials"))
		if err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
	// Generate JWT token
	fmt.Println("User password: ", user.Password)
	fmt.Println("Salted password: ", req.Password)
	token, _ := generateJWTToken(user.ID)
	fmt.Println("Token: ", token)
	// TODO: Send token to client
}

func generateJWTToken(userID uint) (string, error) {
	// TODO: Implement JWT token generation
	// Return dummy token for now
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjF9.8Za0Z", nil
}

// func HashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	return string(bytes), err
// }
