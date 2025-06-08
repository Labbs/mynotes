<script setup lang="ts">
import { computed, onMounted, watch, ref } from 'vue';
import { useRoute } from 'vue-router';
import { useDocumentStore } from '../stores/document';
import { useFavoriteStore } from '../stores/favorite';
import Header from '../components/document/Header.vue';
import ExcalidrawEditor from '../components/document/excalidraw/ExcalidrawEditor.vue'
import TiptapEditor from '../components/document/tiptap/TiptapEditor.vue'

const route = useRoute()
const documentStore = useDocumentStore();
const favoritesStore = useFavoriteStore();
const currentDocument = computed(() => documentStore.currentDocument)
const isEditingTitle = ref(false)
const editableTitle = ref('')
const configSidebarVisible = ref(false)

// Function to update document icon
const updateDocumentIcon = async (icon: string) => {
  if (!currentDocument.value) return;
  
  currentDocument.value.config.icon = icon;
  
  try {
    await documentStore.updateDocument({
      id: currentDocument.value.id,
      name: currentDocument.value.name,
      content: currentDocument.value.content,
      space_id: currentDocument.value.space_id,
      config: currentDocument.value.config
    });
    await favoritesStore.fetchFavorites();
  } catch (error) {
    console.error('Failed to update document icon:', error);
  }
};

const lockDocument = async (lock: boolean) => {
  if (!currentDocument.value) return;
  
  currentDocument.value.config.lock = lock;
  
  try {
    await documentStore.updateDocument({
      id: currentDocument.value.id,
      name: currentDocument.value.name,
      content: currentDocument.value.content,
      space_id: currentDocument.value.space_id,
      config: currentDocument.value.config
    });
  } catch (error) {
    console.error('Failed to update document lock state:', error);
    if (currentDocument.value) {
      currentDocument.value.config.lock = !lock;
    }
  }
};

const fullWidth = async (fullWidth: boolean) => {
  if (!currentDocument.value) return;
  
  currentDocument.value.config.full_width = fullWidth;
  
  try {
    await documentStore.updateDocument({
      id: currentDocument.value.id,
      name: currentDocument.value.name,
      content: currentDocument.value.content,
      space_id: currentDocument.value.space_id,
      config: currentDocument.value.config
    });
  } catch (error) {
    console.error('Failed to update document full width state:', error);
  }
};

// Function to toggle editor editable state
const toggleEditorEditable = () => {
  // Cette fonction n'est plus nécessaire car les éditeurs gèrent leur état d'édition
  console.log('toggleEditorEditable called but not needed anymore');
};

const favorite = async () => {
  if (favoritesStore.favorites.some(f => f.document?.id === currentDocument.value?.id)) {
    await favoritesStore.unFavorite(currentDocument.value?.id as string);
  } else {
    await favoritesStore.addFavorite(currentDocument.value?.id as string);
  }
}

// Improved fetching when the component is mounted
onMounted(async () => {
  try {
    const slug = route.params.slug as string;
    if (slug) {
      await documentStore.fetchDocument(slug);
    }
  } catch (error) {
    console.error('Error initializing document view:', error);
  }
})

watch(() => route.params.slug, async (newSlug) => {
  if (newSlug && newSlug !== currentDocument.value?.slug) {
    await documentStore.fetchDocument(newSlug as string)
  }
}, { immediate: true })

watch(() => currentDocument.value?.name, (newName) => {
  if (newName) {
    editableTitle.value = newName
  }
}, { immediate: true })

const updateTitle = async () => {
  if (currentDocument.value && editableTitle.value !== currentDocument.value.name) {
    const oldSlug = currentDocument.value.slug;
    const currentContent = currentDocument.value.content;
    
    try {
      await documentStore.updateDocument({
        id: currentDocument.value.id,
        name: editableTitle.value,
        content: currentContent,
        space_id: currentDocument.value.space_id,
        config: currentDocument.value.config
      });
      
      isEditingTitle.value = false;
      
      if (currentDocument.value.slug !== oldSlug) {
        window.history.replaceState(null, '', `/d/${currentDocument.value.slug}`);
        
        if (currentDocument.value) {
          currentDocument.value.content = currentContent;
        }

        await favoritesStore.fetchFavorites();
      }
    } catch (error) {
      console.error('Failed to update document title:', error);
    }
  } else {
    isEditingTitle.value = false;
  }
};

// Function to save document content
const saveDocument = async (content: string) => {
  if (!currentDocument.value) return;
  
  try {
    currentDocument.value.content = content;
    
    await documentStore.updateDocument({
      id: currentDocument.value.id,
      name: currentDocument.value.name,
      content: content,
      space_id: currentDocument.value.space_id,
      config: currentDocument.value.config
    });
  } catch (error) {
    console.error('Failed to save document content:', error);
  }
};
</script>

