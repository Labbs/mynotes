import { VueRenderer } from "@tiptap/vue-3";
import { Editor, type Range } from "@tiptap/core";
import CommandsList from "../components/tiptap/CommandsList.vue";
import {
  CodeIcon,
  Heading1Icon,
  Heading2Icon,
  Heading3Icon,
  ListIcon,
  ListOrderedIcon,
  QuoteIcon,
  TextIcon,
  ImageIcon,
  MinusIcon
} from "lucide-vue-next";

interface SuggestionItem {
  name: string;
  description: string;
  icon: any;
  command: (props: { editor: Editor; range: Range }) => void;
}

interface SuggestionProps {
  query: string;
  editor: Editor;
  range: Range;
  clientRect: () => DOMRect;
}

const items: SuggestionItem[] = [
  {
    name: "Text",
    description: "Just start writing with plain text.",
    icon: TextIcon,
    command: ({ editor, range }) => {
      editor.chain().focus().deleteRange(range).setParagraph().run();
    },
  },
  {
    name: "Heading 1",
    description: "Big section heading.",
    icon: Heading1Icon,
    command: ({ editor, range }) => {
      editor
        .chain()
        .focus()
        .deleteRange(range)
        .setNode("heading", { level: 1 })
        .run();
    },
  },
  {
    name: "Heading 2",
    description: "Medium section heading.",
    icon: Heading2Icon,
    command: ({ editor, range }) => {
      editor
        .chain()
        .focus()
        .deleteRange(range)
        .setNode("heading", { level: 2 })
        .run();
    },
  },
  {
    name: "Heading 3",
    description: "Small section heading.",
    icon: Heading3Icon,
    command: ({ editor, range }) => {
      editor
        .chain()
        .focus()
        .deleteRange(range)
        .setNode("heading", { level: 3 })
        .run();
    },
  },
  {
    name: "Bullet List",
    description: "Create a simple bullet list.",
    icon: ListIcon,
    command: ({ editor, range }) => {
      editor.chain().focus().deleteRange(range).toggleBulletList().run();
    },
  },
  {
    name: "Numbered List",
    description: "Create a list with numbering.",
    icon: ListOrderedIcon,
    command: ({ editor, range }) => {
      editor.chain().focus().deleteRange(range).toggleOrderedList().run();
    },
  },
  {
    name: "Quote",
    description: "Capture a quote.",
    icon: QuoteIcon,
    command: ({ editor, range }) => {
      editor.chain().focus().deleteRange(range).toggleBlockquote().run();
    },
  },
  {
    name: "Code Block",
    description: "Add a code block with syntax highlighting",
    icon: CodeIcon,
    command: ({ editor, range }) => {
      editor.chain().focus().deleteRange(range).setNode("codeBlock").run();
    },
  },
  {
    name: "Image",
    description: "Edit an image",
    icon: ImageIcon,
    command: ({ editor, range }) => {
      editor
        .chain()
        .focus()
        .deleteRange(range)
        .insertContent({
          type: "imageEditor",
          attrs: {},
        })
        .run();
    },
  },
  {
    name: "Divider",
    description: "Insert a horizontal divider.",
    icon: MinusIcon,
    command: ({ editor, range }) => {
      editor.chain().focus().deleteRange(range).setHorizontalRule().run();
    },
  },
];

export const suggestion = {
  items: ({ query }: { query: string }): SuggestionItem[] => {
    return items.filter((item) =>
      item.name.toLowerCase().startsWith(query.toLowerCase())
    );
  },

  render: () => {
    let component: VueRenderer;
    let popupElement: HTMLElement | null = null;

    return {
      onStart: (props: SuggestionProps) => {
        component = new VueRenderer(CommandsList, {
          props,
          editor: props.editor,
        });

        if (!props.clientRect) {
          return;
        }

        // Create our own popup element
        popupElement = document.createElement('div');
        popupElement.className = 'suggestion-popup';
        
        // Style the popup to match tippy's behavior
        Object.assign(popupElement.style, {
          position: 'absolute',
          zIndex: '9999',
          backgroundColor: 'white',
          borderRadius: '4px',
          boxShadow: '0 0 0 1px rgba(0,0,0,.05), 0 4px 8px rgba(0,0,0,.1)',
          maxWidth: '400px',
          maxHeight: '400px',
          overflow: 'auto'
        });
        
        // Add the component element to our popup
        if (component.element) {
          popupElement.appendChild(component.element);
        }
        document.body.appendChild(popupElement);
        
        // Position the popup using the clientRect
        const rect = props.clientRect();
        updatePopupPosition(popupElement, rect);
      },

      onUpdate(props: SuggestionProps) {
        component?.updateProps(props);

        if (!props.clientRect || !popupElement) {
          return;
        }

        // Update the position of our popup
        const rect = props.clientRect();
        updatePopupPosition(popupElement, rect);
      },

      onKeyDown(props: { event: KeyboardEvent }) {
        if (props.event.key === "Escape") {
          if (popupElement && popupElement.parentElement) {
            popupElement.parentElement.removeChild(popupElement);
            popupElement = null;
          }
          return true;
        }

        return component?.ref?.onKeyDown(props.event);
      },

      onExit() {
        if (popupElement && popupElement.parentElement) {
          popupElement.parentElement.removeChild(popupElement);
          popupElement = null;
        }
        
        if (component) {
          component.destroy();
        }
      },
    };
  },
};

// Helper function to position the popup
function updatePopupPosition(element: HTMLElement, rect: DOMRect) {
  // Position below the cursor
  element.style.top = `${rect.bottom + window.scrollY}px`;
  element.style.left = `${rect.left + window.scrollX}px`;
  
  // Ensure the popup doesn't go out of the viewport
  const viewportWidth = window.innerWidth;
  const viewportHeight = window.innerHeight;
  const popupRect = element.getBoundingClientRect();
  
  if (rect.left + popupRect.width > viewportWidth) {
    element.style.left = `${viewportWidth - popupRect.width - 10}px`;
  }
  
  if (rect.bottom + popupRect.height > viewportHeight) {
    // Show above if not enough space below
    element.style.top = `${rect.top - popupRect.height + window.scrollY}px`;
  }
}

export default suggestion;
