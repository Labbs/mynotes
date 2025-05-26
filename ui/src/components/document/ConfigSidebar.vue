<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { type Document } from '../../api/document';

const props = defineProps<{
  visible: boolean;
  currentDocument: Document | null;
}>();

const emit = defineEmits<{
  (e: 'close'): void;
  (e: 'setDocumentIcon', icon: string): void;
  (e: 'setLock', lock: boolean): void;
  (e: 'setFullWidth', fullWidth: boolean): void;
}>();

const activeTab = ref('emoji');
const isUpdating = ref(false);
const activeAccordions = ref(new Set<string>());

const currentIcon = computed(() => {
  return props.currentDocument?.config?.icon || null;
});

const lockDocument = computed(() => {
  return props.currentDocument?.config?.lock || false;
});

const fullWidthDocument = computed(() => {
  return props.currentDocument?.config?.full_width || false;
});

// État local pour le verrouillage qui peut être modifié
const isDocumentLocked = ref(false);
// État local pour le mode pleine largeur qui peut être modifié
const isFullWidth = ref(false);

// Synchroniser isDocumentLocked avec l'état actuel du document
watch(() => lockDocument.value, (newValue) => {
  isDocumentLocked.value = newValue;
}, { immediate: true });

// Synchroniser isFullWidth avec l'état actuel du document
watch(() => fullWidthDocument.value, (newValue) => {
  isFullWidth.value = newValue;
}, { immediate: true });

const toggleAccordion = (accordionId: string) => {
  if (activeAccordions.value.has(accordionId)) {
    activeAccordions.value.delete(accordionId);
  } else {
    activeAccordions.value.add(accordionId);
  }
};

const applyIcon = (icon: string) => {
  isUpdating.value = true;
  emit('setDocumentIcon', icon);
  
  // Simuler un court délai pour l'effet visuel
  setTimeout(() => {
    isUpdating.value = false;
    // Fermer automatiquement la sidebar après avoir sélectionné une icône
    if (icon !== '') { // Ne pas fermer si on supprime l'icône
      emit('close');
    }
  }, 300);
};

const lockDocumentChange = () => {
  // Utiliser l'état local et l'inverser
  isDocumentLocked.value = !isDocumentLocked.value;
  emit('setLock', isDocumentLocked.value);
};

const fullWidthDocumentChange = () => {
  // Utiliser l'état local et l'inverser
  isFullWidth.value = !isFullWidth.value;
  emit('setFullWidth', isFullWidth.value);
};
</script>

