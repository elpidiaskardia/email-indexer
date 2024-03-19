<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useSearchStore } from '../stores/searchStore';
import { usePopupStore } from '../stores/popupStore';
import SearchBar from "@/components/SearchBar.vue";
import PopupMessage from "@/components/PopupMessage.vue";
import { emailService } from '../services/apiEmail';
import { typePopup } from '@/enum/typePopup';
const popupStore = usePopupStore();
//const
const selectedRow = ref<number | null>(null);
const emails = ref([
  { subject: "", from: "", to: "", content: "" },
]);
const totalItems = ref(15);
const sortBy = ref<string | null>(null);
const sortDesc = ref(false);
const hoveredColumn = ref<string | null>(null);
const searchParameter = ref('');
const searchStore = useSearchStore()

//functions
const getAllEmails = async () => {
  popupStore.showPopup(typePopup.Loading, "", "Loading...");
  await emailService.getAllEmails()
      .then(result =>  {
        if (result.status == "Success") {
          emails.value = result.result?.data || [];
          totalItems.value = result.result.total;
          popupStore.closePopup();
        } else {
          popupStore.showPopup(typePopup.Information, result.message);
        }
      }).catch(error => {
        popupStore.showPopup(typePopup.Error, error.message, "Error");
      })
};


const getEmailByParameter = async () => {
  if (searchStore.searchParameter.trim() != "") {
      popupStore.showPopup(typePopup.Loading, "", "Loading...");
      await emailService.getEmailByParameter()
          .then(result =>  {
            if (result.status == "Success") {
              emails.value = result.result?.data || [];
              totalItems.value = result.result.total;
              popupStore.closePopup();
            } else {
              popupStore.showPopup(typePopup.Information, result.message);
            }
          }).catch(error => {
            popupStore.showPopup(typePopup.Error, error.message, "Error");
          })
    }else{
      getAllEmails();
    }
};
function boldContent(content: string) {
  const regex = new RegExp(searchStore.searchParameter, 'gi');
  const contentReplace = content.replace(regex, '<strong>$&</strong>')
  return contentReplace;
}

//arrow functions

const selectRow = (index: number) => {
  selectedRow.value = index;
};
const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value) {
    searchStore.paginationParams.current_page = page;
    selectedRow.value = null;
    if (searchStore.searchParameter == "") {
      getAllEmails();

    } else {
      getEmailByParameter();

    }
  }
};
const sortByColumn = (column: string) => {
  if (sortBy.value === column) {
    sortDesc.value = !sortDesc.value;
  } else {
    sortBy.value = column;
    sortDesc.value = false;
  }
  sortTable();
};
const sortTable = () => {
  if (sortBy.value !== null && emails.value ) {
    emails.value.sort((a , b ) => {
      const modifier = sortDesc.value ? -1 : 1;
      const valA = String(a[sortBy.value as keyof typeof a]).toLowerCase();
      const valB = String(b[sortBy.value as keyof typeof a]).toLowerCase();
      if (valA < valB) return -1 * modifier;
      if (valA > valB) return 1 * modifier;
      return 0;
    });
  }
};
const setHoveredColumn = (column: string | null) => {
  hoveredColumn.value = column;
};

const isHovered = (column: string) => {
  return hoveredColumn.value === column;
};

const getRowClasses = computed(() => {
  return (index: number) => ({
    'bg-selected': index === selectedRow.value,
    'bg-white': index % 2 === 0,
    'bg-gray-100': index % 2 !== 0,
  });
});


const totalPages = computed(() => {
  if (totalItems.value > 0) {
    return Math.ceil(totalItems.value / searchStore.paginationParams.page_size);
  }
  return 1;
});
onMounted(() => {
  getAllEmails()

});

</script>

