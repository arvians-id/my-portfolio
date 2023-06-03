CREATE TABLE IF NOT EXISTS skills (
   id                   SERIAL,
   category_skill_id    INTEGER NOT NULL,
   name                 VARCHAR(250) NOT NULL,
   icon                 VARCHAR(250),
   PRIMARY KEY (id),
   FOREIGN KEY (category_skill_id) REFERENCES category_skills(id) ON DELETE CASCADE ON UPDATE CASCADE
)