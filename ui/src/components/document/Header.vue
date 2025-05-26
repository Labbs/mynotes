<script setup lang="ts">
import { type Directive, ref } from 'vue';
import { useFavoriteStore } from '../../stores/favorite';
import { type Document } from '../../api/document';
import { AdjustmentsHorizontalIcon } from "@heroicons/vue/24/solid";
import ConfigSidebar from './ConfigSidebar.vue';

const favoritesStore = useFavoriteStore();

defineProps<{
  currentDocument: Document | null;
  isEditingTitle: boolean;
  editableTitle: string;
}>();

const emit = defineEmits<{
  (e: 'updateTitle'): void;
  (e: 'favorite'): void;
  (e: 'update:isEditingTitle', value: boolean): void;
  (e: 'update:editableTitle', value: string): void;
  (e: 'setDocumentIcon', icon: string): void;
  (e: 'toggleEditorEditable', editable: boolean): void;
  (e: 'iconConfigVisibility', visible: boolean): void;
  (e: 'toggleLock', lock: boolean): void;
  (e: 'setFullWidth', fullWidth: boolean): void;
}>();

// État pour la sidebar de configuration
const showConfigSidebar = ref(false);

const toggleConfigSidebar = () => {
  showConfigSidebar.value = !showConfigSidebar.value;
  // Désactiver l'éditeur quand la sidebar est ouverte
  emit('toggleEditorEditable', !showConfigSidebar.value);
  // Notifier le parent de la visibilité de la sidebar
  emit('iconConfigVisibility', showConfigSidebar.value);
};

// Fonction pour appliquer une icône
const applyIcon = (icon: string) => {
  emit('setDocumentIcon', icon);
};

// Fonction pour basculer le verrouillage du document
const toggleLock = (lock: boolean) => {
  emit('toggleLock', lock);
};

const fullWidth = (fullWidth: boolean) => {
  emit('setFullWidth', fullWidth);
};

// Directive focus
const vFocus: Directive = {
  mounted: (el) => el.focus()
}
</script>

<template>
  <div class="sticky top-0 border-b border-e bg-white">
    <div class="flex justify-between items-center h-16 px-8">
      <div class="flex-1" />
      <div v-if="isEditingTitle" class="flex-1">
        <input 
          :value="editableTitle" 
          type="text" 
          class="w-full text-2xl font-medium text-gray-600 bg-transparent focus:outline-none text-center" 
          @blur="emit('updateTitle')" 
          @keyup.enter="($event.target as HTMLInputElement).blur()" 
          @input="$event.target && ($event.target as HTMLInputElement).value !== undefined && emit('update:editableTitle', ($event.target as HTMLInputElement).value)" 
          v-focus 
        />
      </div>
      <div 
        v-else 
        class="flex-1 flex flex-col items-center justify-center relative group"
      >
        <h1 
          class="text-2xl font-medium text-gray-600 text-center flex items-center justify-center"
          @click="!currentDocument?.config?.lock && emit('update:isEditingTitle', true)"
        >
          <span v-if="currentDocument?.config?.icon" class="mr-2">{{ currentDocument.config.icon }}</span>
          <span>{{ currentDocument?.name }}</span>
        </h1>
      </div>
      <div class="flex-1 flex items-center justify-end gap-2">
        <div v-if="currentDocument?.config?.lock" class="top-0 right-0 m-4 flex items-center px-3 py-1.5 rounded-md bg-red-50 text-red-500 border border-red-200">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4 mr-1.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M16.5 10.5V6.75a4.5 4.5 0 10-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 002.25-2.25v-6.75a2.25 2.25 0 00-2.25-2.25H6.75a2.25 2.25 0 00-2.25 2.25v6.75a2.25 2.25 0 002.25 2.25z" />
          </svg>
                <span class="text-xs font-medium">Document locked</span>
        </div>
        <button 
          class="text-gray-400 hover:text-yellow-500 transition-colors p-2 rounded-lg hover:bg-gray-100"
          @click="emit('favorite')"
          :class="{
            'text-yellow-500': favoritesStore.favorites.some(f => f.document?.id === currentDocument?.id)
          }"
        >
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="M11.48 3.499a.562.562 0 0 1 1.04 0l2.125 5.111a.563.563 0 0 0 .475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 0 0-.182.557l1.285 5.385a.562.562 0 0 1-.84.61l-4.725-2.885a.562.562 0 0 0-.586 0L6.982 20.54a.562.562 0 0 1-.84-.61l1.285-5.386a.562.562 0 0 0-.182-.557l-4.204-3.602a.562.562 0 0 1 .321-.988l5.518-.442a.563.563 0 0 0 .475-.345L11.48 3.5Z"></path>
          </svg>
        </button>
        <button
          class="text-gray-400 hover:text-blue-500 transition-colors p-2 rounded-lg hover:bg-gray-100"
          @click="toggleConfigSidebar"
          >
          <AdjustmentsHorizontalIcon class="size-6" />
        </button>
      </div>
    </div>
    
    <!-- Config Sidebar -->
    <ConfigSidebar 
      :visible="showConfigSidebar" 
      :currentDocument="currentDocument"
      @close="toggleConfigSidebar"
      @setDocumentIcon="applyIcon"
      @setLock="toggleLock"
      @setFullWidth="fullWidth"
    />
  </div>
</template>

<style scoped>
.emoji-btn {
  cursor: pointer !important;
  user-select: none;
  transition: transform 0.1s ease;
  pointer-events: auto !important;
  z-index: 150;
  -webkit-touch-callout: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  position: relative;
}

.emoji-btn:hover {
  transform: scale(1.15);
  background-color: #f3f4f6;
}
</style>