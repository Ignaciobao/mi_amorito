import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { Character } from './types';
import { apiClient } from './api/client';
import CharacterList from './components/CharacterList';
import ChatRoom from './components/ChatRoom';
import './App.css';

function App() {
  const [characters, setCharacters] = useState<Character[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const initializeApp = async () => {
      try {
        // 创建用户（如果不存在）
        await apiClient.createUser();
        
        // 获取角色列表
        const charactersData = await apiClient.getCharacters();
        setCharacters(charactersData);
      } catch (err) {
        console.error('Failed to initialize app:', err);
        setError('应用初始化失败，请检查网络连接');
      } finally {
        setLoading(false);
      }
    };

    initializeApp();
  }, []);

  if (loading) {
    return (
      <div className="app-loading">
        <div className="loading-spinner" />
        <p>正在加载...</p>
      </div>
    );
  }

  if (error) {
    return (
      <div className="app-error">
        <h2>出错了</h2>
        <p>{error}</p>
        <button onClick={() => window.location.reload()}>
          重新加载
        </button>
      </div>
    );
  }

  return (
    <Router>
      <div className="app">
        <Routes>
          <Route 
            path="/" 
            element={<CharacterList characters={characters} />} 
          />
          <Route 
            path="/chat/:characterId" 
            element={<ChatRoom characters={characters} />} 
          />
          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;