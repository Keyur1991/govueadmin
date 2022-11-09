<template>
    <v-layout row wrap>
        <v-flex xs6 md6 sm6 class="pt-1">Rows Per Page:</v-flex>
        <v-flex xs6 md2 sm2 class="per-page-input"><v-select class="pt-0 mt-0" v-model="PerPage" :items="items"  @input="changeValue"></v-select></v-flex>
    </v-layout>
</template>
<script>
export default {
    data() {
        return {
            PerPage: 10,
            items: [5, 10, 20, 50]
        }
    },

    methods: {
        changeValue() {
            console.log(this.PerPage)
            this.$emit('select', this.PerPage)
        }
    }
}
</script>
