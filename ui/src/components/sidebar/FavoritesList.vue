<script setup lang="ts">
import { onMounted } from 'vue';
import { useFavoriteStore } from "../../stores/favorite";

const { isCollapsed, isHovered } = defineProps<{
  isCollapsed: boolean,
  isHovered: boolean
}>()

const favoritesStore = useFavoriteStore();

const capitalizeFirst = (str: string | undefined) => {
  if (!str) return ''
  return str.charAt(0).toUpperCase() + str.slice(1)
}

onMounted(async() => {
  await favoritesStore.fetchFavorites();
});
</script>

<template>
  <div class="pt-2 pb-1 text-[13px]" v-if=favoritesStore.favorites.length>
    <router-link
      v-for="favorite in favoritesStore.favorites"
      :key="favorite.id"
      :to="`/d/${favorite.document?.slug || ''}`"
      class="flex items-center gap-x-2 rounded-lg px-2 py-1 text-[13px] text-gray-500 hover:bg-gray-100 hover:text-gray-700"
      :class="{'justify-center': isCollapsed && !isHovered}">
      <div v-if="!favorite.document?.config.icon" class="size-4 rounded bg-gray-100 flex items-center justify-center opacity-75 text-xs font-medium">
        {{ favorite.document?.name[0]?.toUpperCase() }}
      </div>
      <div v-else>
        {{ favorite.document?.config.icon }}
      </div>
      <span v-show="!isCollapsed || isHovered" class="flex-grow text-left">{{ capitalizeFirst(favorite.document?.name) }}</span>
    </router-link>
  </div>
</template>