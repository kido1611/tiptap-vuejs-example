<script setup lang="ts">
import { onMounted, ref, useTemplateRef } from "vue";
import { useDropZone } from "@vueuse/core";
import Dialog from "./TheDialog.vue";
import type { Image as ImageData } from "./../types";

const serverUrl = import.meta.env.VITE_IMAGE_SERVER_URL;
const isPreview = import.meta.env.VITE_DEPLOY_PREVIEW as boolean;

const emit = defineEmits<{
  insert: [url: string];
}>();

const open = defineModel("open", {
  default: false,
});

const dropZoneRef = useTemplateRef<HTMLElement>("dropZoneRef");
const { isOverDropZone } = useDropZone(dropZoneRef, {
  dataTypes: ["image/jpeg", "image/jpg", "image/png"],
  onDrop: onDropImage,
});

const imageListRef = ref<ImageData[]>([]);

async function onDropImage(files: File[] | null) {
  if (!files || files.length === 0) {
    return;
  }

  await Promise.all(
    files.map((file) => {
      const formData = new FormData();
      formData.append("file", file);

      return fetch(`${serverUrl}/files`, {
        method: "POST",
        body: formData,
      });
    }),
  );

  await loadData();
}

async function loadData() {
  const response = await fetch(`${serverUrl}/files`, {
    method: "GET",
  });
  imageListRef.value = await response.json();
}

function insertImage(url: string) {
  emit("insert", url);
  open.value = false;
}

onMounted(async () => {
  if (!isPreview) {
    await loadData();
  } else {
    imageListRef.value.push({
      name: "System76-Fractal_Mountains-by_Kate_Hazen_of_System76.png",
      url: "/examples/System76-Fractal_Mountains-by_Kate_Hazen_of_System76.png",
    });

    imageListRef.value.push({
      name: "System76-Geometric-adapted_by_Kate_Hazen_of_System76.png",
      url: "/examples/System76-Geometric-adapted_by_Kate_Hazen_of_System76.png",
    });

    imageListRef.value.push({
      name: "System76-Robot-by_Kate_Hazen_of_System76.png",
      url: "/examples/System76-Robot-by_Kate_Hazen_of_System76.png",
    });

    imageListRef.value.push({
      name: "System76-Unleash_Your_Robot_Blue-by_Kate_Hazen_of_System76.png",
      url: "/examples/System76-Unleash_Your_Robot_Blue-by_Kate_Hazen_of_System76.png",
    });

    imageListRef.value.push({
      name: "System76-Unleash_Your_Robot-by_Kate_Hazen_of_System76.png",
      url: "/examples/System76-Unleash_Your_Robot-by_Kate_Hazen_of_System76.png",
    });
  }
});
</script>

<template>
  <Dialog title="Pilih Gambar" v-model:open="open">
    <div class="px-5 py-4">
      <div
        v-if="imageListRef && imageListRef.length > 0"
        class="grid grid-cols-5 gap-2 mb-4"
      >
        <button
          @click="insertImage(image.url)"
          type="button"
          v-for="(image, key) in imageListRef"
          :key="image.name"
          class="rounded-md border border-gray-300 p-1 hover:bg-neutral-200"
        >
          <img
            :alt="`Photo ${key}`"
            :src="image.url"
            class="aspect-square object-contain object-center"
          />
        </button>
      </div>

      <div
        ref="dropZoneRef"
        v-if="!isPreview"
        class="h-32 flex items-center justify-center border border-dashed border-neutral-300 rounded-lg p-4"
      >
        <p v-if="!isOverDropZone" class="text-sm text-neutral-700">
          Drop images here
        </p>
        <p v-else class="text-sm text-neutral-700 font-semibold">Drop here</p>
      </div>

      <div
        v-else
        class="h-32 flex items-center justify-center border border-dashed border-red-300 rounded-lg p-4 bg-red-100"
      >
        <p class="text-sm text-red-700 font-medium">
          Upload image is disabled on this preview site
        </p>
      </div>
    </div>
  </Dialog>
</template>
