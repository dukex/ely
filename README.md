# ELY

ELY is a toy project, please does not use it in production yer, the main idea is run a api service inside the database, a generic golang server is up but the functions to handle the endpoint is running in the any language supported by postgresql

## Installation

Install ELY command package

```
$ go install github.com/dukex/ely/cmd/ely@latest
```

Now you can use `ely` command

```
$ ely
ELY is a tool for create a API service using database functions

Usage:
  ely [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  db          Manage ELY database
  deploy      Deploy functions to database
  help        Help about any command
  server      Run the ELY server

Flags:
  -h, --help   help for ely

Use "ely [command] --help" for more information about a command.
```

## Usage

The ELY requires a up and running postgresql database, all commands depends of `DATABASE_URL` env proper configured.  

```
export DATABASE_URL=postgresql://localhost:5433/ely?sslmode=disable
```

### Create the workdir

Let's to create your project directory

```
$ mkdir -p awesome-db-api/functions
```

### Setup database

Now we need setup the database, we will create the ELY types

```
$ ely db setup
```

### Create functions

Let's to create the first and simpler function, the healthy status endpoint

```
$ touch functions/healthy.sql
```

Open the `functions/healthy.sql` file and put the follow content:

```
CREATE OR REPLACE FUNCTION healthy(req jsonb)
  RETURNS http_response
AS $$
    import json

    return [200, json.dumps({ "healthy": True }), { "Content-Type": "application/json" }]
$$ LANGUAGE plpython3u;
```

This functions basically will responds the HTTP Status 200, with a json `{ "healthy": true }`.

Now let's to deploy the function to the database:

```
$ ely deploy
```

Done! Your functions is ready to be used.

### Setup an endpoint

Let's to create a YAML to configure your API

```
$ touch ely.yml
```

This will map the endpoints with the database functions, given we created the `healthy` previously let's to map this function to a endpoint. Edit the `ely.yml` as:

```yaml
endpoint:
  - path: "/health_status"
    function: 'healthy'
```

### Up and running server

To finish, we should up the server

```
$ ely server
Server starring at 9001
```

Now you can access [localhost:9001/health_status](http://localhost:9001/health_status)

## Complex examples

**You can see more examples in the [examples](examples/) directory.**

Let's to create a users index endpoint

```
$ touch functions/userindex.sql
```

Open the `functions/userindex.sql` file and put the follow content:

```
CREATE OR REPLACE FUNCTION usersindex(req jsonb)
  RETURNS http_response
AS $$
    import json
    global req
    request = json.loads(req)

    enabled = request.get("params", dict({})).get("enabled")

    plan = plpy.prepare("SELECT * FROM users WHERE enabled = $1", [ "boolean" ])

    users = plpy.execute(plan, [ enabled != 'false' ])

    body = [dict(user) for user in users]

    return [200, json.dumps(body), json.dumps({ "Content-Type": "application/json" })]
$$ LANGUAGE plpython3u;
```

And in the `ely.yml` file:

```
endpoint:
  - path: "/users"
    function: 'usersindex'
````

Run server:

```
$ ely server
```

Access the endpoint [localhost:9001/users](http://localhost:9001/users). You can send the enabled query string params to see diferents results [localhost:9001/users?enabled=false](http://localhost:9001/users?enabled=false)


