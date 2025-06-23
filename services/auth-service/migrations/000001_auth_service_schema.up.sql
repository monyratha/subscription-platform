-- Users table
CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY,
    email      VARCHAR(191) NOT NULL UNIQUE,
    password   VARCHAR(191) NOT NULL,
    name       VARCHAR(100),
    role       VARCHAR(20) DEFAULT 'user',
    status     SMALLINT    DEFAULT 1,
    created_at TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
);

-- Casbin rule table
CREATE TABLE casbin_rule
(
    id    BIGSERIAL PRIMARY KEY,
    ptype VARCHAR(100),
    v0    VARCHAR(100),
    v1    VARCHAR(100),
    v2    VARCHAR(100),
    v3    VARCHAR(100),
    v4    VARCHAR(100),
    v5    VARCHAR(100)
);

-- Global Settings table
CREATE TABLE settings
(
    id            BIGSERIAL PRIMARY KEY,
    setting_key   VARCHAR(191) NOT NULL UNIQUE,
    setting_value TEXT,
    description   TEXT,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