<template>
  <!-- Sidebar avec transition -->
  <Transition name="sidebar">
    <div 
      v-if="visible" 
      class="fixed right-0 top-0 bottom-0 w-72 bg-white border-l border-gray-200 shadow-lg z-40 overflow-y-auto transform"
      style="height: 100%;"
    >
    <div class="sticky top-0 p-4 border-b border-gray-200 flex justify-between items-center bg-white z-10">
      <h2 class="text-lg font-medium text-gray-800">Configuration</h2>
      <button 
        @click="emit('close')" 
        class="p-1 rounded-md hover:bg-gray-100 text-gray-500 hover:text-gray-700"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
        </svg>
      </button>
    </div>
    <div class="sticky top-0 p-4 border-b border-gray-200 flex justify-between items-center bg-white z-10">
      <span class="text-xs text-gray-400">Last updated: {{ props.currentDocument?.updated_at ? new Date(props.currentDocument.updated_at).toLocaleString() : 'N/A' }}</span>
    </div>
    <div class="sticky top-0 p-4 border-b border-gray-200 flex justify-between items-center bg-white z-10">
      <label class="inline-flex items-center me-5 cursor-pointer">
        <input type="checkbox" value="" class="sr-only peer" 
          :checked="isDocumentLocked" 
          @change="lockDocumentChange()"
        >
        <div class="relative w-11 h-6 bg-gray-200 rounded-full peer peer-focus:ring-4 peer-focus:ring-red-300 dark:peer-focus:ring-red-800 dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-0.5 after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-red-600 dark:peer-checked:bg-red-600"></div>
        <span class="ms-3 text-sm font-medium text-gray-900 dark:text-gray-500">Lock document</span>
      </label>
    </div>
    <div class="sticky top-0 p-4 border-b border-gray-200 flex justify-between items-center bg-white z-10">
      <label class="inline-flex items-center me-5 cursor-pointer">
        <input type="checkbox" value="" class="sr-only peer" 
          :checked="isFullWidth" 
          @change="fullWidthDocumentChange()"
        >
        <div class="relative w-11 h-6 bg-gray-200 rounded-full peer peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-0.5 after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600 dark:peer-checked:bg-blue-600"></div>
        <span class="ms-3 text-sm font-medium text-gray-900 dark:text-gray-500">Full width</span>
      </label>
    </div>
    <div class="p-4 border-b border-gray-200 bg-white">
      <div class="accordion">
        <div class="accordion-item">
          <button 
            @click="toggleAccordion('icon')" 
            class="w-full flex justify-between items-center py-2 px-3 rounded-md hover:bg-gray-50 transition-colors"
          >
            <div class="flex items-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              <span class="font-medium text-gray-500">Choose your icon</span>
            </div>
            <svg 
              class="h-5 w-5 text-gray-500 transition-transform duration-200" 
              :class="{ 'transform rotate-180': activeAccordions.has('icon') }"
              xmlns="http://www.w3.org/2000/svg" 
              fill="none" 
              viewBox="0 0 24 24" 
              stroke="currentColor"
            >
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
          </button>
          <div 
            v-if="activeAccordions.has('icon')" 
            class="px-3 py-2 mt-1 bg-gray-50 rounded-md overflow-visible transition-all"
          >
            <!-- Current Icon -->
            <div v-if="currentIcon" class="mb-4 p-2 bg-white rounded-md border border-gray-200">
              <div class="flex justify-between items-center">
                <div class="flex items-center">
                  <span class="text-2xl mr-2">{{ currentIcon }}</span>
                  <span class="text-sm text-gray-600">Current icon</span>
                </div>
                <button 
                  @click="applyIcon('')" 
                  :disabled="isUpdating"
                  class="px-2 py-1 text-xs text-white bg-red-500 hover:bg-red-600 rounded-md transition-colors flex items-center"
                >
                  <svg v-if="isUpdating" class="animate-spin h-3 w-3 mr-1" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  <span>Remove</span>
                </button>
              </div>
            </div>
            
            <!-- Tabs -->
            <div class="flex mb-3 border border-gray-200 rounded-md overflow-hidden bg-white">
              <button 
                @click="activeTab = 'emoji'" 
                class="flex-1 py-2 text-sm font-medium transition-colors"
                :class="activeTab === 'emoji' ? 'bg-blue-50 text-blue-600' : 'text-gray-500 hover:text-gray-700 hover:bg-gray-50'"
              >
                Émojis
              </button>
              <button 
                @click="activeTab = 'icon'" 
                class="flex-1 py-2 text-sm font-medium transition-colors"
                :class="activeTab === 'icon' ? 'bg-blue-50 text-blue-600' : 'text-gray-500 hover:text-gray-700 hover:bg-gray-50'"
              >
                Icônes
              </button>
            </div>
            
            <!-- Icons Content -->
            <div class="bg-white p-2 rounded-md border border-gray-200 max-h-80 overflow-y-auto">
              <!-- Émojis -->
              <div v-if="activeTab === 'emoji'" class="grid grid-cols-5 gap-2">
                <button 
                  v-for="emoji in ['😀', '😁', '😂', '🤣', '😃', '😄', '😅', '😆', '😉', '😊', '😋', '😎', '😍', '😘', '🥰', '😗', '😙', '😚', '🙂', '🤗', '🤔', '🤨', '😐', '😑', '😶', '🙄', '😏', '😣', '😥', '😮', '🤐', '😯', '😪', '😫', '😴', '😌', '😛', '😜', '😝', '🤤', '😒', '😓', '😔', '😕', '🙃', '🤑', '😲', '☹️', '🙁', '😖', '😞', '😟', '😤', '😢', '😭', '😦', '😧', '😨', '😩', '🤯', '😬', '😰', '😱', '🥵', '🥶', '😳', '🤪', '😵', '😡', '😠', '🤬', '😷', '🤒', '🤕', '🤢', '🤮', '🤧', '😇', '🥳', '🥴', '🥺', '🤠', '🤡', '🤥', '🤫', '🤭', '🧐', '🤓', '😈', '👻', '👽', '🤖', '💩', '😺', '😸', '😹', '😻', '😼', '😽', '🙀', '😿', '😾']" 
                  :key="emoji"
                  class="emoji-btn w-8 h-8 text-lg flex items-center justify-center hover:bg-gray-100 rounded-md transition-all"
                  :class="{ 'bg-blue-50 ring-2 ring-blue-300': currentIcon === emoji }"
                  @click="applyIcon(emoji)"
                >
                  {{ emoji }}
                </button>
              </div>

              <!-- Icônes -->
              <div v-else class="grid grid-cols-5 gap-2">
                <button 
                  v-for="icon in ['📝', '📚', '🔍', '📊', '📈', '📉', '💻', '🖥️', '📱', '⌨️', '🖱️', '🖨️', '💾', '💿', '📀', '🎬', '📷', '🔬', '🔭', '📡', '🔋', '🔌', '🧰', '🧲', '⚙️', '🔧', '🔨', '⛓️', '📌', '📍', '✂️', '🗑️', '🔒', '🔓', '🔑', '🗝️', '🔨', '⛏️', '⚒️', '🛠️', '🗡️', '⚔️', '🔫', '🏹', '🛡️', '🔮', '💣', '⌛', '⏳', '⌚', '⏰', '🧭', '🌡️', '💉', '💊', '🩹', '🩺', '🚪', '🛏️', '🛋️', '🪑', '🚽', '🚿', '🛁', '🧴', '🧷', '🧹', '🧺', '🧻', '🧼', '🧽', '🧯', '🚬', '⚰️', '⚱️', '🗿', '🏧', '🚮', '🚰', '♿', '🚹', '🚺', '🚻', '🚼', '🚾', '🛂', '🛃', '🛄', '🛅', '⚠️', '🚸', '⛔', '🚫', '🚳', '🚭', '🚯', '🚱', '🚷', '📵', '🔞', '☢️', '☣️', '⬆️', '↗️', '➡️', '↘️', '⬇️', '↙️', '⬅️', '↖️', '↕️', '↔️', '↩️', '↪️', '⤴️', '⤵️', '🔃', '🔄', '🔙', '🔚', '🔛', '🔜', '🔝']" 
                  :key="icon"
                  class="emoji-btn w-8 h-8 text-lg flex items-center justify-center hover:bg-gray-100 rounded-md transition-all"
                  :class="{ 'bg-blue-50 ring-2 ring-blue-300': currentIcon === icon }"
                  @click="applyIcon(icon)"
                >
                  {{ icon }}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  </Transition>
</template>

<style scoped>
.emoji-btn {
  cursor: pointer;
  user-select: none;
  transition: all 0.15s ease;
  position: relative;
}

.emoji-btn:hover {
  transform: scale(1.15);
  background-color: #f3f4f6;
  z-index: 5;
}

.emoji-btn:active {
  transform: scale(0.95);
}

/* Transition pour la sidebar */
.sidebar-enter-active,
.sidebar-leave-active {
  transition: transform 0.3s ease-in-out;
}

.sidebar-enter-from,
.sidebar-leave-to {
  transform: translateX(100%);
}

/* Styles pour l'accordéon */
.accordion-item {
  transition: all 0.3s ease;
}

.accordion-item > div {
  max-height: 0;
  transition: max-height 0.5s ease, padding 0.3s ease, margin 0.3s ease;
}

.accordion-item > div[v-if="true"] {
  max-height: 1000px; /* Augmenter la valeur pour accommoder tous les emojis/icônes */
}
</style>
