<template>
  <main>
    <form id="recipeForm" class="space-y-4" @submit.prevent="submitRecipe">
      <label class="form-control w-full max-w-xs">
        <div class="label">
          <span class="label-text">Enter recipe URL</span>
        </div>
        <input
          type="url"
          placeholder="Recipe URL"
          class="input input-bordered w-full max-w-xs"
          v-model="url"
          required
        />
      </label>
    </form>
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
  if (url.value) {
    const storage = localStorage.getItem('pocketbase_auth')
    if (storage) {
      const json = JSON.parse(storage)
      const token = json.token
      const res = await fetch(import.meta.env.VITE_BACKEND_URL + '/recipe', {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ url: url.value }),
      })

      const newRecipe: Recipe = await res.json()

      recipes.value.push(newRecipe)
      const added = await addRecipe(newRecipe)
      console.log('added', added)
      openRecipe(added.id)
    }
  }
}
</script>
