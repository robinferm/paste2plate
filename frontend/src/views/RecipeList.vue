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
        <button class="btn">Add</button>
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
      @delete="deleteRecipe"
    ></RecipeCard>
  </div>
</template>

<script setup lang="ts">
import RecipeCard from '@/components/RecipeCard.vue'
import router from '@/router'
import { ref, onMounted, computed } from 'vue'
import { useRecipeStore } from '@/stores/recipe'

const recipeStore = useRecipeStore()
const recipes = computed(() => recipeStore.recipes)
onMounted(async () => {
  recipeStore.fetchRecipes()
})

const deleteRecipe = (id: string) => {
  recipeStore.deleteRecipe(id)
}
const openRecipe = (id: string) => {
  router.push({ name: 'RecipeDetail', params: { id } })
}
const url = ref<string>('')

const submitRecipe = async () => {
  if (!url.value) {
    throw new Error('Empty input url')
  }
  const addedId = await recipeStore.addRecipe(url.value)
  openRecipe(addedId)
}
</script>
