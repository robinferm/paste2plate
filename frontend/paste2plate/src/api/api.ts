import type { Recipe } from '@/types/recipe'

export const getRecipes = async (): Promise<Recipe[]> => {
  const res = await fetch('http://localhost:3000/recipes')
  const recipes: Recipe[] = await res.json()
  return recipes
}
