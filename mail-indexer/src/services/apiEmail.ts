
import { useSearchStore } from '../stores/searchStore';
import type { EmailData } from '@/interfaces/emailInterfaces';


const HEADER = {
    'Content-Type': 'application/json',
}

export const emailService = {
    async getAllEmails(): Promise<EmailData> {
        const searchStore = useSearchStore();
        const response = await fetch(`${import.meta.env.VITE_API_URL}getEmails`, {
            method: 'POST',
            headers: HEADER,
            body: JSON.stringify(searchStore.paginationParams),
        });
        const data: EmailData = await response.json();
        return data;
    },

    async getEmailByParameter(): Promise<EmailData> {
        const searchStore = useSearchStore();
        const response = await fetch(`${import.meta.env.VITE_API_URL}getEmailsByParameter?searchParameter=${searchStore.searchParameter}`, {
            method: 'POST',
            headers: HEADER,
            body: JSON.stringify(searchStore.paginationParams),
        });
        const data: EmailData = await response.json();
        return data;
    },
};