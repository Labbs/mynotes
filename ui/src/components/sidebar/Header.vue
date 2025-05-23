<script setup lang="ts">
const { isCollapsed, isHovered } = defineProps<{
  isCollapsed: boolean,
  isHovered: boolean
}>()
</script>

<template>
  <div class="text-[13px]">
    <router-link to="/" class="flex items-center gap-x-2 text-gray-500">
      <div
        class="size-5 rounded-lg bg-gray-100 flex items-center justify-center"
      >
        <span class="font-medium">L</span>
      </div>
      <span v-show="!isCollapsed || isHovered" class="font-medium">Labbs</span>
    </router-link>
  </div>
</template>
