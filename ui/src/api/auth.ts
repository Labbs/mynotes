import api from './axios'

export interface AuthResponse {
  token: string
  session_id: string
  message: string
}

export const authApi = {
  login: (email: string, password: string) => {
    return api.post<AuthResponse>('/auth/login', {
      email,
      password
    })
  },

  register: (email: string, password: string, name: string) => {
    return api.post<AuthResponse>('/auth/register', {
      name: name,
      email: email,
      password: password
    })
  }
}