CREATE TABLE IF NOT EXISTS guest_testimonials (
   id           SERIAL,
   name         VARCHAR(250) NOT NULL,
   email        VARCHAR(250) NOT NULL UNIQUE,
   password     VARCHAR(255) NOT NULL,
   github_id    VARCHAR(255),
   google_id    VARCHAR(255),
   message      VARCHAR(250) NOT NULL,
   is_published BOOLEAN DEFAULT FALSE,
   created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (id)
)