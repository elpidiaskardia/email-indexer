<script setup lang="ts">
import { usePopupStore } from '../stores/popupStore';
import { typePopup } from '@/enum/typePopup';
const popupStore = usePopupStore();

</script>

<template>
  <div v-if="popupStore.isVisible" class="fixed inset-0 flex items-center justify-center">

    <div class="bg-white p-8 rounded shadow-xl text-center">
      <div v-if="popupStore.type == typePopup.Loading">
        <div class="flex items-center justify-center">
          <div class="w-14 h-14 border-t-4 border-blue border-solid rounded-full animate-spin"></div>
        </div>
      </div>
      <div v-if="popupStore.type == typePopup.Information">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
          class="w-16 h-16 text-blue mx-auto mb-4">
          <path stroke-linecap="round" stroke-linejoin="round"
            d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z" />
        </svg>
      </div>
      <div v-if="popupStore.type==typePopup.Success">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
          class="w-16 h-16 text-green-500 mx-auto mb-4">
          <path stroke-linecap="round" stroke-linejoin="round"
            d="M9 12.75 11.25 15 15 9.75M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
        </svg>
      </div>
      <div v-if="popupStore.type==typePopup.Error">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
          class="w-16 h-16 text-red-500 mx-auto mb-4">
          <path stroke-linecap="round" stroke-linejoin="round"
            d="m9.75 9.75 4.5 4.5m0-4.5-4.5 4.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
        </svg>
      </div>
      <p class="text-xl font-bold mb-2">{{ popupStore.title }}</p>
      <p>{{ popupStore.message }}</p>
      <button v-if="popupStore.type != typePopup.Loading" @click="popupStore.closePopup" class="mt-4 bg-darkblue text-white py-2 px-4 rounded">Continue</button>
    </div>
  </div>
</template>
