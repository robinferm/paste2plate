<template>
  <button class="btn" @click="goBack">back</button>
  <div v-if="recipe">
    <h2>{{ recipe.title }}</h2>
    <ul>
      <li v-for="ingredient in recipe.ingredients" :key="ingredient.name">
        {{ ingredient.name }} - {{ ingredient.amount }}
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { getRecipe } from '@/api/api'
import type { Recipe } from '@/types/recipe'
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

const props = defineProps<{ id: string }>()
const recipe = ref<Recipe | null>(null)

const fetchRecipe = async (id: string) => {
  recipe.value = await getRecipe(Number(id))
  console.log(recipe.value)
}

const router = useRouter()
const goBack = () => {
  router.back()
}

onMounted(() => {
  fetchRecipe(props.id)
})
</script>