<template>
  <main class="flex-1 flex flex-col overflow-hidden bg-white">
    <!-- Afficher un indicateur de chargement -->
    <div v-if="documentStore.loadingDocument" class="flex justify-center items-center h-screen">
      <div class="animate-pulse text-xl text-gray-500">Loading the document</div>
    </div>
    
    <!-- Afficher le contenu une fois chargé -->
    <template v-else-if="currentDocument">
      <div class="flex flex-col h-full">
        <div 
          :class="{ 'mr-72': configSidebarVisible }" 
          class="transition-all duration-300 ease-in-out w-full flex flex-col flex-1"
        >
          <Header 
            :current-document="currentDocument" 
            :is-editing-title="isEditingTitle" 
            :editable-title="editableTitle"
            @update-title="updateTitle"
            @favorite="favorite"
            @update:is-editing-title="(value) => { 
              if (!currentDocument?.config?.lock || !value) {
                isEditingTitle = value;
              }
            }"
            @update:editable-title="editableTitle = $event"
            @setDocumentIcon="updateDocumentIcon"
            @toggleEditorEditable="toggleEditorEditable"
            @iconConfigVisibility="configSidebarVisible = $event"
            @toggleLock="lockDocument"
            @setFullWidth="fullWidth"
          />
          
          <!-- Affichage conditionnel selon le type de document -->
          <div 
            class="editor-container flex-1 relative" 
            :class="{ 
              'px-8 py-4 max-w-7xl mx-auto': !currentDocument.config?.full_width && currentDocument.type !== 'excalidraw',
              'h-full': currentDocument.type === 'excalidraw'
            }"
          >
            <ExcalidrawEditor 
              v-if="currentDocument.type === 'excalidraw'"
              :document="currentDocument" 
              :onSave="saveDocument"
              class="h-full"
            />
            
            <TiptapEditor 
              v-else
              :document="currentDocument" 
              :lock="!!currentDocument.config?.lock" 
              :onSave="saveDocument"
            />
          </div>
        </div>
      </div>
    </template>
    
    <!-- Message si aucun document n'est trouvé -->
    <div v-else class="flex justify-center items-center h-screen">
      <div class="text-xl text-gray-500">Document not found</div>
    </div>
  </main>
</template>

<style>
/* Basic editor styles */
.ProseMirror {
  min-height: 300px;
  padding: 1rem;
  border-radius: 0.25rem;
  outline: none;
}

/* Style pour l'éditeur verrouillé */
.ProseMirror[contenteditable="false"] {
  position: relative;
  color: inherit !important; /* Forcer la même couleur de texte que l'éditeur normal */
  opacity: 1 !important; /* Forcer l'opacité complète */
}

/* S'assurer que tous les éléments de texte à l'intérieur gardent leur style normal */
.ProseMirror[contenteditable="false"] p,
.ProseMirror[contenteditable="false"] h1,
.ProseMirror[contenteditable="false"] h2,
.ProseMirror[contenteditable="false"] h3,
.ProseMirror[contenteditable="false"] ul,
.ProseMirror[contenteditable="false"] ol,
.ProseMirror[contenteditable="false"] li,
.ProseMirror[contenteditable="false"] blockquote,
.ProseMirror[contenteditable="false"] pre,
.ProseMirror[contenteditable="false"] code {
  color: inherit !important;
  opacity: 1 !important;
  filter: none !important;
}

/* Styles spécifiques pour les liens et le texte formaté */
.ProseMirror[contenteditable="false"] a {
  color: #3b82f6 !important; /* Bleu standard pour les liens */
  text-decoration: underline !important;
  opacity: 1 !important;
}

.ProseMirror[contenteditable="false"] strong {
  font-weight: bold !important;
  opacity: 1 !important;
}

.ProseMirror[contenteditable="false"] em {
  font-style: italic !important;
  opacity: 1 !important;
}

/* Curseur approprié pour indiquer que l'édition est désactivée */
.ProseMirror[contenteditable="false"] * {
  cursor: default !important;
}

.ProseMirror p.is-editor-empty:first-child::before {
  content: attr(data-placeholder);
  float: left;
  color: #adb5bd;
  pointer-events: none;
  height: 0;
}

/* Table styles */
.ProseMirror table {
  border-collapse: collapse;
  table-layout: fixed;
  width: 100%;
  margin: 0;
  overflow: hidden;
}

/* Styles for the Excalidraw container */
.editor-container.h-full {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.editor-container.h-full .excalidraw-container {
  flex: 1;
  height: 100%;
}

.ProseMirror td,
.ProseMirror th {
  min-width: 1em;
  border: 2px solid #ced4da;
  padding: 3px 5px;
  vertical-align: top;
  box-sizing: border-box;
  position: relative;
}

.ProseMirror th {
  font-weight: bold;
  background-color: #f1f3f5;
}
</style>
