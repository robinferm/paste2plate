import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import RecipeList from '@/views/RecipeList.vue'
import RecipeDetail from '@/views/RecipeDetail.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: RecipeList,
      meta: { requiresAuth: true },
    },
    { path: '/recipe/:id', name: 'RecipeDetail', component: RecipeDetail, props: true },
  ],
})

export default router
