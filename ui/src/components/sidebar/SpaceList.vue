<script setup lang="ts">
import { onMounted, onUnmounted, ref, computed, watch } from 'vue'
import { useSpaceStore } from '../../stores/space'
import { useDocumentStore } from '../../stores/document'
import { useRouter, useRoute } from 'vue-router'
import DocumentList from './DocumentList.vue'

// Props pour le composant
defineProps<{
  isCollapsed?: boolean,
  isHovered?: boolean
}>()

const spaceStore = useSpaceStore()
const documentStore = useDocumentStore()
// Utiliser des arrays plutôt que des Set pour éviter les problèmes de typage
const expandedSpaceIds = ref<string[]>([])
const expandedDocumentIds = ref<string[]>([])
const activeCreateMenu = ref<string | null>(null)
// Ajouter un Set pour suivre les documents survolés
const hoveredDocIds = ref<Set<string>>(new Set())
const router = useRouter()
const route = useRoute()
const currentSlug = computed(() => route.params.slug as string)

// Chargement des espaces développés depuis le localStorage
const loadExpandedItems = () => {
  try {
    const savedExpandedSpaces = localStorage.getItem('mynotes_expanded_spaces')
    if (savedExpandedSpaces) {
      expandedSpaceIds.value = JSON.parse(savedExpandedSpaces)
    }
    
    const savedExpandedDocuments = localStorage.getItem('mynotes_expanded_documents')
    if (savedExpandedDocuments) {
      expandedDocumentIds.value = JSON.parse(savedExpandedDocuments)
    }
  } catch (err) {
    console.error('Failed to load expanded items from localStorage:', err)
  }
}

// Sauvegarde des espaces développés dans le localStorage
const saveExpandedSpaces = () => {
  try {
    localStorage.setItem(
      'mynotes_expanded_spaces',
      JSON.stringify(expandedSpaceIds.value)
    )
  } catch (err) {
    console.error('Failed to save expanded spaces to localStorage:', err)
  }
}

// Sauvegarde des documents développés dans le localStorage
const saveExpandedDocuments = () => {
  try {
    localStorage.setItem(
      'mynotes_expanded_documents',
      JSON.stringify(expandedDocumentIds.value)
    )
  } catch (err) {
    console.error('Failed to save expanded documents to localStorage:', err)
  }
}

// Observer les changements dans expandedSpaceIds et expandedDocumentIds
watch(expandedSpaceIds, saveExpandedSpaces, { deep: true })
watch(expandedDocumentIds, saveExpandedDocuments, { deep: true })

// Vérifier les documents parents après chargement des espaces
watch(() => documentStore.documentsBySpace, (newVal) => {
  // Pour chaque espace
  for (const spaceId in newVal) {
    // Parcourir les documents pour trouver ceux avec parent_id
    for (const doc of newVal[spaceId]) {
      if (doc.parent_id && doc.parent_id !== '') {
        // Marquer le parent comme ayant des enfants
        documentStore.documentsWithChildren.add(doc.parent_id)
      }
    }
  }
}, { deep: true })

const capitalizeFirst = (str: string) => {
  if (!str) return ''
  return str.charAt(0).toUpperCase() + str.slice(1)
}

// Fonctions pour gérer le hover des documents
// const setHovered = (docId: string, isHovered: boolean) => {
//   if (isHovered) {
//     hoveredDocIds.value.add(docId)
//   } else {
//     hoveredDocIds.value.delete(docId)
//   }
// }

const isDocHovered = (docId: string) => {
  return hoveredDocIds.value.has(docId)
}

const isSpaceExpanded = (spaceId: string) => {
  return expandedSpaceIds.value.includes(spaceId)
}

const isDocumentExpanded = (documentId: string) => {
  return expandedDocumentIds.value.includes(documentId)
}

const toggleSpace = async (spaceId: string) => {
  if (isSpaceExpanded(spaceId)) {
    expandedSpaceIds.value = expandedSpaceIds.value.filter(id => id !== spaceId)
  } else {
    expandedSpaceIds.value.push(spaceId)
    await documentStore.fetchDocumentsBySpace(spaceId)
  }
}

const toggleDocument = async (spaceId: string, documentId: string) => {
  if (isDocumentExpanded(documentId)) {
    expandedDocumentIds.value = expandedDocumentIds.value.filter(id => id !== documentId)
  } else {
    expandedDocumentIds.value.push(documentId)
    
    // Toujours recharger les sous-documents lorsqu'on étend un document
    await documentStore.fetchDocumentsByParentDocument(spaceId, documentId)
  }
}

