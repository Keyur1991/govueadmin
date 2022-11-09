<template>
    <v-text-field hide-details outlined dense class="pt-0" v-model="GlobalSearch" @keyup.native="applySearch" :label="`${$vuetify.lang.t('$vuetify.Search')}`" append-icon="fa-search" @click:clear="clearSearch"></v-text-field>
</template>
<script>
export default {
    data() {
        return {
            GlobalSearch: ''
        }
    },

    methods: {
        clearSearch() {
            this.GlobalSearch = ''
            this.applySearch()
        },

        applySearch(e) {
            if (e.which == 13) this.$emit('filter', this.GlobalSearch)
        }
    }
}
</script>
