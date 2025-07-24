<template>
  <Main>
    <div class="md:flex md:justify-between gap-4 pt-5">
      <form v-on:submit.prevent="onSubmit" class="flex-1">
        <div>
          <label for="email" class="block mb-2 text-sm font-medium text-gray-900">Email</label>
          <input ref="inputEl" type="email" v-model="email" id="email" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" placeholder="violet@example.com" required />
        </div>

        <div class="mt-3">
          <label for="subject" class="block mb-2 text-sm font-medium text-gray-900">Subject</label>
          <input type="text" v-model="subject" id="subject" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" placeholder="Some subject" />
        </div>

        <div class="mt-3">
          <label for="message" class="block mb-2 text-sm font-medium text-gray-900">Body Message</label>
          <textarea v-model="message" id="message" rows="4"
            class="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500"
            placeholder="Write your thoughts here..."></textarea>
        </div>
  
        <button type="submit"
          class="text-gray-900 mt-3 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2">Generate</button>
      </form>
  
      <QrPreview :base64Result="base64Result" :variant="variant" @variant="variant = $event" />
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
const email = ref('')
const subject = ref('')
const message = ref('')
const variant = ref('email')
const inputEl = ref<HTMLInputElement | null>(null)

onMounted(() => {
  inputEl.value?.focus()
})

const debouncedQuery = useDebounce(email, 500)
const debouncedSubject = useDebounce(subject, 500)
const debouncedMessage = useDebounce(message, 500)

watch(debouncedQuery, async (val) => {
  if (val) {
    onSubmit()
  } else {
    base64Result.value = ''
  }
})

watch(debouncedSubject, async (val) => {
  onSubmit()
})

watch(debouncedMessage, async (val) => {
    onSubmit()
})

watch(variant, async (val) => {
  onSubmit()
})

const onSubmit = async () => {
  if (!email.value) {
    base64Result.value = ''

    return
  }

  try {
    const response = await axios.post('/generate-qr', {
      content: `mailto:${email.value}?subject=${subject.value}&body=${message.value}`,
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