<script setup lang="ts">
import { computed, onMounted, ref, onUnmounted } from "vue";
import { useRouter } from "vue-router";
import { useSpaceStore } from "../stores/space";
import { useSidebarStore } from "../stores/sidebar";
import { useAuthStore } from "../stores/auth";
import { useFavoriteStore } from "../stores/favorite";
import { usePreferencesStore } from "../stores/preferences";

import Header from "./sidebar/Header.vue";
import CommonMenu from "./sidebar/CommonMenu.vue";
import SpaceList from "./sidebar/SpaceList.vue";
import FavoritesList from "./sidebar/FavoritesList.vue";

const router = useRouter();
const spaceStore = useSpaceStore();
const sidebarStore = useSidebarStore();
const authStore = useAuthStore();
const favoritesStore = useFavoriteStore();
const preferencesStore = usePreferencesStore();

onMounted(async () => {
  // Charger d'abord les préférences si l'utilisateur est connecté
  if (authStore.isAuthenticated) {
    await preferencesStore.loadPreferences();
    // Puis initialiser le sidebar avec les préférences
    sidebarStore.initializeFromPreferences();
  }
  
  spaceStore.fetchSpaces();
  favoritesStore.fetchFavorites();
});

const isResizing = ref(false);

const sidebarWidth = computed(() => {
  if (sidebarStore.isCollapsed && !sidebarStore.isHovering) {
    return { width: '0px', minWidth: '0px' }; // Complètement masqué quand collapse
  }
  return { width: `${sidebarStore.width}px`, minWidth: '212px' };
});

function startResize(_event: MouseEvent) {
  isResizing.value = true;
  document.body.classList.add("resizing");
  document.addEventListener("mousemove", handleResize);
  document.addEventListener("mouseup", stopResize);
}

function handleResize(event: MouseEvent) {
  if (!isResizing.value) return;
  requestAnimationFrame(() => {
    sidebarStore.setWidth(event.clientX);
  });
}

function stopResize() {
  isResizing.value = false;
  document.body.classList.remove("resizing");
  document.removeEventListener("mousemove", handleResize);
  document.removeEventListener("mouseup", stopResize);
}

onMounted(() => {
  document.addEventListener("mouseup", stopResize);
});

onUnmounted(() => {
  document.removeEventListener("mouseup", stopResize);
});

const handleLogout = () => {
  authStore.logout();
  router.push("/auth/login");
};

const handleMouseEnter = () => {
  sidebarStore.setHovering(true);
};

const handleMouseLeave = () => {
  sidebarStore.setHovering(false);
};
</script>

<template>
  <div class="flex overflow-hidden"
    :class="{
      'w-0': sidebarStore.isCollapsed && !sidebarStore.isHovering,
      'w-auto': !sidebarStore.isCollapsed || sidebarStore.isHovering,
    }">
    <div
      :style="sidebarStore.isCollapsed && !sidebarStore.isHovering ? { width: '0px', minWidth: '0px' } : sidebarWidth"
      :class="[
        'flex h-screen flex-col justify-between border-e border-gray-200 bg-white transition-all duration-300',
        { 'transition-none': isResizing },
      ]"
      @mouseenter="handleMouseEnter"
      @mouseleave="handleMouseLeave"
    >
      <div>
        <div class="flex items-center justify-between p-2">
          <Header
            :isCollapsed="sidebarStore.isCollapsed"
            :isHovered="sidebarStore.isHovering"
          />
          <button
            v-show="!sidebarStore.isCollapsed"
            @click="sidebarStore.toggleCollapse"
            class="p-2 hover:bg-gray-100 rounded-lg"
          >
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
                d="M11 19l-7-7 7-7M19 19l-7-7 7-7"
              />
            </svg>
          </button>
        </div>

        <div class="border-t border-gray-100">
          <div class="px-2">
            <CommonMenu
              :isCollapsed="sidebarStore.isCollapsed"
              :isHovered="sidebarStore.isHovering"
            />
          </div>
        </div>

        <div class="border-t border-gray-100" v-if="favoritesStore.favorites.length">
          <div class="px-2">
            <FavoritesList
              :isCollapsed="sidebarStore.isCollapsed"
              :isHovered="sidebarStore.isHovering"
            />
          </div>
        </div>

        <div class="border-t border-gray-100">
          <div class="px-2">
            <SpaceList
              :isCollapsed="sidebarStore.isCollapsed"
              :isHovered="sidebarStore.isHovering"
            />
          </div>
        </div>
      </div>

      <div
        class="sticky inset-x-0 bottom-0 border-t border-gray-100 bg-white p-2"
      >
        <form action="#">
          <button
            @click="handleLogout"
            class="flex w-full items-center gap-x-2 rounded-lg px-2 py-2 text-gray-500 hover:bg-gray-100 hover:text-gray-700"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="size-5 opacity-75"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              stroke-width="2"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
              />
            </svg>

            <span
              v-show="!sidebarStore.isCollapsed || sidebarStore.isHovering"
              class="text-[13px]">Logout</span>
          </button>
        </form>
      </div>
    </div>

    <!-- Resizer -->
    <div
      v-show="!sidebarStore.isCollapsed || sidebarStore.isHovering"
      class="w-[1px] hover:w-2 bg-gray-200 hover:bg-gray-200 cursor-col-resize transition-all"
      @mousedown.prevent="startResize"
    />
  </div>
</template>

<style scoped>
.transition-all {
  transition-property: all;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 300ms;
}

/* Désactiver la sélection de texte pendant le redimensionnement */
*::selection {
  user-select: none;
}

/* Ajouter un curseur de redimensionnement pendant le drag */
:deep(*) {
  cursor: inherit;
}

:global(body.resizing) {
  cursor: col-resize !important;
  user-select: none;
}
</style>
