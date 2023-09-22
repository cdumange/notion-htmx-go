CREATE TABLE IF NOT EXISTS tasks (
    id uuid PRIMARY KEY,
    category_id uuid,
    title text,
    creation_date timestamp with time zone
);

CREATE INDEX CONCURRENTLY index_tasks_category_id ON tasks(category_id);