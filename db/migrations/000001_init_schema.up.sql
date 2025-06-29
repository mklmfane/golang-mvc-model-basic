CREATE TABLE accounts_bank (
  id SERIAL PRIMARY KEY,
  full_name VARCHAR NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now()
);
