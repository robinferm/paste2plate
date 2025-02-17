import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import RecipeList from '@/views/RecipeList.vue'
import RecipeDetail from '@/views/RecipeDetail.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: HomeView,
    },
    {
      path: '/recipes',
      name: 'RecipeList',
      component: RecipeList,
    },
    {
      path: '/recipe/:id',
      name: 'RecipeDetail',
      component: RecipeDetail,
      props: true,
    },
  ],
})

export default router
