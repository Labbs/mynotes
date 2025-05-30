<script setup lang="ts">
import { useSidebarStore } from "../stores/sidebar";

const sidebarStore = useSidebarStore();
</script>

<template>
  <div 
    v-if="sidebarStore.isCollapsed" 
    class="sidebar-toggle"
    @click="sidebarStore.expand()"
  >
    <span class="icon">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-5 w-5"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M4 6h16M4 12h16M4 18h16"
        />
      </svg>
    </span>
  </div>
</template>

<style scoped>
.sidebar-toggle {
  position: fixed;
  top: 12px;
  left: 12px;
  width: 32px;
  height: 32px;
  background-color: white;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  z-index: 101;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
  transition: background-color 0.2s;
}

.sidebar-toggle:hover {
  background-color: #f3f4f6;
}
</style>
