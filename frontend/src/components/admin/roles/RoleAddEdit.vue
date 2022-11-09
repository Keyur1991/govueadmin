<template>
    <v-container fluid >
        <v-row dense>
            <v-col>
                <v-dialog v-model="dialog" max-width="600px">
                    <v-card>
                        <v-card-title class="headline primary lighten-2 pa-3" primary-title>
                            <span class="headline white--text">{{ title }}</span>
                        </v-card-title>

                        <v-form :action="createEditActionUrl" method="post" ref="createEditRoleForm" @submit.prevent="createOrEditRole()">
                            <v-card-text>
                                <v-row dense wrap v-if="containsErrors || serverValidationErrors.length > 0">
                                    <v-col>
                                        <validation-error-alert :contains-errors="containsErrors"></validation-error-alert>
                                        <server-validation-messages :validation-messages="serverValidationErrors"></server-validation-messages>
                                    </v-col>
                                </v-row>
                                <v-row dense>
                                    <v-col>
                                        <v-text-field outlined dense :label="`${$vuetify.lang.t('$vuetify.Roles.Fields.Name')}*`"
                                            v-model="RoleName"
                                            :error-messages="nameErrors" >
                                        </v-text-field>
                                    </v-col>
                                </v-row>
                            </v-card-text>
                            <v-card-actions>
                                <v-btn color="primary" type="submit">{{ $vuetify.lang.t('$vuetify.SaveBtn') }}</v-btn>
                                <v-spacer></v-spacer>
                                <v-btn  type="button" text @click="$emit('close')">{{ $vuetify.lang.t('$vuetify.CancelBtn') }}</v-btn>
                            </v-card-actions>
                        </v-form>
                    </v-card>
                </v-dialog>
            </v-col>
        </v-row>
    </v-container>
    
</template>

<script>
import Vue from "vue"

import Vuelidate from 'vuelidate'

Vue.use(Vuelidate)

import { validationMixin } from "vuelidate";

import {
  required,
  maxLength
} from "vuelidate/lib/validators";

var ServerValidationMessages = require('../ServerValidationMessages.vue').default

var ValidationErrorAlert = require('../ValidationErrorAlert.vue').default
export default {
    props: ['roleItem', 'title', 'edit', 'roleItemModal'],
    data() {
        return {
            RoleName: '',
            serverValidationErrors: []
        }
    },
    computed: {
        dialog: {
            get() {
                if (this.edit) {
                    this.RoleName = this.roleItem.name
                } else {
                    this.RoleName = ''
                }
                this.serverValidationErrors = []
                return this.roleItemModal
            }
        },
        ShowProgress() {
            return this.$store.state.ShowProgress
        },
        createEditActionUrl() {
            if (this.edit == true) {
                return this.$URLs.ROLES_LIST + '/' + this.roleItem.id
            } else {
                return this.$URLs.ROLES_LIST
            }
        },

        nameErrors() {
            const errors = [];

            if (!this.$v.RoleName.$dirty) return errors;

            !this.$v.RoleName.required && errors.push(this.$vuetify.lang.t('$vuetify.Roles.Validations.RoleRequired'));

            !this.$v.RoleName.maxLength &&
                errors.push(this.$vuetify.lang.t('$vuetify.Roles.Validations.RoleMaxLength'));

            return errors;
        },

        containsErrors() {
            return this.nameErrors.length > 0;
        }
    },
    methods: {
        async createOrEditRole() {
            if (this.$Utils.isValidForm.call(this)) {
                this.$store.dispatch('showProgress', true)

                let data = {'RoleName' : this.RoleName};
                await this.$axios.get(this.$URLs.SANCTUM_CSRF)
                await this.$axios({
                    url: this.createEditActionUrl,
                    method: "PUT",
                    data: data
                }).then(response => {
                    this.$emit('reload');

                    this.resetDialog()

                    this.$emit('close')

                    this.$store.dispatch('showSnackbarMessage',{message: response.data.message, code : '1'})

                    this.$store.dispatch('showProgress', false);
                }).catch(e => {
                    this.$store.dispatch('showProgress', false)

                    this.serverValidationErrors = []

                    if (typeof e.response != 'undefined' && e.response.status == 422) {
                        if (typeof e.response.data.errors.RoleName != 'undefined') {
                            e.response.data.errors.RoleName.forEach((value, index, array) => {
                                this.serverValidationErrors.push(value);
                            });
                        }
                    } else {
                        this.$store.dispatch('serverError', e)
                    }
                });
            }

            return false;
        },

        resetDialog() {
            this.RoleName = ''
            this.$v.$reset()
        }
    },

    validations: {
        RoleName: {
            required,
            maxLength: maxLength(20)
        }
    },

    components: {
        'server-validation-messages' : ServerValidationMessages,
        'validation-error-alert': ValidationErrorAlert
    },

    created() {
        
    }

}
</script>
