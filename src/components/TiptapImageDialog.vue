<template>
  <Dialog title="Pilih Gambar" :show="show" @close="closeDialog">
    <div
      class="rounded-lg border border-red-300 bg-red-100 p-5 text-sm text-gray-700"
    >
      Upload feature on this site has been disabled
    </div>
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

const imageListRef = ref<ImageData[]>([
  {
    Name: "System76-Fractal_Mountains-by_Kate_Hazen_of_System76.png",
    Url: "/uploaded/System76-Fractal_Mountains-by_Kate_Hazen_of_System76.png",
  },
  {
    Name: "System76-Geometric-adapted_by_Kate_Hazen_of_System76.png",
    Url: "/uploaded/System76-Geometric-adapted_by_Kate_Hazen_of_System76.png",
  },
  {
    Name: "System76-Robot-by_Kate_Hazen_of_System76.png",
    Url: "/uploaded/System76-Robot-by_Kate_Hazen_of_System76.png",
  },
])

function closeDialog() {
  emit("close")
}

function onDropImage(acceptedFiles: any[]) {
  if (acceptedFiles.length === 0) {
    return
  }

  // Removed
}

function loadData() {
  // Removed
}

function insertImage(url: string) {
  emit("insert", url)
  closeDialog()
}

onMounted(() => {
  loadData()
})
</script>
