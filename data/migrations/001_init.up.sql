CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    name       TEXT                     NOT NULL,
    email      TEXT                     NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
