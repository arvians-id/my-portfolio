CREATE TABLE IF NOT EXISTS users (
    id          SERIAL,
    name        VARCHAR(250) NOT NULL,
    email       VARCHAR(250) NOT NULL UNIQUE,
    password    VARCHAR(255) NOT NULL,
    bio         TEXT,
    pronouns    VARCHAR(100) NOT NULL,
    country     VARCHAR(100) NOT NULL,
    job_title   VARCHAR(100) NOT NULL,
    image       VARCHAR(250),
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
)