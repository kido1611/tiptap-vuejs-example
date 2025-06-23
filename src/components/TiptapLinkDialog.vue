<template>
  <Dialog v-model:open="open" title="Tautan">
    <form @submit.prevent="update">
      <div class="flex flex-col space-y-5 px-5 py-4">
        <InputContainer>
          <Label for="input-link-url">Tautan</Label>
          <Input type="url" id="input-link-url" v-model="inputLinkRef" />
        </InputContainer>

        <div class="flex flex-row justify-end space-x-3">
          <button
            type="submit"
            class="rounded-md bg-blue-700 px-4 py-3 text-sm font-medium text-white hover:bg-opacity-80"
          >
            Simpan
          </button>
        </div>
      </div>
    </form>
  </Dialog>
</template>

<script setup lang="ts">
import { ref } from "vue";
import Dialog from "./TheDialog.vue";
import Label from "./Label.vue";
import Input from "./Input.vue";
import InputContainer from "./InputContainer.vue";

const open = defineModel<boolean>("open", {
  default: false,
});

const emit = defineEmits<{
  update: [value: string];
}>();

defineExpose({
  setLink,
});
const inputLinkRef = ref<string>("");

function update() {
  open.value = false;
  emit("update", inputLinkRef.value);
}

function setLink(url: string) {
  inputLinkRef.value = url;
}
</script>
