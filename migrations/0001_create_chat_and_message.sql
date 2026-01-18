-- +goose Up
CREATE TABLE chats (
    id         BIGSERIAL PRIMARY KEY,
    title      VARCHAR(200) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE messages (
    id         BIGSERIAL PRIMARY KEY,
    chat_id    BIGINT NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    text       TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_messages_chat_id_created_at ON messages (chat_id, created_at DESC);

-- +goose Down
DROP TABLE messages;
DROP TABLE chats;
