CREATE TABLE IF NOT EXISTS educations (
   id               SERIAL,
   institution      VARCHAR(250) NOT NULL,
   degree           VARCHAR(250) NOT NULL,
   field_of_study   VARCHAR(250) NOT NULL,
   grade            DECIMAL(3,2) NOT NULL,
   description      TEXT,
   start_date       DATE NOT NULL,
   end_date         DATE,
   PRIMARY KEY (id)
)