<script setup lang="ts">
import { BubbleMenu } from '@tiptap/vue-3';
import type { Editor } from '@tiptap/core';
import '@tiptap/extension-underline';
import {
  BoldIcon,
  ItalicIcon,
  UnderlineIcon,
  StrikethroughIcon,
  CodeIcon,
  Heading1Icon,
  Heading2Icon,
  Heading3Icon,
  PaletteIcon,
  HighlighterIcon,
  ChevronDownIcon,
  ListIcon,
  ListOrderedIcon,
  QuoteIcon,
  TextIcon
} from "lucide-vue-next";
import { computed, ref } from 'vue';
import type { EditorState } from "@tiptap/pm/state";
import type { EditorView } from "@tiptap/pm/view";

// Définir les props du composant
const props = defineProps({
  editor: {
    type: Object as () => Editor,
    required: true
  }
});

interface Color {
  name: string;
  value: string;
}

// État pour gérer l'affichage des menus
const showContentTypeMenu = ref(false);
const showColorMenu = ref(false);
const showBgColorMenu = ref(false);

// Couleurs disponibles pour le texte et le fond
const colors: Color[] = [
  { name: "Default", value: "inherit" },
  { name: "Gray", value: "#6B7280" },
  { name: "Brown", value: "#92400E" },
  { name: "Orange", value: "#EA580C" },
  { name: "Yellow", value: "#CA8A04" },
  { name: "Green", value: "#16A34A" },
  { name: "Blue", value: "#2563EB" },
  { name: "Purple", value: "#9333EA" },
  { name: "Pink", value: "#DB2777" },
  { name: "Red", value: "#DC2626" },
];

// Fonction pour déterminer si le menu doit être affiché
const shouldShow = (params: {
  editor: any;
  view: EditorView;
  state: EditorState;
  oldState?: EditorState;
  from: number;
  to: number;
}) => {
  const { state, from } = params;
  const { doc, selection } = state;
  const { empty } = selection;

  // Ne pas afficher si la sélection est vide
  if (empty) return false;

  // Ne pas afficher si l'éditeur n'est pas éditable (document verrouillé)
  if (!props.editor.isEditable) return false;

  // Ne pas afficher pour les images
  const nodeAtSelection = doc.nodeAt(from);
  if (nodeAtSelection && nodeAtSelection.type.name === "imageEditor") {
    return false;
  }

  // Ne pas afficher pour les blocs de code
  if (props.editor.isActive('codeBlock')) {
    return false;
  }

  return true;
};

// Types de contenu disponibles
const contentTypes = [
  {
    name: "paragraph",
    label: "Text",
    icon: TextIcon,
    command: () => props.editor.chain().focus().setParagraph().run(),
    attrs: undefined
  },
  {
    name: "heading",
    label: "Heading 1",
    icon: Heading1Icon,
    command: () => props.editor.chain().focus().toggleHeading({ level: 1 }).run(),
    attrs: { level: 1 }
  },
  {
    name: "heading",
    label: "Heading 2",
    icon: Heading2Icon,
    command: () => props.editor.chain().focus().toggleHeading({ level: 2 }).run(),
    attrs: { level: 2 }
  },
  {
    name: "heading",
    label: "Heading 3",
    icon: Heading3Icon,
    command: () => props.editor.chain().focus().toggleHeading({ level: 3 }).run(),
    attrs: { level: 3 }
  },
  {
    name: "bulletList",
    label: "Bullet List",
    icon: ListIcon,
    command: () => props.editor.chain().focus().toggleBulletList().run(),
    attrs: undefined
  },
  {
    name: "orderedList",
    label: "Numbered List",
    icon: ListOrderedIcon,
    command: () => props.editor.chain().focus().toggleOrderedList().run(),
    attrs: undefined
  },
  {
    name: "codeBlock",
    label: "Code Block",
    icon: CodeIcon,
    command: () => props.editor.chain().focus().toggleCodeBlock().run(),
    attrs: undefined
  },
  {
    name: "blockquote",
    label: "Quote",
    icon: QuoteIcon,
    command: () => props.editor.chain().focus().toggleBlockquote().run(),
    attrs: undefined
  },
];

// Actions de formatage de texte
const textActions = [
  {
    name: "bold",
    label: "Bold",
    icon: BoldIcon,
    command: () => props.editor.chain().focus().toggleBold().run()
  },
  {
    name: "italic",
    label: "Italic",
    icon: ItalicIcon,
    command: () => props.editor.chain().focus().toggleItalic().run()
  },
  {
    name: "underline",
    label: "Underline",
    icon: UnderlineIcon,
    command: () => props.editor.chain().focus().toggleUnderline().run()
  },
  {
    name: "strike",
    label: "Strikethrough",
    icon: StrikethroughIcon,
    command: () => props.editor.chain().focus().toggleStrike().run()
  }
];

// Type de contenu actuel
const currentContentType = computed(() => {
  if (props.editor.isActive("heading", { level: 1 })) return "Heading 1";
  if (props.editor.isActive("heading", { level: 2 })) return "Heading 2";
  if (props.editor.isActive("heading", { level: 3 })) return "Heading 3";
  if (props.editor.isActive("bulletList")) return "Bullet List";
  if (props.editor.isActive("orderedList")) return "Numbered List";
  if (props.editor.isActive("codeBlock")) return "Code Block";
  if (props.editor.isActive("blockquote")) return "Quote";
  return "Text";
});

