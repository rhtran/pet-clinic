-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS vets (
  id SERIAL,
  first_name VARCHAR(30),
  last_name VARCHAR(30),
  created_at timestamptz not null default current_timestamp,
  updated_at timestamptz not null default current_timestamp,
  deleted_at timestamptz,
  CONSTRAINT pk_vets PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_vets_last_name ON vets (last_name);

ALTER SEQUENCE vets_id_seq RESTART WITH 100;


CREATE TABLE IF NOT EXISTS specialties (
  id SERIAL,
  name VARCHAR(80),
  created_at timestamptz not null default current_timestamp,
  updated_at timestamptz not null default current_timestamp,
  deleted_at timestamptz,
  CONSTRAINT pk_specialties PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_specialties_name ON specialties (name);

ALTER SEQUENCE specialties_id_seq RESTART WITH 100;


CREATE TABLE IF NOT EXISTS vet_specialties (
  vet_id INT NOT NULL,
  specialty_id INT NOT NULL,
  created_at timestamptz not null default current_timestamp,
  updated_at timestamptz not null default current_timestamp,
  deleted_at timestamptz,
  FOREIGN KEY (vet_id) REFERENCES vets(id),
  FOREIGN KEY (specialty_id) REFERENCES specialties(id),
  CONSTRAINT unique_ids UNIQUE (vet_id,specialty_id)
);

INSERT INTO vets VALUES (1, 'James', 'Carter') ON CONFLICT DO NOTHING;
INSERT INTO vets VALUES (2, 'Helen', 'Leary') ON CONFLICT DO NOTHING;
INSERT INTO vets VALUES (3, 'Linda', 'Douglas') ON CONFLICT DO NOTHING;
INSERT INTO vets VALUES (4, 'Rafael', 'Ortega') ON CONFLICT DO NOTHING;
INSERT INTO vets VALUES (5, 'Henry', 'Stevens') ON CONFLICT DO NOTHING;
INSERT INTO vets VALUES (6, 'Sharon', 'Jenkins') ON CONFLICT DO NOTHING;

INSERT INTO specialties VALUES (1, 'Radiology') ON CONFLICT DO NOTHING;
INSERT INTO specialties VALUES (2, 'Surgery') ON CONFLICT DO NOTHING;
INSERT INTO specialties VALUES (3, 'Dentistry') ON CONFLICT DO NOTHING;

INSERT INTO vet_specialties VALUES (2, 1) ON CONFLICT DO NOTHING;
INSERT INTO vet_specialties VALUES (3, 2) ON CONFLICT DO NOTHING;
INSERT INTO vet_specialties VALUES (3, 3) ON CONFLICT DO NOTHING;
INSERT INTO vet_specialties VALUES (4, 2) ON CONFLICT DO NOTHING;
INSERT INTO vet_specialties VALUES (5, 1) ON CONFLICT DO NOTHING;


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.