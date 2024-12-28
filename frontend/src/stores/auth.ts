import { defineStore } from 'pinia'
import pb from '@/services/pocketbase'
import { useStorage } from '@vueuse/core'
import type { RecordModel } from 'pocketbase'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: useStorage('user', null as RecordModel | null),
    token: useStorage('token', null as string | null),
  }),
  actions: {
    async loginWithGoogle() {
      try {
        const authData = await pb.collection('users').authWithOAuth2({
          provider: 'google',
        })

        this.user = authData.record
        this.token = pb.authStore.token

        return authData
      } catch (err) {
        console.error('Google login failed:', err)
        throw err
      }
    },
    logout() {
      pb.authStore.clear()
      this.user = null
      this.token = null
    },
    isLoggedIn() {
      return this.user !== null
    },
    async authRefresh() {
      const authData = await pb.collection('users').authRefresh()
      this.user = authData.record
      this.token = pb.authStore.token
    },
  },
})
