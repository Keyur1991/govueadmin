<template>
    <div v-if="loggedin">
        <h4>You are logged in </h4>
    </div>
</template>
<script setup>
import { computed, onMounted, inject } from "vue"
import { useStore } from "vuex"

const store = useStore()

/**
 * Check if user is logged in
 */
const loggedin = computed(() => {
    return store.getters.loggedin
})

onMounted(() => {
    store.dispatch('isLoggedIn')
})

</script>
