import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useSidebarStore = defineStore('sidebar', () => {
  const MIN_WIDTH = 150 // 16rem (w-64)
  const MAX_WIDTH = 600
  const isCollapsed = ref(localStorage.getItem('sidebarCollapsed') === 'true')
  const width = ref(parseInt(localStorage.getItem('sidebarWidth') || String(MIN_WIDTH)))
  const isHovering = ref(false)

  function toggleCollapse() {
    isCollapsed.value = !isCollapsed.value
    localStorage.setItem('sidebarCollapsed', isCollapsed.value.toString())
  }

  function setWidth(newWidth: number) {
    width.value = Math.max(MIN_WIDTH, Math.min(MAX_WIDTH, newWidth))
    localStorage.setItem('sidebarWidth', width.value.toString())
  }
  
  function setHovering(value: boolean) {
    isHovering.value = value
  }
  
  function collapse() {
    isCollapsed.value = true
    localStorage.setItem('sidebarCollapsed', 'true')
  }
  
  function expand() {
    isCollapsed.value = false
    localStorage.setItem('sidebarCollapsed', 'false')
  }

  return {
    isCollapsed,
    isHovering,
    width,
    MIN_WIDTH,
    MAX_WIDTH,
    toggleCollapse,
    setHovering,
    collapse,
    expand,
    setWidth
  }
}) 