export interface Recipe {
  id: string
  title: string
  ingredients: ingredient[]
  instructions: string[]
  image: string
}

export interface ingredient {
  name: string
  amount: string
}
