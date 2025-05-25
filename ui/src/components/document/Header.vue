<script setup lang="ts">
import { type Directive, ref } from 'vue';
import { useFavoriteStore } from '../../stores/favorite';
import { type Document } from '../../api/document';
import { FaceSmileIcon, PhotoIcon } from "@heroicons/vue/24/solid";
import IconSidebar from './IconSidebar.vue';

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
  (e: 'iconSidebarVisibility', visible: boolean): void;
}>();

// État pour la sidebar d'icônes
const showIconSidebar = ref(false);

// Fonction pour ouvrir/fermer la sidebar d'icônes
const toggleIconSidebar = () => {
  showIconSidebar.value = !showIconSidebar.value;
  // Désactiver l'éditeur quand la sidebar est ouverte
  emit('toggleEditorEditable', !showIconSidebar.value);
  // Notifier le parent de la visibilité de la sidebar
  emit('iconSidebarVisibility', showIconSidebar.value);
};

// Fonction pour appliquer une icône
const applyIcon = (icon: string) => {
  emit('setDocumentIcon', icon);
};

// Directive focus
const vFocus: Directive = {
  mounted: (el) => el.focus()
}
</script>

<template>
  <div class="sticky top-0 border-b border-e bg-white">
    <div class="flex justify-between items-center h-24 px-8">
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
          class="text-2xl font-medium text-gray-600 text-center cursor-text flex items-center justify-center"
          @click="emit('update:isEditingTitle', true)"
        >
          <span v-if="currentDocument?.config?.icon" class="mr-2">{{ currentDocument.config.icon }}</span>
          <span>{{ currentDocument?.name }}</span>
        </h1>
        <div class="opacity-0 group-hover:opacity-100 transition-opacity mt-2 flex items-center">
          <div 
            class="text-gray-400 cursor-pointer hover:bg-gray-100 hover:text-gray-500 rounded-lg px-1 flex items-center"
            @click="toggleIconSidebar"
          >
            <div class="flex-shrink-0 mr-1">
              <FaceSmileIcon class="size-4 opacity-75" />
            </div>
            <span class="flex-grow">{{ currentDocument?.config?.icon ? 'Change icon' : 'Add an icon' }}</span>
          </div>
          <div class="text-gray-400 cursor-pointer hover:bg-gray-100 hover:text-gray-500 rounded-lg px-1 flex items-center ml-2">
            <div class="flex-shrink-0 mr-1">
              <PhotoIcon class="size-4 opacity-75" />
            </div>
            <span class="flex-grow">Add a background</span>
          </div>
        </div>
      </div>
      <div class="flex-1 flex items-center justify-end gap-2">
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
      </div>
    </div>
    
    <!-- IconSidebar component -->
    <IconSidebar 
      :visible="showIconSidebar" 
      :currentDocument="currentDocument"
      @close="toggleIconSidebar"
      @setDocumentIcon="applyIcon"
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