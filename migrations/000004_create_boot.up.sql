CREATE OR REPLACE FUNCTION boot(fn text, req jsonb)
RETURNS http_response
AS $$
    import json;
    from collections import defaultdict
    
    render = lambda body, code, headers : [code, body, json.dumps(headers)];
    render_json = lambda body, code = 200, headers = {} : render(json.dumps(body), code, headers | { "Content-Type": "application/json" })
    
    def safe_navigation_dict(d, keys):
        for key in keys:
            if key in d:
                d = d[key]
            else:
                return None
        return d
    
    request = json.loads(req);
    
    GD['ely'] = {
      "render": render,
      "render_json": render_json,
      "params": lambda *keys : safe_navigation_dict(request.get("params"), keys),
      "request": request
    }
    
    return plpy.execute("SELECT status, body, headers FROM " + fn + "()", 1)[0];
$$ LANGUAGE plpython3u;
