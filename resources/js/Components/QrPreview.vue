<template>
  <div>
    <p>Preview</p>

    <img
      v-if="base64Result"
      class="w-full md:w-[200px] border"
      :src="base64Result"
      alt="Preview"
    />

    <div v-else class="w-full h-[200px] md:w-[200px] flex items-center justify-center border">
      <span class="text-gray-500 text-sm">Preview not available</span>
    </div>

    <div class="flex gap-2 my-2">
      <button
          @click="changeVariant('default')"
          class="text-gray-900 bg-white border border-gray-300 font-medium rounded-lg overflow-hidden p-2"
          v-bind:class="variant !== 'circle-shape' ? 'ring-2 ring-offset-2 ring-blue-500' : ''"
        >
          <span class="block w-[30px] h-[30px] bg-gray-800"></span>
      </button>
      <button
          @click="changeVariant('circle-shape')"
          class="text-gray-900 bg-white border border-gray-300 font-medium rounded-lg overflow-hidden p-2"
          v-bind:class="variant === 'circle-shape' ? 'ring-2 ring-offset-2 ring-blue-500' : ''"
        >
        <span class="block w-[30px] h-[30px] bg-gray-800 rounded-full"></span>
      </button>
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
  variant: string | undefined
}>()

const emit = defineEmits(['variant'])

function changeVariant(v: string) {
  emit('variant', v)
}
</script>
