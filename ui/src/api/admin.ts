import api from './axios'
import type { User, Group, Space } from './interface'

const adminApiUrl = '/v1/admin'

export const adminApi = {
  getUsers: () => {
    return api.get<User[]>(`${adminApiUrl}/users`)
  },

  getGroups: () => {
    return api.get<Group[]>(`${adminApiUrl}/groups`)
  },

  getSpaces: () => {
    return api.get<Space[]>(`${adminApiUrl}/spaces`)
  }
}
