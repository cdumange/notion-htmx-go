CREATE TABLE IF NOT EXISTS categories (
    id uuid PRIMARY KEY,
    title text NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks (
    id uuid PRIMARY KEY,
    category_id uuid,
    title text,
    creation_date timestamp with time zone,
    CONSTRAINT fk_category
      FOREIGN KEY(category_id) 
        REFERENCES categories(id)
);

CREATE INDEX index_tasks_category_id ON tasks(category_id);