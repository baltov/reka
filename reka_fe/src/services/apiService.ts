import axios from 'axios';
import type { Question } from '../types';

const API_URL = import.meta.env.VITE_API_URL + '/questions';

export async function fetchQuestions(): Promise<Question[]> {
    try {
        const response = await axios.get<Question[]>(API_URL);
        console.log('Fetched questions:', response.data);
        return response.data;
    } catch (error) {
        console.error('Failed to fetch questions:', error);
        throw error;
    }
}