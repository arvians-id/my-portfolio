CREATE TABLE IF NOT EXISTS projects (
   id           SERIAL,
   category     VARCHAR(250) NOT NULL,
   title        VARCHAR(250) NOT NULL,
   description  TEXT,
   url          VARCHAR(250),
   is_featured  BOOLEAN DEFAULT FALSE,
   date         DATE NOT NULL,
   working_type VARCHAR(250) NOT NULL,
   created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (id)
)