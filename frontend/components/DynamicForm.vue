<template>
  <Form @submit="submitHandler" class="bg-white dark:bg-gray-800 w-[380px] p-8 rounded-xl shadow-lg transform transition-all duration-300">
    <div
      v-for="{
        as,
        name,
        label,
        placeholder,
        class: fieldClass,
        children,
        ...attrs
      } in schema.fields"
      :key="name"
      :class="['mb-5', fieldClass?.wrapper]"
    >
      <label :for="name" :class="['block text-sm font-medium text-gray-600 dark:text-gray-400 mb-1', fieldClass?.label]">
        {{ label }}
      </label>
      <Field
        :as="as"
        :id="name"
        :name="name"
        :placeholder="placeholder"
        :class="[
          'w-full px-4 py-3 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-gray-300 transition duration-300',
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
      <ErrorMessage :name="name" :class="['text-red-500 text-sm mt-1', fieldClass?.error]" />
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
