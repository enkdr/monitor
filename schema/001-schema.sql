-- schema.sql
CREATE TABLE if not exists fs_stats  (
    id SERIAL PRIMARY KEY,
    fs_stats_json JSONB NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
