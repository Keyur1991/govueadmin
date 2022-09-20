<template>
  <v-app :theme="theme">
    <v-app-bar>
      <v-app-bar-title>Go + Vue Admin</v-app-bar-title>
      <v-btn @click="toggleTheme" variant="plain"><v-icon >mdi-brightness-4</v-icon></v-btn>
      <v-btn v-if="!isLoggedIn" variant="plain" small to="/login">Login</v-btn>
      <v-btn v-else @click="logout">Logout</v-btn>
    </v-app-bar>
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useStore } from 'vuex'

const theme = ref('light')
const store = useStore()

const toggleTheme = () => {
  theme.value = theme.value === 'light' ? 'dark' : 'light'
}

const isLoggedIn = computed(() => {
  return store.getters.loggedin
})

const logout = () => {
  store.dispatch("logout")
}
</script>
