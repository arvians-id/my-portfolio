CREATE TABLE IF NOT EXISTS project_images (
   id          SERIAL,
   project_id  INTEGER NOT NULL,
   image       VARCHAR(250) NOT NULL,
   PRIMARY KEY (id),
   FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE ON UPDATE CASCADE
)