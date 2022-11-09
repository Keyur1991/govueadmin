<template>
    <v-dialog
        v-model="UserViewDialog"
        width="500"
        >

        <v-card v-if="user != null" color="blue-grey darken-2" class="white--text">
            <v-row dense>
                <v-col md="5" sm="5">
                    <v-img
                        :src="user.profile_pic"
                        height="150px"
                        width="150px"
                        :lazy-src="LazyLoader"
                    >
                    </v-img>
                </v-col>

                <v-col md="7" sm="7">
                    <v-card-title primary-title class="pb-0" >
                        <div>
                            <div class="headline">{{ user.name }}</div>

                        </div>
                    </v-card-title>


                    <v-card-text>
                        <div class="body-1">{{ user.email }}</div>
                    </v-card-text>
                </v-col>
            </v-row>
            <v-divider></v-divider>
            <v-card-actions>
                <v-spacer></v-spacer>

                <v-btn v-if="hasEditAccess"
                    color="primary"
                    :to="editLink"
                >
                    <v-icon small >fa-edit</v-icon>&nbsp;Edit
                </v-btn>
                <v-btn
                    color="secondary"
                    @click="UserViewDialog = false"
                >
                    Close
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
export default {
    props: ['user', 'hasEditAccess'],
    data() {
        return {
            UserViewDialog: false
        }
    },

    computed: {
        LazyLoader() {
            return this.$store.state.LazyLoader
        },

        editLink() {
            return '/admin/users/edit/' + this.user.id
        }
    }
}
</script>
