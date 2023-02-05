<template>
  <Dialog title="Pilih Gambar" :show="show" @close="closeDialog">
    <div v-bind="getRootProps()" class="rounded-lg border border-gray-300 p-4">
      <input v-bind="getInputProps()" />
      <div
        v-if="imageListRef?.length > 0"
        class="grid grid-cols-3 gap-3 sm:grid-cols-4"
      >
        <button
          @click="insertImage(image.Url)"
          type="button"
          v-for="image in imageListRef"
          :key="image.Name"
          class="rounded-md border border-gray-300 p-1"
        >
          <img
            alt=""
            :src="image.Url"
            class="aspect-square object-scale-down object-center"
          />
        </button>
      </div>
      <div
        :class="[
          isDragActive ? 'bg-gray-100' : '',
          imageListRef?.length > 0 ? 'mt-4 ' : '',
        ]"
        class="rounded-lg border border-dashed border-gray-300 px-8 py-12"
      >
        <p
          v-if="isDragActive"
          class="text-center text-sm font-medium text-gray-700"
        >
          Drop the files here ...
        </p>
        <p v-else class="text-center text-sm font-medium text-gray-700">
          Drag 'n' drop some image file here, or click to select file
        </p>
      </div>
    </div>
  </Dialog>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue"
import { useDropzone } from "vue3-dropzone"
import Dialog from "./Dialog.vue"
import axios from "axios"
import type ImageData from "@/models/image"

defineProps<{
  show: boolean
}>()
const emit = defineEmits<{
  (e: "close"): void
  (e: "insert", url: string): void
}>()

const { getRootProps, getInputProps, isDragActive } = useDropzone({
  accept: "image/png,image/jpeg",
  multiple: false,
  onDrop: onDropImage,
  noClick: true,
})

const imageListRef = ref<ImageData[]>([])

function closeDialog() {
  emit("close")
}

function onDropImage(acceptedFiles: any[]) {
  if (acceptedFiles.length === 0) {
    return
  }

  const formData = new FormData()
  formData.append("file", acceptedFiles[0])

  axios
    .post("http://localhost:8080/files", formData, {
      headers: {
        "Content-type": "multipart/form-data",
      },
    })
    .then(() => {
      loadData()
    })
}

function loadData() {
  axios.get("http://localhost:8080/files").then((result) => {
    imageListRef.value = result.data
  })
}

function insertImage(url: string) {
  emit("insert", url)
  closeDialog()
}

onMounted(() => {
  loadData()
})
</script>
