CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE base_table (
                            created_at TIMESTAMP NOT NULL,
                            updated_at TIMESTAMP NOT NULL
);

CREATE TABLE user_account (
                              id uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
                              email VARCHAR(255) UNIQUE NOT NULL,
                              password_hash VARCHAR(255) NOT NULL
) INHERITS (base_table);

CREATE TABLE post (
                      id uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
                      title VARCHAR(255) NOT NULL,
                      body TEXT,
                      author_id uuid,
                      FOREIGN KEY (author_id) REFERENCES user_account (id) ON DELETE CASCADE
) INHERITS (base_table);