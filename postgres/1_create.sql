CREATE TABLE users (
  id SERIAL NOT NULL,
  email VARCHAR(20) NOT NULL,
  name VARCHAR(80) NOT NULL,
  displayed_name VARCHAR(80) NOT NULL,
  encrypted_password TEXT NOT NULL,
  profile_picture_path TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id)
);

CREATE TABLE channels (
  id SERIAL NOT NULL,
  name VARCHAR(100) NOT NULL,
  description VARCHAR(255),
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  dm_flag boolean NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE messages (
  id SERIAL NOT NULL,
  message TEXT NOT NULL,
  user_id int8 NOT NULL,
  channel_id int8 NOT NULL,
  send_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id)
)
