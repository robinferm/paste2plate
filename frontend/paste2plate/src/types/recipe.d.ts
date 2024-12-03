export interface Recipe {
  id: number
  title: string
  ingredients: ingredient[]
  instructions: string[]
  image: string
}

export interface ingredient {
  name: string
  amount: string
}
