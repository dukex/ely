CREATE OR REPLACE FUNCTION usersindex()
  RETURNS http_response
AS $$
    ely = GD['ely']
    render_json = ely['render_json']
    params = ely['params']

    enabled = params("enabled")

    plan = plpy.prepare("SELECT * FROM users WHERE enabled = $1", [ "boolean" ])

    users = plpy.execute(plan, [ enabled != 'false' ])

    return render_json([dict(user) for user in users]) 
$$ LANGUAGE plpython3u;
