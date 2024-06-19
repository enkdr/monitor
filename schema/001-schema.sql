-- schema.sql
CREATE TABLE if not exists fs_stats  (
    id SERIAL PRIMARY KEY,
    stats_json JSONB NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE if not exists cpu_stats  (
    id SERIAL PRIMARY KEY,
    stats_json JSONB NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE if not exists process_stats  (
    id SERIAL PRIMARY KEY,
    stats_json JSONB NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
