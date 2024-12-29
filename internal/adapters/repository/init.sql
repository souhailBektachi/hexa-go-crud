DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'something') THEN
        CREATE DATABASE something 
            WITH OWNER = postgres
                 ENCODING = 'UTF8'
                 CONNECTION LIMIT = -1;
    END IF;
END
$$;

CREATE TABLE IF NOT EXISTS something (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);


