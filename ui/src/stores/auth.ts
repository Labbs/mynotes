import { defineStore } from "pinia"
import { ref } from "vue"
import { authApi } from "../api/auth"

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('token'))
  const isAuthenticated = ref(!!token.value)

  async function login(email: string, password: string) {
    try {
      const { data } = await authApi.login(email, password)
      token.value = data.token
      isAuthenticated.value = true
      localStorage.setItem('token', data.token)
      localStorage.setItem('session_id', data.session_id)
      return data
    } catch (error) {
      throw error
    }
  }

  async function register(name: string, email: string, password: string) {
    try {
      const { data } = await authApi.register(name, email, password)
      token.value = data.token
      isAuthenticated.value = false
      return data
    } catch (error) {
      throw error
    }
  }

  function logout() {
    token.value = null
    isAuthenticated.value = false
    localStorage.removeItem('token')
    localStorage.removeItem('session_id')
  }

  return {
    token,
    isAuthenticated,
    login,
    register,
    logout
  }
})