<script setup lang="ts">
import { reactive } from "vue";
import Dialog from "./TheDialog.vue";
import type { Table } from "./../types";
import InputContainer from "./InputContainer.vue";
import Label from "./Label.vue";
import Input from "./Input.vue";
import { SwitchRoot, SwitchThumb } from "reka-ui";

const open = defineModel("open", {
  default: false,
});

const emit = defineEmits<{
  insert: [table: Table];
}>();

const state = reactive({
  columns: 3,
  rows: 3,
  withHeader: true,
});

function onSubmit() {
  emit("insert", state);
  open.value = false;

  // Reset form
  state.columns = 3;
  state.rows = 3;
  state.withHeader = true;
}
</script>

<template>
  <Dialog title="Tambah Tabel" v-model:open="open">
    <form @submit.prevent="onSubmit">
      <div class="flex flex-col space-y-5 px-5 py-4">
        <div class="flex flex-row space-x-5">
          <InputContainer class="w-full flex-1">
            <Label for="input-table-columns">Kolom</Label>
            <Input
              v-model="state.columns"
              id="input-table-columns"
              required
              type="number"
              min="1"
              class="w-full"
            />
          </InputContainer>
          <InputContainer class="w-full flex-1">
            <Label for="input-table-rows">Baris</Label>
            <Input
              v-model="state.rows"
              id="input-table-rows"
              required
              type="number"
              min="1"
              class="w-full"
            />
          </InputContainer>
        </div>
        <div class="flex flex-row items-center space-x-3">
          <SwitchRoot
            id="switch-header"
            v-model="state.withHeader"
            class="relative inline-flex h-6 w-11 shrink-0 items-center rounded-full transition-colors focus:outline-hidden focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 data-[state=unchecked]:bg-gray-200 data-[state=checked]:bg-blue-600"
          >
            <SwitchThumb
              class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform translate-x-1 data-[state=checked]:translate-x-6"
            />
          </SwitchRoot>
          <label for="switch-header" class="select-none text-sm text-gray-600">
            Tabel Header
          </label>
        </div>
        <div class="flex flex-row justify-end space-x-3">
          <button
            type="submit"
            class="rounded-md bg-blue-700 px-4 py-3 text-sm font-medium text-white hover:bg-opacity-80"
          >
            Tambah
          </button>
        </div>
      </div>
    </form>
  </Dialog>
</template>
