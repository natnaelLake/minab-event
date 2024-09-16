<template>
  <div class="chart-container bg-white p-6 rounded-lg shadow-lg">
    <h2 class="text-xl font-semibold mb-4 text-gray-700">{{ title }}</h2>
    <div class="h-64">
      <BarChart v-if="chartType === 'bar'" :chart-data="chartData" />
      <LineChart v-if="chartType === 'line'" :chart-data="chartData" />
      <PieChart v-if="chartType === 'pie'" :chart-data="chartData" />
    </div>
  </div>
</template>

<script setup>
import { BarChart, LineChart, PieChart } from 'vue-chart-3';
import { Chart, registerables } from 'chart.js';
Chart.register(...registerables);

const props = defineProps({
  title: String,
  chartType: {
    type: String,
    required: true,
    default: 'bar',
  },
  chartData: {
    type: Object,
    required: true,
  }
});
</script>

<style scoped>
.chart-container {
  margin-bottom:20px;
  height: 500px;
  background-color: #fff;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: box-shadow 0.3s ease;
}

.chart-container:hover {
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15);
}

h2 {
  color: #333;
}
</style>
