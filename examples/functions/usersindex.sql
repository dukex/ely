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
