<script setup lang="ts">
import { ref } from "vue";
import { SwitchRoot, SwitchThumb } from "reka-ui";

import TiptapEditor from "./components/TiptapEditor.vue";
import example from "./assets/example.txt?raw";

const content = ref<string>(example);

const showRaw = ref<boolean>(true);
</script>

<template>
  <div class="flex flex-col space-y-6 max-w-screen-xl mx-auto py-8">
    <header>
      <h1 class="text-2xl font-bold">Tiptap Editor</h1>
    </header>
    <div class="grid grid-cols-1 gap-10 xl:grid-cols-2">
      <TiptapEditor v-model="content" />

      <div>
        <div class="flex flex-row items-center justify-between gap-x-4">
          <h2 class="text-2xl font-semibold mb-3">Output</h2>

          <div class="flex flex-row items-center space-x-3">
            <SwitchRoot
              id="switch-header"
              v-model="showRaw"
              class="relative inline-flex h-6 w-11 shrink-0 items-center rounded-full transition-colors focus:outline-hidden focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 data-[state=unchecked]:bg-gray-200 data-[state=checked]:bg-blue-600"
            >
              <SwitchThumb
                class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform translate-x-1 data-[state=checked]:translate-x-6"
              />
            </SwitchRoot>
            <label
              for="switch-header"
              class="select-none text-sm text-gray-600"
            >
              Raw output
            </label>
          </div>
        </div>
        <div
          class="text-sm text-gray-700 max-h-[70vh] overflow-y-auto border border-neutral-300 p-4 rounded-md"
        >
          <div v-if="showRaw" class="break-all">
            {{ content }}
          </div>

          <div v-else class="tiptap-prose" v-html="content"></div>
        </div>
      </div>
    </div>
  </div>
</template>
