import api from './axios'

export interface Document {
  id: string
  name: string
  space_id: string
  type: 'document' | 'folder'
  parent_id?: string
  slug?: string
  // properties: Property[]
  config: Config
  content?: string
  is_favorite?: boolean
  created_at?: string
  updated_at?: string
}

export interface Config {
  full_width?: boolean
  icon?: string
  lock?: boolean
  header_background?: string
}

export interface CreateDocumentParams {
  name: string
  space_id: string
}

export const documentApi = {
  getSpaceDocuments: (spaceId: string) =>
    api.get<Document[]>(`/v1/document/space/${spaceId}`),

  getDocumentBySlug: (slug: string) =>
    api.get<Document>(`/v1/document/slug/${slug}`),

  getDocumentsBySpace: (spaceId: string) =>
    api.get<Document[]>(`/v1/document/space/${spaceId}`),

  updateDocumentConfig: (documentId: string, config: Config) =>
    api.patch<Document>(`/v1/document/${documentId}/config`, { config }),

  create: (params: CreateDocumentParams) =>
    api.post<Document>('/v1/document', params),

  updateDocument: (documentId: string, params: Partial<Document>) =>
    api.put<Document>(`/v1/document/${documentId}`, params),
}
