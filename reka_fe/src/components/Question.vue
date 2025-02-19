<template>
  <div class="p-6 bg-white shadow-md rounded-lg mb-5">
    <h3 class="text-xl font-semibold mb-4" style="color: black">
      {{ questionNumber }}. {{ question.text }}
      <!-- <span class="text-sm text-gray-500 ml-2">[Number: {{ question.numberField }}]</span> -->
    </h3>

    <div v-if="question.images.length > 0" class="mb-4 answer-image">
      <img
          :src="getImageUrl(question)"
          alt="Картинка от въпроса"
      />
    </div>

    <div class="flex flex-col space-y-3">
      <AnswerItem
          v-for="answer in question.answers"
          :key="answer.id"
          :answer="answer"
          :isSelected="isSelected(answer.id)"
          :uniqueId="`question-${question.id}-answer-${answer.id}`"
          :disabled="false"
          :hasSubmitted="hasSubmitted"
          @update:selection="handleSelection(answer.id, $event)"
          class="w-full"
      />
    </div>

    <div v-if="hasSubmitted" class="mt-4" >
      <span
          v-if="isCorrect"
          class="text-green-600 font-medium flex items-center"
      >
        Вярно! <span class="ml-2">✅</span>
      </span>
      <span
          v-else
          class="text-red-600 font-medium flex items-center"
      >
        Грешно <span class="ml-2">❌</span>
      </span>
    </div>

    <div v-if="hasSubmitted" class="mt-4">
      <h4 class="text-lg font-semibold mb-2" style="color: black">Верният отговор е:</h4>
      <ul class="list-disc list-inside space-y-1">
        <li v-for="ans in correctAnswers" :key="ans.id" class="text-gray-700">
          {{ ans.text }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, computed } from 'vue';
import AnswerItem from './AnswerItem.vue';
import type { Question } from '../types';

export default defineComponent({
  name: 'Question',
  components: {
    AnswerItem,
  },
  props: {
    question: {
      type: Object as () => Question,
      required: true,
    },
    questionNumber: {
      type: Number,
      required: true,
    },
  },
  methods: {
    getImageUrl(question: Question): string {
      const imageId = question.images[0].id;
      return `/api/images/${imageId}`;
    },
  },
  emits: ['question-answered'],

  setup(props, { emit }) {
    const state = reactive({
      selectedAnswers: new Set<number>(),
      hasSubmitted: false,
      isCorrect: false,
    });

    const correctAnswers = computed(() => {
      return props.question.answers.filter((a) => a.correct);
    });

    const handleSelection = (answerId: number, isSelected: boolean) => {
      if (isSelected) {
        state.selectedAnswers.add(answerId);
      } else {
        state.selectedAnswers.delete(answerId);
      }

      if (!state.hasSubmitted) {
        submitAnswer();
      }
    };

    const isSelected = (answerId: number) => {
      return state.selectedAnswers.has(answerId);
    };

    const submitAnswer = () => {
      const correctAnswerIds = props.question.answers
          .filter((a) => a.correct)
          .map((a) => a.id);

      const allCorrectSelected = correctAnswerIds.every((id) =>
          state.selectedAnswers.has(id)
      );

      const anyIncorrectSelected = props.question.answers
          .filter((a) => !a.correct)
          .some((a) => state.selectedAnswers.has(a.id));

      state.isCorrect = allCorrectSelected && !anyIncorrectSelected;
      state.hasSubmitted = true;

      emit('question-answered', {
        questionId: props.question.id,
        isCorrect: state.isCorrect,
      });
    };

    return {
      handleSelection,
      isSelected,
      hasSubmitted: computed(() => state.hasSubmitted),
      isCorrect: computed(() => state.isCorrect),
      correctAnswers,
    };
  },
});
</script>