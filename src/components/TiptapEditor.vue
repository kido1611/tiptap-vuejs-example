<script setup lang="ts">
import { onBeforeUnmount, ref, useTemplateRef } from "vue";
import { EditorContent, useEditor } from "@tiptap/vue-3";
import { tiptapExtensions } from "../utils/tiptap";

import IconArrowBackUp from "~icons/tabler/arrow-back-up";
import IconArrowForwardUp from "~icons/tabler/arrow-forward-up";
import IconBlockquote from "~icons/tabler/blockquote";
import IconH1 from "~icons/tabler/h1";
import IconH2 from "~icons/tabler/h2";
import IconH3 from "~icons/tabler/h3";
import IconH4 from "~icons/tabler/h4";
import IconH5 from "~icons/tabler/h5";
import IconH6 from "~icons/tabler/h6";
import IconBold from "~icons/tabler/bold";
import IconItalic from "~icons/tabler/italic";
import IconStrikethrough from "~icons/tabler/strikethrough";
import IconUnderline from "~icons/tabler/underline";
import IconLink from "~icons/tabler/link";
import IconListDetails from "~icons/tabler/list-details";
import IconListNumbers from "~icons/tabler/list-numbers";
import IconMovie from "~icons/tabler/movie";
import IconPhoto from "~icons/tabler/photo";
import IconMinus from "~icons/tabler/minus";
import IconTable from "~icons/mdi/table";
import IconTableRemove from "~icons/mdi/table-remove";
import IconTableColumnPlusBefore from "~icons/mdi/table-column-plus-before";
import IconTableColumnPlusAfter from "~icons/mdi/table-column-plus-after";
import IconTableColumnRemove from "~icons/mdi/table-column-remove";
import IconTableRowPlusAfter from "~icons/mdi/table-row-plus-after";
import IconTableRowPlusBefore from "~icons/mdi/table-row-plus-before";
import IconTableRowRemove from "~icons/mdi/table-row-remove";
import IconTableMergeCells from "~icons/mdi/table-merge-cells";

import TiptapToolbarButton from "./TiptapToolbarButton.vue";
import TiptapToolbarGroup from "./TiptapToolbarGroup.vue";

import TiptapLinkDialog from "./TiptapLinkDialog.vue";
import TiptapVideoDialog from "./TiptapVideoDialog.vue";
import TiptapTableDialog from "./TiptapTableDialog.vue";
import TiptapImageDialog from "./TiptapImageDialog.vue";

import type { Table as DataTable } from "./../types";

const model = defineModel<string>({
  required: true,
  default: "",
});

const editorInstance = useEditor({
  content: model.value,
  editorProps: {
    attributes: {
      class: "tiptap-prose",
    },
  },
  extensions: tiptapExtensions(),
  onUpdate: ({ editor }) => {
    model.value = editor.getHTML();
  },
});

const {
  isOpen: isLinkDialogVisible,
  open: openLinkDialog,
  updateLink,
} = useTiptapLink();

const {
  isOpen: isVideoDialogVisible,
  open: openVideoDialog,
  insertVideo,
} = useTiptapVideo();

const {
  isOpen: isTableDialogVisible,
  open: openTableDialog,
  insertTable,
} = useTiptapTable();

const {
  isOpen: isImageDialogVisible,
  open: openImageDialog,
  insertImage,
} = useTiptapImage();

onBeforeUnmount(() => {
  editorInstance.value?.destroy();
});

function useTiptapLink() {
  const isOpen = ref<boolean>(false);

  type LinkDialog = InstanceType<typeof TiptapLinkDialog>;
  const dialog = useTemplateRef<LinkDialog>("link-dialog");

  function open() {
    isOpen.value = true;
    dialog.value?.setLink(
      editorInstance.value?.getAttributes("link").href ?? "",
    );
  }

  function updateLink(value?: string) {
    if (!value) {
      editorInstance.value
        ?.chain()
        .focus()
        .extendMarkRange("link")
        .unsetLink()
        .run();
      return;
    }

    editorInstance.value
      ?.chain()
      .focus()
      .extendMarkRange("link")
      .setLink({ href: value })
      .run();
  }

  return {
    isOpen,
    open,
    updateLink,
  };
}

