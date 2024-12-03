<template>
  <main>
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

const recipes = ref<Recipe[]>([])
onMounted(async () => {
  const fetchedRecipes = await getRecipes()
  console.log(fetchedRecipes)
  recipes.value = fetchedRecipes
})

const openRecipe = (id: number) => {
  router.push({ name: 'RecipeDetail', params: { id } })
}
</script>
