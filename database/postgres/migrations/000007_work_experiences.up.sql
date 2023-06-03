CREATE TABLE IF NOT EXISTS work_experiences (
   id           SERIAL,
   role         VARCHAR(250) NOT NULL,
   company      VARCHAR(250) NOT NULL,
   description  TEXT,
   start_date   DATE NOT NULL,
   end_date     DATE,
   job_type     VARCHAR(250) NOT NULL,
   created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (id)
)