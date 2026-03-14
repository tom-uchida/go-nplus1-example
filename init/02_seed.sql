INSERT INTO users (name)
SELECT 'user_' || generate_series
FROM generate_series(1,10000);

INSERT INTO posts (user_id, title)
SELECT (random()*999 + 1)::int, 'post'
FROM generate_series(1,100000);