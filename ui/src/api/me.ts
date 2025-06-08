import api from './axios'
import type { Favorite, Space, User, UserPreferences } from './interface'

const meApiUrl = '/v1/me'

interface MyFavoritesResponse {
  favorites: Favorite[]
}

interface MySpacesResponse {
  spaces: Space[]
}

export const meApi = {
  getMySpaces: () => {
    return api.get<MySpacesResponse>(`${meApiUrl}/spaces`)
  },

  getMyFavorites: () => {
    return api.get<MyFavoritesResponse>(`${meApiUrl}/favorites`)
  },

  getProfile: () => {
    return api.get<User>(`${meApiUrl}/profile`)
  },

  addFavorite: (documentId: string) => {
    return api.post(`${meApiUrl}/favorites/${documentId}`, {})
  },

  unFavorite: (documentId: string) => {
    return api.delete(`${meApiUrl}/favorites/${documentId}`)
  },

  getPreferences: () => {
    return api.get<UserPreferences>(`${meApiUrl}/preferences`)
  },

  updatePreferences: (preferences: UserPreferences) => {
    return api.put(`${meApiUrl}/preferences`, preferences)
  },
}