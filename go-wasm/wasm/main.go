package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strings"
	"syscall/js"

	utils "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-wasm/utils"
)

// Try to connect to the server and do a ping request
func connectToServer(this js.Value, args []js.Value) interface{} {
	go func() {
		resp, err := http.Get("http://127.0.0.1:" + os.Getenv("LOAD_BALANCER_PORT") + "/api/v1/ping")
		if err != nil {
			fmt.Printf("GET request failed: %s\n", err)
			return
		}
		defer resp.Body.Close()
		// Print the response status code and a success message
		fmt.Printf("Server is up and running with status code: %d and message: %s\n", resp.StatusCode, string(utils.ReadResponseBody(resp.Body)))
	}()
	return nil
}

func registerUserHTTP(this js.Value, args []js.Value) interface{} {
	go func() {
		if len(args) != 2 {
			fmt.Println("Invalid number of arguments passed")
			js.Global().Call("regUserError")
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
			js.Global().Call("regUserError")
			return
		}
		// Make a POST request to the server with the JSON payload
		resp, err := http.Post("http://127.0.0.1:8000/api/v1/register-user", "application/json", strings.NewReader(string(data)))
		if err != nil {
			fmt.Printf("POST request failed: %s\n", err)
			js.Global().Call("regUserError")
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != 201 {
			fmt.Printf("User registration failed, Username already exists\n")
			js.Global().Call("regUserError")
			return
		}
		// Print the response status code and a success message
		fmt.Printf("User registered with status code: %d\n", resp.StatusCode)
		js.Global().Call("regUserSuccess")
	}()
	return nil
}

func loginUserHTTP(this js.Value, args []js.Value) interface{} {
	go func() {
		if len(args) != 2 {
			fmt.Println("Invalid number of arguments passed")
			js.Global().Call("loginError")
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
			js.Global().Call("loginError")
			return
		}
		// Do something for Private and Public key here
		respPrecheck, err := http.Post("http://127.0.0.1:8000/api/v1/login-precheck", "application/json", strings.NewReader(string(dataPrecheck)))
		if err != nil {
			fmt.Printf("POST request failed: %s\n", err)
			js.Global().Call("loginError")
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
			js.Global().Call("loginError")
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
			js.Global().Call("loginError")
			return
		}
		// Make a POST request to the server with the JSON payload
		resp, err := http.Post("http://127.0.0.1:8000/api/v1/login-user", "application/json", strings.NewReader(string(data)))
		if err != nil {
			fmt.Printf("POST request failed: %s\n", err)
			js.Global().Call("loginError")
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			fmt.Printf("User login failed, Invalid credentials\n")
			js.Global().Call("loginError")
			return
		}
		// Print the response status code and a success message
		fmt.Printf("User successfully logged in with status code: %d\n", resp.StatusCode)

		bodyLogin := utils.ReadResponseBody(resp.Body)
		// Unmarshal the response body into a map
		var resultLogin map[string]interface{}
		err = json.Unmarshal(bodyLogin, &resultLogin)
		if err != nil {
			fmt.Printf("Error unmarshaling JSON: %s\n", err)
			js.Global().Call("loginError")
			return
		}
		// Get the token from the map
		token := resultLogin["token"].(string)
		// Store the token in the browser's local storage
		// js.Global().Get("localStorage").Call("setItem", "token", token)
		fmt.Printf("Token: %s\n", token)
		js.Global().Call("loginSuccess", token)
	}()
	return nil
}

func getImageURL(this js.Value, args []js.Value) interface{} {
	go func() {
		token := args[0].String()
		choice := args[1].String()
		privateKeyBytes, _ := hex.DecodeString("9c285f0cc6dbe2a3ef7db9cce7d64045bf38b150b430448c2bd0d421034ae915")
		// Convert the private key bytes to an ECDSA private key
		privKey := new(ecdsa.PrivateKey)
		privKey.Curve = elliptic.P256()
		privKey.D = new(big.Int).SetBytes(privateKeyBytes)
		privKey.PublicKey.Curve = privKey.Curve
		privKey.PublicKey.X, privKey.PublicKey.Y = privKey.Curve.ScalarBaseMult(privKey.D.Bytes())
		// Create a new claims for the additional Token part
		claims := map[string]interface{}{
			"WasmSignature": "Signed by Go-WASM Client, Globe&Citizen",
		}
		// Serialize the new claims
		encodedClaims := base64.RawURLEncoding.EncodeToString([]byte(fmt.Sprintf(`{"WasmSignature":"%s"}`, claims["WasmSignature"])))
		// Hash the data to be signed
		hash := sha256.Sum256([]byte(encodedClaims))
		// Compute the ECDSA signature
		r, s, err := ecdsa.Sign(rand.Reader, privKey, hash[:])
		if err != nil {
			fmt.Println("Error signing:", err)
			return
		}
		SignedToken := fmt.Sprintf(".%s.%s", base64.RawURLEncoding.EncodeToString(r.Bytes()), base64.RawURLEncoding.EncodeToString(s.Bytes()))
		DoubleSignedToken := fmt.Sprintf("%s%s", token, SignedToken)
		fmt.Printf("DoubleSignedToken: %s\n", DoubleSignedToken)

		choicePayload := struct {
			Choice string `json:"choice"`
			Token  string `json:"token"`
		}{
			Choice: choice,
			Token:  DoubleSignedToken,
		}
		// Marshal the payload to JSON
		dataChoice, err := json.Marshal(choicePayload)
		if err != nil {
			fmt.Printf("Error marshaling JSON: %s\n", err)
			js.Global().Call("loginError")
			return
		}
		// Make a POST request to the server with the JSON payload
		respChoice, err := http.Post("http://127.0.0.1:8000/api/v1/get-content", "application/json", strings.NewReader(string(dataChoice)))
		if err != nil {
			fmt.Printf("POST request failed: %s\n", err)
			js.Global().Call("loginError")
			return
		}
		defer respChoice.Body.Close()
		if respChoice.StatusCode == 401 {
			fmt.Printf("User not authorized\n")
			js.Global().Call("notAuthorized")
			return
		}
		// Read the response body
		bodyChoice := utils.ReadResponseBody(respChoice.Body)
		// Unmarshal the response body into a map
		mapData := make(map[string]interface{})
		err = json.Unmarshal(bodyChoice, &mapData)
		if err != nil {
			fmt.Printf("Error unmarshaling JSON: %s\n", err)
			js.Global().Call("loginError")
			return
		}
		ImgURL := mapData["url"].(string)
		js.Global().Call("displayImage", ImgURL)
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
	// Register the getImageUrl function to the global namespace
	js.Global().Set("getImageURL", js.FuncOf(getImageURL))
	// Keep the program running
	<-make(chan bool)
}
