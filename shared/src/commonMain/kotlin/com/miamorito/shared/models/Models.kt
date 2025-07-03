package com.miamorito.shared.models

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class User(
    @SerialName("id") val id: String,
    @SerialName("device_id") val deviceId: String,
    @SerialName("nickname") val nickname: String,
    @SerialName("avatar") val avatar: String,
    @SerialName("created_at") val createdAt: String,
    @SerialName("updated_at") val updatedAt: String
)

@Serializable
data class Character(
    @SerialName("id") val id: String,
    @SerialName("character_id") val characterId: String,
    @SerialName("name") val name: String,
    @SerialName("avatar") val avatar: String,
    @SerialName("description") val description: String,
    @SerialName("personality") val personality: String,
    @SerialName("background") val background: String,
    @SerialName("greeting") val greeting: String,
    @SerialName("tags") val tags: List<String>,
    @SerialName("is_active") val isActive: Boolean,
    @SerialName("sort_order") val sortOrder: Int,
    @SerialName("created_at") val createdAt: String,
    @SerialName("updated_at") val updatedAt: String
)

@Serializable
data class ChatMessage(
    @SerialName("id") val id: String,
    @SerialName("session_id") val sessionId: String,
    @SerialName("user_id") val userId: String,
    @SerialName("character_id") val characterId: String,
    @SerialName("role") val role: String, // "user" 或 "assistant"
    @SerialName("content") val content: String,
    @SerialName("timestamp") val timestamp: String,
    @SerialName("created_at") val createdAt: String
)

@Serializable
data class ChatSession(
    @SerialName("id") val id: String,
    @SerialName("session_id") val sessionId: String,
    @SerialName("user_id") val userId: String,
    @SerialName("character_id") val characterId: String,
    @SerialName("title") val title: String,
    @SerialName("last_message") val lastMessage: String,
    @SerialName("created_at") val createdAt: String,
    @SerialName("updated_at") val updatedAt: String
)

// API 请求模型
@Serializable
data class CreateUserRequest(
    @SerialName("device_id") val deviceId: String,
    @SerialName("nickname") val nickname: String? = null
)

@Serializable
data class SendMessageRequest(
    @SerialName("session_id") val sessionId: String? = null,
    @SerialName("character_id") val characterId: String,
    @SerialName("content") val content: String
)

// API 响应模型
@Serializable
data class ApiResponse<T>(
    @SerialName("data") val data: T? = null,
    @SerialName("error") val error: String? = null,
    @SerialName("code") val code: String? = null
)

@Serializable
data class UserResponse(
    @SerialName("user") val user: User
)

@Serializable
data class SendMessageResponse(
    @SerialName("session_id") val sessionId: String,
    @SerialName("message") val message: ChatMessage,
    @SerialName("reply") val reply: ChatMessage
)

@Serializable
data class GetCharactersResponse(
    @SerialName("characters") val characters: List<Character>
)

@Serializable
data class GetChatHistoryResponse(
    @SerialName("messages") val messages: List<ChatMessage>,
    @SerialName("has_more") val hasMore: Boolean
)

@Serializable
data class GetChatSessionsResponse(
    @SerialName("sessions") val sessions: List<ChatSession>
)

@Serializable
data class CharacterResponse(
    @SerialName("character") val character: Character
)