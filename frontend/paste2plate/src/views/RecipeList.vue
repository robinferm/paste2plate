<template>
  <main>
    <button class="btn" @click="addRecipeClick()">Add</button>
    <div class="grid lg:grid-cols-3 gap-4">
      <RecipeCard
        v-for="recipe in recipes"
        :key="recipe.id"
        :recipe="recipe"
        @click="openRecipe(recipe.id)"
      ></RecipeCard>
    </div>
  </main>
</template>

<script setup lang="ts">
import { getRecipes } from '@/api/api'
import RecipeCard from '@/components/RecipeCard.vue'
import router from '@/router'
import type { Recipe } from '@/types/recipe'
import { onMounted, ref } from 'vue'
import { addRecipe } from '@/api/api'

const recipes = ref<Recipe[]>([])
onMounted(async () => {
  const fetchedRecipes = await getRecipes()
  console.log(fetchedRecipes)
  recipes.value = fetchedRecipes
})

const openRecipe = (id: number) => {
  router.push({ name: 'RecipeDetail', params: { id } })
}

const addRecipeClick = async () => {
  const newRecipe: Recipe = {
    id: 6,
    title: 'Tomato Soup',
    ingredients: [
      { name: 'Tomatoes', amount: '500 g' },
      { name: 'Onion', amount: '1' },
      { name: 'Garlic', amount: '2 cloves' },
    ],
    instructions: [
      'Chop the onions and garlic.',
      'Saute them in a pot.',
      'Add tomatoes and cook for 20 minutes.',
      'Blend and serve.',
    ],
    image: 'https://img.koket.se/standard-giant/snabb-pasta-och-tomatsas-med-burrata.jpg.webp',
  }
  recipes.value.push(newRecipe)
  await addRecipe(newRecipe)
}
</script>
