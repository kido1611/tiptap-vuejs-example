<template>
  <input
    v-model="model"
    class="rounded-md border border-gray-400 px-3 py-2 text-gray-800 focus:ring-blue-800"
  />
</template>

<script setup lang="ts">
const model = defineModel<string>()
</script>
