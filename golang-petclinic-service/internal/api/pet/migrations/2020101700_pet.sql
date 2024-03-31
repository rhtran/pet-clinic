-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS types (
  id SERIAL,
  name VARCHAR(80),
  created_at timestamptz not null default current_timestamp,
  updated_at timestamptz not null default current_timestamp,
  deleted_at timestamptz,
  CONSTRAINT pk_types PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_types_name ON types (name);

ALTER SEQUENCE types_id_seq RESTART WITH 100;

CREATE TABLE IF NOT EXISTS owners (
  id SERIAL,
  first_name VARCHAR(30),
  last_name VARCHAR(30),
  address VARCHAR(255),
  city VARCHAR(80),
  telephone VARCHAR(20),
  created_at timestamptz not null default current_timestamp,
  updated_at timestamptz not null default current_timestamp,
  deleted_at timestamptz,
  CONSTRAINT pk_owners PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_owners_last_name ON owners (last_name);

ALTER SEQUENCE owners_id_seq RESTART WITH 100;


CREATE TABLE IF NOT EXISTS pets (
  id SERIAL,
  name VARCHAR(30),
  birth_date DATE,
  type_id INT NOT NULL,
  owner_id INT NOT NULL,
  created_at timestamptz default now(),
  updated_at timestamptz default now(),
  deleted_at timestamptz,
  FOREIGN KEY (owner_id) REFERENCES owners(id),
  FOREIGN KEY (type_id) REFERENCES types(id),
  CONSTRAINT pk_pets PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_pets_name ON pets (name);

ALTER SEQUENCE pets_id_seq RESTART WITH 100;


INSERT INTO types VALUES (1, 'cat') ON CONFLICT DO NOTHING;
INSERT INTO types VALUES (2, 'dog') ON CONFLICT DO NOTHING;
INSERT INTO types VALUES (3, 'lizard') ON CONFLICT DO NOTHING;
INSERT INTO types VALUES (4, 'snake') ON CONFLICT DO NOTHING;
INSERT INTO types VALUES (5, 'bird') ON CONFLICT DO NOTHING;
INSERT INTO types VALUES (6, 'hamster') ON CONFLICT DO NOTHING;

INSERT INTO owners VALUES (1, 'George', 'Franklin', '110 W. Liberty St.', 'Madison', '6085551023') ON CONFLICT DO NOTHING;
INSERT INTO owners VALUES (2, 'Betty', 'Davis', '638 Cardinal Ave.', 'Sun Prairie', '6085551749') ON CONFLICT DO NOTHING;
INSERT INTO owners VALUES (3, 'Eduardo', 'Rodriquez', '2693 Commerce St.', 'McFarland', '6085558763') ON CONFLICT DO NOTHING;
INSERT INTO owners VALUES (4, 'Harold', 'Davis', '563 Friendly St.', 'Windsor', '6085553198') ON CONFLICT DO NOTHING;
INSERT INTO owners VALUES (5, 'Peter', 'McTavish', '2387 S. Fair Way', 'Madison', '6085552765') ON CONFLICT DO NOTHING;
INSERT INTO owners VALUES (6, 'Jean', 'Coleman', '105 N. Lake St.', 'Monona', '6085552654') ON CONFLICT DO NOTHING;
INSERT INTO owners VALUES (7, 'Jeff', 'Black', '1450 Oak Blvd.', 'Monona', '6085555387') ON CONFLICT DO NOTHING;
INSERT INTO owners VALUES (8, 'Maria', 'Escobito', '345 Maple St.', 'Madison', '6085557683') ON CONFLICT DO NOTHING;
INSERT INTO owners VALUES (9, 'David', 'Schroeder', '2749 Blackhawk Trail', 'Madison', '6085559435') ON CONFLICT DO NOTHING;
INSERT INTO owners VALUES (10, 'Carlos', 'Estaban', '2335 Independence La.', 'Waunakee', '6085555487') ON CONFLICT DO NOTHING;

INSERT INTO pets VALUES (1, 'Leo', '2000-09-07', 1, 1) ON CONFLICT DO NOTHING;
INSERT INTO pets VALUES (2, 'Basil', '2002-08-06', 6, 2) ON CONFLICT DO NOTHING;
INSERT INTO pets VALUES (3, 'Rosy', '2001-04-17', 2, 3) ON CONFLICT DO NOTHING;
INSERT INTO pets VALUES (4, 'Jewel', '2000-03-07', 2, 3) ON CONFLICT DO NOTHING;
INSERT INTO pets VALUES (5, 'Iggy', '2000-11-30', 3, 4) ON CONFLICT DO NOTHING;
INSERT INTO pets VALUES (6, 'George', '2000-01-20', 4, 5) ON CONFLICT DO NOTHING;
INSERT INTO pets VALUES (7, 'Samantha', '1995-09-04', 1, 6) ON CONFLICT DO NOTHING;
INSERT INTO pets VALUES (8, 'Max', '1995-09-04', 1, 6) ON CONFLICT DO NOTHING;
INSERT INTO pets VALUES (9, 'Lucky', '1999-08-06', 5, 7) ON CONFLICT DO NOTHING;
INSERT INTO pets VALUES (10, 'Mulligan', '1997-02-24', 2, 8) ON CONFLICT DO NOTHING;
INSERT INTO pets VALUES (11, 'Freddy', '2000-03-09', 5, 9) ON CONFLICT DO NOTHING;
INSERT INTO pets VALUES (12, 'Lucky', '2000-06-24', 2, 10) ON CONFLICT DO NOTHING;
INSERT INTO pets VALUES (13, 'Sly', '2002-06-08', 1, 10) ON CONFLICT DO NOTHING;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.