import { defineStore } from "pinia"
import { ref } from "vue"
import { authApi } from "../api/auth"
import { usePreferencesStore } from "./preferences"

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('token'))
  const isAuthenticated = ref(!!token.value)

  // Initialiser les préférences au démarrage si l'utilisateur est déjà connecté
  async function initializeAuth() {
    if (token.value && isAuthenticated.value) {
      const preferencesStore = usePreferencesStore()
      try {
        await preferencesStore.loadPreferences()
      } catch (error) {
        // Si on ne peut pas charger les préférences, charger depuis localStorage
        preferencesStore.loadFromLocalStorage()
      }
    }
  }

  async function login(email: string, password: string) {
    try {
      const { data } = await authApi.login(email, password)
      token.value = data.token
      isAuthenticated.value = true
      localStorage.setItem('token', data.token)
      localStorage.setItem('session_id', data.session_id)
      
      // Charger les préférences utilisateur après la connexion
      const preferencesStore = usePreferencesStore()
      await preferencesStore.loadPreferences()
      
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
    
    // Nettoyer les préférences utilisateur au logout
    const preferencesStore = usePreferencesStore()
    preferencesStore.clearLocalStorage()
  }

  return {
    token,
    isAuthenticated,
    login,
    register,
    logout,
    initializeAuth
  }
})