import { defineStore } from 'pinia'
import { ref } from 'vue'
import { usePreferencesStore } from './preferences'

export const useSidebarStore = defineStore('sidebar', () => {
  const MIN_WIDTH = 150 // 16rem (w-64)
  const MAX_WIDTH = 600
  const preferencesStore = usePreferencesStore()
  
  const isCollapsed = ref(false)
  const width = ref(MIN_WIDTH)
  const isHovering = ref(false)

  // Initialiser les valeurs depuis les préférences
  function initializeFromPreferences() {
    isCollapsed.value = preferencesStore.preferences.ui.sidebarCollapsed
    width.value = preferencesStore.preferences.ui.sidebarWidth
  }

  function toggleCollapse() {
    isCollapsed.value = !isCollapsed.value
    preferencesStore.updatePreference('sidebarCollapsed', isCollapsed.value)
  }

  function setWidth(newWidth: number) {
    const clampedWidth = Math.max(MIN_WIDTH, Math.min(MAX_WIDTH, newWidth))
    width.value = clampedWidth
    preferencesStore.updatePreference('sidebarWidth', clampedWidth)
  }
  
  function setHovering(value: boolean) {
    isHovering.value = value
  }
  
  function collapse() {
    isCollapsed.value = true
    preferencesStore.updatePreference('sidebarCollapsed', true)
  }
  
  function expand() {
    isCollapsed.value = false
    preferencesStore.updatePreference('sidebarCollapsed', false)
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
    setWidth,
    initializeFromPreferences
  }
}) 