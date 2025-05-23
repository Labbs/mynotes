import api from './axios'

const meApiUrl = '/v1/me'

export interface MySpace {
  id: string
  name: string
  description?: string
  slug?: string
  icon?: string
  icon_color?: string
  created_at?: string
  updated_at?: string
}

export interface MyFavorite {
  id: string
  user_id: string
  document_id?: string
  database_id?: string
  position?: string
  created_at?: string
}

interface MyFavoritesResponse {
  favorites: MyFavorite[]
}

interface MySpacesResponse {
  spaces: MySpace[]
}

export const meApi = {
  getMySpaces: () => {
    return api.get<MySpacesResponse>(`${meApiUrl}/spaces`)
  },

  getMyFavorites: () => {
    return api.get<MyFavoritesResponse>(`${meApiUrl}/favorites`)
  },
}