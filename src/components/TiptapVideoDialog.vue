<template>
  <Dialog title="Tambah Video Youtube" v-model:open="open">
    <form @submit.prevent="onSubmit">
      <div class="flex flex-col space-y-5 px-5 py-4">
        <InputContainer>
          <Label for="input-add-youtube-url">Tautan Youtube</Label>
          <Input
            type="url"
            id="input-add-youtube-url"
            v-model="inputYoutubeUrlRef"
            required
          />
        </InputContainer>
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

<script setup lang="ts">
import { ref } from "vue";
import Dialog from "./TheDialog.vue";
import InputContainer from "./InputContainer.vue";
import Input from "./Input.vue";
import Label from "./Label.vue";

const emit = defineEmits<{
  insert: [value: string];
}>();

const open = defineModel<boolean>("open", {
  default: false,
});

const inputYoutubeUrlRef = ref<string>("");

function onSubmit() {
  emit("insert", inputYoutubeUrlRef.value);
  inputYoutubeUrlRef.value = "";

  open.value = false;
}
</script>
