import { defineStore } from 'pinia'
import { ref } from 'vue'
import { meApi } from '../api/me'

export interface UserPreferences {
  ui: {
    expanded_documents: string[]
    expanded_spaces: string[]
    sidebarCollapsed: boolean
    sidebarWidth: number
  }
}

export const usePreferencesStore = defineStore('preferences', () => {
  const preferences = ref<UserPreferences>({
    ui: {
      expanded_documents: [],
      expanded_spaces: [],
      sidebarCollapsed: false,
      sidebarWidth: 256
    }
  })

  const loading = ref(false)
  const error = ref<string | null>(null)

  // Initialiser avec localStorage au démarrage du store
  loadFromLocalStorage()

  // Charger les préférences depuis l'API
  async function loadPreferences() {
    loading.value = true
    error.value = null

    try {
      const { data } = await meApi.getPreferences()
      
      // Si on a des préférences, les utiliser, sinon garder les valeurs par défaut
      if (data && data.ui) {
        preferences.value = {
          ui: {
            expanded_documents: data.ui.expanded_documents || [],
            expanded_spaces: data.ui.expanded_spaces || [],
            sidebarCollapsed: data.ui.sidebarCollapsed || false,
            sidebarWidth: data.ui.sidebarWidth || 256
          }
        }
      }

      // Synchroniser avec localStorage
      syncToLocalStorage()
    } catch (err) {
      console.error('Failed to load user preferences:', err)
      // En cas d'erreur, charger depuis localStorage
      loadFromLocalStorage()
    } finally {
      loading.value = false
    }
  }

  // Sauvegarder les préférences vers l'API
  async function savePreferences() {
    try {
      await meApi.updatePreferences(preferences.value)
      // Synchroniser avec localStorage après sauvegarde réussie
      syncToLocalStorage()
    } catch (err) {
      console.error('Failed to save user preferences:', err)
      error.value = 'Failed to save preferences'
    }
  }

  // Mettre à jour une préférence spécifique
  function updatePreference<K extends keyof UserPreferences['ui']>(
    key: K, 
    value: UserPreferences['ui'][K]
  ) {
    preferences.value.ui[key] = value
    
    // Sauvegarder immédiatement vers localStorage pour une réponse rapide
    syncToLocalStorage()
    
    // Sauvegarder vers l'API (de manière asynchrone)
    savePreferences()
  }

  // Charger depuis localStorage (fallback)
  function loadFromLocalStorage() {
    try {
      const expandedSpaces = localStorage.getItem('mynotes_expanded_spaces')
      const expandedDocuments = localStorage.getItem('mynotes_expanded_documents')
      const sidebarCollapsed = localStorage.getItem('sidebarCollapsed')
      const sidebarWidth = localStorage.getItem('sidebarWidth')

      if (expandedSpaces) {
        preferences.value.ui.expanded_spaces = JSON.parse(expandedSpaces)
      }
      if (expandedDocuments) {
        preferences.value.ui.expanded_documents = JSON.parse(expandedDocuments)
      }
      if (sidebarCollapsed) {
        preferences.value.ui.sidebarCollapsed = sidebarCollapsed === 'true'
      }
      if (sidebarWidth) {
        preferences.value.ui.sidebarWidth = parseInt(sidebarWidth)
      }
    } catch (err) {
      console.error('Failed to load preferences from localStorage:', err)
    }
  }

  // Synchroniser vers localStorage
  function syncToLocalStorage() {
    try {
      localStorage.setItem('mynotes_expanded_spaces', JSON.stringify(preferences.value.ui.expanded_spaces))
      localStorage.setItem('mynotes_expanded_documents', JSON.stringify(preferences.value.ui.expanded_documents))
      localStorage.setItem('sidebarCollapsed', preferences.value.ui.sidebarCollapsed.toString())
      localStorage.setItem('sidebarWidth', preferences.value.ui.sidebarWidth.toString())
    } catch (err) {
      console.error('Failed to sync preferences to localStorage:', err)
    }
  }

  // Nettoyer localStorage (au logout)
  function clearLocalStorage() {
    try {
      localStorage.removeItem('mynotes_expanded_spaces')
      localStorage.removeItem('mynotes_expanded_documents')
      localStorage.removeItem('sidebarCollapsed')
      localStorage.removeItem('sidebarWidth')
      
      // Réinitialiser les préférences
      preferences.value = {
        ui: {
          expanded_documents: [],
          expanded_spaces: [],
          sidebarCollapsed: false,
          sidebarWidth: 256
        }
      }
    } catch (err) {
      console.error('Failed to clear localStorage:', err)
    }
  }

  return {
    preferences,
    loading,
    error,
    loadPreferences,
    savePreferences,
    updatePreference,
    loadFromLocalStorage,
    syncToLocalStorage,
    clearLocalStorage
  }
})
