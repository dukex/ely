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
