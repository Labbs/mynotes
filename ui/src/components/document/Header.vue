<script setup lang="ts">
const props = defineProps<{
  id: string
  title: string
}>()
</script>

<template>
  <div class="sticky top-0 bg-[#1a1a1a] border-b border-gray-800/50 mb-8">
    <div class="flex justify-between items-center h-16 px-8">
      <div class="flex-1" />
      <div class="flex-1 text-center">
        <h1 class="text-2xl font-medium text-gray-600">{{ props.title }}</h1>
      </div>
    </div>
  </div>
</template>