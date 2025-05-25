<script setup lang="ts">
import { computed, ref } from 'vue';
import { type Document } from '../../api/document';
import { Transition } from 'vue';

const props = defineProps<{
  visible: boolean;
  currentDocument: Document | null;
}>();

const emit = defineEmits<{
  (e: 'close'): void;
  (e: 'setDocumentIcon', icon: string): void;
}>();

const activeTab = ref('emoji');
const isUpdating = ref(false);

const currentIcon = computed(() => {
  return props.currentDocument?.config?.icon || null;
});

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
</script>

<template>
  <!-- Sidebar avec transition -->
  <Transition name="sidebar">
    <div 
      v-if="visible" 
      class="fixed right-0 top-0 bottom-0 w-72 bg-white border-l border-gray-200 shadow-lg z-40 overflow-y-auto transform"
      style="height: 100%;"
    >
    <!-- Header -->
    <div class="sticky top-0 p-4 border-b border-gray-200 flex justify-between items-center bg-white z-10">
      <h2 class="text-lg font-medium text-gray-800">Choose your icon</h2>
      <button 
        @click="emit('close')" 
        class="p-1 rounded-md hover:bg-gray-100 text-gray-500 hover:text-gray-700"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
        </svg>
      </button>
    </div>

    <!-- Tabs -->
    <div class="sticky top-16 flex border-b border-gray-200 bg-white z-10">
      <button 
        @click="activeTab = 'emoji'" 
        class="flex-1 py-3 text-sm font-medium transition-colors"
        :class="activeTab === 'emoji' ? 'text-blue-600 border-b-2 border-blue-600' : 'text-gray-500 hover:text-gray-700'"
      >
        Émojis
      </button>
      <button 
        @click="activeTab = 'icon'" 
        class="flex-1 py-3 text-sm font-medium transition-colors"
        :class="activeTab === 'icon' ? 'text-blue-600 border-b-2 border-blue-600' : 'text-gray-500 hover:text-gray-700'"
      >
        Icônes
      </button>
    </div>

    <!-- Current Icon -->
    <div v-if="currentIcon" class="sticky top-28 p-4 border-b border-gray-200 bg-white z-10">
      <div class="flex justify-between items-center">
        <div class="flex items-center">
          <span class="text-2xl mr-2">{{ currentIcon }}</span>
          <span class="text-sm text-gray-600">Current Icon</span>
        </div>
        <button 
          @click="applyIcon('')" 
          :disabled="isUpdating"
          class="px-3 py-1.5 text-sm text-white bg-red-500 hover:bg-red-600 rounded-md transition-colors flex items-center gap-1"
        >
          <svg v-if="isUpdating" class="animate-spin h-4 w-4 mr-1" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <span>Remove Icon</span>
        </button>
      </div>
    </div>

    <!-- Content -->
    <div class="p-4 pt-2 pb-20 overflow-y-auto">
      <!-- Émojis -->
      <div v-if="activeTab === 'emoji'" class="grid grid-cols-5 gap-3">
        <button 
          v-for="emoji in ['😀', '😁', '😂', '🤣', '😃', '😄', '😅', '😆', '😉', '😊', '😋', '😎', '😍', '😘', '🥰', '😗', '😙', '😚', '🙂', '🤗', '🤔', '🤨', '😐', '😑', '😶', '🙄', '😏', '😣', '😥', '😮', '🤐', '😯', '😪', '😫', '😴', '😌', '😛', '😜', '😝', '🤤', '😒', '😓', '😔', '😕', '🙃', '🤑', '😲', '☹️', '🙁', '😖', '😞', '😟', '😤', '😢', '😭', '😦', '😧', '😨', '😩', '🤯', '😬', '😰', '😱', '🥵', '🥶', '😳', '🤪', '😵', '😡', '😠', '🤬', '😷', '🤒', '🤕', '🤢', '🤮', '🤧', '😇', '🥳', '🥴', '🥺', '🤠', '🤡', '🤥', '🤫', '🤭', '🧐', '🤓', '😈', '👻', '👽', '🤖', '💩', '😺', '😸', '😹', '😻', '😼', '😽', '🙀', '😿', '😾']" 
          :key="emoji"
          class="emoji-btn w-10 h-10 text-xl flex items-center justify-center hover:bg-gray-100 rounded-md transition-all"
          :class="{ 'bg-blue-50 ring-2 ring-blue-300': currentIcon === emoji }"
          @click="applyIcon(emoji)"
        >
          {{ emoji }}
        </button>
      </div>

      <!-- Icônes -->
      <div v-else class="grid grid-cols-5 gap-3">
        <button 
          v-for="icon in ['📝', '📚', '🔍', '📊', '📈', '📉', '💻', '🖥️', '📱', '⌨️', '🖱️', '🖨️', '💾', '💿', '📀', '🎬', '📷', '🔬', '🔭', '📡', '🔋', '🔌', '🧰', '🧲', '⚙️', '🔧', '🔨', '⛓️', '📌', '📍', '✂️', '🗑️', '🔒', '🔓', '🔑', '🗝️', '🔨', '⛏️', '⚒️', '🛠️', '🗡️', '⚔️', '🔫', '🏹', '🛡️', '🔮', '💣', '⌛', '⏳', '⌚', '⏰', '🧭', '🌡️', '💉', '💊', '🩹', '🩺', '🚪', '🛏️', '🛋️', '🪑', '🚽', '🚿', '🛁', '🧴', '🧷', '🧹', '🧺', '🧻', '🧼', '🧽', '🧯', '🚬', '⚰️', '⚱️', '🗿', '🏧', '🚮', '🚰', '♿', '🚹', '🚺', '🚻', '🚼', '🚾', '🛂', '🛃', '🛄', '🛅', '⚠️', '🚸', '⛔', '🚫', '🚳', '🚭', '🚯', '🚱', '🚷', '📵', '🔞', '☢️', '☣️', '⬆️', '↗️', '➡️', '↘️', '⬇️', '↙️', '⬅️', '↖️', '↕️', '↔️', '↩️', '↪️', '⤴️', '⤵️', '🔃', '🔄', '🔙', '🔚', '🔛', '🔜', '🔝']" 
          :key="icon"
          class="emoji-btn w-10 h-10 text-xl flex items-center justify-center hover:bg-gray-100 rounded-md transition-all"
          :class="{ 'bg-blue-50 ring-2 ring-blue-300': currentIcon === icon }"
          @click="applyIcon(icon)"
        >
          {{ icon }}
        </button>
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
</style>
