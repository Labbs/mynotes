<script setup lang="ts">
import { computed, onMounted, watch, ref, onUnmounted, nextTick, type Directive } from 'vue';
import { useRoute } from 'vue-router';
import { useDocumentStore } from '../stores/document';
import { useFavoriteStore } from '../stores/favorite';
import { useEditor, EditorContent } from '@tiptap/vue-3';
import Header from '../components/document/Header.vue';
import StarterKit from '@tiptap/starter-kit';
import Placeholder from '@tiptap/extension-placeholder';
import { Color } from "@tiptap/extension-color";
import TextStyle from "@tiptap/extension-text-style";
import Image from "@tiptap/extension-image";
import BubbleMenu from '../components/tiptap/BubbleMenu.vue';
import CodeBlock from '@tiptap/extension-code-block';
import Commands from '../extensions/Commands';
import suggestion from '../extensions/Suggestion';
import Highlight from '@tiptap/extension-highlight';

const route = useRoute()
const documentStore = useDocumentStore();
const favoritesStore = useFavoriteStore();
const currentDocument = computed(() => documentStore.currentDocument)
const isEditingTitle = ref(false)
const editableTitle = ref('')
const saveTimeout = ref<number | null>(null)
const isEditing = ref(false)
const editorInitialized = ref(false)
const isEditorUpdating = ref(false)
const configSidebarVisible = ref(false) // Pour suivre l'état de visibilité de la sidebar de configuration

// Function to update document icon
const updateDocumentIcon = async (icon: string) => {
  console.log('updateDocumentIcon called with:', icon);
  if (!currentDocument.value) return;
  
  currentDocument.value.config.icon = icon;
  console.log('Current document before update:', currentDocument.value);
  
  try {
    await documentStore.updateDocument({
      id: currentDocument.value.id,
      name: currentDocument.value.name,
      content: currentDocument.value.content,
      space_id: currentDocument.value.space_id,
      config: currentDocument.value.config
    });
    await favoritesStore.fetchFavorites(); // Refresh favorites after update
    console.log('Document icon updated successfully');
  } catch (error) {
    console.error('Failed to update document icon:', error);
  }
};

const lockDocument = async (lock: boolean) => {
  if (!currentDocument.value) return;
  
  // Mettre à jour l'état de verrouillage du document
  currentDocument.value.config.lock = lock;
  console.log('Locking document:', lock ? 'locked' : 'unlocked');
  
  // Utiliser la méthode intégrée de TipTap pour rendre l'éditeur non éditable
  if (editor.value) {
    editor.value.setEditable(!lock);
  }
  
  try {
    // Sauvegarder l'état de verrouillage dans la base de données
    await documentStore.updateDocument({
      id: currentDocument.value.id,
      name: currentDocument.value.name,
      content: currentDocument.value.content,
      space_id: currentDocument.value.space_id,
      config: currentDocument.value.config
    });
    console.log('Document lock state updated successfully');
  } catch (error) {
    // En cas d'erreur, rétablir l'état précédent
    console.error('Failed to update document lock state:', error);
    if (currentDocument.value) {
      currentDocument.value.config.lock = !lock;
      if (editor.value) {
        editor.value.setEditable(lock);
      }
    }
  }
};

const fullWidth = async (fullWidth: boolean) => {
  if (!currentDocument.value) return;
  
  // Mettre à jour l'état de largeur complète du document
  currentDocument.value.config.full_width = fullWidth;
  
  try {
    await documentStore.updateDocument({
      id: currentDocument.value.id,
      name: currentDocument.value.name,
      content: currentDocument.value.content,
      space_id: currentDocument.value.space_id,
      config: currentDocument.value.config
    });
    console.log('Document full width state updated successfully');
  } catch (error) {
    console.error('Failed to update document full width state:', error);
  }
};

// Function to toggle editor editable state
const toggleEditorEditable = (editable: boolean) => {
  if (editor.value) {
    console.log('Setting editor editable state to:', editable);
    editor.value.setEditable(editable);
  }
};

