<script setup lang="ts">
import {
  DialogRoot,
  DialogContent,
  DialogOverlay,
  DialogPortal,
  DialogTitle,
  DialogTrigger,
  DialogClose,
} from "reka-ui";
import IconX from "~icons/tabler/x";

const { title } = defineProps<{
  title: string;
}>();

const open = defineModel<boolean>("open", {
  default: false,
});
</script>

<template>
  <DialogRoot v-model:open="open">
    <DialogTrigger as-child>
      <slot name="trigger" />
    </DialogTrigger>

    <DialogPortal>
      <DialogOverlay class="bg-black/50 fixed inset-0 z-30" />

      <DialogContent
        class="bg-white fixed top-[50%] left-[50%] -translate-x-1/2 -translate-y-1/2 z-100 max-w-[450px] rounded-lg flex max-h-[calc(100dvh-2.5rem)] w-full flex-col shadow"
      >
        <div
          class="relative px-5 py-3 flex flex-row items-center gap-x-4 justify-between border-b border-neutral-300"
        >
          <DialogTitle class="text-base font-semibold">
            {{ title }}
          </DialogTitle>
          <DialogClose
            aria-label="Close dialog"
            class="size-9 flex items-center justify-center hover:bg-neutral-200 text-neutral-600 rounded-md -mr-2 -mt-0.5"
          >
            <IconX class="size-4" stroke-width="2" />
          </DialogClose>
        </div>
        <div class="overflow-y-auto h-full">
          <slot />
        </div>
      </DialogContent>
    </DialogPortal>
  </DialogRoot>
</template>
