<script setup lang="ts">
import { computed, ref } from 'vue'
import { useDocumentStore } from '../../stores/document'
import { useRoute } from 'vue-router'

const props = defineProps<{
  documents: any[]
  spaceId: string
  expandedDocumentIds: string[]
  isCollapsed: boolean
  isHovered: boolean
}>()

const emit = defineEmits<{
  (e: 'toggleDocument', spaceId: string, documentId: string): void
  (e: 'createDocument', spaceId: string, parentId: string): void
}>()

const documentStore = useDocumentStore()
const route = useRoute()
const currentSlug = computed(() => route.params.slug as string)

// Garder une trace des documents qui sont survolés
const hoveredDocIds = ref<Set<string>>(new Set())

// Fonctions pour gérer le hover
const setHovered = (docId: string, isHovered: boolean) => {
  if (isHovered) {
    hoveredDocIds.value.add(docId)
  } else {
    hoveredDocIds.value.delete(docId)
  }
}

const isHovered = (docId: string) => {
  return hoveredDocIds.value.has(docId)
}

const capitalizeFirst = (str: string) => {
  if (!str) return ''
  return str.charAt(0).toUpperCase() + str.slice(1)
}

const isDocumentExpanded = (documentId: string) => {
  const isExpanded = props.expandedDocumentIds.includes(documentId)
  if (isExpanded && !documentStore.documentsByParent[documentId]) {
    documentStore.fetchDocumentsByParentDocument(props.spaceId, documentId)
  }
  return isExpanded
}

const toggleDocument = (documentId: string, event: Event) => {
  if (event) {
    event.preventDefault()
    event.stopPropagation()
  }
  emit('toggleDocument', props.spaceId, documentId)
}

const createDocument = (parentId: string, event: Event) => {
  if (event) {
    event.preventDefault()
    event.stopPropagation()
  }
  emit('createDocument', props.spaceId, parentId)
}
</script>

<template>
  <div class="space-y-1">
    <div 
      v-for="doc in documents" 
      :key="doc.id"
      class="relative"
    >
      <div 
        class="relative flex items-center w-full group hover:bg-gray-50 rounded-lg"
        @mouseenter="setHovered(doc.id, true)"
        @mouseleave="setHovered(doc.id, false)"
      >
        <div class="w-4 flex items-center justify-center flex-shrink-0" v-show="hoveredDocIds.has(doc.id)">
          <button
            @click="(e) => toggleDocument(doc.id, e)"
            class="flex flex-1 items-center gap-x-2 rounded-lg px-1 py-1 text-gray-500 hover:bg-gray-100 hover:text-gray-700"
            title="Toggle sub-documents"
          >
            <svg
              class="inset-0 h-5 w-5 opacity-0 group-hover:opacity-100 transition-opacity"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
            >
              <path 
                :d="isDocumentExpanded(doc.id) ? 'M19 9l-7 7-7-7' : 'M9 5l7 7-7 7'"
                stroke-width="2" 
                stroke-linecap="round" 
                stroke-linejoin="round"
              />
            </svg>
          </button>
        </div>
        
        <router-link
          :to="`/d/${doc.slug}`"
          class="flex items-center gap-x-2 rounded-lg px-2 py-1 text-[13px] text-gray-500 hover:bg-gray-100 hover:text-gray-700 overflow-hidden w-full"
          :class="{ 
            'bg-blue-50 text-blue-600 font-medium': doc.slug === currentSlug,
          }"
        >
          <div class="flex-shrink-0" v-show="!hoveredDocIds.has(doc.id)">
            <!-- Document icon -->
            <div v-if="!doc.config.icon"
              class="size-4 rounded bg-gray-100 flex items-center justify-center opacity-75 text-xs font-medium"
              :class="{ 'bg-blue-100': doc.slug === currentSlug }"
            >
              {{ doc.name?.[0]?.toUpperCase() }}
            </div>
            <div v-else>
              {{ doc.config.icon }}
            </div>
          </div>
          <span class="truncate overflow-hidden text-ellipsis block min-w-0 flex-grow">
            {{ capitalizeFirst(doc.name) }}
          </span>
        </router-link>
        
        <!-- Action buttons pour document (en position absolue) -->
        <div 
          v-if="!isCollapsed || isHovered"
          class="absolute right-1 flex items-center h-full transition-opacity"
          :class="{ 'opacity-100': isHovered(doc.id), 'opacity-0': !isHovered(doc.id) }"
        >
          <button
            class="p-1 rounded hover:bg-gray-100 text-gray-400 hover:text-gray-600 bg-white bg-opacity-90 shadow-sm"
            @click="(e) => createDocument(doc.id, e)"
            title="Create sub-document"
          >
            <svg class="w-3 h-3" viewBox="0 0 24 24" fill="none" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
            </svg>
          </button>
        </div>
      </div>
      
      <!-- Sous-documents (récursif) -->
      <div 
        v-if="isDocumentExpanded(doc.id)"
        class="ml-2 mt-1"
      >
        <div v-if="!documentStore.documentsByParent?.[doc.id]?.length" class="text-[13px] text-gray-500 px-2">
          No documents
        </div>
        <DocumentList
          v-else
          :documents="documentStore.documentsByParent[doc.id]"
          :space-id="spaceId"
          :expanded-document-ids="expandedDocumentIds"
          :is-collapsed="isCollapsed"
          :is-hovered="isHovered"
          @toggle-document="(spaceId, docId) => $emit('toggleDocument', spaceId, docId)"
          @create-document="(spaceId, parentId) => $emit('createDocument', spaceId, parentId)"
        />
      </div>
    </div>
  </div>
</template>
