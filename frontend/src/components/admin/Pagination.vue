<template>
    <v-row class="pa-2">
        <v-col md="5" sm="5" xs="12">
            <v-layout row wrap>
                <v-flex xs6 md3 sm3 class="pt-1">{{ $vuetify.lang.t('$vuetify.RowsPerPage') }}:</v-flex>
                <v-flex xs6 md2 sm2 class="per-page-input"><v-select class="pt-0 mt-0" v-model="PerPage" :items="items"></v-select></v-flex>
            </v-layout>
        </v-col>
        <v-col md="7" sm="7" xs="12">
            <v-pagination
                prev-icon="fa-angle-left"
                next-icon="fa-angle-right"
                v-model="CurrentPage"
                :length="pages"
                total-visible="7"
            ></v-pagination>
        </v-col>
    </v-row>
</template>
<script>
export default {
    props: ['pages'],
    data() {
        return {
            PerPage: 10,
            items: [5, 10, 20, 50],
            CurrentPage: 1
        }
    },
    watch: {
        CurrentPage() {
            this.fetchData()
        },

        PerPage() {
            this.fetchData()
        }
    },
    methods: {
        fetchData() {
            this.$emit('update', this.CurrentPage, this.PerPage)
        }
    }
}
</script>
