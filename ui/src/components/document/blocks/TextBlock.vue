<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { type Block } from '../../../api/document'
import { useEditor, EditorContent } from '@tiptap/vue-3'
import StarterKit from '@tiptap/starter-kit'
import TextStyle from '@tiptap/extension-text-style'
import Color from '@tiptap/extension-color'
import Highlight from '@tiptap/extension-highlight'
import {
  BoldIcon,
  ItalicIcon,
  StrikethroughIcon,
  Code2Icon,
  Heading1Icon,
  Heading2Icon,
  Heading3Icon,
  PaletteIcon,
  HighlighterIcon,
} from 'lucide-vue-next'

type Level = 1 | 2 | 3 | 4 | 5 | 6

const props = defineProps<{
  block: Block
}>()

const editor = useEditor({
  content: props.block.content ?? '',
  extensions: [
    StarterKit.configure({
      heading: {
        levels: [1, 2, 3]
      }
    }),
    TextStyle,
    Color.configure({
      types: ['textStyle']
    }),
    Highlight.configure({ 
      multicolor: true,
    }),
  ],
  editorProps: {
    attributes: {
      class: 'prose-sm focus:outline-none',
    },
    handleKeyDown: (view, event) => {
      if (event.key === 'Enter') {
        if (event.shiftKey) {
          view.dispatch(view.state.tr.replaceSelectionWith(
            view.state.schema.nodes.hardBreak.create()
          ))
          return true
        } else {
          const { from, to } = view.state.selection
          const currentContent = view.state.doc.textBetween(from, view.state.doc.content.size)
          const beforeContent = view.state.doc.textBetween(0, from)
          
          view.dispatch(view.state.tr.delete(from, view.state.doc.content.size))
          
          emit('create-block', props.block.id, currentContent)
          return true
        }
      }
      return false
    }
  },
})

const showColorPicker = ref(false)
const showHighlightPicker = ref(false)

const colors = [
  { name: 'Default', color: '#000000' },
  { name: 'Gray', color: '#666666' },
  { name: 'Red', color: '#EF4444' },
  { name: 'Orange', color: '#F97316' },
  { name: 'Yellow', color: '#EAB308' },
  { name: 'Green', color: '#22C55E' },
  { name: 'Blue', color: '#3B82F6' },
  { name: 'Purple', color: '#A855F7' },
  { name: 'Pink', color: '#EC4899' },
]

const highlights = [
  { name: 'Yellow', color: '#FEF9C3' },
  { name: 'Orange', color: '#FFEDD5' },
  { name: 'Green', color: '#DCFCE7' },
  { name: 'Blue', color: '#DBEAFE' },
  { name: 'Purple', color: '#F3E8FF' },
  { name: 'Pink', color: '#FCE7F3' },
]

const isEditing = ref(false)
const hasSelection = ref(false)
const blockRef = ref<HTMLElement | null>(null)

// Attendre l'initialisation de l'éditeur
watch(editor, (newEditor) => {
  if (newEditor) {
    newEditor.on('selectionUpdate', () => {
      const selection = window.getSelection()
      hasSelection.value = selection ? !selection.isCollapsed : false
    })
  }
})

// Gérer le focus/blur de l'éditeur
const handleBlockClick = (event: MouseEvent) => {
  if (event.target === blockRef.value || blockRef.value?.contains(event.target as Node)) {
    editor.value?.chain().focus().run()
    isEditing.value = true
  }
}

const emit = defineEmits<{
  'create-block': [afterId: string, content: string]
}>()
</script>

