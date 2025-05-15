<script setup lang="ts">
import { ref } from 'vue'
import { type Block } from '../../api/document'
import TextBlock from './blocks/TextBlock.vue'
import BlockTypeMenu from './BlockTypeMenu.vue'

const props = defineProps<{
  blocks: Block[]
}>()

const newBlockPosition = ref<number | null>(null)
const showMenuAtIndex = ref<number | null>(null)

const showBlockTypeMenu = (position: number) => {
  newBlockPosition.value = position
}

const handleKeyDown = (event: KeyboardEvent, index: number) => {
  if (event.key === 'Enter' && !event.shiftKey) {
    event.preventDefault()
    showBlockTypeMenu(index + 1)
  }
}

// Bloc vide par défaut
const emptyBlock: Block = {
  id: 'empty',
  type: 'text',
  content: { markdown: '' },
  document_id: props.blocks[0]?.document_id || '',
  space_id: props.blocks[0]?.space_id || ''
}
</script>

<template>
  <div class="space-y-2">
    <div 
      v-for="(block, index) in [...blocks, emptyBlock]" 
      :key="block.id"
      class="group relative flex items-start gap-2"
    >
      <!-- Block type selector -->
      <button 
        class="p-1 text-gray-400 hover:text-gray-600 opacity-0 group-hover:opacity-100 transition-opacity self-center"
        @click="showMenuAtIndex = index"
      >
        <span v-if="block.type === 'text'">¶</span>
        <span v-else-if="block.type === 'heading-1'">H1</span>
        <span v-else-if="block.type === 'heading-2'">H2</span>
        <span v-else-if="block.type === 'bullet-list'">•</span>
        <span v-else-if="block.type === 'table'">⊞</span>
      </button>

      <!-- Block content -->
      <div class="flex-1">
        <component
          :is="block.type === 'text' ? TextBlock : null"
          :block="block"
          @keydown="handleKeyDown($event, index)"
        />
      </div>

      <!-- Block type menu -->
      <BlockTypeMenu
        v-if="showMenuAtIndex === index"
        :position="index"
        @select="(type, pos) => { /* TODO: Implement block type change */ showMenuAtIndex = null }"
        @close="showMenuAtIndex = null"
      />
    </div>
  </div>
</template> 