<template>
  <SearchBar @search="getEmailByParameter" />
  <div class="lg:flex lg:flex-row flex-col ">
    <!-- table section -->
    <div class="g:w-1/2 p-1 overflow-auto flex-grow  ">
      <div class="p-3 ml-3 items-center justify-between">
        <div class="max-h-screen overflow-auto  shadow-md sm:rounded-lg">
          <table class="h-[48rem] w-full text-sm text-left rtl:text-right text-gray-700 bg-white">
            <thead class="sticky top-0 text-md text-gray-50  bg-darkblue">
              <tr>
                <th class="py-3 px-6 text-left  relative" @mouseenter="setHoveredColumn('subject')"
                  @mouseleave="setHoveredColumn(null)" @click="sortByColumn('subject')">
                  <div class="flex items-center">
                    <span class="mr-2">Subject</span>
                    <span v-if="isHovered('subject')" class="sort-icon">{{ sortDesc ? '▼' : '▲' }}</span>
                  </div>
                </th>
                <th class="py-3 px-6 text-left   relative" @mouseenter="setHoveredColumn('from')"
                  @mouseleave="setHoveredColumn(null)" @click="sortByColumn('from')">
                  <div class="flex items-center">
                    <span class="mr-2">From</span>
                    <span v-if="isHovered('from')" class="sort-icon">{{ sortDesc ? '▼' : '▲' }}</span>
                  </div>
                </th>
                <th class="py-3 px-6 text-left" @mouseenter="setHoveredColumn('to')" @mouseleave="setHoveredColumn(null)"
                  @click="sortByColumn('to')">
                  <div class="flex items-center">
                    <span class="mr-2">To</span>
                    <span v-if="isHovered('to')" class="sort-icon">{{ sortDesc ? '▼' : '▲' }}</span>
                  </div>
                </th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(email, index) in emails" :key="index" :class="getRowClasses(index)"
                @click="selectRow(index)">
                <td class="py-2 px-4   border-darkblue  max-w-xs overflow-hidden">{{ email.subject }}</td>
                <td class="py-2 px-4   border-darkblue max-w-xs overflow-hidden">{{ email.from }}</td>
                <td class="py-2 px-4 max-w-xs overflow-hidden ">{{ email.to }}</td>
              </tr>
            </tbody>
          </table>
        </div>
       
        <div class="w-full p-3 ">
        <strong>Total:</strong> {{totalItems}}
        <div class="flex justify-center items-center">  
          <button class="px-1 py-1 bg-gray-200 text-gray-700 rounded-md mr-2"
            :disabled="searchStore.paginationParams.current_page === 1" @click="goToPage(1)">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
              stroke="currentColor" class="w-6 h-6">
              <path stroke-linecap="round" stroke-linejoin="round" d="m18.75 4.5-7.5 7.5 7.5 7.5m-6-15L5.25 12l7.5 7.5" />
            </svg>

          </button>
          <button class="px-1 py-1 bg-gray-200 text-gray-700 rounded-md mr-2"
            :disabled="searchStore.paginationParams.current_page === 1"
            @click="goToPage(searchStore.paginationParams.current_page - 1)">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
              stroke="currentColor" class="w-6 h-6">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 19.5 8.25 12l7.5-7.5" />
            </svg>
          </button>
          <span class="text-gray-700">Page {{ searchStore.paginationParams.current_page }} of {{ totalPages }}</span>
          <button class="px-1 py-1 bg-gray-200 text-gray-700 rounded-md ml-2"
            :disabled="searchStore.paginationParams.current_page === totalPages"
            @click="goToPage(searchStore.paginationParams.current_page + 1)">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
              stroke="currentColor" class="w-6 h-6">
              <path stroke-linecap="round" stroke-linejoin="round" d="m8.25 4.5 7.5 7.5-7.5 7.5" />
            </svg>

          </button>
          <button class="px-1 py-1 bg-gray-200 text-gray-700 rounded-md ml-2"
            :disabled="searchStore.paginationParams.current_page === totalPages" @click="goToPage(totalPages)">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
              stroke="currentColor" class="w-6 h-6">
              <path stroke-linecap="round" stroke-linejoin="round" d="m5.25 4.5 7.5 7.5-7.5 7.5m6-15 7.5 7.5-7.5 7.5" />
            </svg>

          </button>

        </div>
      </div>
      </div>
    </div>

    <div class="lg:w-1/2 p-3 container mx-auto">
      <!-- Content section -->
      <div class="p-1 items-center justify-between">
        <div class="p-3 border border-gray-200 sm:rounded-lg shadow-md bg-white h-[48rem] overflow-auto">
          <div v-if="selectedRow !== null">
            <h2 class="text-xl font-bold mb-2 text-gray-800">{{ emails[selectedRow].subject }}</h2>
            {{ searchParameter }}
            <p class="text-gray-700" style="white-space: pre-line" v-html="boldContent(emails[selectedRow].content)"></p>
          </div>
        </div>
      </div>
    </div>
    <PopupMessage />
  </div>
</template>

<style>
@media (max-width: 639px) {
  .lg:flex-row {
    flex-direction: column;
  }
}

.bg-selected {
  background-color: #5B9FFF;
}

.overflow-auto::-webkit-scrollbar {
  width: 12px;
}

.overflow-auto::-webkit-scrollbar-thumb {
  background-color: #7c7c7c;

  border-radius: 6px;
}

.overflow-auto::-webkit-scrollbar-track {
  background-color: #f1f1f1;
}

.overflow-auto::-webkit-scrollbar-thumb:hover {
  background-color: #bababa;
}

.sort-icon {
  font-size: 12px;
  vertical-align: middle;
}
</style>