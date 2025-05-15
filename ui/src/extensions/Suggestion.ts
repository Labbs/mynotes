import { VueRenderer } from "@tiptap/vue-3";
import tippy, { type Instance as TippyInstance } from "tippy.js";
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
    let popup: TippyInstance[];

    return {
      onStart: (props: SuggestionProps) => {
        component = new VueRenderer(CommandsList, {
          props,
          editor: props.editor,
        });

        if (!props.clientRect) {
          return;
        }

        popup = tippy("body", {
          getReferenceClientRect: props.clientRect,
          appendTo: () => document.body,
          content: component.element,
          showOnCreate: true,
          interactive: true,
          trigger: "manual",
          placement: "bottom-start",
        });
      },

      onUpdate(props: SuggestionProps) {
        component.updateProps(props);

        if (!props.clientRect) {
          return;
        }

        popup[0].setProps({
          getReferenceClientRect: props.clientRect,
        });
      },

      onKeyDown(props: { event: KeyboardEvent }) {
        if (props.event.key === "Escape") {
          popup[0].hide();
          return true;
        }

        return component.ref?.onKeyDown(props.event);
      },

      onExit() {
        popup[0].destroy();
        component.destroy();
      },
    };
  },
};

export default suggestion;
