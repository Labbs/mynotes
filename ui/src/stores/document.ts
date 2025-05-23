import { defineStore } from 'pinia'
import { ref } from 'vue'
import { documentApi, type Document, type CreateDocumentParams, type Config } from '../api/document'

export const useDocumentStore = defineStore('document', () => {
  const currentDocument = ref<Document | null>(null)
  const documentsBySpace = ref<Record<string, Document[]>>({})
  const loadingSpaces = ref<Set<string>>(new Set())
  const loadingDocument = ref(false)
  const error = ref<string | null>(null)

  async function fetchDocument(slug: string) {
    loadingDocument.value = true;
    error.value = null;

    try {
      console.log('Fetching document by slug:', slug);
      const { data } = await documentApi.getDocumentBySlug(slug);
      
      // Debug logging to check what's coming from the server
      console.log('Server response for document:', {
        id: data.id,
        name: data.name,
        hasContent: !!data.content, 
        contentLength: data.content ? data.content.length : 0
      });
      
      // Make sure to create a new object to trigger reactivity
      currentDocument.value = { ...data };
      
      // Force editor update if available
      if (data.content) {
        console.log('Document content loaded successfully, length:', data.content.length);
      } else {
        console.warn('Document loaded without content!');
      }
    } catch (err) {
      error.value = 'Failed to fetch document';
      console.error('Error fetching document:', err);
    } finally {
      loadingDocument.value = false;
    }
  }

  async function fetchDocumentsBySpace(spaceId: string, forceRefresh = false) {
    if (!spaceId) return
    if (documentsBySpace.value[spaceId] && !forceRefresh) return // Return if already loaded and not forcing refresh
    
    loadingSpaces.value.add(spaceId)
    try {
      const { data } = await documentApi.getDocumentsBySpace(spaceId)
      documentsBySpace.value[spaceId] = data
    } catch (err) {
      console.error('Error fetching documents:', err)
      error.value = 'Failed to fetch documents'
    } finally {
      loadingSpaces.value.delete(spaceId)
    }
  }

  async function fetchDocumentById(documentId: string) {
    try {
      const { data } = await documentApi.getDocumentBySlug(documentId)
      currentDocument.value = data
    } catch (err) {
      console.error('Error fetching document by ID:', err)
      error.value = 'Failed to fetch document by ID'
    }
  }

  async function fetchDocumentBySlug(slug: string) {
    try {
      const { data } = await documentApi.getDocumentBySlug(slug)
      currentDocument.value = data
    } catch (err) {
      console.error('Error fetching document by slug:', err)
      error.value = 'Failed to fetch document by slug'
    }
  }

  async function updateDocumentConfig(config: Config) {
    if (!currentDocument.value) return

    try {
      const { data } = await documentApi.updateDocumentConfig(currentDocument.value.id, config)
      currentDocument.value = data
    } catch (err) {
      console.error('Error updating document config:', err)
      error.value = 'Failed to update document config'
    }
  }

  async function createDocument(params: CreateDocumentParams) {
    try {
      const { data } = await documentApi.create(params)
      return data
    } catch (error) {
      throw error
    }
  }

  async function updateDocument({ id, name, content, space_id }: Partial<Document>) {
    if (!id) return null;
    
    console.log('Store: Updating document:', { id, name, contentChanged: !!content, space_id });
    
    // Mise à jour optimiste maintenant que nous sommes sûrs qu'il y a des changements
    if (currentDocument.value && currentDocument.value.id === id) {
      if (content !== undefined) {
        // Pour éviter de perdre la référence, mise à jour en profondeur
        currentDocument.value = {
          ...currentDocument.value,
          content: content,
          name: name !== undefined ? name : currentDocument.value.name
        };
      } else if (name !== undefined) {
        currentDocument.value.name = name;
        // Ajout important : si on met à jour seulement le nom, s'assurer de préserver le contenu
        const documentToUpdate = {
          id,
          name,
          space_id: space_id || currentDocument.value?.space_id,
          // Toujours inclure le contenu actuel, même lors d'une mise à jour de nom uniquement
          content: currentDocument.value.content
        };
        
        try {
          const { data } = await documentApi.updateDocument(id, documentToUpdate);
          
          // Mise à jour dans la liste des documents
          if (data.space_id) {
            const spaceId = data.space_id;
            const index = documentsBySpace.value[spaceId]?.findIndex(d => d.id === id);
            if (index !== -1 && documentsBySpace.value[spaceId]) {
              documentsBySpace.value[spaceId][index] = data;
            }
          }
          
          // Mettre à jour le document actuel tout en préservant le contenu
          if (currentDocument.value && currentDocument.value.id === id) {
            if (data.slug !== currentDocument.value.slug) {
              currentDocument.value.slug = data.slug;
            }
            
            // Important: préserver le contenu
            console.log('Store: Document updated with server response');
          }
          
          return data;
        } catch (error) {
          console.error('Store: Error updating document:', error);
          throw error;
        }
        
        // Retourner pour éviter d'exécuter le reste de la fonction
        return;
      }
    }
    
    try {
      const { data } = await documentApi.updateDocument(id, { 
        id,
        name,
        space_id: space_id || currentDocument.value?.space_id,
        content,
      });
      
      // Mise à jour dans la liste des documents
      if (data.space_id) {
        const spaceId = data.space_id;
        const index = documentsBySpace.value[spaceId]?.findIndex(d => d.id === id);
        if (index !== -1 && documentsBySpace.value[spaceId]) {
          documentsBySpace.value[spaceId][index] = data;
        }
      }
      
      // NE PAS remplacer tout le document actif, juste mettre à jour les propriétés nécessaires
      if (currentDocument.value && currentDocument.value.id === id) {
        // Préserver la référence pour éviter un re-rendu complet
        if (data.slug !== currentDocument.value.slug) {
          currentDocument.value.slug = data.slug;
        }
        // Ne pas mettre à jour le contenu ici si nous sommes en train d'éditer
        console.log('Store: Document updated with server response');
      }
      
      return data;
    } catch (error) {
      console.error('Store: Error updating document:', error);
      throw error;
    }
  }

  return {
    currentDocument,
    documentsBySpace,
    loadingSpaces,
    error,
    fetchDocument,
    fetchDocumentsBySpace,
    loadingDocument,
    updateDocumentConfig,
    createDocument,
    updateDocument,
    fetchDocumentById,
    fetchDocumentBySlug,
  }
})