CREATE TABLE IF NOT EXISTS urls (
    id BIGSERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_code VARCHAR(32) UNIQUE,
    click_count BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_urls_short_code
ON urls(short_code);

CREATE INDEX IF NOT EXISTS idx_urls_created_at
ON urls(created_at DESC);

CREATE INDEX IF NOT EXISTS idx_urls_click_count
ON urls(click_count DESC);