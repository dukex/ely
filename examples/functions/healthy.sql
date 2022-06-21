CREATE OR REPLACE FUNCTION healthy()
  RETURNS http_response
AS $$
    ely = GD['ely']
    render_json = ely['render_json']

    return render_json({ "healthy": True });
$$ LANGUAGE plpython3u;
