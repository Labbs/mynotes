import { defineStore } from 'pinia'
import { ref } from 'vue'
import { documentApi, type Document, type CreateDocumentParams, type Config } from '../api/document'

export const useDocumentStore = defineStore('document', () => {
  const currentDocument = ref<Document | null>(null)
  const documentsBySpace = ref<Record<string, Document[]>>({})
  const documentsByParent = ref<Record<string, Document[]>>({}) // Pour stocker les sous-documents par parent
  const documentsWithChildren = ref<Set<string>>(new Set()) // Pour stocker les documents qui ont des enfants
  const listExcalidrawLibs = ref<string[]>([]) // Pour stocker les bibliothèques Excalidraw
  const loadingSpaces = ref<Set<string>>(new Set())
  const loadingDocument = ref(false)
  const error = ref<string | null>(null)
  
  // Pour vérifier si un document a des enfants
  function hasChildren(documentId: string): boolean {
    return documentsWithChildren.value.has(documentId)
  }
  
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
        type: data.type,
        hasContent: !!data.content, 
        contentLength: data.content ? data.content.length : 0
      });
      
      // Make sure to create a new object to trigger reactivity
      currentDocument.value = { ...data };
      
      // Ensure content is properly initialized based on document type
      if (data.type === 'excalidraw') {
        if (!data.content || data.content === 'null') {
          // Initialize empty Excalidraw content with proper structure
          currentDocument.value.content = JSON.stringify({ 
            elements: [], 
            appState: { 
              viewBackgroundColor: '#ffffff',
              currentItemFontFamily: 1,
              gridSize: 20
            }, 
            files: {} 
          });
          console.log('Initialized empty Excalidraw content');
          
          // Save the initialized content to the server
          await updateDocument({
            id: data.id,
            content: currentDocument.value.content,
            name: data.name,
            space_id: data.space_id,
            config: data.config
          });
          console.log('Saved initialized Excalidraw content to server');
        } else {
          try {
            // Validate the existing content format
            const parsed = JSON.parse(data.content);
            const isValid = parsed && 
                          (Array.isArray(parsed.elements)) && 
                          (typeof parsed.appState === 'object') &&
                          (typeof parsed.files === 'object');
            
            if (!isValid) {
              // Fix malformed content
              console.warn('Malformed Excalidraw content detected, fixing structure');
              currentDocument.value.content = JSON.stringify({ 
                elements: Array.isArray(parsed?.elements) ? parsed.elements : [], 
                appState: typeof parsed?.appState === 'object' ? parsed.appState : { 
                  viewBackgroundColor: '#ffffff',
                  currentItemFontFamily: 1,
                  gridSize: 20
                }, 
                files: typeof parsed?.files === 'object' ? parsed.files : {} 
              });
            }
          } catch (e) {
            console.error('Error parsing existing Excalidraw content, resetting:', e);
            currentDocument.value.content = JSON.stringify({ 
              elements: [], 
              appState: { 
                viewBackgroundColor: '#ffffff',
                currentItemFontFamily: 1,
                gridSize: 20
              }, 
              files: {} 
            });
          }
        }
      } else if (!data.content) {
        console.warn('Document loaded without content!');
        // For non-Excalidraw documents, initialize with empty content
        currentDocument.value.content = '';
      } else {
        console.log('Document content loaded successfully, length:', data.content.length);
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

  async function fetchDocumentsByParentDocument(spaceId: string, documentId: string) {
    if (!spaceId || !documentId) return

    try {
      const { data } = await documentApi.getDocumentsByParentDocument(spaceId, documentId)
      documentsByParent.value[documentId] = data
      
      // Mettre à jour documentsWithChildren
      if (data && data.length > 0) {
        documentsWithChildren.value.add(documentId)
      } else {
        documentsWithChildren.value.delete(documentId)
      }
    } catch (err) {
      console.error('Error fetching documents by parent document:', err)
      error.value = 'Failed to fetch documents by parent document'
    }
  }

  async function updateDocumentConfig(config: Config) {
    if (!currentDocument.value) return

    try {
      console.log('Store: Updating document config:', config)
      // Mise à jour du document en utilisant updateDocument au lieu de updateDocumentConfig
      const updatedDocument = {
        id: currentDocument.value.id,
        name: currentDocument.value.name,
        space_id: currentDocument.value.space_id,
        config: config
      }
      const result = await updateDocument(updatedDocument)
      console.log('Store: Document config updated successfully:', result)
      // currentDocument.value est déjà mis à jour dans updateDocument
    } catch (err) {
      console.error('Error updating document config:', err)
      error.value = 'Failed to update document config'
    }
  }

  async function createDocument(params: CreateDocumentParams) {
    try {
      // Initialize content for Excalidraw documents
      let initialContent = undefined;
      if (params.type === 'excalidraw') {
        // Create proper Excalidraw initial structure with required appState properties
        initialContent = JSON.stringify({ 
          elements: [], 
          appState: { 
            viewBackgroundColor: '#ffffff',
            currentItemFontFamily: 1,
            gridSize: 20
          }, 
          files: {} 
        });
        console.log('Initializing new Excalidraw document with proper empty content');
      }
      
      const { data } = await documentApi.create({
        ...params,
        content: initialContent
      });
      
      return data;
    } catch (error) {
      throw error;
    }
  }

  async function updateDocument({ id, name, content, space_id, config }: Partial<Document>) {
    if (!id) return null;
    
    console.log('Store: Updating document:', { id, name, contentChanged: !!content, space_id });
    
    try {
      const { data } = await documentApi.updateDocument(id, { 
        id,
        name,
        space_id: space_id || currentDocument.value?.space_id,
        content,
        config
      });
      
      // Update document in space list if needed
      if (data.space_id) {
        const spaceId = data.space_id;
        const index = documentsBySpace.value[spaceId]?.findIndex(d => d.id === id);
        if (index !== -1 && documentsBySpace.value[spaceId]) {
          // Only update non-content properties in the space list
          const { content: _, ...rest } = data;
          documentsBySpace.value[spaceId][index] = {
            ...documentsBySpace.value[spaceId][index],
            ...rest
          };
        }
      }
      
      // Only update non-content properties for current document to prevent refresh
      if (currentDocument.value && currentDocument.value.id === id) {
        const { content: serverContent, ...otherProps } = data;
        
        // Update metadata properties but keep local content
        Object.assign(currentDocument.value, {
          ...otherProps,
          content: currentDocument.value.content // Preserve local content
        });
      }
      
      return data;
    } catch (error) {
      console.error('Store: Error updating document:', error);
      throw error;
    }
  }

  async function deleteDocument(documentId: string) {
    if (!documentId) return null;

    try {
      await documentApi.deleteDocument(documentId);
      // Supprimer le document de la liste des documents par espace
      Object.keys(documentsBySpace.value).forEach(spaceId => {
        documentsBySpace.value[spaceId] = documentsBySpace.value[spaceId].filter(doc => doc.id !== documentId);
      });
      // Supprimer le document de la liste des documents par parent
      Object.keys(documentsByParent.value).forEach(parentId => {
        documentsByParent.value[parentId] = documentsByParent.value[parentId].filter(doc => doc.id !== documentId);
      });
      // Supprimer le document de currentDocument si c'est le même
      if (currentDocument.value && currentDocument.value.id === documentId) {
        currentDocument.value = null;
      }
    } catch (error) {
      console.error('Store: Error deleting document:', error);
      throw error;
    }
  }

  // Fetch Excalidraw libraries
  async function fetchExcalidrawLibs() {
    try {
      const { data } = await documentApi.listExcalidrawLibs();
      listExcalidrawLibs.value = data;
    } catch (error) {
      console.error('Error fetching Excalidraw libraries:', error);
      throw error;
    }
  }

  return {
    currentDocument,
    documentsBySpace,
    documentsByParent,
    documentsWithChildren,
    hasChildren,
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
    fetchDocumentsByParentDocument,
    deleteDocument,
    fetchExcalidrawLibs,
  }
})