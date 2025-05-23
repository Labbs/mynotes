<script setup lang="ts">
import { onMounted, ref, computed, watch } from 'vue'
import { useSpaceStore } from '../../stores/space'
import { useDocumentStore } from '../../stores/document'
import { useRouter, useRoute } from 'vue-router'

// Use defineProps without assigning to a variable
const { isCollapsed, isHovered } = defineProps<{
  isCollapsed: boolean,
  isHovered: boolean
}>()

const spaceStore = useSpaceStore()
const documentStore = useDocumentStore()
const expandedSpaces = ref<Set<string>>(new Set())
const router = useRouter()
const route = useRoute()
const currentSlug = computed(() => route.params.slug as string)

// Chargement des espaces développés depuis le localStorage
const loadExpandedSpaces = () => {
  try {
    const savedExpandedSpaces = localStorage.getItem('mynotes_expanded_spaces')
    if (savedExpandedSpaces) {
      expandedSpaces.value = new Set(JSON.parse(savedExpandedSpaces))
    }
  } catch (err) {
    console.error('Failed to load expanded spaces from localStorage:', err)
  }
}

// Sauvegarde des espaces développés dans le localStorage
const saveExpandedSpaces = () => {
  try {
    localStorage.setItem(
      'mynotes_expanded_spaces',
      JSON.stringify([...expandedSpaces.value])
    )
  } catch (err) {
    console.error('Failed to save expanded spaces to localStorage:', err)
  }
}

// Observer les changements dans expandedSpaces et les sauvegarder
watch(expandedSpaces.value, saveExpandedSpaces, { deep: true })

const capitalizeFirst = (str: string) => {
  if (!str) return ''
  return str.charAt(0).toUpperCase() + str.slice(1)
}

const toggleSpace = async (spaceId: string) => {
  if (expandedSpaces.value.has(spaceId)) {
    expandedSpaces.value.delete(spaceId)
  } else {
    expandedSpaces.value.add(spaceId)
    await documentStore.fetchDocumentsBySpace(spaceId)
  }
  saveExpandedSpaces() // Sauvegarder après la modification
}

const createDocument = async (spaceId: string) => {
  try {
    const doc = await documentStore.createDocument({
      name: 'New document',
      space_id: spaceId
    })
    
    // Ouvrir le space s'il n'est pas déjà ouvert
    if (!expandedSpaces.value.has(spaceId)) {
      expandedSpaces.value.add(spaceId)
      saveExpandedSpaces() // Sauvegarder après la modification
    }
    
    // On force le rafraîchissement des documents
    await documentStore.fetchDocumentsBySpace(spaceId, true)
    
    // Rediriger vers le nouveau document
    router.push(`/d/${doc.slug}`)
  } catch (err) {
    console.error('Failed to create document:', err)
  }
}

const showSpaceMenu = (spaceId: string) => {
  // TODO: Implement space menu
  console.log('Space menu clicked for space:', spaceId)
}

onMounted(async () => {
  await spaceStore.fetchSpaces()
  loadExpandedSpaces() // Charger les espaces développés au montage du composant
  
  // Charger les documents pour tous les espaces déjà ouverts
  const loadDocumentsForExpandedSpaces = async () => {
    const promises = Array.from(expandedSpaces.value).map(spaceId => 
      documentStore.fetchDocumentsBySpace(spaceId)
    )
    await Promise.all(promises)
  }
  
  await loadDocumentsForExpandedSpaces()
  
  // Ouvrir automatiquement le space du document actuellement ouvert
  const openCurrentDocumentSpace = async () => {
    if (currentSlug.value) {
      try {
        await documentStore.fetchDocumentBySlug(currentSlug.value)
        if (documentStore.currentDocument && documentStore.currentDocument.space_id) {
          if (!expandedSpaces.value.has(documentStore.currentDocument.space_id)) {
            expandedSpaces.value.add(documentStore.currentDocument.space_id)
            await documentStore.fetchDocumentsBySpace(documentStore.currentDocument.space_id)
            saveExpandedSpaces()
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
      <div 
        class="group flex items-center"
      >
        <button
          @click="toggleSpace(space.id)"
          class="flex flex-1 items-center gap-x-2 rounded-lg px-1 py-1 text-gray-500 hover:bg-gray-100 hover:text-gray-700"
          :class="{
            'justify-center': isCollapsed && !isHovered
          }"
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
              v-show="(!isCollapsed || isHovered)"
              class="absolute inset-0 h-5 w-5 opacity-0 group-hover:opacity-100 transition-opacity"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
            >
              <path 
                :d="expandedSpaces.has(space.id) ? 'M19 9l-7 7-7-7' : 'M9 5l7 7-7 7'"
                stroke-width="2" 
                stroke-linecap="round" 
                stroke-linejoin="round"
              />
            </svg>
          </div>

          <span 
            v-show="!isCollapsed || isHovered"
            class="flex-grow text-left"
          >
            {{ capitalizeFirst(space.name) }}
          </span>
        </button>

        <!-- Action buttons -->
        <div 
          v-show="!isCollapsed || isHovered"
          class="flex gap-1 px-1 opacity-0 group-hover:opacity-100 transition-opacity"
        >
          <button
            class="p-1 rounded hover:bg-gray-100 text-gray-400 hover:text-gray-600"
            @click="createDocument(space.id)"
          >
            <svg class="w-3 h-3" viewBox="0 0 24 24" fill="none" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
            </svg>
          </button>
          
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
        v-if="expandedSpaces.has(space.id) && (!isCollapsed || isHovered)"
        class="ml-4 mt-1 space-y-1"
      >
        <div v-if="documentStore.loadingSpaces.has(space.id)" class="text-[14px] text-gray-500 px-2">
          Loading...
        </div>
        <div v-else-if="!documentStore.documentsBySpace[space.id]?.length" class="text-[13px] text-gray-500 px-2">
          No documents
        </div>
        <router-link
          v-else
          v-for="doc in documentStore.documentsBySpace[space.id]"
          :key="doc.id"
          :to="`/d/${doc.slug}`"
          class="flex items-center gap-x-2 rounded-lg px-2 py-1 text-[13px] text-gray-500 hover:bg-gray-100 hover:text-gray-700"
          :class="{ 
            'bg-blue-50 text-blue-600 font-medium': doc.slug === currentSlug 
          }"
        >
          <div 
            class="size-4 rounded bg-gray-100 flex items-center justify-center opacity-75 text-xs font-medium"
            :class="{ 'bg-blue-100': doc.slug === currentSlug }"
          >
            {{ doc.name?.[0]?.toUpperCase() }}
          </div>
          <span>{{ capitalizeFirst(doc.name) }}</span>
        </router-link>
      </div>
    </div>
  </div>
</template>