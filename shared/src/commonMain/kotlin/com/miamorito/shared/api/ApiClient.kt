package com.miamorito.shared.api

import com.miamorito.shared.models.*
import io.ktor.client.*
import io.ktor.client.call.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.client.plugins.logging.*
import io.ktor.client.request.*
import io.ktor.http.*
import io.ktor.serialization.kotlinx.json.*
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import kotlinx.serialization.json.Json

class ApiClient(
    private val baseUrl: String = "http://localhost:8080",
    private val deviceId: String
) {
    private val httpClient = HttpClient {
        install(ContentNegotiation) {
            json(Json {
                ignoreUnknownKeys = true
                isLenient = true
            })
        }
        install(Logging) {
            level = LogLevel.INFO
        }
    }

    // 用户相关API
    suspend fun createUser(nickname: String? = null): Result<User> {
        return try {
            val response: UserResponse = httpClient.post("$baseUrl/api/v1/users") {
                contentType(ContentType.Application.Json)
                headers {
                    append("X-Device-ID", deviceId)
                }
                setBody(CreateUserRequest(deviceId, nickname))
            }.body()
            
            Result.success(response.user)
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun getUserProfile(): Result<User> {
        return try {
            val response: UserResponse = httpClient.get("$baseUrl/api/v1/users/profile") {
                headers {
                    append("X-Device-ID", deviceId)
                }
            }.body()
            
            Result.success(response.user)
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    // 角色相关API
    suspend fun getCharacters(): Result<List<Character>> {
        return try {
            val response: GetCharactersResponse = httpClient.get("$baseUrl/api/v1/characters").body()
            Result.success(response.characters)
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun getCharacter(characterId: String): Result<Character> {
        return try {
            val response: CharacterResponse = httpClient.get("$baseUrl/api/v1/characters/$characterId").body()
            Result.success(response.character)
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    // 聊天相关API
    suspend fun sendMessage(
        characterId: String,
        content: String,
        sessionId: String? = null
    ): Result<SendMessageResponse> {
        return try {
            val response: SendMessageResponse = httpClient.post("$baseUrl/api/v1/chat/send") {
                contentType(ContentType.Application.Json)
                headers {
                    append("X-Device-ID", deviceId)
                }
                setBody(SendMessageRequest(sessionId, characterId, content))
            }.body()
            
            Result.success(response)
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun getChatHistory(
        sessionId: String,
        limit: Int = 20,
        skip: Int = 0
    ): Result<GetChatHistoryResponse> {
        return try {
            val response: GetChatHistoryResponse = httpClient.get("$baseUrl/api/v1/chat/history/$sessionId") {
                headers {
                    append("X-Device-ID", deviceId)
                }
                parameter("limit", limit)
                parameter("skip", skip)
            }.body()
            
            Result.success(response)
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun getChatSessions(): Result<List<ChatSession>> {
        return try {
            val response: GetChatSessionsResponse = httpClient.get("$baseUrl/api/v1/chat/sessions") {
                headers {
                    append("X-Device-ID", deviceId)
                }
            }.body()
            
            Result.success(response.sessions)
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    fun close() {
        httpClient.close()
    }
}

// 流式响应扩展
fun ApiClient.getChatHistoryFlow(
    sessionId: String,
    limit: Int = 20
): Flow<List<ChatMessage>> = flow {
    var skip = 0
    var hasMore = true
    val allMessages = mutableListOf<ChatMessage>()
    
    while (hasMore) {
        val result = getChatHistory(sessionId, limit, skip)
        result.fold(
            onSuccess = { response ->
                allMessages.addAll(response.messages)
                emit(allMessages.toList())
                hasMore = response.hasMore
                skip += limit
            },
            onFailure = {
                hasMore = false
            }
        )
    }
}