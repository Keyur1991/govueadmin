<template>
    <v-layout row wrap>
        <v-flex xs12 sm12 md12>

        </v-flex>
    </v-layout>
</template>

<script>
    export default {
        data() {
            return {

            }
        },

        created() {
            this.$Utils.checkUserLoggedIn.call(this)
        }
    }
</script>
