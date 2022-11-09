<style>
.white-bg {
    background: #fff !important;
}
</style>
<template>
    <v-container fluid>
        <v-row>
            <v-col>
                <v-dialog scrollable v-model="permissionsModalDialog" content-class="white-bg" max-width="600px">
                    <v-card>
                        <v-card-title class="headline primary lighten-2 pa-3" primary-title>
                            <span class="headline white--text">{{ $vuetify.lang.t('$vuetify.Roles.Listing.PermissionsBtn') }} / {{ roleItem.name }}</span>
                        </v-card-title>
                        <v-card-text class="pt-2">
                            <v-expansion-panels>
                                <v-expansion-panel v-for="(items, key) in permissions" :key="key">
                                    <v-expansion-panel-header>
                                        {{ key }}
                                    </v-expansion-panel-header>
                                    <v-expansion-panel-content  >
                                        <v-list-item dense v-for="(item, index) in items" :key="index">
                                            <v-list-item-action>
                                                <v-checkbox v-model="item.isSelected" hide-details></v-checkbox>
                                            </v-list-item-action>

                                            <v-list-item-content >
                                                <v-list-item-title>{{ item.feature_name }} </v-list-item-title>
                                                <v-list-item-subtitle>{{ item.description }}</v-list-item-subtitle>
                                            </v-list-item-content>
                                        </v-list-item>
                                    </v-expansion-panel-content>
                                </v-expansion-panel>
                            </v-expansion-panels>
                        </v-card-text>
                        <v-card-actions>
                            <v-btn color="primary" type="button" @click="savePermissions">{{ $vuetify.lang.t('$vuetify.SaveBtn') }}</v-btn>
                            <v-spacer></v-spacer>
                            <v-btn type="button" text @click="resetDialog()">{{ $vuetify.lang.t('$vuetify.CancelBtn') }}</v-btn>
                        </v-card-actions>
                        
                    </v-card>
                </v-dialog>
            </v-col>
        </v-row>
    </v-container>

</template>

<script>
export default {
    props: ['roleItem', 'permissionsSaveModal', 'permissions'],
    data() {
        return {
            ExpansionPanel: true,
            allowedPermissions: []
        }
    },
    computed: {
        permissionsSaveUrl() {
            return this.$URLs.ROLES_PERMISSIONS  + '/' + this.roleItem.id
        },
        permissionsModalDialog: {
            set(value) {
                this.$emit('close')
            },

            get() {
                return this.permissionsSaveModal
            }
        }
    },
    methods: {
        resetDialog() {
            this.roleName = ''
            this.permissionsModalDialog = false
        },

        async savePermissions() {
            this.$store.dispatch('showProgress', true)

            this.allowedPermissions = [];

            for (let p in this.permissions) {
                this.permissions[p].forEach((item) => {
                    if (item.isSelected == true) {
                        this.allowedPermissions.push(item.id)
                    }
                })
            }


            let dataBody = {
                permissions: this.allowedPermissions,
                _method: 'PUT'
            }
            
            await this.$axios.get(this.$URLs.SANCTUM_CSRF)
            await this.$axios({
                url: this.permissionsSaveUrl,
                method: 'POST',
                data: dataBody
            }).then((response) => {
                this.$store.dispatch('showSnackbarMessage',{message: response.data.message, code : '1'})

                this.$store.dispatch('showProgress', false);

                this.permissionsModalDialog = false
            }).catch((e) => {
                this.$store.dispatch('showProgress', false)
                this.$store.dispatch('serverError', e)
            });
        }
    }
}
</script>
