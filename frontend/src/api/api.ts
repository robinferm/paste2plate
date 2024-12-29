import type { Recipe } from '@/types/recipe'
import pb from '@/services/pocketbase'
import { useAuthStore } from '@/stores/auth'

export const getRecipes = async (): Promise<Recipe[]> => {
  const recipes = await pb.collection('recipes').getFullList<Recipe>()
  return recipes
}

export const getRecipe = async (id: string): Promise<Recipe> => {
  const recipe = await pb.collection<Recipe>('recipes').getOne(id)
  return recipe
}

export const deleteRecipe = async (id: string) => {
  await pb.collection('recipes').delete(id)
}

export const addRecipe = async (url: string): Promise<Recipe> => {
  const authStore = useAuthStore()

  if (!authStore.user?.id) {
    throw new Error("Can't add recipe, user is not logged in")
  }

  const res = await fetch(import.meta.env.VITE_BACKEND_URL + '/api/recipe', {
    method: 'POST',
    headers: {
      Authorization: `Bearer ${authStore.token}`,
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ url: url }),
  })

  const recipe: Recipe = await res.json()
  recipe.user = authStore.user?.id
  console.log('adaoiwdjadw', recipe)
  const addedRecipe = await pb.collection('recipes').create<Recipe>(recipe)
  console.log('added', addedRecipe)
  return addedRecipe
}
