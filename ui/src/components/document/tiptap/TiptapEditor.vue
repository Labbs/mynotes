<template>
  <div class="relative">
    <editor-content v-if="editor" :editor="editor" class="prose max-w-none" />
    <div 
      v-if="lock" 
      class="absolute inset-0 bg-gray-50 bg-opacity-20 z-10"
      style="pointer-events: none;"
    ></div>
    <BubbleMenu :editor="editor" v-if="editor" />
    <div v-if="!editor" class="min-h-[300px] border rounded p-4 flex items-center justify-center text-gray-400">
      Loading editor...
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, nextTick, onUnmounted } from 'vue'
import { useEditor, EditorContent } from '@tiptap/vue-3'
import StarterKit from '@tiptap/starter-kit'
import Placeholder from '@tiptap/extension-placeholder'
import { Color } from "@tiptap/extension-color"
import TextStyle from "@tiptap/extension-text-style"
import Image from "@tiptap/extension-image"
import CommandMenuExtension from '../../../extensions/Commands'
import suggestion from '../../../extensions/Suggestion'
import Highlight from '@tiptap/extension-highlight'
import BubbleMenu from '../../../components/tiptap/BubbleMenu.vue'
import type { Document } from '../../../api/document'

const props = defineProps<{
  document: Document
  lock: boolean
  onSave: (content: string) => void
}>()

const isEditorUpdating = ref(false)
const editor = useEditor({
  content: props.document.content || '<p></p>',
  editable: !props.lock,
  extensions: [
    StarterKit.configure({ 
      heading: { levels: [1, 2, 3] },
      // Désactiver le codeBlock de StarterKit pour éviter les doublons
      codeBlock: false 
    }),
    Placeholder.configure({ placeholder: "Type '/' for commands" }),
    Color,
    TextStyle,
    Image,
    CommandMenuExtension.configure({ suggestion }),
    Highlight.configure({ multicolor: true }),
  ],
  onUpdate: ({ editor }) => {
    if (!editor.isEditable || isEditorUpdating.value) return
    const content = editor.getHTML()
    props.onSave(content)
  },
})

watch(() => props.lock, (isLocked) => {
  if (editor.value) editor.value.setEditable(!isLocked)
})

watch(() => props.document.content, (newContent) => {
  if (editor.value && newContent && !isEditorUpdating.value) {
    isEditorUpdating.value = true
    nextTick(() => {
      try {
        editor.value?.commands.setContent(newContent)
      } finally {
        setTimeout(() => { isEditorUpdating.value = false }, 50)
      }
    })
  }
})

onUnmounted(() => {
  editor.value?.destroy()
})
</script>
