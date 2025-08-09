<template>
  <Main>
    <div class="md:flex md:justify-between gap-4 pt-5">
      <form @submit.prevent="onSubmit" class="flex-1">
        <div>
          <InputLabel for="link">Link</InputLabel>
          <Input id="link" type="url" v-model="link" placeholder="https://example.com" required autofocus />
        </div>

        <div class="mt-3">
          <InputLabel for="tab">Tab</InputLabel>
          <Input id="tab" type="text" v-model="tab" placeholder="Some tab" required />
        </div>

        <div class="mt-3 flex gap-3">
          <Button type="submit">Generate</Button>
          <Button type="button" @click="reset">Reset</Button>
        </div>
      </form>

      <QrPreview :base64Result="base64Result" :variant="variant" @variant="variant = $event" :filename="filename" />
    </div>
  </Main>
</template>

<script setup lang="ts">
import { ref, watch, computed } from 'vue'
import { useDebounce } from '@vueuse/core'
import Main from '../../Components/Layouts/Main.vue'
import QrPreview from '../../Components/QrPreview.vue'
import Button from '../../Components/Button.vue'
import InputLabel from '../../Components/InputLabel.vue'
import Input from '../../Components/Input.vue'
import { useQrGenerator } from '../../Composables/useQrGenerator'

const link = ref<string>('')
const tab = ref<string>('')
const variant = ref<'email' | string>('email')
const filename = computed(() =>
  tab.value ? `qr-tab-${tab.value}.png` : 'qrcode.png'
)

const { base64Result, generateQr } = useQrGenerator()

const debouncedLink = useDebounce(link, 500)

watch([debouncedLink, variant], async ([newLink]) => {
  if (newLink) {
    await generateQr(newLink, variant.value)
  } else {
    base64Result.value = ''
  }
})

const onSubmit = async () => {
  await generateQr(link.value, variant.value)
}

const reset = () => {
  link.value = ''
  tab.value = ''
  base64Result.value = ''
}
</script>
