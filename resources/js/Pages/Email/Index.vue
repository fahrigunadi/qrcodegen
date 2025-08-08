<template>
  <Main>
    <div class="md:flex md:justify-between gap-4 pt-5">
      <form v-on:submit.prevent="onSubmit" class="flex-1">
        <div>
          <InputLabel for="email">Email</InputLabel>
          <Input type="email" v-model="email" id="email" placeholder="email@example.com" required autofocus />
        </div>

        <div class="mt-3">
          <InputLabel for="subject">Subject</InputLabel>
          <Input type="text" v-model="subject" id="subject" placeholder="Some subject" required />
        </div>

        <div class="mt-3">
          <InputLabel for="message">Body Message</InputLabel>
          <Textarea v-model="message" id="message" rows="4" placeholder="Write your thoughts here..." />
        </div>
  
        <Button type="submit" class="mt-3">Generate</Button>
      </form>
  
      <QrPreview :base64Result="base64Result" :variant="variant" @variant="variant = $event" />
    </div>
  </Main>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useDebounce } from '@vueuse/core'
import Main from '../../Components/Layouts/Main.vue'
import QrPreview from '../../Components/QrPreview.vue'
import InputLabel from '../../Components/InputLabel.vue'
import Textarea from '../../Components/Textarea.vue'
import Input from '../../Components/Input.vue'
import Button from '../../Components/Button.vue'
import { useQrGenerator } from '../../Composables/useQrGenerator'

const { base64Result, generateQr } = useQrGenerator()
const email = ref('')
const subject = ref('')
const message = ref('')
const variant = ref('email')

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
  await generateQr(`mailto:${email.value}?subject=${subject.value}&body=${message.value}`, variant.value)
}
</script>