const editor = useEditor({
  content: '<p></p>', // Start with empty content
  editable: currentDocument.value ? !currentDocument.value.config?.lock : true, // État initial de l'éditeur basé sur l'état de verrouillage du document
  extensions: [
    StarterKit.configure({
      heading: {
        levels: [1, 2, 3],
      },
    }),
    Placeholder.configure({
      placeholder: "Type '/' for commands",
    }),
    Color,
    TextStyle,
    Image,
    CodeBlock,
    Commands.configure({
      suggestion,
    }),
    Highlight.configure({
      multicolor: true,
    }),
  ],
  onUpdate: ({ editor }) => {
    // Si l'éditeur n'est pas éditable (document verrouillé), ne rien faire
    if (!editor.isEditable || !currentDocument.value || isEditorUpdating.value) {
      return;
    }
    
    // Sauvegarder le contenu lorsqu'il change
    const content = editor.getHTML();
    currentDocument.value.content = content;
    saveContent();
  },
});

// Mark editor as initialized
editorInitialized.value = true;

// Add this watch to ensure editor editable state is synchronized with document lock state
watch(() => currentDocument.value?.config?.lock, (isLocked) => {
  if (editor.value && isLocked !== undefined) {
    console.log(`Document lock state changed to: ${isLocked ? 'locked' : 'unlocked'}`);
    editor.value.setEditable(!isLocked);
  }
}, { immediate: true });

// Add this watch to update the editor when document content changes
watch(() => currentDocument.value?.content, (newContent) => {
  if (editor.value && newContent && !isEditing.value) {
    console.log('Updating editor with document content');
    isEditorUpdating.value = true;
    
    // Wait a tick to ensure the document is fully loaded
    nextTick(() => {
      try {
        editor.value?.commands.setContent(newContent);
        console.log('Editor content updated successfully');
      } catch (error) {
        console.error('Error updating editor content:', error);
      } finally {
        // Reset the flag after a short delay
        setTimeout(() => {
          isEditorUpdating.value = false;
        }, 50);
      }
    });
  }
}, { immediate: true });

// Add this to help with debugging
watch(() => documentStore.loadingDocument, (isLoading) => {
  console.log(`Document loading: ${isLoading}`);
  if (!isLoading && currentDocument.value) {
    console.log('Document loaded:', {
      id: currentDocument.value.id,
      name: currentDocument.value.name,
      contentLength: currentDocument.value.content?.length || 0,
      hasContent: !!currentDocument.value.content
    });
  }
});

const favorite = async () => {
  // check if the document is already in favorites
  if (favoritesStore.favorites.some(f => f.document?.id === currentDocument.value?.id)) {
    // remove from favorites
    await favoritesStore.unFavorite(currentDocument.value?.id as string);
  } else {
    // add to favorites
    await favoritesStore.addFavorite(currentDocument.value?.id as string);
  }
}

// Improve the updateEditorFromDocument function
const updateEditorFromDocument = () => {
  if (editor.value && currentDocument.value) {
    // Désactiver les mises à jour pendant l'opération
    console.log('Updating editor from document, content length:', 
                currentDocument.value.content?.length || 0);
    isEditorUpdating.value = true;
    
    nextTick(() => {
      try {
        // Mettre à jour le contenu
        if (currentDocument.value?.content) {
          editor.value?.commands.setContent(currentDocument.value.content);
        } else {
          editor.value?.commands.clearContent(true);
        }
        
        // Mettre à jour l'état d'édition selon l'état de verrouillage
        const isLocked = !!currentDocument.value?.config?.lock;
        console.log('Document lock state:', isLocked ? 'locked' : 'unlocked');
        
        // Utiliser la méthode native de TipTap pour définir l'état éditable
        if (editor.value && editor.value.isEditable !== !isLocked) {
          editor.value.setEditable(!isLocked);
        }
      } catch (error) {
        console.error('Error updating editor with document content:', error);
      } finally {
        // Réactiver les mises à jour après un court délai
        setTimeout(() => {
          isEditorUpdating.value = false;
        }, 50);
      }
    });
  } else {
    console.warn('Cannot update editor: editor or document is not available');
  }
}

// Improved fetching when the component is mounted
onMounted(async () => {
  try {
    const slug = route.params.slug as string;
    if (slug) {
      console.log('Mounting document view with slug:', slug);
      await documentStore.fetchDocument(slug);
      
      // Ensure we update the editor after the document is loaded
      nextTick(() => {
        updateEditorFromDocument();
      });
    }
  } catch (error) {
    console.error('Error initializing document view:', error);
  }
})

