<template>
  <Main>
    <div class="flex justify-between gap-4">
      <form v-on:submit.prevent="onSubmit" class="mt-5 flex-1">
        <label for="content" class="block mb-2 text-sm font-medium text-gray-900">URL</label>
        <input type="url" v-model="content" id="content" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" placeholder="https://example.com" required />
  
        <button type="submit"
          class="text-gray-900 mt-3 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2">Generate</button>
      </form>
  
      <div>
        <p>Preview</p>
        <img v-if="base64Result" class="w-[200px] border" v-bind:src="base64Result" alt="Preview">
        <div v-else class="w-[200px] h-[200px] flex items-center justify-center border">
          <span>Preview not available</span>
        </div>

        <a v-bind:href="base64Result" download="qr.png" v-bind:disabled="!base64Result" className="mt-6">
            <Button className="text-gray-900 mt-3 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 w-full" v-bind:disabled="!base64Result">Download</Button>
        </a>
      </div>
    </div>
  </Main>
</template>

<script setup lang="ts">
import axios from 'axios'
import { ref, watch } from 'vue'
import { useDebounce } from '@vueuse/core'
import Main from '../../Components/Layouts/Main.vue'

const base64Result = ref('')
const content = ref('')

const debouncedQuery = useDebounce(content, 500)

watch(debouncedQuery, async (val) => {
  if (val) {
    onSubmit()
  } else {
    base64Result.value = ''
  }
})

const onSubmit = async () => {
  try {
    const response = await axios.post('/generate-qr', {
      content: content.value
    })

    base64Result.value = response.data.base64
  } catch (error) {
    console.log(error)
  }
}

defineProps({
  text: String,
  qr: String,
})
</script>