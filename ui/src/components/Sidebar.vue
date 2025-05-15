<script setup lang="ts">
import { computed, onMounted, ref, onUnmounted } from "vue";
import { useRouter } from "vue-router";
import { useSpaceStore } from "../stores/space";
import { useSidebarStore } from "../stores/sidebar";
import { useAuthStore } from "../stores/auth";

import Header from "./sidebar/Header.vue";
import CommonMenu from "./sidebar/CommonMenu.vue";
import SpaceList from "./sidebar/SpaceList.vue";

const router = useRouter();
const spaceStore = useSpaceStore();
const sidebarStore = useSidebarStore();
const authStore = useAuthStore();

const props = defineProps<{
  spaceId?: string;
}>();

onMounted(() => {
  spaceStore.fetchSpaces();
});

const isHovered = ref(false);
const isResizing = ref(false);

const sidebarWidth = computed(() => {
  if (sidebarStore.isCollapsed && !isHovered.value) {
    return "w-16";
  }
  return { width: `${sidebarStore.width}px` };
});

function startResize(event: MouseEvent) {
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
</script>

<template>
  <div class="flex">
    <div
      :style="sidebarStore.isCollapsed && !isHovered ? {} : sidebarWidth"
      :class="[
        sidebarStore.isCollapsed && !isHovered ? 'w-16' : '',
        'flex h-screen flex-col justify-between border-e bg-white transition-all duration-300',
        { 'transition-none': isResizing },
      ]"
      @mouseenter="isHovered = true"
      @mouseleave="isHovered = false"
    >
      <div>
        <div class="flex items-center justify-between p-2">
          <Header
            :isCollapsed="sidebarStore.isCollapsed"
            :isHovered="isHovered"
          />
          <button
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
                :d="
                  sidebarStore.isCollapsed
                    ? 'M13 5l7 7-7 7M5 5l7 7-7 7'
                    : 'M11 19l-7-7 7-7M19 19l-7-7 7-7'
                "
              />
            </svg>
          </button>
        </div>

        <div class="border-t border-gray-100">
          <div class="px-2">
            <CommonMenu
              :isCollapsed="sidebarStore.isCollapsed"
              :isHovered="isHovered"
            />
          </div>
        </div>

        <div class="border-t border-gray-100">
          <div class="px-2">
            <SpaceList
              :isCollapsed="sidebarStore.isCollapsed"
              :isHovered="isHovered"
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
              v-show="!sidebarStore.isCollapsed || isHovered"
              class="text-[13px]"
            >
              Logout
            </span>
          </button>
        </form>
      </div>
    </div>

    <!-- Resizer -->
    <div
      v-show="!sidebarStore.isCollapsed || isHovered"
      class="w-1 hover:w-2 bg-transparent hover:bg-gray-200 cursor-col-resize transition-all"
      @mousedown.prevent="startResize"
    />
  </div>
</template>

<style scoped>
.transition-all {
  transition-property: all;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 10ms;
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
