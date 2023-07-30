# Go-Wasm-To-Layer8-To-DB

A test project for experimenting the connection between a Go standard WASM module > Custom Load Balancer > Layer8 Slaves + DB > Content Server

## Prerequisites

- Go 1.11 or higher
- PostgreSQL 9.6 or higher

## How to run

### Usage

Open multiple terminals (6) and run the following commands in each:

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

**Note:** The `make wasm` command will generate the WASM module and run a File Server on specified Port. The `make load-balancer` will start a custom Load Balancer, that will serve as a proxy between WASM Module and Layer8 Slaves and distribute incoming requests equally. The `make layer8-slave-one` & `make layer8-slave-two` & `make layer8-slave-three` will start 3 Layer8 Slaves, that have same functionality, but different ports. The `make content-server` will start a Content Server, from which Layer8 Slaves will get the content and send it to the WASM module.

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
