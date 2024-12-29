export interface Recipe {
  id: string
  title: string
  ingredients: ingredient[]
  steps: string[]
  imageUrl: string
  user: string
}

export interface ingredient {
  name: string
  amount: string
}
