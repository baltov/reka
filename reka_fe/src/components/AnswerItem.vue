<template>
  <div class="flex items-start mb-3">
    <label :for="uniqueId" class="flex items-center cursor-pointer w-full">
      <input
          type="checkbox"
          :id="uniqueId"
          :value="answer.id"
          :checked="isSelected"
          @change="onSelectionChange"
          :disabled="disabled || hasSubmitted"
          :class="checkboxClasses"
          class="form-checkbox h-5 w-5 transition-colors duration-300"
      />
      <span class="ml-2 text-gray-800 dark:text-gray-200 text-left flex-1">
        {{ answer.text }}
      </span>
    </label>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue';
import type { PropType } from 'vue';
import type { Answer } from '../types';

export default defineComponent({
  name: 'AnswerItem',
  props: {
    answer: {
      type: Object as PropType<Answer>,
      required: true,
    },
    isSelected: {
      type: Boolean,
      required: true,
    },
    uniqueId: {
      type: String,
      required: true,
    },
    disabled: {
      type: Boolean,
      default: false,
    },
    hasSubmitted: {
      type: Boolean,
      default: false,
    },
  },
  emits: ['update:selection'],
  setup(props, { emit }) {
    const onSelectionChange = (event: Event) => {
      const target = event.target as HTMLInputElement;
      emit('update:selection', target.checked);
    };


    const checkboxClasses = computed(() => {
      if (!props.hasSubmitted) {
        return 'text-blue-600';
      }

      return 'text-red-500';
    });

    return {
      onSelectionChange,
      checkboxClasses,
    };
  },
});
</script>

<style scoped>
/* Minimal custom styles; leveraging Tailwind CSS */
</style>