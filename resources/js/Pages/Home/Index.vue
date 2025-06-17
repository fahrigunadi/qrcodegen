<template>
  <Head>
    <title>QR Code Generator</title>
  </Head>
  <div class="relative isolate px-6 pt-14 lg:px-8">
    <div class="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80" aria-hidden="true">
      <div
        class="relative left-[calc(50%-11rem)] aspect-[1155/678] w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-gradient-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]"
        style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)" />
    </div>
    <div class="mx-auto max-w-2xl py-32 sm:py-48 lg:py-56">
      <div class="text-center">
        <h1 class="text-4xl font-bold tracking-tight text-gray-900 sm:text-6xl">{{ text }}</h1>
      </div>

      <form v-on:submit.prevent="onSubmit" class="mt-10">
        <label for="content" class="block mb-2 text-sm font-medium text-gray-900">Your content</label>
        <textarea v-model="form.content" id="content" rows="4"
          class="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500"
          placeholder="Write your thoughts here..."></textarea>

          <button type="submit" class="text-gray-900 mt-3 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2">Generate</button>
      </form>

      <p>{{ qr }}</p>
      <img v-if="base64Result" v-bind:src="base64Result" alt="qr">
    </div>
  </div>
</template>

<script setup lang="ts">
import { Head, useForm } from '@inertiajs/vue3'
import axios from 'axios';
import { ref } from 'vue';

const base64Result = ref('');

const form = useForm({
  content: '',
});

const onSubmit = () => {
  axios.post('/generate-qr', form.data()).then(response => {
    base64Result.value = response.data.base64;
  });
}

defineProps({
  text: String,
  qr: String,
});
</script>