<template>
  <div 
    ref="blockRef"
    class="min-h-[2em] relative group cursor-text"
    @click="handleBlockClick"
  >
    <div class="relative">
      <!-- Toolbar - visible uniquement lors d'une sélection -->
      <div 
        v-show="hasSelection"
        class="absolute -top-10 left-0 flex items-center gap-1 bg-white border rounded-md shadow-sm px-1 py-0.5 transition-all duration-200 pointer-events-auto z-50"
      >
        <!-- Headings -->
        <button
          v-for="level in [1, 2, 3]"
          :key="level"
          class="p-1 rounded hover:bg-gray-100"
          :class="{ 'bg-gray-100': editor?.isActive('heading', { level }) }"
          @click="editor?.chain().focus().toggleHeading({ level: level as Level }).run()"
        >
          <component 
            :is="level === 1 ? Heading1Icon : level === 2 ? Heading2Icon : Heading3Icon"
            class="w-4 h-4"
          />
        </button>

        <div class="w-px h-4 bg-gray-200" />

        <!-- Text formatting -->
        <button
          class="p-1 rounded hover:bg-gray-100"
          :class="{ 'bg-gray-100': editor?.isActive('bold') }"
          @click="editor?.chain().focus().toggleBold().run()"
        >
          <BoldIcon class="w-4 h-4" />
        </button>

        <button
          class="p-1 rounded hover:bg-gray-100"
          :class="{ 'bg-gray-100': editor?.isActive('italic') }"
          @click="editor?.chain().focus().toggleItalic().run()"
        >
          <ItalicIcon class="w-4 h-4" />
        </button>

        <button
          class="p-1 rounded hover:bg-gray-100"
          :class="{ 'bg-gray-100': editor?.isActive('strike') }"
          @click="editor?.chain().focus().toggleStrike().run()"
        >
          <StrikethroughIcon class="w-4 h-4" />
        </button>

        <div class="w-px h-4 bg-gray-200" />

        <!-- Color picker -->
        <div class="relative">
          <button
            class="p-1 rounded hover:bg-gray-100"
            @click="showColorPicker = !showColorPicker; showHighlightPicker = false"
          >
            <PaletteIcon class="w-4 h-4" />
          </button>

          <div 
            v-if="showColorPicker"
            class="absolute top-full left-0 mt-1 py-1 bg-white border rounded-md shadow-lg w-40 z-50"
          >
            <button
              v-for="{ color, name } in colors"
              :key="color"
              class="flex items-center gap-2 w-full px-2 py-1 hover:bg-gray-100"
              @click="editor?.chain().focus().setColor(color).run(); showColorPicker = false"
            >
              <div 
                class="w-4 h-4 rounded-sm border border-gray-200"
                :style="{ backgroundColor: color }"
              />
              <span class="text-sm">{{ name }}</span>
            </button>
          </div>
        </div>

        <!-- Highlight picker -->
        <div class="relative">
          <button
            class="p-1 rounded hover:bg-gray-100"
            @click="showHighlightPicker = !showHighlightPicker; showColorPicker = false"
          >
            <HighlighterIcon class="w-4 h-4" />
          </button>

          <div 
            v-if="showHighlightPicker"
            class="absolute top-full left-0 mt-1 py-1 bg-white border rounded-md shadow-lg w-40 z-50"
          >
            <button
              v-for="{ color, name } in highlights"
              :key="color"
              class="flex items-center gap-2 w-full px-2 py-1 hover:bg-gray-100"
              @click="editor?.chain().focus().toggleHighlight({ color }).run(); showHighlightPicker = false"
            >
              <div 
                class="w-4 h-4 rounded-sm border border-gray-200"
                :style="{ backgroundColor: color }"
              />
              <span class="text-sm">{{ name }}</span>
            </button>
          </div>
        </div>

        <div class="w-px h-4 bg-gray-200" />

        <button
          class="p-1 rounded hover:bg-gray-100"
          :class="{ 'bg-gray-100': editor?.isActive('code') }"
          @click="editor?.chain().focus().toggleCode().run()"
        >
          <Code2Icon class="w-4 h-4" />
        </button>
      </div>

      <EditorContent 
        :editor="editor" 
        class="editor-content focus:outline-none rounded-lg transition-colors"
      />
    </div>
  </div>
</template>

<style>
/* Styles spécifiques pour l'éditeur */
.editor-content {
  font-size: 0.875rem;
  line-height: 1.5;
}

.editor-content h1,
.editor-content h2,
.editor-content h3 {
  font-weight: 600;
  margin-top: 0.5rem;
  margin-bottom: 0.5rem;
}

.editor-content h1 {
  font-size: 1.5rem;
}

.editor-content h2 {
  font-size: 1.25rem;
}

.editor-content h3 {
  font-size: 1.125rem;
}

.editor-content p {
  margin-bottom: 0.5rem;
}

.editor-content code {
  background-color: #f3f4f6;
  padding: 0.125rem 0.25rem;
  border-radius: 0.25rem;
  font-family: ui-monospace, monospace;
}

.ProseMirror {
  padding: 0.75rem 1rem;
  min-height: 2em;
  cursor: text;
}

.ProseMirror:hover {
  background-color: rgb(249 250 251);
}

.ProseMirror:focus {
  background-color: rgb(249 250 251);
  outline: none;
}

.ProseMirror p.is-editor-empty:first-child::before {
  color: #adb5bd;
  content: attr(data-placeholder);
  float: left;
  height: 0;
  pointer-events: none;
}
</style> 