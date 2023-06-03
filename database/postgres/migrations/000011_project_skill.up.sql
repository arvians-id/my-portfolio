CREATE TABLE IF NOT EXISTS project_skill (
   project_id   INTEGER REFERENCES projects (id) ON DELETE CASCADE,
   skill_id     INTEGER NOT NULL REFERENCES skills (id) ON DELETE CASCADE,
   PRIMARY KEY (project_id, skill_id)
)