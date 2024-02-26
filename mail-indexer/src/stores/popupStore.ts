import { typePopup } from '@/enum/typePopup';
import { defineStore } from 'pinia';

export const usePopupStore = defineStore('popupStore', {
  state: () => ({
    isVisible: false,
    type: typePopup.Loading,
    title: '',
    message: '',
  }),
  actions: {
    showPopup(type: typePopup, message: string = '', title: string = ''): void {
      this.isVisible = true;
      this.type = type;
      this.title = title;
      this.message = message;
    },
    closePopup() {
      this.isVisible = false;
    },
  },
});