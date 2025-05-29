<script setup lang="ts">
import { computed } from "vue";
import { useRouter } from "vue-router";
import { HomeIcon, Cog6ToothIcon, MagnifyingGlassIcon } from "@heroicons/vue/24/outline";

const { isCollapsed, isHovered } = defineProps<{
  isCollapsed: boolean,
  isHovered: boolean
}>()

const router = useRouter();

const isHomePage = computed(() => {
  return router.currentRoute.value.path === "/";
});
const isSettingsPage = computed(() => {
  return router.currentRoute.value.path === "/settings";
});
</script>

<template>
  <div class="pt-2 pb-1 text-[13px]">
    <div       class="flex items-center gap-x-2 rounded-lg px-1 py-1 text-gray-500 hover:bg-gray-100 hover:text-gray-700"
      :class="{
        'justify-center': isCollapsed && !isHovered,
      }"
    >
      <div class="flex-shrink-0">
        <MagnifyingGlassIcon class="size-5 opacity-75" />
      </div>
      <span v-show="!isCollapsed || isHovered" class="flex-grow"> Search </span>
    </div>
    <router-link
      to="/"
      class="flex items-center gap-x-2 rounded-lg px-1 py-1 text-gray-500 hover:bg-gray-100 hover:text-gray-700"
      :class="{
        'bg-gray-100 text-gray-700': isHomePage,
        'justify-center': isCollapsed && !isHovered,
      }"
    >
      <div class="flex-shrink-0">
        <HomeIcon class="size-5 opacity-75" />
      </div>
      <span v-show="!isCollapsed || isHovered" class="flex-grow"> Home </span>
    </router-link>

    <router-link
      to="/settings"
      class="flex items-center gap-x-2 rounded-lg px-1 py-1 text-gray-500 hover:bg-gray-100 hover:text-gray-700 mt-1"
      :class="{
        'bg-gray-50 text-gray-700': isSettingsPage,
        'justify-center': isCollapsed && !isHovered,
      }"
    >
      <div class="flex-shrink-0">
        <Cog6ToothIcon class="size-5 opacity-75" />
      </div>
      <span v-show="!isCollapsed || isHovered" class="flex-grow">
        Settings
      </span>
    </router-link>
  </div>
</template>
