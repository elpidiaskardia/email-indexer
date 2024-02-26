<script setup lang="ts">
import IconSearch from './icons/IconSearch.vue';
import { useSearchStore } from '../stores/searchStore';
const searchStore = useSearchStore()
const emit = defineEmits(['update:searchParameter', 'search'])
function handleEnterKey(event: Event) {
  const target = event.target as HTMLInputElement;
  const value = target.value;
  searchStore.searchParameter = value;
  searchStore.paginationParams.current_page = 1;
  emit('search', value);
};
</script>

<template>
  <div class="p-3 ml-3 items-center justify-between">
    <div class="relative rounded-md shadow-sm">
      <IconSearch />
      <input type="text"
        class="block w-full rounded-md border-0 py-1.5 pl-12 pr-20 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400  focus:ring-2  focus:ring-inset focus:ring-blue sm:text-sm sm:leading-8"
        placeholder="Find email..." @keyup.enter="handleEnterKey" />
    </div>
  </div>
</template>
