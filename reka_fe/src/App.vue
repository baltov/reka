<template>
  <div id="app" class="max-w-7xl mx-auto p-8 text-center">
    <h1 class="text-4xl font-semibold mb-8">Листовки на МА за водач на малък кораб</h1>

    <div v-if="isLoading" class="text-lg">Loading questions...</div>
    <div v-else>
      <div class="mb-4 text-lg">Брой въпроси: {{ questions.length }}</div>
      <div v-for="(question, index) in questions" :key="question.id" class="mb-6">
        <h2 class="text-2xl font-medium mb-4">Въпрос {{ index + 1 }}</h2>
        <Question
            :question="question"
            :questionNumber="index + 1"
            @question-answered="handleQuestionAnswered"
        />
      </div>
      <div v-if="summary" class="summary mt-10 p-5 border-t-2 border-gray-800">
        <h2 class="text-2xl font-medium mb-4">Статистика</h2>
        <p class="mb-2">Брой въпроси: {{ summary.total }}</p>
        <p class="mb-2 text-green-500">Правилно отговорени: {{ summary.correct }}</p>
        <p class="text-red-300">Грешно отговорени (max 9): {{ summary.incorrect }}</p>
      </div>
    </div>
    <div v-if="error" class="error mt-5 text-red-500">
      <p>Failed to load questions. Please try again later.</p>
      <p>{{ error }}</p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue';
import Question from './components/Question.vue';
import type { Question as QuestionType } from './types';

import { fetchQuestions } from "./services/apiService.ts";

export default defineComponent({
  name: 'App',
  components: {
    Question,
  },
  setup() {
    const questions = ref<QuestionType[]>([]);

    const state = reactive({
      answeredQuestions: {} as { [key: number]: boolean },
    });

    const isLoading = ref<boolean>(false);
    const error = ref<string | null>(null);
    const summary = ref<{ total: number; correct: number; incorrect: number } | null>(null);

    const loadQuestions = async () => {
      isLoading.value = true;
      try {
        const fetchedQuestions = await fetchQuestions();
        questions.value = fetchedQuestions;
      } catch (err: any) {
        error.value = err.message || 'Unknown error';
      } finally {
        isLoading.value = false;
      }
    };

    const handleQuestionAnswered = (payload: { questionId: number; isCorrect: boolean }) => {
      state.answeredQuestions[payload.questionId] = payload.isCorrect;
      updateSummary();
    };

    const updateSummary = () => {
      const total = questions.value.length;
      const correct = Object.values(state.answeredQuestions).filter(Boolean).length;
      const incorrect = Object.values(state.answeredQuestions).filter((val) => val === false).length;

      summary.value = { total, correct, incorrect };
    };

    onMounted(() => {
      loadQuestions();
    });

    return {
      questions,
      isLoading,
      error,
      handleQuestionAnswered,
      summary,
    };
  },
});
</script>

<style>
/* Removed custom styles in favor of Tailwind CSS */
</style>