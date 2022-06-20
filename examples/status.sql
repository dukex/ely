CREATE OR REPLACE FUNCTION status() RETURNS integer AS  $$
BEGIN
  return 42
END;
$$ LANGUAGE plpgsql;
