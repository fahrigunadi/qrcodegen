<template>
  <Main>
    <div class="md:flex md:justify-between gap-4 pt-5">
      <form v-on:submit.prevent="onSubmit" class="flex-1">
        <div>
          <InputLabel for="link">Link</InputLabel>
          <Input type="url" v-model="link" id="link" placeholder="https://example.com" required autofocus />
        </div>

        <div>
          <InputLabel for="tab">Tab</InputLabel>
          <Input type="text" v-model="tab" id="tab" placeholder="Some tab" required />
        </div>

        <Button type="submit" class="mt-3">Generate</Button>
        <Button type="submit" class="mt-3" v-on:click="reset">Reset</Button>
      </form>

      <QrPreview :base64Result="base64Result" :variant="variant" @variant="variant = $event" :filename="filename" />
    </div>
  </Main>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useDebounce } from '@vueuse/core'
import Main from '../../Components/Layouts/Main.vue'
import QrPreview from '../../Components/QrPreview.vue'
import Button from '../../Components/Button.vue'
import InputLabel from '../../Components/InputLabel.vue'
import Input from '../../Components/Input.vue'
import { useQrGenerator } from '../../Composables/useQrGenerator'

const { base64Result, generateQr } = useQrGenerator()
const link = ref('')
const tab = ref('')
const filename = ref('')
const variant = ref('email')

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

const reset = () => {
  link.value = ''
  tab.value = ''
  base64Result.value = ''
}

const onSubmit = async () => {
  await generateQr(link.value, variant.value)
}
</script>