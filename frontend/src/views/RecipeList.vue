<template>
  <form id="recipeForm" class="space-y-4" @submit.prevent="submitRecipe">
    <label class="form-control w-full max-w-xs">
      <div class="flex gap-1">
        <input
          type="url"
          placeholder="Recipe URL"
          class="input input-bordered w-full max-w-xs"
          v-model="url"
          required
        />
        <button class="btn" @onclick.prevent="submitRecipe">Add</button>
      </div>
    </label>
  </form>
  <div class="divider"></div>
  <div class="grid sm:grid-cols-3 gap-4">
    <RecipeCard
      v-for="recipe in recipes"
      :key="recipe.id"
      :recipe="recipe"
      @click="openRecipe(recipe.id)"
    ></RecipeCard>
  </div>
</template>

<script setup lang="ts">
import RecipeCard from '@/components/RecipeCard.vue'
import router from '@/router'
import type { Recipe } from '@/types/recipe'
import { ref, onMounted } from 'vue'
import { addRecipe, getRecipes } from '@/api/api'

const recipes = ref<Recipe[]>([])
onMounted(async () => {
  const fetchedRecipes = await getRecipes()
  recipes.value = fetchedRecipes
})
const openRecipe = (id: string) => {
  router.push({ name: 'RecipeDetail', params: { id } })
}
const url = ref<string>('')

const submitRecipe = async () => {
  if (!url.value) {
    throw new Error('Empty input url')
  }
  const added = await addRecipe(url.value)
  recipes.value.push(added)
  openRecipe(added.id)
}
</script>
