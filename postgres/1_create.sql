CREATE TABLE users (
  id SERIAL NOT NULL,
  email TEXT UNIQUE NOT NULL,
  name VARCHAR(80) NOT NULL,
  displayed_name VARCHAR(80),
  password TEXT NOT NULL,
  profile_picture TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id)
);

CREATE TABLE channels (
  id SERIAL NOT NULL,
  name VARCHAR(100) NOT NULL,
  description VARCHAR(255),
  member INTEGER[] NOT NULL,
  dm_flag BOOLEAN NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id)
);

CREATE TABLE messages (
  id SERIAL NOT NULL,
  message TEXT NOT NULL,
  user_id int8 NOT NULL,
  channel_id int8 NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) references users(id),
  FOREIGN KEY (channel_id) references users(id)
)
