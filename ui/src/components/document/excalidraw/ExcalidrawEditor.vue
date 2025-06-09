<template>
  <div class="excalidraw-container">
    <div ref="excalidrawContainer" class="excalidraw-editor"></div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted, onBeforeUnmount } from 'vue'
import type { Document } from '../../../api/interface'
import * as React from 'react'
import * as ReactDOM from 'react-dom/client'
import { ExcalidrawWrapper } from './ExcalidrawWrapper'
import { loadLibraries } from './utils/loadLibraries'

const props = defineProps<{
  document: Document
  onSave: (content: string) => void
}>()

const excalidrawContainer = ref<HTMLElement | null>(null)
const excalidrawData = ref<any>({ 
  elements: [], 
  appState: { 
    viewBackgroundColor: '#ffffff',
    currentItemFontFamily: 1
  }, 
  files: {},
  libraryItems: [] // Changed from libraries to libraryItems for Excalidraw compatibility
})
const isFullscreen = ref(false)

// Function to update container size
const updateContainerSize = () => {
  if (excalidrawContainer.value) {
    if (isFullscreen.value) {
      excalidrawContainer.value.style.height = `${window.innerHeight - 50}px`;
    } else {
      const parentHeight = excalidrawContainer.value.parentElement?.clientHeight || 800;
      excalidrawContainer.value.style.height = `${parentHeight - 50}px`;
    }
  }
};

// Initialize from document content
if (props.document.content) {
  try {
    const parsed = JSON.parse(props.document.content);
    
    // Validate the data structure to avoid "Document loaded without content!" error
    if (!parsed || (Array.isArray(parsed.elements) && parsed.elements.length === 0 && !parsed.appState)) {
      console.warn('Empty Excalidraw document content detected, initializing with default state');
      excalidrawData.value = { 
        elements: [], 
        appState: { 
          viewBackgroundColor: '#ffffff',
          currentItemFontFamily: 1
        }, 
        files: {},
        libraryItems: [] // Les bibliothèques seront chargées dans onMounted
      };
    } else {
      // Ensure we have a valid structure with required fields
      excalidrawData.value = {
        elements: parsed.elements || [],
        appState: parsed.appState || { 
          viewBackgroundColor: '#ffffff',
          currentItemFontFamily: 1
        },
        files: parsed.files || {},
        libraryItems: parsed.libraryItems || [] // Conserver les bibliothèques existantes
      };
    }
  } catch (e) {
    console.error('Error parsing initial Excalidraw content:', e);
    // Initialize with a valid default state
    excalidrawData.value = { 
      elements: [], 
      appState: { 
        viewBackgroundColor: '#ffffff',
        currentItemFontFamily: 1
      }, 
      files: {},
      libraryItems: [] // Les bibliothèques seront chargées dans onMounted
    };
  }
} else {
  // Initialize with a valid default state for new documents
  excalidrawData.value = { 
    elements: [], 
    appState: { 
      viewBackgroundColor: '#ffffff',
      currentItemFontFamily: 1
    }, 
    files: {},
    libraryItems: [] // Les bibliothèques seront chargées dans onMounted
  };
}

let reactRoot: ReactDOM.Root | null = null
let lastSavedData = { elements: [], appState: {}, files: {} };
let saveTimeoutId: number | null = null;
let isSaving = false;
let isInitialRender = true;

function onChange(elements: any, appState: any, files: any) {
  // Skip saving during initial render to prevent flicker
  if (isInitialRender) {
    isInitialRender = false;
    return;
  }
  
  // Mettre à jour excalidrawData pour garder la synchronisation
  excalidrawData.value = {
    elements,
    appState: {
      ...excalidrawData.value.appState,
      ...appState
    },
    files,
    libraryItems: excalidrawData.value.libraryItems
  };
  
  // Ne pas démarrer une nouvelle sauvegarde si l'une est déjà en cours
  if (isSaving) {
    return;
  }
  
  // Utiliser un throttle/debounce basique pour l'auto-sauvegarde
  if (saveTimeoutId) {
    clearTimeout(saveTimeoutId);
  }
  
  saveTimeoutId = setTimeout(async () => {
    // Vérifier si des changements significatifs avant de sauvegarder
    const currentElements = JSON.stringify(elements);
    const savedElements = JSON.stringify(lastSavedData.elements); 
    const hasSignificantChanges = elements.length !== lastSavedData.elements.length || 
                                currentElements !== savedElements;
    
    if (hasSignificantChanges) {
      isSaving = true;
      try {
        // Ne sauvegarder que les données essentielles
        const dataToSave = {
          elements,
          appState: excalidrawData.value.appState,
          files,
          libraryItems: excalidrawData.value.libraryItems
        };
        
        await props.onSave(JSON.stringify(dataToSave));
        lastSavedData = {
          elements,
          appState: excalidrawData.value.appState,
          files
        };
      } finally {
        isSaving = false;
      }
    }
  }, 2000) as unknown as number;
}

async function saveExcalidraw() {
  try {
    // Récupérer les données actuelles depuis excalidrawData
    const currentState = {
      elements: excalidrawData.value.elements,
      appState: excalidrawData.value.appState,
      files: excalidrawData.value.files,
      libraryItems: excalidrawData.value.libraryItems
    };
    
    // Save valid data
    await props.onSave(JSON.stringify(currentState));
    lastSavedData = {
      elements: currentState.elements,
      appState: currentState.appState,
      files: currentState.files
    };
  } catch (e) {
    console.error('Error saving Excalidraw content:', e);
    throw e;
  }
}

