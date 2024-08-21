<template>
    <form @submit.prevent="handleSubmit">
      <input v-model="task.title" type="text" placeholder="Task Title" required />
      <textarea v-model="task.description" placeholder="Task Description" required></textarea>
      <button type="submit">{{ isEdit ? 'Update Task' : 'Create Task' }}</button>
    </form>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useMutation } from '@vue/apollo-composable';
  import { CREATE_TASK_MUTATION, UPDATE_TASK_MUTATION } from '~/graphql/mutations';
  
  const props = defineProps({
    task: {
      type: Object,
      default: () => ({
        title: '',
        description: '',
      }),
    },
    isEdit: {
      type: Boolean,
      default: false,
    },
  });
  
  const router = useRouter();
  const task = ref({ ...props.task });
  
  const mutation = useMutation(props.isEdit ? UPDATE_TASK_MUTATION : CREATE_TASK_MUTATION);
  
  const handleSubmit = async () => {
    try {
      await mutation.mutate({
        variables: {
          id: task.value.id,
          title: task.value.title,
          description: task.value.description,
        },
      });
      router.push('/tasks');
    } catch (error) {
      console.error(error);
    }
  };
  </script>
  
  <style scoped>
  /* Add some styles */
  </style>
  