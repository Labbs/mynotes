<template>
  <div class="relative">
    <div 
      class="absolute z-50 mt-1 right-0 w-48 bg-white rounded-md shadow-lg py-1 text-sm"
      @click.stop.prevent
    >
      <button 
        @click.stop.prevent="$emit('select', 'document')"
        class="flex w-full items-center px-4 py-2 text-gray-700 hover:bg-gray-100"
      >
        <svg class="w-4 h-4 mr-2 text-gray-500" viewBox="0 0 24 24" fill="none" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
        </svg>
        Document
      </button>
      

      <button 
        @click.stop.prevent="$emit('select', 'database')"
        class="flex w-full items-center px-4 py-2 text-gray-700 hover:bg-gray-100"
      >
        <svg class="w-4 h-4 mr-2 text-gray-500" viewBox="0 0 24 24" fill="none" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4m0 5c0 2.21-3.582 4-8 4s-8-1.79-8-4" />
        </svg>
        Database
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
defineEmits<{
  (e: 'select', type: 'document' | 'database'): void
}>();
</script>
