import { defineStore } from 'pinia'
import { ref } from 'vue'
import { adminApi } from '../api/admin'
import type { User, Group, Space } from '../api/interface'

export const useAdminStore = defineStore('admin', () => {
  const users = ref<User[]>([])
  const groups = ref<Group[]>([])
  const spaces = ref<Space[]>([])
  
  const loading = ref({
    users: false,
    groups: false,
    spaces: false
  })
  
  const error = ref<string | null>(null)

  async function fetchUsers() {
    loading.value.users = true
    error.value = null
    
    try {
      const { data } = await adminApi.getUsers()
      users.value = data
    } catch (err) {
      error.value = 'Failed to fetch users'
      console.error('Error fetching users:', err)
    } finally {
      loading.value.users = false
    }
  }

  async function fetchGroups() {
    loading.value.groups = true
    error.value = null
    
    try {
      const { data } = await adminApi.getGroups()
      groups.value = data
    } catch (err) {
      error.value = 'Failed to fetch groups'
      console.error('Error fetching groups:', err)
    } finally {
      loading.value.groups = false
    }
  }

  async function fetchSpaces() {
    loading.value.spaces = true
    error.value = null
    
    try {
      const { data } = await adminApi.getSpaces()
      spaces.value = data
    } catch (err) {
      error.value = 'Failed to fetch spaces'
      console.error('Error fetching spaces:', err)
    } finally {
      loading.value.spaces = false
    }
  }

  return {
    users,
    groups,
    spaces,
    loading,
    error,
    fetchUsers,
    fetchGroups,
    fetchSpaces
  }
})
