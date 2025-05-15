import api from './axios'

export interface Space {
  id: string
  name: string
  description?: string
  slug?: string
  icon?: string
  organisation_owner_id?: string
  user_owner_id?: string
  created_at?: string
}

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