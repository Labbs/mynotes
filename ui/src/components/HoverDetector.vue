<script setup lang="ts">
import { useSidebarStore } from "../stores/sidebar";

const sidebarStore = useSidebarStore();
</script>

<template>
  <div 
    v-if="sidebarStore.isCollapsed"
    class="hover-detector"
    @mouseenter="sidebarStore.setHovering(true)"
  ></div>
</template>

<style scoped>
.hover-detector {
  position: fixed;
  left: 0;
  top: 0;
  height: 100vh;
  width: 15px;
  z-index: 99;
}
</style>
