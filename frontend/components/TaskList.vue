<template>
    <div>
      <ul>
        <li v-for="task in tasks" :key="task.id">
          <router-link :to="`/tasks/${task.id}`">{{ task.title }}</router-link>
          <button @click="deleteTask(task.id)">Delete</button>
        </li>
      </ul>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  import { useQuery, useMutation } from '@vue/apollo-composable';
  import { GET_TASKS_QUERY, DELETE_TASK_MUTATION } from '~/graphql/queries';
  
  const tasks = ref([]);
  const { result } = useQuery(GET_TASKS_QUERY);
  
  result.value?.tasks.forEach((task) => tasks.value.push(task));
  
  const deleteTaskMutation = useMutation(DELETE_TASK_MUTATION);
  
  const deleteTask = async (id) => {
    try {
      await deleteTaskMutation.mutate({ id });
      tasks.value = tasks.value.filter(task => task.id !== id);
    } catch (error) {
      console.error(error);
    }
  };
  </script>
  
  <style scoped>
  /* Add some styles */
  </style>
  