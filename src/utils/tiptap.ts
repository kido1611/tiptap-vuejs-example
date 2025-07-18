import Paragraph from "@tiptap/extension-paragraph";
import Document from "@tiptap/extension-document";
import Text from "@tiptap/extension-text";
import { BulletList, OrderedList, ListItem } from "@tiptap/extension-list";
import {
  UndoRedo,
  Dropcursor,
  Gapcursor,
  CharacterCount,
} from "@tiptap/extensions";
import Heading from "@tiptap/extension-heading";
import Bold from "@tiptap/extension-bold";
import Italic from "@tiptap/extension-italic";
import Underline from "@tiptap/extension-underline";
import Strike from "@tiptap/extension-strike";
import Link from "@tiptap/extension-link";
import { Blockquote } from "@tiptap/extension-blockquote";
import { HardBreak } from "@tiptap/extension-hard-break";
import { Youtube } from "@tiptap/extension-youtube";
import { HorizontalRule } from "@tiptap/extension-horizontal-rule";
import {
  Table,
  TableRow,
  TableCell,
  TableHeader,
} from "@tiptap/extension-table";
import Image from "@tiptap/extension-image";
import Code from "@tiptap/extension-code";
import CodeBlock from "@tiptap/extension-code-block";
import type { TiptapExtensionOptions } from "../types";

export const tiptapExtensions = (options?: Partial<TiptapExtensionOptions>) => {
  const tiptapExtension: TiptapExtensionOptions = {
    heading: {
      level: [1, 2, 3, 4, 5, 6],
    },
    dropCursor: {
      color: "#2563eb",
    },

    ...options,
  };

  return [
    Paragraph,
    Document,
    Text,
    UndoRedo,
    Heading.configure({
      levels: tiptapExtension.heading.level,
    }),
    Bold,
    Italic,
    Underline,
    Strike,
    ListItem,
    BulletList,
    OrderedList,
    Link.configure({
      openOnClick: false,
    }),
    HardBreak,
    Blockquote,
    CharacterCount,
    Youtube,
    Dropcursor.configure({
      width: 2,
      color: tiptapExtension.dropCursor.color,
    }),
    HorizontalRule,
    Table.configure({
      resizable: false,
      allowTableNodeSelection: true,
    }),
    TableRow,
    TableHeader,
    TableCell,
    Gapcursor,
    Image,
    Code,
    CodeBlock,
  ];
};