function useTiptapVideo() {
  const isOpen = ref<boolean>(false);

  function open() {
    isOpen.value = true;
  }

  function insertVideo(url: string) {
    editorInstance.value?.commands.setYoutubeVideo({
      src: url,
      width: 400,
      height: 300,
    });
  }

  return {
    isOpen,
    open,
    insertVideo,
  };
}

function useTiptapTable() {
  const isOpen = ref<boolean>(false);

  function open() {
    isOpen.value = true;
  }

  function insertTable(table: DataTable) {
    editorInstance.value
      ?.chain()
      .focus()
      .insertTable({
        rows: table.rows,
        cols: table.columns,
        withHeaderRow: table.withHeader,
      })
      .run();
  }

  return {
    isOpen,
    open,
    insertTable,
  };
}

function useTiptapImage() {
  const isOpen = ref<boolean>(false);

  function open() {
    isOpen.value = true;
  }
  function insertImage(url: string) {
    editorInstance.value?.chain().focus().setImage({ src: url }).run();
  }

  return {
    isOpen,
    open,
    insertImage,
  };
}
</script>

<template>
  <div
    id="tiptap"
    class="divide-y divide-gray-400 rounded-md border border-gray-400 overflow-clip"
  >
    <div
      id="tiptap-toolbar"
      class="divide-x divide-gray-400 sticky top-0 inset-x-0 bg-white z-1"
    >
      <TiptapToolbarGroup>
        <TiptapToolbarButton
          label="Undo"
          @click="editorInstance?.chain().focus().undo().run()"
          :disabled="!editorInstance?.can().chain().focus().undo().run()"
        >
          <IconArrowBackUp class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          label="Redo"
          @click="editorInstance?.chain().focus().redo().run()"
          :disabled="!editorInstance?.can().chain().focus().redo().run()"
        >
          <IconArrowForwardUp class="h-5 w-5" />
        </TiptapToolbarButton>
      </TiptapToolbarGroup>
      <TiptapToolbarGroup>
        <TiptapToolbarButton
          label="Heading 1"
          :is-active="editorInstance?.isActive('heading', { level: 1 })"
          @click="
            editorInstance?.chain().focus().toggleHeading({ level: 1 }).run()
          "
        >
          <IconH1 class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          label="Heading 2"
          :is-active="editorInstance?.isActive('heading', { level: 2 })"
          @click="
            editorInstance?.chain().focus().toggleHeading({ level: 2 }).run()
          "
        >
          <IconH2 class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          label="Heading 3"
          :is-active="editorInstance?.isActive('heading', { level: 3 })"
          @click="
            editorInstance?.chain().focus().toggleHeading({ level: 3 }).run()
          "
        >
          <IconH3 class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          label="Heading 4"
          :is-active="editorInstance?.isActive('heading', { level: 4 })"
          @click="
            editorInstance?.chain().focus().toggleHeading({ level: 4 }).run()
          "
        >
          <IconH4 class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          label="Heading 5"
          :is-active="editorInstance?.isActive('heading', { level: 5 })"
          @click="
            editorInstance?.chain().focus().toggleHeading({ level: 5 }).run()
          "
        >
          <IconH5 class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          label="Heading 6"
          :is-active="editorInstance?.isActive('heading', { level: 6 })"
          @click="
            editorInstance?.chain().focus().toggleHeading({ level: 6 }).run()
          "
        >
          <IconH6 class="h-5 w-5" />
        </TiptapToolbarButton>
      </TiptapToolbarGroup>
      <TiptapToolbarGroup>
        <TiptapToolbarButton
          label="Bold"
          :is-active="editorInstance?.isActive('bold')"
          @click="editorInstance?.chain().focus().toggleBold().run()"
        >
          <IconBold class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          label="Italic"
          :is-active="editorInstance?.isActive('italic')"
          @click="editorInstance?.chain().focus().toggleItalic().run()"
        >
          <IconItalic class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          label="Underline"
          :is-active="editorInstance?.isActive('underline')"
          @click="editorInstance?.chain().focus().toggleUnderline().run()"
        >
          <IconUnderline class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          label="Strikethrough"
          :is-active="editorInstance?.isActive('strike')"
          @click="editorInstance?.chain().focus().toggleStrike().run()"
        >
          <IconStrikethrough class="h-5 w-5" />
        </TiptapToolbarButton>
      </TiptapToolbarGroup>
      <TiptapToolbarGroup>
        <TiptapToolbarButton
          label="Unordered List"
          :is-active="editorInstance?.isActive('bulletList')"
          @click="editorInstance?.chain().focus().toggleBulletList().run()"
        >
          <IconListDetails class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          label="Ordered List"
          :is-active="editorInstance?.isActive('orderedList')"
          @click="editorInstance?.chain().focus().toggleOrderedList().run()"
        >
          <IconListNumbers class="h-5 w-5" />
        </TiptapToolbarButton>
      </TiptapToolbarGroup>
      <TiptapToolbarGroup>
        <TiptapToolbarButton
          label="Link"
          @click="openLinkDialog"
          :is-active="editorInstance?.isActive('link')"
        >
          <IconLink class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton label="Image" @click="openImageDialog">
          <IconPhoto class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          label="Blockquote"
          :is-active="editorInstance?.isActive('blockquote')"
          @click="editorInstance?.chain().focus().toggleBlockquote().run()"
        >
          <IconBlockquote class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton label="Table" @click="openTableDialog">
          <IconTable class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton label="Youtube" @click="openVideoDialog">
          <IconMovie class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          @click="editorInstance?.chain().focus().setHorizontalRule().run()"
          label="Horizontal Line"
        >
          <IconMinus class="h-5 w-5" />
        </TiptapToolbarButton>
      </TiptapToolbarGroup>
      <TiptapToolbarGroup v-if="editorInstance?.isActive('table')">
        <TiptapToolbarButton
          @click="editorInstance?.commands.deleteTable()"
          label="Remove table"
        >
          <IconTableRemove class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          label="Add column before"
          @click="editorInstance?.commands.addColumnBefore()"
        >
          <IconTableColumnPlusBefore class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          label="Add column after"
          @click="editorInstance?.commands.addColumnAfter()"
        >
          <IconTableColumnPlusAfter class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          label="Remove column"
          @click="editorInstance?.commands.deleteColumn()"
        >
          <IconTableColumnRemove class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          label="Add row before"
          @click="editorInstance?.commands.addRowBefore()"
        >
          <IconTableRowPlusBefore class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          @click="editorInstance?.commands.addRowAfter()"
          label="Add row after"
        >
          <IconTableRowPlusAfter class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          label="Remove row"
          @click="editorInstance?.commands.deleteRow()"
        >
          <IconTableRowRemove class="h-5 w-5" />
        </TiptapToolbarButton>
        <TiptapToolbarButton
          label="Merge or split cell"
          @click="editorInstance?.commands.mergeOrSplit()"
        >
          <IconTableMergeCells class="h-5 w-5" />
        </TiptapToolbarButton>
      </TiptapToolbarGroup>
    </div>

    <div class="flex flex-col">
      <div class="max-h-[60vh] overflow-y-auto">
        <EditorContent :editor="editorInstance" />
      </div>

      <div
        class="mx-4 border-t border-gray-300 py-3 text-right text-sm text-gray-500"
      >
        {{ editorInstance?.storage.characterCount.characters() }} characters,
        {{ editorInstance?.storage.characterCount.words() }} words
      </div>
    </div>

    <TiptapLinkDialog
      v-model:open="isLinkDialogVisible"
      ref="link-dialog"
      @update="updateLink"
    />
    <TiptapVideoDialog
      v-model:open="isVideoDialogVisible"
      @insert="insertVideo"
    />
    <TiptapTableDialog
      v-model:open="isTableDialogVisible"
      @insert="insertTable"
    />
    <TiptapImageDialog
      v-model:open="isImageDialogVisible"
      @insert="insertImage"
    />
  </div>
</template>
