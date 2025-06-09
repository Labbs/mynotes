import { defineStore } from 'pinia'
import { ref } from 'vue'
import { meApi } from '../api/me'
import type { User } from '../api/interface'

export const useProfileStore = defineStore('profile', () => {
  const profile = ref<User | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchProfile() {
    loading.value = true
    error.value = null
    
    try {
      const { data } = await meApi.getProfile()
      profile.value = data
      return data
    } catch (err) {
      error.value = 'Failed to fetch profile'
      console.error('Error fetching profile:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  return {
    profile,
    loading,
    error,
    fetchProfile
  }
})
