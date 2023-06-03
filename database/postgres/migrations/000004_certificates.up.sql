CREATE TABLE IF NOT EXISTS certificates (
    id              SERIAL,
    name            VARCHAR(250) NOT NULL,
    organization    VARCHAR(250) NOT NULL,
    issue_date      DATE NOT NULL,
    expiration_date DATE,
    credential_id   VARCHAR(250),
    image           VARCHAR(250),
    PRIMARY KEY (id)
)