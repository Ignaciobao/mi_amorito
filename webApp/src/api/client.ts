import axios, { AxiosInstance } from 'axios';
import { v4 as uuidv4 } from 'uuid';
import {
  User,
  Character,
  ChatMessage,
  ChatSession,
  SendMessageRequest,
  SendMessageResponse,
  ApiError
} from '../types';

class ApiClient {
  private client: AxiosInstance;
  private deviceId: string;

  constructor(baseURL: string = 'http://localhost:8080') {
    this.deviceId = this.getOrCreateDeviceId();
    
    this.client = axios.create({
      baseURL: `${baseURL}/api/v1`,
      headers: {
        'Content-Type': 'application/json',
        'X-Device-ID': this.deviceId,
      },
    });

    // 响应拦截器
    this.client.interceptors.response.use(
      (response) => response,
      (error) => {
        console.error('API Error:', error.response?.data || error.message);
        return Promise.reject(error);
      }
    );
  }

  private getOrCreateDeviceId(): string {
    let deviceId = localStorage.getItem('device_id');
    if (!deviceId) {
      deviceId = uuidv4();
      localStorage.setItem('device_id', deviceId);
    }
    return deviceId!;
  }

  // 用户相关API
  async createUser(nickname?: string): Promise<User> {
    const response = await this.client.post('/users', {
      device_id: this.deviceId,
      nickname: nickname || '匿名用户',
    });
    return response.data.user;
  }

  async getUserProfile(): Promise<User> {
    const response = await this.client.get('/users/profile');
    return response.data.user;
  }

  // 角色相关API
  async getCharacters(): Promise<Character[]> {
    const response = await this.client.get('/characters');
    return response.data.characters;
  }

  async getCharacter(characterId: string): Promise<Character> {
    const response = await this.client.get(`/characters/${characterId}`);
    return response.data.character;
  }

  // 聊天相关API
  async sendMessage(
    characterId: string,
    content: string,
    sessionId?: string
  ): Promise<SendMessageResponse> {
    const request: SendMessageRequest = {
      character_id: characterId,
      content,
      session_id: sessionId,
    };
    
    const response = await this.client.post('/chat/send', request);
    return response.data;
  }

  async getChatHistory(
    sessionId: string,
    limit: number = 20,
    skip: number = 0
  ): Promise<{ messages: ChatMessage[]; has_more: boolean }> {
    const response = await this.client.get(`/chat/history/${sessionId}`, {
      params: { limit, skip },
    });
    return response.data;
  }

  async getChatSessions(): Promise<ChatSession[]> {
    const response = await this.client.get('/chat/sessions');
    return response.data.sessions;
  }

  getDeviceId(): string {
    return this.deviceId;
  }
}

export const apiClient = new ApiClient();
export default ApiClient;