CREATE TABLE
  account (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    deleted_at TIMESTAMP
  );

CREATE TABLE
  account_information (
    id UUID PRIMARY KEY,
    display_name TEXT NOT NULL,
    favorite_number INTEGER,
    homeworld_realm TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    deleted_at TIMESTAMP,
    account_id UUID NOT NULL UNIQUE REFERENCES account (id)
  );

CREATE TABLE
  user_role (
    id UUID PRIMARY KEY,
    role_title TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    deleted_at TIMESTAMP
  );

CREATE TABLE
  account_role (
    id UUID PRIMARY KEY,
    account_id UUID NOT NULL REFERENCES account (id),
    user_role_id UUID NOT NULL REFERENCES user_role (id)
  );

CREATE TABLE
  project (
    id UUID PRIMARY KEY,
    title VARCHAR(24) NOT NULL,
    sector VARCHAR(24) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    deleted_at TIMESTAMP
  );

CREATE TABLE
  project_account_assignment (
    id UUID PRIMARY KEY,
    project_id UUID NOT NULL REFERENCES project (id),
    account_id UUID NOT NULL REFERENCES account (id)
  );

CREATE TABLE
  invitation_status (
    id UUID PRIMARY KEY,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    deleted_at TIMESTAMP
  );