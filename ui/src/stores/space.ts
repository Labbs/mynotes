import { defineStore } from 'pinia'
import { spaceApi } from '../api/space'
import type { Space } from '../api/interface'

interface SpaceState {
  spaces: Space[]
  loading: boolean
  error: string | null
}

export const useSpaceStore = defineStore('space', {
  state: (): SpaceState => ({
    spaces: [],
    loading: false,
    error: null
  }),
  
  actions: {
    async fetchSpaces() {
      this.loading = true
      this.error = null
      
      try {
        const { data } = await spaceApi.list()
        console.log('data:', data)
        this.spaces = Array.isArray(data) ? data : data.spaces || []
      } catch (error) {
        this.error = 'Failed to fetch spaces'
        console.error('Error fetching spaces:', error)
      } finally {
        this.loading = false
      }
    },
    
    async createSpace(params: { name: string, description?: string, private: boolean }) {
      try {
        const { data } = await spaceApi.create(params)
        this.spaces.push(data)
        return data
      } catch (error) {
        console.error('Error creating space:', error)
        throw error
      }
    }
  }
})