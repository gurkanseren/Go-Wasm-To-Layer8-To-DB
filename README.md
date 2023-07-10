# Go-Wasm-To-Layer8-To-DB

A test project for experimenting the connection between a Go standard WASM module > Go Standard HTTP Layer8 Server > PostgreSQL DB

## Prerequisites

- Go 1.11 or higher
- PostgreSQL 9.6 or higher

## How to run

### Simple Usage

Open two terminals and run the following commands:

```bash
# Terminal 1
make wasm
```

```bash
# Terminal 2
make layer8
```

**Note:** The `make wasm` command will generate the WASM module and the `make layer8` command will start the Layer8 server. Make sure that before testing the connection between the WASM module and the Layer8 server, you should have local db set up according to the env variables that you will set. Use the SQL given below for generating the table in a local PostgreSQL DB.

```sql
CREATE TABLE IF NOT EXISTS public.users
(
    id SERIAL PRIMARY KEY NOT NULL,
    username character varying NOT NULL,
    password character varying NOT NULL,
    salt character varying(255)
);
```

Go to `http://localhost:9090` and open the browser console. You should see the following output:

```bash
Go Web Assemly Demo
```

Here you can test the connection my using the following commands:

```bash
connectToServer(); // Make a ping request to the layer8 server
```

```bash
registerUser('Name', 'Password'); // Register a user
```

```bash
loginUser('Name', 'Password'); // Login a user
```
