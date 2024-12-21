export interface Recipe {
  id: string
  title: string
  ingredients: ingredient[]
  steps: string[]
  imageurl: string
}

export interface ingredient {
  name: string
  amount: string
}
