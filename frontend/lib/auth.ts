import axios from 'axios';
import Cookies from 'js-cookie';
import type { User, AuthResponse, LoginCredentials, RegisterCredentials } from '@/types/auth';

// Helper for adding auth token to requests
export const api = axios.create();

// Configure axios interceptor to add authorization headers
api.interceptors.request.use((config) => {
  // Get token from httpOnly cookie (handled by the browser)
  // or from localStorage as fallback
// Helper for adding auth token to requests
export const api = axios.create();

// Configure axios interceptor to add authorization headers
api.interceptors.request.use((config) => {
  // Get token from httpOnly cookie (handled by the browser)
  const token = Cookies.get('token');
  
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  
  return config;
});

// Auth helper functions
export const auth = {
  // Check if user is authenticated
  isAuthenticated: () => {
    return !!Cookies.get('token');
  },
  
  // Login user
  login: async (credentials: LoginCredentials): Promise<AuthResponse> => {
    const response = await axios.post('/api/auth/login', credentials);
    
    if (response.data.token) {
  
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  
  return config;
});

// Auth helper functions
export const auth = {
  // Check if user is authenticated
  isAuthenticated: () => {
    return !!localStorage.getItem('token');
  },
  
  // Login user
// Login user
  login: async (credentials: LoginCredentials): Promise<AuthResponse> => {
    try {
      const response = await axios.post('/api/auth/login', credentials);
      
      if (response.data.token) {
        localStorage.setItem('token', response.data.token);
      }
      
      return response.data;
    } catch (error) {
      console.error('Login failed:', error);
      throw error;
    }
  },

  // Register user
    const response = await axios.post('/api/auth/login', credentials);
    
    if (response.data.token) {
// Import CryptoJS for encryption
// CryptoJS is used to encrypt sensitive data before storing in localStorage
import CryptoJS from 'crypto-js';

  login: async (credentials: LoginCredentials): Promise<AuthResponse> => {
    const response = await axios.post('/api/auth/login', credentials);
    
    if (response.data.token) {
      const encryptedToken = CryptoJS.AES.encrypt(response.data.token, process.env.ENCRYPTION_KEY).toString();
      localStorage.setItem('token', encryptedToken);
    }
    
    return response.data;
    }
    
    return response.data;
  },

  // Register user
  register: async (userData: RegisterCredentials) => {
    const response = await axios.post('/api/auth/register', userData);
    return response.data;
  },
  
  // Logout user
  logout: async () => {
    localStorage.removeItem('token');
    // Also try to clear httpOnly cookie by calling backend
    try {
      await axios.post('/api/auth/logout');
    } catch (error) {
      console.error('Error during logout:', error);
    }
  },
  
  // Get current user info
  getCurrentUser: async (): Promise<User | null> => {
    if (!auth.isAuthenticated()) {
      return null;
    }
    
    try {
      const response = await api.get('/api/auth/me');
      return response.data.user;
    } catch (error) {
      console.error('Failed to get current user:', error);
      return null;
    }
  }
};