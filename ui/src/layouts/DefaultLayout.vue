<script setup lang="ts">
import Sidebar from '../components/Sidebar.vue'
import SidebarToggle from '../components/SidebarToggle.vue'
import HoverDetector from '../components/HoverDetector.vue'
import { useSidebarStore } from '../stores/sidebar'

const sidebarStore = useSidebarStore()
</script>

<template>
  <div class="flex">
    <Sidebar />
    <SidebarToggle />
    <HoverDetector />
    <router-view class="flex-1" :class="{ 'pl-0': sidebarStore.isCollapsed && !sidebarStore.isHovering }"></router-view>
  </div>
</template> 