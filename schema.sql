CREATE TABLE players (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  experience VARCHAR(255),
  playing_style VARCHAR(255)
);

CREATE TABLE courts (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  location VARCHAR(255) NOT NULL,
  availability VARCHAR(255) NOT NULL,
  price DECIMAL(10, 2) NOT NULL
);

CREATE TABLE matches (
  id SERIAL PRIMARY KEY,
  player1_id INTEGER NOT NULL REFERENCES players(id),
  player2_id INTEGER NOT NULL REFERENCES players(id),
  court_id INTEGER NOT NULL REFERENCES courts(id),
  match_date DATE NOT NULL,
  score VARCHAR(255) NOT NULL
);

CREATE TABLE bookings (
  id SERIAL PRIMARY KEY,
  court_id INTEGER NOT NULL REFERENCES courts(id),
  player_id INTEGER NOT NULL REFERENCES players(id),
  booking_date DATE NOT NULL,
  start_time TIME NOT NULL,
  end_time TIME NOT NULL
);

CREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  player_id INTEGER NOT NULL REFERENCES players(id),
  content TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comments (
  id SERIAL PRIMARY KEY,
  post_id INTEGER NOT NULL REFERENCES posts(id),
  player_id INTEGER NOT NULL REFERENCES players(id),
  content TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
