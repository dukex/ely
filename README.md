# ELY

ELY is a toy project, please does not use it in production yer, the main idea is run a api service inside the database, a generic golang server is up but the functions to handle the endpoint is running in the any language supported by postgresql

## Examples

See the [Makefile](Makefile) the task `build_and_run` and the [examples](examples/) directory.

Create the `usersindex` function as:

```
CREATE OR REPLACE FUNCTION usersindex(req jsonb)
  RETURNS http_response
AS $$
    import json
    global req
    params = json.loads(req)

    q = params.get("query", dict({})).get("q")

    users = plpy.execute("SELECT * FROM users", 5)

    body = []
    for u in users:
        body.append(dict(u))

    return [200, json.dumps(body), { "Content-Type": "application/json" }]
$$ LANGUAGE plpython3u;
```

Setup the endpoint:

```
endpoint:
  - path: "/users"
    function: 'usersindex'
````

Run server:

```
ely server -p 3000 -c ./examples/ely.yaml
```

Access the endpoint:

[localhost:3000/users](https://localhost:3000/users)
