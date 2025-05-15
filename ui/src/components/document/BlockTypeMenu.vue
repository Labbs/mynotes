<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'

const props = defineProps<{
  position: number
}>()

const emit = defineEmits<{
  (e: 'select', type: string, position: number): void
  (e: 'close'): void
}>()

const menuRef = ref<HTMLDivElement | null>(null)

const handleClickOutside = (event: MouseEvent) => {
  if (menuRef.value && !menuRef.value.contains(event.target as Node)) {
    emit('close')
  }
}

onMounted(() => {
  document.addEventListener('mousedown', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('mousedown', handleClickOutside)
})

const blockTypes = [
  { id: 'text', label: 'Text', icon: '¶' },
  { id: 'heading-1', label: 'Heading 1', icon: 'H1' },
  { id: 'heading-2', label: 'Heading 2', icon: 'H2' },
  { id: 'bullet-list', label: 'Bullet List', icon: '•' },
  { type: 'divider' },
  { id: 'quote', label: 'Quote', icon: '"' },
  { id: 'callout', label: 'Callout', icon: '💡' },
  { id: 'code', label: 'Code Block', icon: '</>' },
  { type: 'divider' },
  { id: 'two-columns', label: '2 Columns Layout', icon: '◫◫' },
  { id: 'three-columns', label: '3 Columns Layout', icon: '◫◫◫' },
  { type: 'divider' },
  { id: 'action-button', label: 'Action Button', icon: '⚡' },
  { id: 'table', label: 'Table', icon: '⊞' },
  { id: 'divider', label: 'Divider', icon: '—' },
]
</script>

<template>
  <div 
    ref="menuRef"
    class="absolute left-0 z-10 mt-1 w-60 rounded-lg border border-gray-200 bg-white shadow-lg"
  >
    <div class="p-2 space-y-1">
      <template v-for="(type, index) in blockTypes" :key="index">
        <div 
          v-if="type.type === 'divider'"
          class="h-px bg-gray-200 my-1"
        />
        
        <button
          v-else
          class="flex w-full items-center gap-x-2 rounded-md px-2 py-1 text-sm text-gray-600 hover:bg-gray-100"
          @click="emit('select', type.id, position)"
        >
          <span class="flex-shrink-0 text-gray-400">{{ type.icon }}</span>
          <span>{{ type.label }}</span>
        </button>
      </template>
    </div>
  </div>
</template> 