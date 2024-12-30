import { defineStore } from 'pinia'
import type { Recipe } from '@/types/recipe'
import { getRecipes, deleteRecipe, addRecipe } from '@/api/api'

export const useRecipeStore = defineStore('recipe', {
  state: () => ({
    recipes: [] as Recipe[],
  }),
  actions: {
    async fetchRecipes() {
      try {
        const fetchedRecipes = await getRecipes()
        this.recipes = fetchedRecipes
      } catch {}
    },
    async addRecipe(url: string): Promise<string> {
      try {
        const added = await addRecipe(url)
        this.recipes.push(added)
        return added.id
      } catch {
        throw new Error('Failed to add recipe')
      }
    },
    async deleteRecipe(id: string) {
      try {
        await deleteRecipe(id)
        this.recipes = this.recipes.filter((recipe) => recipe.id !== id)
      } catch {}
    },
  },
})
