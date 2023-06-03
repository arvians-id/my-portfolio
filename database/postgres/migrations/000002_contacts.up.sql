CREATE TABLE IF NOT EXISTS contacts (
    id       SERIAL,
    platform VARCHAR(250) NOT NULL,
    url      VARCHAR(250) NOT NULL,
    icon     VARCHAR(250),
    PRIMARY KEY (id)
)