const createDocument = async (spaceId: string, parentId: string, type: 'document' | 'database' = 'document') => {
  try {
    const doc = await documentStore.createDocument({
      name: `New ${type}`,
      space_id: spaceId,
      parent_id: parentId,
      type
    })
    
    // Ouvrir le space s'il n'est pas déjà ouvert
    if (!isSpaceExpanded(spaceId)) {
      expandedSpaceIds.value.push(spaceId)
    }
    
    // Si c'est un sous-document, ouvrir le parent s'il n'est pas déjà ouvert
    if (parentId && !isDocumentExpanded(parentId)) {
      expandedDocumentIds.value.push(parentId)
      
      // Marquer le parent comme ayant des enfants
      documentStore.documentsWithChildren.add(parentId)
    }
    
    // Forcer le rafraîchissement des documents
    await documentStore.fetchDocumentsBySpace(spaceId, true)
    
    // Si c'est un sous-document, charger les sous-documents du parent
    if (parentId) {
      await documentStore.fetchDocumentsByParentDocument(spaceId, parentId)
    }
    
    // Rediriger vers le nouveau document
    router.push(`/d/${doc.slug}`)
  } catch (err) {
    console.error('Failed to create document:', err)
  }
}

const showSpaceMenu = (spaceId: string) => {
  // Code pour afficher le menu de l'espace
  console.log('Show space menu for', spaceId)
}

// Fonction pour gérer la création de différents types de documents
const handleCreateDocumentType = (spaceId: string, parentId: string, type: 'document' | 'database') => {
  activeCreateMenu.value = null
  createDocument(spaceId, parentId, type)
}

// Fermer le menu si on clique ailleurs dans la page
const handleDocumentClick = () => {
  activeCreateMenu.value = null
}

// Ajouter et supprimer les écouteurs d'événements quand le composant est monté/démonté
onMounted(() => {
  document.addEventListener('click', handleDocumentClick)
})

onUnmounted(() => {
  document.removeEventListener('click', handleDocumentClick)
})

onMounted(async () => {
  await spaceStore.fetchSpaces()
  loadExpandedItems() // Charger les espaces et documents développés
  
  // Charger les documents pour tous les espaces déjà ouverts
  const loadDocumentsForExpandedSpaces = async () => {
    const promises = expandedSpaceIds.value.map(spaceId => 
      documentStore.fetchDocumentsBySpace(spaceId)
    )
    await Promise.all(promises)
  }
  
  await loadDocumentsForExpandedSpaces()
  
  // Charger les sous-documents pour tous les documents déjà ouverts
  const loadDocumentsForExpandedDocuments = async () => {
    const promises = []
    
    // Pour chaque espace ouvert
    for (const spaceId of expandedSpaceIds.value) {
      // Pour chaque document ouvert dans cet espace
      if (documentStore.documentsBySpace[spaceId]) {
        for (const doc of documentStore.documentsBySpace[spaceId]) {
          if (expandedDocumentIds.value.includes(doc.id)) {
            // Force le rechargement des sous-documents pour chaque document expanded
            promises.push(documentStore.fetchDocumentsByParentDocument(spaceId, doc.id))
          }
        }
      }
    }
    
    await Promise.all(promises)
  }
  
  await loadDocumentsForExpandedDocuments()
  
  // Ouvrir automatiquement le space du document actuellement ouvert
  const openCurrentDocumentSpace = async () => {
    if (currentSlug.value) {
      try {
        await documentStore.fetchDocumentBySlug(currentSlug.value)
        if (documentStore.currentDocument && documentStore.currentDocument.space_id) {
          const spaceId = documentStore.currentDocument.space_id
          
          // Ouvrir l'espace si nécessaire
          if (!isSpaceExpanded(spaceId)) {
            expandedSpaceIds.value.push(spaceId)
            await documentStore.fetchDocumentsBySpace(spaceId)
          }
          
          // Si le document actuel a un parent_id, on ouvre ce parent
          if (documentStore.currentDocument.parent_id) {
            const parentId = documentStore.currentDocument.parent_id
            if (!isDocumentExpanded(parentId)) {
              expandedDocumentIds.value.push(parentId)
              await documentStore.fetchDocumentsByParentDocument(spaceId, parentId)
            }
          }
        }
      } catch (err) {
        console.error('Error opening current document space:', err)
      }
    }
  }
  
  await openCurrentDocumentSpace()
})
</script>

