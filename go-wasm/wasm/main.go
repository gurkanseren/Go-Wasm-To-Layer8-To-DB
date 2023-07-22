package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"syscall/js"
	"time"

	utils "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-wasm/utils"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// Try to connect to the server and do a ping request
func connectToServer(this js.Value, args []js.Value) interface{} {
	go func() {
		// Make a GET request to the server
		resp, err := http.Get("http://127.0.0.1:8080/api/v1/ping")
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
		// Generate a random salt
		rmSalt := utils.GenerateRandomSalt(utils.SaltSize)
		HashedAndSaltedPassword := utils.SaltAndHashPassword(password, rmSalt)
		// Create a JSON payload with name and age
		payload := struct {
			// Email    string `json:"email"`
			Username string `json:"username"`
			Password string `json:"password"`
			Salt     string `json:"salt"`
		}{
			// Email:    email,
			Username: username,
			Password: HashedAndSaltedPassword,
			Salt:     rmSalt,
		}
		// Marshal the payload to JSON
		data, err := json.Marshal(payload)
		if err != nil {
			fmt.Printf("Error marshaling JSON: %s\n", err)
			return
		}
		// Make a POST request to the server with the JSON payload
		resp, err := http.Post("http://127.0.0.1:8080/api/v1/register-user", "application/json", strings.NewReader(string(data)))
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
		respPrecheck, err := http.Post("http://127.0.0.1:8080/api/v1/login-precheck", "application/json", strings.NewReader(string(dataPrecheck)))
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
		// Salt the password using the salt from the database
		HashedAndSaltedPassword := utils.SaltAndHashPassword(password, salt)
		// Create a JSON payload with name and age
		payloadLogin := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{
			Username: username,
			Password: HashedAndSaltedPassword,
		}
		// Marshal the payload to JSON
		data, err := json.Marshal(payloadLogin)
		if err != nil {
			fmt.Printf("Error marshaling JSON: %s\n", err)
			return
		}
		// Make a POST request to the server with the JSON payload
		resp, err := http.Post("http://127.0.0.1:8080/api/v1/login-user", "application/json", strings.NewReader(string(data)))
		if err != nil {
			fmt.Printf("POST request failed: %s\n", err)
			return
		}
		defer resp.Body.Close()
		// Print the response status code and a success message
		fmt.Printf("Response status code: %d\n", resp.StatusCode)
		fmt.Printf("Response body: %s\n", string(utils.ReadResponseBody(resp.Body)))
		utils.UpgradeConnToWebSocket()
	}()
	return nil
}

func connectToWebSocket(this js.Value, args []js.Value) interface{} {
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		c, _, err := websocket.Dial(ctx, "ws://localhost:8080/ws", nil)
		if err != nil {
			log.Fatal(err)
		}
		defer c.Close(websocket.StatusInternalError, "the sky is falling")

		err = wsjson.Write(ctx, c, "hi")
		if err != nil {
			log.Fatal(err)
		}

		c.Close(websocket.StatusNormalClosure, "")
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
	// Register the connectToWebSocket function to the global namespace
	js.Global().Set("connectToWebSocket", js.FuncOf(connectToWebSocket))
	// Keep the program running
	<-make(chan bool)
}
