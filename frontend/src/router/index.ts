import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', name: 'home', component: HomeView },
    { path: '/coin', name: 'coin', component: () => import('../views/CoinView.vue') },
    { path: '/dice', name: 'dice', component: () => import('../views/DiceView.vue') },
    { path: '/card', name: 'card', component: () => import('../views/CardView.vue') },
    { path: '/number', name: 'number', component: () => import('../views/NumberView.vue') },
  ],
})

export default router