function toggleFullscreen() {
  isFullscreen.value = !isFullscreen.value;
  
  // Attendre que la transition de classe soit appliquée puis mettre à jour la taille
  setTimeout(() => {
    updateContainerSize();
    // Forcer Excalidraw à se redimensionner
    window.dispatchEvent(new Event('resize'));
  }, 100);
}

function renderExcalidraw() {
  if (!excalidrawContainer.value) return
  
  if (reactRoot) {
    reactRoot.unmount()
  }
  
  // Make sure we have valid data to render
  if (!excalidrawData.value || !excalidrawData.value.elements) {
    console.warn('Invalid Excalidraw data - reinitializing with default structure');
    excalidrawData.value = { 
      elements: [], 
      appState: { 
        viewBackgroundColor: '#ffffff',
        currentItemFontFamily: 1,
        gridSize: 20
      }, 
      files: {} 
    };
  }
  
  // S'assurer que le conteneur a la bonne taille avant de rendre
  updateContainerSize();
  window.addEventListener('resize', updateContainerSize);
  
  reactRoot = ReactDOM.createRoot(excalidrawContainer.value);
  reactRoot.render(
    React.createElement(ExcalidrawWrapper, {
      initialData: excalidrawData.value,
      onChange: onChange
    })
  );
}

watch(() => props.document.content, (newContent) => {
  if (newContent && newContent !== 'null') {
    try {
      const parsed = JSON.parse(newContent);
      
      // Validate the data structure to avoid "Document loaded without content!" error
      if (!parsed || (Array.isArray(parsed.elements) && parsed.elements.length === 0 && !parsed.appState)) {
        console.warn('Empty Excalidraw document content detected, initializing with default state');
        excalidrawData.value = { 
          elements: [], 
          appState: { 
            viewBackgroundColor: '#ffffff',
            currentItemFontFamily: 1
          }, 
          files: {} 
        };
      } else {
        // Ensure we have a valid structure with required fields
        excalidrawData.value = {
          elements: parsed.elements || [],
          appState: parsed.appState || { 
            viewBackgroundColor: '#ffffff',
            currentItemFontFamily: 1
          },
          files: parsed.files || {}
        };
      }
    } catch (e) {
      console.error('Error parsing Excalidraw content in watch:', e);
      // Initialize with a valid default state
      excalidrawData.value = { 
        elements: [], 
        appState: { 
          viewBackgroundColor: '#ffffff',
          currentItemFontFamily: 1
        }, 
        files: {} 
      };
    }
  } else {
    // Initialize with a valid default state for new documents
    excalidrawData.value = { 
      elements: [], 
      appState: { 
        viewBackgroundColor: '#ffffff',
        currentItemFontFamily: 1
      }, 
      files: {} 
    };
  }
  renderExcalidraw();
})

// Raccourcis clavier pour le mode plein écran
const handleKeyDown = (e: KeyboardEvent) => {
  if (e.key === 'F11' || (e.key === 'f' && e.ctrlKey)) {
    e.preventDefault();
    toggleFullscreen();
  }
  
  // Enregistrer avec Ctrl+S
  if (e.key === 's' && (e.ctrlKey || e.metaKey)) {
    e.preventDefault();
    saveExcalidraw();
  }
};

onMounted(async () => {
  // Charger les bibliothèques au démarrage
  try {
    const libraries = await loadLibraries();
    // Fusionner tous les libraryItems des bibliothèques chargées
    excalidrawData.value.libraryItems = libraries.flatMap(lib => lib.libraryItems);
    console.log('Loaded Excalidraw libraries:', libraries.length);
  } catch (error) {
    console.error('Failed to load Excalidraw libraries:', error);
  }
  
  renderExcalidraw();
  window.addEventListener('keydown', handleKeyDown);
})

onBeforeUnmount(() => {
  if (reactRoot) {
    reactRoot.unmount()
  }
  // Nettoyer les écouteurs d'événements
  window.removeEventListener('resize', updateContainerSize);
  window.removeEventListener('keydown', handleKeyDown);
})
</script>

<style scoped>
.excalidraw-container {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  position: relative;
}

.excalidraw-header {
  text-align: right;
  margin-bottom: 8px;
  padding: 4px 8px;
  display: flex;
  justify-content: flex-end;
  background-color: #f9fafb;
  border-radius: 6px;
}

.save-button {
  padding: 6px 12px;
  background-color: #3b82f6;
  color: white;
  border: none;
  border-radius: 4px;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.2s;
  display: flex;
  align-items: center;
}

.save-button:hover {
  background-color: #2563eb;
}

.tool-button {
  padding: 6px;
  background-color: #f3f4f6;
  color: #4b5563;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
}

.tool-button:hover {
  background-color: #e5e7eb;
  color: #1f2937;
}

.excalidraw-editor {
  height: calc(100% - 40px);
  width: 100%;
  border: 1px solid #eee;
  border-radius: 4px;
  overflow: hidden;
  position: relative;
  flex-grow: 1;
  transition: all 0.3s ease;
}

.fullscreen {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  height: 100vh;
  width: 100vw;
  z-index: 9999;
  border-radius: 0;
  border: none;
}

/* Ensure Excalidraw's UI is properly sized */
.excalidraw-editor :deep(.excalidraw) {
  height: 100%;
  width: 100%;
}

/* Fix for potential UI issues */
.excalidraw-editor :deep(.excalidraw .App-menu) {
  z-index: 10;
}
</style>
