# Go-Wasm-To-Layer8-To-DB

A test project for experimenting the connection between a Go standard WASM module > Custom Load Balancer > Layer8 Slaves + DB > Content Server

## Prerequisites

- Go 1.11 or higher
- PostgreSQL 9.6 or higher

## How to run

### Usage

**Note:** Make sure to run`go mod tidy` before running the project for the first time and make a .env file with the following variables (You can change the values if you want and use your own local DB Name and Password):

```bash
LOAD_BALANCER_PORT=8000
LAYER8_SLAVE_ONE_PORT=8001
LAYER8_SLAVE_TWO_PORT=8002
LAYER8_SLAVE_THREE_PORT=8003
WASM_SERVER_PORT=9000
LAYER8_MASTER_PORT=9001
CONTENT_SERVER_PORT=9002
DB_USER=postgres
DB_PASS=
DB_NAME=
DB_HOST=localhost
DB_PORT=5432
SSL_MODE=disable
JWT_SECRET=secret
```

Open multiple terminals (7) and run the following commands in each:

```bash
make wasm
```

```bash
make load-balancer
```

```bash
make layer8-slave-one
```

```bash
make layer8-slave-two
```

```bash
make layer8-slave-three
```

```bash
make content-server
```

```bash
make layer8-master
```

**Note:** The `make wasm` command will generate the WASM module and run a File Server on specified Port. The `make load-balancer` will start a custom Load Balancer, that will serve as a proxy between WASM Module and Layer8 Slaves and distribute incoming requests equally. The `make layer8-slave-one` & `make layer8-slave-two` & `make layer8-slave-three` will start 3 Layer8 Slaves, that have same functionality, but different ports. The `make content-server` will start a Content Server, from which Layer8 Slaves will get the content and send it to the WASM module and the `make layer8-master` will start a Layer8 Master, from where the Slaves will get the Secret Key for JWT.

### Database

**Note:** For first time running, you need to create a database and a table. You can do that with the following SQL Query:

```sql
CREATE TABLE IF NOT EXISTS public.users
(
    id SERIAL PRIMARY KEY NOT NULL,
    username character varying UNIQUE NOT NULL,
    password character varying NOT NULL,
    salt character varying(255)
);
```

### Running

Go to `http://localhost:9000` and try Login and Register. The WASM module will send the request to the Layer8 server and the Layer8 server will send the request to the DB and return the response to the WASM module.