<template>
  <div class="pt-2 pb-1 text-[13px]">
    <div v-for="space in spaceStore.spaces" :key="space.id">
      <div class="group flex items-center">
        <button
          @click.stop.prevent="toggleSpace(space.id)"
          class="flex flex-1 items-center gap-x-2 rounded-lg px-1 py-1 text-gray-500 hover:bg-gray-100 hover:text-gray-700"
        >
          <div class="relative flex-shrink-0">
            <!-- Space icon (visible by default, hidden on hover) -->
            <div class="group-hover:opacity-0 transition-opacity">
              <img 
                v-if="space.icon"
                :src="space.icon" 
                class="size-5 opacity-75"
                :alt="space.name" 
              />
              <div 
                v-else 
                class="size-5 rounded-lg bg-gray-100 flex items-center justify-center opacity-75"
              >
                {{ space.name?.[0]?.toUpperCase() }}
              </div>
            </div>

            <!-- Chevron icon (hidden by default, visible on hover) -->
            <svg
              class="absolute inset-0 h-5 w-5 opacity-0 group-hover:opacity-100 transition-opacity"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
            >
              <path 
                :d="isSpaceExpanded(space.id) ? 'M19 9l-7 7-7-7' : 'M9 5l7 7-7 7'"
                stroke-width="2" 
                stroke-linecap="round" 
                stroke-linejoin="round"
              />
            </svg>
          </div>

          <span 
            class="flex-grow text-left truncate overflow-hidden text-ellipsis"
          >
            {{ capitalizeFirst(space.name) }}
          </span>
        </button>

        <!-- Action buttons -->
        <div 
          class="flex gap-1 px-1 opacity-0 group-hover:opacity-100 transition-opacity relative"
        >
          <button
            class="p-1 rounded hover:bg-gray-100 text-gray-400 hover:text-gray-600"
            @click.stop.prevent="activeCreateMenu === space.id ? activeCreateMenu = null : activeCreateMenu = space.id"
          >
            <svg class="w-3 h-3" viewBox="0 0 24 24" fill="none" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
            </svg>
          </button>
          
          <!-- Menu pour créer différents types de documents -->
          <div 
            v-if="activeCreateMenu === space.id" 
            class="absolute z-50 top-full right-0 w-40 bg-white rounded-md shadow-lg py-1 text-sm"
            @click.stop.prevent
          >
            <button 
              @click.stop.prevent="handleCreateDocumentType(space.id, '', 'document')"
              class="flex w-full items-center px-4 py-2 text-gray-700 hover:bg-gray-100"
            >
              <svg class="w-4 h-4 mr-2 text-gray-500" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
              </svg>
              Document
            </button>
            

            <button 
              @click.stop.prevent="handleCreateDocumentType(space.id, '', 'database')"
              class="flex w-full items-center px-4 py-2 text-gray-700 hover:bg-gray-100"
            >
              <svg class="w-4 h-4 mr-2 text-gray-500" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4m0 5c0 2.21-3.582 4-8 4s-8-1.79-8-4" />
              </svg>
              Database
            </button>
          </div>
          
          <button
            class="p-1 rounded hover:bg-gray-100 text-gray-400 hover:text-gray-600"
            @click="showSpaceMenu(space.id)"
          >
            <svg class="w-3 h-3" viewBox="0 0 24 24" fill="none" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"/>
            </svg>
          </button>
        </div>
      </div>

      <!-- Documents list -->
      <div 
        v-if="isSpaceExpanded(space.id)"
        class="ml-2 mt-1"
      >
        <div v-if="documentStore.loadingSpaces.has(space.id)" class="text-[14px] text-gray-500 px-2">
          Loading...
        </div>
        <div v-else-if="!documentStore.documentsBySpace[space.id]?.length" class="text-[13px] text-gray-500 px-2">
          No documents
        </div>
        <DocumentList
          v-else
          :documents="documentStore.documentsBySpace[space.id].filter(doc => !doc.parent_id || doc.parent_id === '')"
          :space-id="space.id"
          :expanded-document-ids="expandedDocumentIds"
          :is-hovered="isDocHovered"
          @toggle-document="toggleDocument"
          @create-document="createDocument"
        />
      </div>
    </div>
  </div>
</template>