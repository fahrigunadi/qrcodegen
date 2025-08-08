<template>
  <Main>
    <div class="md:flex md:justify-between gap-4 pt-5">
      <form v-on:submit.prevent="onSubmit" class="flex-1">
        <InputLabel for="content">URL</InputLabel>
        <Input type="url" v-model="content" id="content" placeholder="https://example.com" required autofocus />

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
import Input from '../../Components/Input.vue'
import Button from '../../Components/Button.vue'
import { useQrGenerator } from '../../Composables/useQrGenerator'

const { base64Result, generateQr } = useQrGenerator()
const content = ref('')
const variant = ref('default')

const debouncedQuery = useDebounce(content, 500)

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

const onSubmit = async () => {
  await generateQr(content.value, variant.value)
}
</script>