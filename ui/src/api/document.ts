import api from './axios'
import type { Document } from './interface'

export interface CreateDocumentParams {
  name: string
  space_id: string
  parent_id?: string
  type?: 'document' | 'excalidraw' | 'database'
  content?: string
}

export const documentApi = {
  getSpaceDocuments: (spaceId: string) =>
    api.get<Document[]>(`/v1/document/space/${spaceId}`),

  getDocumentBySlug: (slug: string) =>
    api.get<Document>(`/v1/document/slug/${slug}`),

  getDocumentsBySpace: (spaceId: string) =>
    api.get<Document[]>(`/v1/document/space/${spaceId}`),

  create: (params: CreateDocumentParams) =>
    api.post<Document>('/v1/document', params),

  updateDocument: (documentId: string, params: Partial<Document>) =>
    api.put<Document>(`/v1/document/${documentId}`, params),

  getDocumentsByParentDocument: (spaceId: string, documentId: string) =>
    api.get<Document[]>(`/v1/document/space/${spaceId}/parent/${documentId}`),

  deleteDocument: (documentId: string) =>
    api.delete(`/v1/document/${documentId}`),

  listExcalidrawLibs: () => 
    api.get<string[]>('/v1/document/excalidraw/list/libs'),

  getExcalidrawLib: (name: string) =>
    api.get<{ libraryItems: any[] }>(`/v1/document/excalidraw/libraries/${name}`, {
      headers: {
        'Accept': 'application/json'
      },
      transformResponse: [(data) => {
        // Assurez-vous que la réponse est bien du JSON
        try {
          return typeof data === 'string' ? JSON.parse(data) : data;
        } catch (e) {
          console.error('Failed to parse library data:', e);
          return { libraryItems: [] };
        }
      }]
    }),
}
