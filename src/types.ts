import type { Level } from "@tiptap/extension-heading";

type TiptapExtensionHeadingOptions = {
  level: Level[];
};

type TiptapExtensionDropCursorOptions = {
  color: string;
};

export type TiptapExtensionOptions = {
  heading: TiptapExtensionHeadingOptions;
  dropCursor: TiptapExtensionDropCursorOptions;
};

export type Table = {
  columns: number;
  rows: number;
  withHeader: boolean;
};

export type Image = {
  name: string;
  url: string;
};
