// MongoDB 初始化脚本
db = db.getSiblingDB('mi_amorito');

// 创建用户集合
db.createCollection('users');
db.users.createIndex({ "device_id": 1 }, { unique: true });
db.users.createIndex({ "created_at": 1 });

// 创建角色集合
db.createCollection('characters');
db.characters.createIndex({ "character_id": 1 }, { unique: true });
db.characters.createIndex({ "is_active": 1 });

// 创建聊天消息集合
db.createCollection('chats');
db.chats.createIndex({ "user_id": 1, "character_id": 1 });
db.chats.createIndex({ "user_id": 1, "created_at": -1 });
db.chats.createIndex({ "session_id": 1 });

// 创建聊天会话集合
db.createCollection('chat_sessions');
db.chat_sessions.createIndex({ "user_id": 1, "updated_at": -1 });
db.chat_sessions.createIndex({ "session_id": 1 }, { unique: true });

print('MongoDB initialization completed!');