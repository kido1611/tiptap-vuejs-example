<template>
  <input
    v-model="proxyModel"
    class="rounded-md border border-gray-400 px-3 py-2 text-gray-800 focus:ring-blue-800"
  />
</template>

<script setup lang="ts">
import { computed } from "vue"
const props = defineProps<{
  modelValue: string | number | undefined
}>()

const emit = defineEmits<{
  (e: "update:modelValue", value: string | number): void
}>()

const proxyModel = computed({
  get: () => props.modelValue ?? "",
  set: (value: string | number): void => emit("update:modelValue", value),
})
</script>
