<script setup lang="ts">
import { ref } from 'vue'

const greeting = ref(`Good ${getTimeOfDay()}, ${getUserFirstName()}`)

function getTimeOfDay() {
  const hour = new Date().getHours()
  if (hour < 12) return 'morning'
  if (hour < 18) return 'afternoon'
  return 'evening'
}

function getUserFirstName() {

}
</script>

<template>
  <main class="flex-1 p-10 overflow-y-auto">
    <h1 class="text-2xl font-medium text-gray-600">{{ greeting }}</h1>
  </main>
</template>