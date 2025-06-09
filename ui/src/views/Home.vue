<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { meApi } from '../api/me'

const greeting = ref('Good ' + getTimeOfDay())
const isLoading = ref(true)

function getTimeOfDay() {
  const hour = new Date().getHours()
  if (hour < 12) return 'morning'
  if (hour < 18) return 'afternoon'
  return 'evening'
}

async function loadProfile() {
  try {
    isLoading.value = true
    const { data } = await meApi.getProfile()
    console.log('Profile data:', data)
    if (data.name) {
      greeting.value = `Good ${getTimeOfDay()}, ${data.name}`
    }
  } catch (error) {
    console.error('Failed to load profile:', error)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  loadProfile()
})
</script>

<template>
  <main class="flex-1 p-10 overflow-y-auto flex flex-col items-center bg-white">
    <div v-if="isLoading" class="animate-pulse">
      <div class="h-8 w-64 bg-gray-200 rounded"></div>
    </div>
    <h1 v-else class="text-2xl font-medium text-gray-600">{{ greeting }}</h1>
  </main>
</template>