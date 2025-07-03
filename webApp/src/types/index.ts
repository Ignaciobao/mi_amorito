export interface User {
  id: string;
  device_id: string;
  nickname: string;
  avatar: string;
  created_at: string;
  updated_at: string;
}

export interface Character {
  id: string;
  character_id: string;
  name: string;
  avatar: string;
  description: string;
  personality: string;
  background: string;
  greeting: string;
  tags: string[];
  is_active: boolean;
  sort_order: number;
  created_at: string;
  updated_at: string;
}

export interface ChatMessage {
  id: string;
  session_id: string;
  user_id: string;
  character_id: string;
  role: 'user' | 'assistant';
  content: string;
  timestamp: string;
  created_at: string;
}

export interface ChatSession {
  id: string;
  session_id: string;
  user_id: string;
  character_id: string;
  title: string;
  last_message: string;
  created_at: string;
  updated_at: string;
}

export interface SendMessageRequest {
  session_id?: string;
  character_id: string;
  content: string;
}

export interface SendMessageResponse {
  session_id: string;
  message: ChatMessage;
  reply: ChatMessage;
}

export interface ApiError {
  error: string;
  code: string;
}