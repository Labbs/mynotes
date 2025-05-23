<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  name?: string
  icon?: string
}>()

const initial = computed(() => {
  return props.name ? props.name.charAt(0).toUpperCase() : '?'
})
</script>

<template>
  <div v-if="!icon" class="w-4 h-4 rounded bg-gray-700 flex items-center justify-center text-xs font-medium text-gray-300">
    {{ initial }}
  </div>
  <img v-else :src="icon" class="w-4 h-4 rounded" :alt="name || ''" />
</template>