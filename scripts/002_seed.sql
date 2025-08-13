INSERT INTO events (id, title, description, start_time, end_time)
VALUES
    (gen_random_uuid(), 'Sample Event', 'seed data', now() + interval '1 hour', now() + interval '2 hour')
    ON CONFLICT DO NOTHING;
