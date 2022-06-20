CREATE OR REPLACE FUNCTION healthy(req jsonb)
  RETURNS http_response
AS $$
    import json

    return [200, json.dumps({ "healthy": True }), { "Content-Type": "application/json" }]
$$ LANGUAGE plpython3u;
