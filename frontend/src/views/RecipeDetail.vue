<template>
  <div v-if="recipe">
    <h2 class="text-2xl my-2">{{ recipe.title }}</h2>
    <figure>
      <img :src="recipe.imageUrl" />
    </figure>
    <div class="divider"></div>
    <h2 class="text-2xl mb-2">Ingredients</h2>
    <ul>
      <li class="columns-2" v-for="ingredient in recipe.ingredients" :key="ingredient.name">
        {{ ingredient.name }}
        <div class="divider-horizontal"></div>
        {{ ingredient.amount }}
      </li>
    </ul>
    <div class="divider"></div>
    <h2 class="text-2xl mb-2">Steps</h2>
    <ol class="flex flex-col gap-6">
      <li
        :class="`card card-bordere p-3 flex flex-row gap-2 ${index % 2 === 0 ? 'dark:bg-gray-80 bg-gray-300' : ''}`"
        v-for="(instruction, index) in recipe.steps"
        :key="instruction"
      >
        <div>{{ index + 1 }}.</div>
        <div>{{ instruction }}</div>
      </li>
    </ol>
  </div>
</template>

<script setup lang="ts">
import { getRecipe } from '@/api/api'
import type { Recipe } from '@/types/recipe'
import { onMounted, ref } from 'vue'

const props = defineProps<{ id: string }>()
const recipe = ref<Recipe | null>(null)

const fetchRecipe = async (id: string) => {
  recipe.value = await getRecipe(id)
}

onMounted(() => {
  fetchRecipe(props.id)
})
</script>
