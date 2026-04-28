DROP INDEX IF EXISTS idx_invite_tokens_token_hash;
DROP INDEX IF EXISTS idx_chat_participants_user_id;
DROP INDEX IF EXISTS idx_messages_created_at;
DROP INDEX IF EXISTS idx_messages_chat_id;

DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS chat_participants;
DROP TABLE IF EXISTS chats;
DROP TABLE IF EXISTS invite_tokens;
DROP TABLE IF EXISTS users;

DROP EXTENSION IF EXISTS "uuid-ossp";