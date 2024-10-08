<template>
  <Form
    @submit="submitHandler"
    class="bg-white dark:bg-gray-200 w-[380px] p-8 rounded-xl shadow-lg transform transition-all duration-300"
  >
    <div
      v-for="{
        as,
        name,
        label,
        placeholder,
        type,
        class: fieldClass,
        children,
        ...attrs
      } in schema.fields"
      :key="name"
      :class="['mb-5', fieldClass?.wrapper]"
    >
      <!-- Row Configuration for Date and Time -->
      <div
        v-if="['event_start_time', 'event_end_time', 'is_free'].includes(name)"
        class="flex flex-wrap gap-4 mb-5"
      >
        <!-- Time Picker Handling for Start Time -->
        <div v-if="name === 'event_start_time'" class="flex-1 min-w-[200px]">
          <label
            :for="name"
            :class="[
              'block text-sm font-medium text-gray-600 dark:text-gray-800 mb-1',
              fieldClass?.label,
            ]"
          >
            {{ label }}
          </label>
          <Field
            :id="name"
            :name="name"
            type="datetime-local"
            :class="[
              'w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-300 transition duration-300',
              fieldClass?.input,
            ]"
            v-bind="attrs"
          />
          <ErrorMessage
            :name="name"
            :class="['text-red-500 text-sm mt-1', fieldClass?.error]"
          />
        </div>

        <!-- Time Picker Handling for End Time -->
        <div v-if="name === 'event_end_time'" class="flex-1 min-w-[200px]">
          <label
            :for="name"
            :class="[
              'block text-sm font-medium text-gray-600 dark:text-gray-800 mb-1',
              fieldClass?.label,
            ]"
          >
            {{ label }}
          </label>
          <Field
            :id="name"
            :name="name"
            type="datetime-local"
            :class="[
              'w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-300 transition duration-300',
              fieldClass?.input,
            ]"
            v-bind="attrs"
          />
          <ErrorMessage
            :name="name"
            :class="['text-red-500 text-sm mt-1', fieldClass?.error]"
          />
        </div>
        <div v-if="name === 'is_free'" class="flex-1 min-w-[200px]">
          <label
            :for="name"
            :class="[
              'block text-sm font-medium text-gray-600 dark:text-gray-800 mb-1',
              fieldClass?.label,
            ]"
          >
            {{ label }}
          </label>
          <Field
            type="checkbox"
            :id="name"
            :name="name"
            :class="[
              'focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-800 transition duration-300',
              fieldClass?.input,
            ]"
            v-bind="attrs"
          />
          <ErrorMessage
            :name="name"
            :class="['text-red-500 text-sm mt-1', fieldClass?.error]"
          />
        </div>
      </div>

      <!-- Handling for File Input and Other Types -->
      <template v-else>
        <label
          :for="name"
          :class="[
            'block text-sm font-medium text-gray-600 dark:text-gray-800 mb-1',
            fieldClass?.label,
          ]"
        >
          {{ label }}
        </label>

        <!-- File Input Handling -->
        <Field
          v-if="type === 'file'"
          :id="name"
          :name="name"
          type="file"
          :multiple="attrs.multiple || true"
          :class="[
            'w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-300 transition duration-300',
            fieldClass?.input,
          ]"
          v-bind="attrs"
        />
        <Field
          v-else
          :as="as"
          :id="name"
          :name="name"
          :placeholder="placeholder"
          :type="type"
          :class="[
            'w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-200 dark:text-gray-300 transition duration-300',
            fieldClass?.input,
          ]"
          v-bind="attrs"
        >
          <template v-if="children && children.length">
            <component
              v-for="({ tag, text, ...childAttrs }, idx) in children"
              :key="idx"
              :is="tag"
              v-bind="childAttrs"
            >
              {{ text }}
            </component>
          </template>
        </Field>
        <ErrorMessage
          :name="name"
          :class="['text-red-500 text-sm mt-1', fieldClass?.error]"
        />
        <!-- File Input Handling -->
      </template>
    </div>
    <slot />
  </Form>
</template>

<script setup>
import { Form, Field, ErrorMessage } from "vee-validate";
import { defineProps } from "vue";

const props = defineProps({
  schema: {
    type: Object,
    required: true,
  },
  submitHandler: {
    type: Function,
    required: true,
  },
});
</script>
