<template>
  <Main>
    <div class="md:flex md:justify-between gap-4 pt-5">
      <form v-on:submit.prevent="onSubmit" class="flex-1">
        <div>
          <label for="link" class="block mb-2 text-sm font-medium text-gray-900">link</label>
          <input ref="inputEl" type="url" v-model="link" id="link" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" placeholder="https://example.com" required />
        </div>

        <div class="mt-3">
          <label for="tab" class="block mb-2 text-sm font-medium text-gray-900">Tab</label>
          <input type="text" v-model="tab" id="tab" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" placeholder="Some tab" required />
        </div>
  
        <button type="submit"
          class="text-gray-900 mt-3 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2">Generate</button>
      </form>
  
      <QrPreview :base64Result="base64Result" :variant="variant" @variant="variant = $event" :filename="filename" />
    </div>
  </Main>
</template>

<script setup lang="ts">
import axios from 'axios'
import { onMounted, ref, watch } from 'vue'
import { useDebounce } from '@vueuse/core'
import Main from '../../Components/Layouts/Main.vue'
import QrPreview from '../../Components/QrPreview.vue'

const base64Result = ref('')
const link = ref('')
const tab = ref('')
const filename = ref('')
const variant = ref('email')
const inputEl = ref<HTMLInputElement | null>(null)

onMounted(() => {
  inputEl.value?.focus()
})

const debouncedQuery = useDebounce(link, 500)

watch(debouncedQuery, async (val) => {
  if (val) {
    onSubmit()
  } else {
    base64Result.value = ''
  }
})

watch(variant, async (val) => {
  onSubmit()
})

watch(tab, async (val) => {
    if (val) {
        filename.value = `qr-tab-${val}.png`
    } else {
        filename.value = 'qrcode.png'
    }
})

const onSubmit = async () => {
  if (!link.value) {
    base64Result.value = ''

    return
  }

  try {
    const response = await axios.post('/generate-qr', {
      content: link.value,
      variant: variant.value,
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