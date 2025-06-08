import api from './axios'
import type { Space } from './interface'

interface SpacesResponse {
  spaces: Space[]
}

interface CreateSpaceParams {
  name: string
  description?: string
  private: boolean
}

export const spaceApi = {
  create: (params: CreateSpaceParams) => 
    api.post<Space>('/space/create', params),
    
  list: () => 
    api.get<SpacesResponse>('/v1/me/spaces')
}