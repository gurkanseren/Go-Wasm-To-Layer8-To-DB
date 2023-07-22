# Go-Wasm-To-Layer8-To-DB

A test project for experimenting the connection between a Go standard WASM module > Go Standard HTTP Layer8 Server > PostgreSQL DB

## Prerequisites

- Go 1.11 or higher
- PostgreSQL 9.6 or higher

## How to run

### Simple Usage

Open a terminal and run the following command:

```bash
make run # This will generate the builds for layer8 and wasm and start the single port multi-server
```

**Note:** The `make build-wasm` command will generate the WASM module and the `make build-layer8` command will generate the Layer8 server build, and `make run-server` will start the multi-server which serve WASM Module on the given Port and also the Layer8 APIs on the same Port (this is made on single port in order to solve the CORS issues and WASM to Websocket support). Make sure that before testing the connection between the WASM module and the Layer8 server, you should have local db set up according to the env variables that you will set. Use the SQL given below for generating the table in a local PostgreSQL DB.

```sql
CREATE TABLE IF NOT EXISTS public.users
(
    id SERIAL PRIMARY KEY NOT NULL,
    username character varying NOT NULL,
    password character varying NOT NULL,
    salt character varying(255)
);
```

Go to `http://localhost:8080` and open the browser console. You should see the following output:

```bash
Go Web Assemly Demo
```

Here you can test the connection my using the following commands:

```bash
connectToServer(); // Make a ping request to the layer8 server to test the connection
```
- Expected Output:
```json
Response Status Code: 200
Response Body: Ping successful
```

```bash
registerUser('Name', 'Password'); // Register a user
```

```bash
loginUser('Name', 'Password'); // Login a user
```
