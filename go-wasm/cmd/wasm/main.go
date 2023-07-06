package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"syscall/js"

	utils "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-wasm/cmd/utils"
)

// Try to connect to the server and do a ping request
func connectToServer(this js.Value, args []js.Value) interface{} {
	go func() {
		// Make a GET request to the server
		resp, err := http.Get("http://127.0.0.1:3000/api/v1/ping")
		if err != nil {
			fmt.Printf("GET request failed: %s\n", err)
			return
		}
		defer resp.Body.Close()
		// Print the response status code and a success message
		fmt.Printf("Response status code: %d\n", resp.StatusCode)
		fmt.Printf("Response body: %s\n", string(utils.ReadResponseBody(resp.Body)))
	}()
	return nil
}

func registerUserHTTP(this js.Value, args []js.Value) interface{} {
	go func() {
		if len(args) != 2 {
			fmt.Println("Invalid number of arguments passed")
			return
		}
		username := args[0].String()
		password := args[1].String()
		// email := args[2].String()
		// Generate a random salt
		rmSalt := utils.GenerateRandomSalt(utils.SaltSize)
		// Hash the salted password using bcrypt
		// hashedPassword, err := utils.HashPassword(password)
		// Salt the hashed password using the random salt
		saltedPassword := utils.SaltPassword(password, []byte(rmSalt))
		// Create a JSON payload with name and age
		payload := struct {
			// Email    string `json:"email"`
			Username string `json:"username"`
			Password string `json:"password"`
			Salt     string `json:"salt"`
		}{
			// Email:    email,
			Username: username,
			Password: saltedPassword,
			Salt:     rmSalt,
		}
		// Marshal the payload to JSON
		data, err := json.Marshal(payload)
		if err != nil {
			fmt.Printf("Error marshaling JSON: %s\n", err)
			return
		}
		// Make a POST request to the server with the JSON payload
		resp, err := http.Post("http://127.0.0.1:3000/api/v1/register-user", "application/json", strings.NewReader(string(data)))
		if err != nil {
			fmt.Printf("POST request failed: %s\n", err)
			return
		}
		defer resp.Body.Close()
		// Print the response status code and a success message
		fmt.Printf("Response status code: %d\n", resp.StatusCode)
		fmt.Printf("Successfully registered user\n")
	}()
	return nil
}

func loginUserHTTP(this js.Value, args []js.Value) interface{} {
	go func() {
		if len(args) != 2 {
			fmt.Println("Invalid number of arguments passed")
			return
		}
		username := args[0].String()
		password := args[1].String()
		// Get the user salt from the database
		payloadPrecheck := struct {
			Username string `json:"username"`
		}{
			Username: username,
		}
		// Marshal the payload to JSON
		dataPrecheck, err := json.Marshal(payloadPrecheck)
		if err != nil {
			fmt.Printf("Error marshaling JSON: %s\n", err)
			return
		}
		respPrecheck, err := http.Post("http://127.0.0.1:3000/api/v1/login-precheck", "application/json", strings.NewReader(string(dataPrecheck)))
		if err != nil {
			fmt.Printf("POST request failed: %s\n", err)
			return
		}
		defer respPrecheck.Body.Close()
		// Read the response body
		body := utils.ReadResponseBody(respPrecheck.Body)
		// Unmarshal the response body into a map
		var result map[string]interface{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Printf("Error unmarshaling JSON: %s\n", err)
			return
		}
		// Get the salt from the map
		salt := result["salt"].(string)
		fmt.Printf("Salt: %s\n", salt)
		// Salt the password using the salt from the database
		saltedPassword := utils.SaltPassword(password, []byte(salt))
		// Create a JSON payload with name and age
		payloadLogin := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{
			Username: username,
			Password: saltedPassword,
		}
		// Marshal the payload to JSON
		data, err := json.Marshal(payloadLogin)
		if err != nil {
			fmt.Printf("Error marshaling JSON: %s\n", err)
			return
		}
		// Make a POST request to the server with the JSON payload
		resp, err := http.Post("http://127.0.0.1:3000/api/v1/login-user", "application/json", strings.NewReader(string(data)))
		if err != nil {
			fmt.Printf("POST request failed: %s\n", err)
			return
		}
		defer resp.Body.Close()
		// Print the response status code and a success message
		fmt.Printf("Response status code: %d\n", resp.StatusCode)
		fmt.Printf("Successfully logged in user\n")
	}()
	return nil
}

func main() {
	fmt.Println("Go Web Assembly Demo")
	// Register the connectToServer function to the global namespace
	js.Global().Set("connectToServer", js.FuncOf(connectToServer))
	// Register the registerUser function to the global namespace
	js.Global().Set("registerUser", js.FuncOf(registerUserHTTP))
	// Register the loginUser function to the global namespace
	js.Global().Set("loginUser", js.FuncOf(loginUserHTTP))
	// Keep the program running
	<-make(chan bool)
}
