ALTER TABLE users 
    ALTER COLUMN role TYPE INTEGER;
ALTER TABLE users
    RENAME COLUMN role TO role_id;