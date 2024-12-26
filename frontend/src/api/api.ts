import type { Recipe } from '@/types/recipe'
import pb from '@/services/pocketbase'
import { useAuthStore } from '@/stores/auth'

export const getRecipes = async (): Promise<Recipe[]> => {
  const recipes = await pb.collection('recipes').getFullList<Recipe>()
  console.log('recipe', recipes[0])
  return recipes
}

export const getRecipe = async (id: string): Promise<Recipe> => {
  const recipe = await pb.collection<Recipe>('recipes').getOne(id)
  return recipe
}

export const deleteRecipe = async (id: string) => {
  await pb.collection('recipes').delete(id)
}

export const addRecipe = async (recipe: Recipe): Promise<Recipe> => {
  const authStore = useAuthStore()

  if (!authStore.user?.id) {
    throw new Error("Can't add recipe, user is not logged in")
  }
  recipe.user = authStore.user?.id

  const addedRecipe = await pb.collection('recipes').create<Recipe>(recipe)
  return addedRecipe
}
