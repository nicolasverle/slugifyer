CREATE TABLE IF NOT EXISTS slugs (
  url varchar(250) NOT NULL,
  slug char(6) NOT NULL,
  PRIMARY KEY (slug)
);
