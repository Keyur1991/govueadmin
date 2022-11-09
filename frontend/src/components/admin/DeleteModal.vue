<template>
    <v-dialog
        v-model="deleteItemDialog"
        width="500"

    >
        <v-card>
        <v-card-title class="headline">
            Are you sure you want to delete ?
        </v-card-title>
        <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" @click="DeleteRecord()">Yes</v-btn>
            <v-btn text @click="deleteItemDialog = false">No</v-btn>
        </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
export default {
    props: ["item", "deleteItemModal", "url"],
    computed: {
        deleteItemDialog: {
            set(value) {
                this.$emit('close')
            },

            get() {
                return this.deleteItemModal
            }
        }
    },
    methods: {
        async DeleteRecord() {
            this.deleteItemDialog = false;

            this.$store.dispatch('showProgress', true)
            await this.$axios.get(this.$URLs.SANCTUM_CSRF)
            let options = {
                _method: 'DELETE'
            }

            await this.$axios({
                url: this.url,
                method: "POST",
                data:options
            }).then((response) => {
                this.$store.dispatch('showProgress', false)

                this.$store.dispatch('showSnackbarMessage', {message: response.data.message, code: '1'})

                this.$emit('reload');

            }).catch((e) => {
                this.$store.dispatch('showProgress', false)
                this.$store.dispatch('serverError', e)
            })
        }
    }
}
</script>