// Fonctions pour gérer les menus
const toggleContentTypeMenu = () => {
  showContentTypeMenu.value = !showContentTypeMenu.value;
  showColorMenu.value = false;
  showBgColorMenu.value = false;
};

const toggleColorMenu = () => {
  showColorMenu.value = !showColorMenu.value;
  showContentTypeMenu.value = false;
  showBgColorMenu.value = false;
};

const toggleBgColorMenu = () => {
  showBgColorMenu.value = !showBgColorMenu.value;
  showContentTypeMenu.value = false;
  showColorMenu.value = false;
};

const setContentType = (command: () => void) => {
  command();
  showContentTypeMenu.value = false;
};

const setTextColor = (color: string) => {
  props.editor.chain().focus().setColor(color).run();
  showColorMenu.value = false;
};

const setBackgroundColor = (color: string) => {
  props.editor.chain().focus().setHighlight({ color }).run();
  showBgColorMenu.value = false;
};

const isTextColorActive = (color: string) => {
  return props.editor.isActive('textStyle', { color });
};

const isBackgroundColorActive = (color: string) => {
  return props.editor.isActive('highlight', { color });
};
</script>

<template>
  <BubbleMenu
    v-if="props.editor"
    :editor="props.editor"
    :should-show="shouldShow"
    class="flex flex-col bg-white rounded-md shadow-md p-1"
  >
    <div class="flex items-center">
      <!-- Menu type de contenu -->
      <div class="flex items-center relative">
        <button
          @click="toggleContentTypeMenu"
          :class="{ 'bg-gray-100/80': showContentTypeMenu }"
          class="flex items-center text-sm px-2.5 py-1.5 rounded hover:bg-gray-100/80"
        >
          {{ currentContentType }}
          <ChevronDownIcon class="w-3.5 h-3.5 ml-1" />
        </button>
        <div
          v-if="showContentTypeMenu"
          class="absolute top-full left-0 bg-white rounded-md shadow-md z-10 py-1.5 min-w-[140px]"
        >
          <button
            v-for="type in contentTypes"
            :key="type.name + (type.attrs ? JSON.stringify(type.attrs) : '')"
            @click="setContentType(type.command)"
            :class="{
              'bg-gray-100/80': props.editor.isActive(type.name, type.attrs),
            }"
            class="flex items-center w-full text-left px-3 py-2 hover:bg-gray-100/80 rounded"
          >
            <component :is="type.icon" class="w-4 h-4 mr-2" />
            <span class="text-sm">{{ type.label }}</span>
          </button>
        </div>
      </div>

      <div class="w-px my-auto h-8 bg-gray-200/70 mx-1"></div>

      <!-- Actions de formatage -->
      <div class="flex items-center space-x-1">
        <button
          v-for="action in textActions"
          :key="action.name"
          @click="action.command()"
          :class="{
            'bg-gray-100/80 text-primary': props.editor.isActive(action.name),
          }"
          :title="action.label"
          class="rounded hover:bg-gray-100/80 w-8 h-8 grid place-items-center"
        >
          <component :is="action.icon" class="w-4 h-4" />
        </button>
      </div>

      <!-- Menu couleur de texte -->
      <div class="flex items-center relative ml-1">
        <button
          @click="toggleColorMenu"
          :class="{ 'bg-gray-100/80': showColorMenu }"
          class="rounded hover:bg-gray-100/80 w-8 h-8 grid place-items-center"
          title="Text color"
        >
          <PaletteIcon class="w-4 h-4" />
        </button>
        <div
          v-if="showColorMenu"
          class="absolute top-full left-0 bg-white rounded-md shadow-md z-10 py-1.5"
        >
          <button
            v-for="color in colors"
            :key="color.name"
            @click="setTextColor(color.value)"
            :class="{
              'bg-gray-100/80': isTextColorActive(color.value),
            }"
            class="flex items-center w-full text-left px-3 py-2 hover:bg-gray-100/80 rounded"
          >
            <span
              class="w-4 h-4 rounded-full mr-2 border-2 border-gray-200"
              :style="{ backgroundColor: color.value }"
            ></span>
            <span class="text-sm font-medium text-gray-600">{{
              color.name
            }}</span>
          </button>
        </div>
      </div>

      <!-- Menu couleur de fond -->
      <div class="flex items-center relative">
        <button
          @click="toggleBgColorMenu"
          :class="{ 'bg-gray-100/80': showBgColorMenu }"
          class="rounded hover:bg-gray-100/80 w-8 h-8 grid place-items-center"
          title="Background color"
        >
          <HighlighterIcon class="w-4 h-4" />
        </button>
        <div
          v-if="showBgColorMenu"
          class="absolute top-full left-0 bg-white rounded-md shadow-md z-10 py-1.5"
        >
          <button
            v-for="color in colors"
            :key="color.name"
            @click="setBackgroundColor(color.value)"
            :class="{
              'bg-gray-100/80': isBackgroundColorActive(color.value),
            }"
            class="flex items-center w-full text-left px-3 py-2 hover:bg-gray-100/80 rounded"
          >
            <span
              class="w-4 h-4 rounded-full mr-2 border-2 border-gray-200"
              :style="{ backgroundColor: color.value }"
            ></span>
            <span class="text-sm font-medium text-gray-600">{{
              color.name
            }}</span>
          </button>
        </div>
      </div>
    </div>
  </BubbleMenu>
</template>

<style scoped>
/* Les styles Tailwind sont appliqués directement dans le template */
</style>