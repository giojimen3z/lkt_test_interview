CREATE TABLE IF NOT EXISTS events (
                                      id UUID PRIMARY KEY,
                                      title VARCHAR(100) NOT NULL,
    description TEXT,
    start_time TIMESTAMPTZ NOT NULL,
    end_time   TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
    );

CREATE INDEX IF NOT EXISTS idx_events_start_time ON events (start_time);
CREATE INDEX IF NOT EXISTS idx_events_title ON events (title);
