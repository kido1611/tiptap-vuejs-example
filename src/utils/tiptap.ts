import Paragraph from "@tiptap/extension-paragraph";
import Document from "@tiptap/extension-document";
import Text from "@tiptap/extension-text";
import History from "@tiptap/extension-history";
import Heading from "@tiptap/extension-heading";
import Bold from "@tiptap/extension-bold";
import Italic from "@tiptap/extension-italic";
import Underline from "@tiptap/extension-underline";
import Strike from "@tiptap/extension-strike";
import ListItem from "@tiptap/extension-list-item";
import BulletList from "@tiptap/extension-bullet-list";
import OrderedList from "@tiptap/extension-ordered-list";
import Link from "@tiptap/extension-link";
import { Blockquote } from "@tiptap/extension-blockquote";
import { HardBreak } from "@tiptap/extension-hard-break";
import { CharacterCount } from "@tiptap/extension-character-count";
import { Youtube } from "@tiptap/extension-youtube";
import Dropcursor from "@tiptap/extension-dropcursor";
import { HorizontalRule } from "@tiptap/extension-horizontal-rule";
import { Table } from "@tiptap/extension-table";
import { TableHeader } from "@tiptap/extension-table-header";
import { TableRow } from "@tiptap/extension-table-row";
import { TableCell } from "@tiptap/extension-table-cell";
import Gapcursor from "@tiptap/extension-gapcursor";
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
    History,
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
