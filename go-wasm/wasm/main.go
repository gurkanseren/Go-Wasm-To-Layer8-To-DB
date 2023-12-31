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

	"github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-wasm/models"
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
		privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		// Serialize the private key
		privKeyBytes := privKey.D.Bytes()
		// Convert the private key bytes to a hex string
		privKeyHex := hex.EncodeToString(privKeyBytes)
		// Generate a public key from the private key
		PubKeyHex := utils.GenPubKeyHex(privKeyHex)
		// Store the private key in the browser's local storage
		// js.Global().Get("localStorage").Call("setItem", "privKey", privKeyHex)
		// Store the private key in the browser's memory
		js.Global().Call("makePrivKeyInMemory", privKeyHex)
		payloadLogin := struct {
			Username string `json:"username"`
			Password string `json:"password"`
			PubKey   string `json:"public_key"`
		}{
			Username: username,
			Password: HashedAndSaltedPassword,
			PubKey:   PubKeyHex,
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
		// Store the token in the browser's memory
		js.Global().Call("loginSuccess", token)
		fmt.Printf("Token: %s\n", token)
	}()
	return nil
}

func getImageURL(this js.Value, args []js.Value) interface{} {
	go func() {
		token := args[0].String()
		choice := args[1].String()
		// Get private key from the browser's local storage
		// privKeyHex := js.Global().Get("localStorage").Call("getItem", "privKey").String()
		// Get private key from the browser's memory
		privKeyHex := js.Global().Call("getPrivKeyFromMemory").String()
		// Convert the private key hex string to bytes
		privateKeyBytes, _ := hex.DecodeString(privKeyHex)
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

func initializeECDHTunnel(this js.Value, args []js.Value) interface{} {
	go func() {
		token := args[0].String()
		privKeyHex := js.Global().Call("getPrivKeyFromMemory").String()
		// Convert the private key hex string to bytes
		privateKeyBytes, _ := hex.DecodeString(privKeyHex)
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

		privKeyWasm, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		pubKeyWasmX, pubKeyWasmY := privKeyWasm.PublicKey.Curve.ScalarBaseMult(privKeyWasm.D.Bytes())
		pubKeyWasm := &ecdsa.PublicKey{
			Curve: elliptic.P256(),
			X:     pubKeyWasmX,
			Y:     pubKeyWasmY,
		}
		Payload := struct {
			Token       string   `json:"token"`
			PubKeyWasmX *big.Int `json:"pub_key_wasm_x"`
			PubKeyWasmY *big.Int `json:"pub_key_wasm_y"`
		}{
			Token:       DoubleSignedToken,
			PubKeyWasmX: pubKeyWasm.X,
			PubKeyWasmY: pubKeyWasm.Y,
		}
		// Marshal the payload to JSON
		data, err := json.Marshal(Payload)
		if err != nil {
			fmt.Printf("Error marshaling JSON: %s\n", err)
			js.Global().Call("loginError")
			return
		}
		// Make a POST request to the server with the JSON payload
		respChoice, err := http.Post("http://127.0.0.1:8000/api/v1/initialize-ecdh-tunnel", "application/json", strings.NewReader(string(data)))
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
		Respbody := utils.ReadResponseBody(respChoice.Body)

		var pubKeyXY models.ECDHKeyExchangeOutput
		err = json.Unmarshal(Respbody, &pubKeyXY)
		if err != nil {
			fmt.Printf("Error unmarshaling JSON: %s\n", err)
			js.Global().Call("loginError")
			return
		}

		pubKeyServerX := pubKeyXY.PubKeyServerX
		pubKeyServerY := pubKeyXY.PubKeyServerY

		sharedX, sharedY := elliptic.P256().ScalarMult(pubKeyServerX, pubKeyServerY, privKeyWasm.D.Bytes())
		sharedKeyWasm := &ecdsa.PublicKey{
			Curve: elliptic.P256(),
			X:     sharedX,
			Y:     sharedY,
		}
		fmt.Printf("\nShared key (Wasm) (%x, %x)\n", sharedKeyWasm.X, sharedKeyWasm.Y)
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
	// Register the initializeECDHTunnel function to the global namespace
	js.Global().Set("initializeECDHTunnel", js.FuncOf(initializeECDHTunnel))
	// Keep the program running
	<-make(chan bool)
}
