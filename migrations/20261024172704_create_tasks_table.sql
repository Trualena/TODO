-- +goose Up
CREATE TABLE tasks(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL, 
    description TEXT,
    done BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE
);
CREATE INDEX  idx_tasks_user_id ON tasks(user_id);
-- +goose Down
DROP TABLE tasks;
