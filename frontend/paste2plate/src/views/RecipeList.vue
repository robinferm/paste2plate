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
import { getRecipes } from '@/api/api'
import RecipeCard from '@/components/RecipeCard.vue'
import router from '@/router'
import type { Recipe } from '@/types/recipe'
import { onMounted, ref } from 'vue'
import { addRecipe } from '@/api/api'

const recipes = ref<Recipe[]>([])
onMounted(async () => {
  const fetchedRecipes = await getRecipes()
  recipes.value = fetchedRecipes
})

const openRecipe = (id: number) => {
  router.push({ name: 'RecipeDetail', params: { id } })
}
const url = ref<string>('')

const submitRecipe = async () => {
  if (url.value) {
    const res = await fetch('http://localhost:5000/recipe', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ url: url.value }),
    })

    const newRecipe: Recipe = await res.json()

    // const newRecipe: Recipe = {
    //   id: (recipes.value.length + 1).toString(),
    //   title: 'Tomato Soup',
    //   ingredients: [
    //     { name: 'Tomatoes', amount: '500 g' },
    //     { name: 'Onion', amount: '1' },
    //     { name: 'Garlic', amount: '2 cloves' },
    //   ],
    //   instructions: [
    //     'Chop the onions and garlic.',
    //     'Saute them in a pot.',
    //     'Add tomatoes and cook for 20 minutes.',
    //     'Blend and serve.',
    //   ],
    //   image: 'https://img.koket.se/standard-giant/snabb-pasta-och-tomatsas-med-burrata.jpg.webp',
    // }

    newRecipe.id = (recipes.value.length + 1).toString()

    recipes.value.push(newRecipe)
    const added = await addRecipe(newRecipe)
    openRecipe(Number(added.id))
  }
}
</script>
