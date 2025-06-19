<template>
  <div>
    <p>Preview</p>

    <img
      v-if="base64Result"
      class="w-[200px] border"
      :src="base64Result"
      alt="Preview"
    />

    <div v-else class="w-[200px] h-[200px] flex items-center justify-center border">
      <span class="text-gray-500 text-sm">Preview not available</span>
    </div>

    <a
      :href="base64Result"
      download="qr.png"
      :disabled="!base64Result"
      class="mt-1 block"
    >
      <button
        class="text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 w-full"
        :disabled="!base64Result"
      >
        Download
      </button>
    </a>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  base64Result: string | undefined
}>()
</script>
