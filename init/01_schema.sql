CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT
);

CREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  user_id INT,
  title TEXT
);