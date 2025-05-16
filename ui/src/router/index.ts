import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import DefaultLayout from '../layouts/DefaultLayout.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/auth',
      children: [
        {
          path: 'login',
          name: 'login',
          component: () => import('../views/Login.vue')
        },
        {
          path: 'register',
          name: 'register',
          component: () => import('../views/Register.vue')
        }
      ]
    },
    {
      path: '/',
      component: DefaultLayout,
      meta: { requiresAuth: true },
      children: [
        {
          path: '',
          name: 'home',
          component: () => import('../views/Home.vue')
        },
        {
          path: '/settings',
          name: 'settings',
          component: () => import('../views/Settings.vue')
        },
        {
          path: '/d/:slug',
          name: 'document',
          component: () => import('../views/Document.vue')
        }
      ]
    }
  ]
})

router.beforeEach((to, _from, next) => {
  const auth = useAuthStore()

  if (to.meta.requiresAuth && !auth.isAuthenticated) {
    next({ name: 'login' })
  } else if (auth.isAuthenticated && to.path.startsWith('/auth')) {
    next({ name: 'home' })
  } else {
    next()
  }
})

export default router