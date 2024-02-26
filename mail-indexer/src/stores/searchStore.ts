import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useSearchStore = defineStore('searchStore', () => {
  const paginationParams = {
    current_page: ref(1),
    page_size: ref(15),
    total_records: 0, 
  };
  const searchParameter = ref('');
  return { paginationParams, searchParameter }
})
