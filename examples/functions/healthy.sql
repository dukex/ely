CREATE OR REPLACE FUNCTION healthy() RETURNS http_response AS $$
BEGIN
    return (200, '{ "healthy": true }'::jsonb,'{ "Content-Type": "application/json" }'::jsonb );
END;
$$ LANGUAGE plpgsql;
