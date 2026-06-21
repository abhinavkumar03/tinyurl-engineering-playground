CREATE TABLE IF NOT EXISTS url_events (

    id BIGSERIAL PRIMARY KEY,

    url_id BIGINT NOT NULL,

    ip_address VARCHAR(255),

    user_agent TEXT,

    referrer TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_url_events_url_id
    FOREIGN KEY (url_id)
    REFERENCES urls(id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_url_events_url_id
ON url_events(url_id);

CREATE INDEX IF NOT EXISTS idx_url_events_created_at
ON url_events(created_at DESC);