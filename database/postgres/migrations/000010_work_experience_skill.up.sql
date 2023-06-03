CREATE TABLE IF NOT EXISTS work_experience_skill (
   work_experience_id   INTEGER REFERENCES work_experiences (id) ON DELETE CASCADE,
   skill_id             INTEGER NOT NULL REFERENCES skills (id) ON DELETE CASCADE,
   PRIMARY KEY (work_experience_id, skill_id)
)