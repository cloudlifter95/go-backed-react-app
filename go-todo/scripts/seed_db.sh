#!/bin/sh
set -e

DB_PATH=$1

sqlite3 $DB_PATH <<EOF
INSERT INTO todos (title, completed, created_at, updated_at)
VALUES 
  ('Buy groceries', 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('Learn Go', 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('Build a full-stack app', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
EOF

echo "Database seeded successfully!"
