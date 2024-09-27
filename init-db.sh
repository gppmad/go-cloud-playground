#!/bin/bash

# Connect to the database
psql -h localhost -U golang -d golangdatabase -c "
  CREATE TABLE IF NOT EXISTS People (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    birthdate DATE
  );
"