watch(() => route.params.slug, async (newSlug) => {
  if (newSlug && newSlug !== currentDocument.value?.slug) {
    await documentStore.fetchDocument(newSlug as string)
    
    // Réinitialiser l'éditeur avec le nouveau document
    if (editorInitialized.value) {
      updateEditorFromDocument()
    }
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
      // Make sure to explicitly send the current content with the title update
      await documentStore.updateDocument({
        id: currentDocument.value.id,
        name: editableTitle.value,
        content: currentContent, // Explicitly include the content
        space_id: currentDocument.value.space_id,
        config: currentDocument.value.config
      });
      
      isEditingTitle.value = false;
      
      // Check if the slug changed after update
      if (currentDocument.value.slug !== oldSlug) {
        // Update URL without causing navigation
        window.history.replaceState(null, '', `/d/${currentDocument.value.slug}`);
        
        // Important: After URL change, make sure content is preserved in local state
        if (currentDocument.value) {
          currentDocument.value.content = currentContent;
        }
        
        // Force update editor content to ensure it stays intact
        if (editor.value && currentContent) {
          editor.value.commands.setContent(currentContent);
        }

        await favoritesStore.fetchFavorites();
      }
      
      console.log('Document title updated successfully, content preserved');
    } catch (error) {
      console.error('Failed to update document title:', error);
    }
  } else {
    isEditingTitle.value = false;
  }
};

// Function to save editor content with debounce
const saveContent = async () => {
  if (saveTimeout.value) {
    clearTimeout(saveTimeout.value);
  }
  
  saveTimeout.value = window.setTimeout(async () => {
    // Vérifier si le document est verrouillé avant de sauvegarder
    if (!currentDocument.value || !editor.value || currentDocument.value.config?.lock) {
      console.log('Not saving: document is null, editor is null, or document is locked');
      return;
    }
    
    try {
      // Marquer que nous sommes en train d'éditer pour éviter de re-rendre
      isEditing.value = true;
            
      await documentStore.updateDocument({
        id: currentDocument.value.id,
        name: currentDocument.value.name,
        content: currentDocument.value.content,
        space_id: currentDocument.value.space_id,
        config: currentDocument.value.config
      });
      
      console.log('Content auto-saved');
    } catch (error) {
      console.error('Failed to auto-save content:', error);
    } finally {
      // Réinitialiser le flag après la sauvegarde
      isEditing.value = false;
    }
  }, 1000); // Debounce de 1 seconde
};

onUnmounted(() => {
  if (saveTimeout.value) {
    clearTimeout(saveTimeout.value);
  }
  editor.value?.destroy();
});
</script>

<template>
  <main class="flex-1 overflow-y-auto relative">
    <!-- Afficher un indicateur de chargement -->
    <div v-if="documentStore.loadingDocument" class="flex justify-center items-center h-screen">
      <div class="animate-pulse text-xl text-gray-500">Loading the document</div>
    </div>
    
    <!-- Afficher le contenu une fois chargé -->
    <template v-else-if="currentDocument">
      <div class="flex relative h-full">
        <div 
          :class="{ 'mr-72': configSidebarVisible }" 
          class="transition-all duration-300 ease-in-out w-full"
        >
          <Header 
            :current-document="currentDocument" 
            :is-editing-title="isEditingTitle" 
            :editable-title="editableTitle"
            @update-title="updateTitle"
            @favorite="favorite"
            @update:is-editing-title="(value) => { 
              // N'autoriser l'édition du titre que si le document n'est pas verrouillé
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
          
          <!-- This is where the editor should be rendered -->
          <div class="editor-container px-8 py-4 relative"
            :class="{ 'max-w-7xl mx-auto': !currentDocument?.config?.full_width }">
            <div class="relative">
              <editor-content v-if="editor" :editor="editor" class="prose max-w-none" />
              <!-- Overlay semi-transparent sur tout l'éditeur quand verrouillé -->
              <div 
                v-if="currentDocument?.config?.lock" 
                class="absolute inset-0 bg-gray-50 bg-opacity-20 z-10"
                style="pointer-events: none;"
              ></div>
              <BubbleMenu :editor="editor" v-if="editor" />
              <!-- Indicateur de verrouillage -->
            </div>
            <div v-if="!editor" class="min-h-[300px] border rounded p-4 flex items-center justify-center text-gray-400">
              Loading editor...
            </div>
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