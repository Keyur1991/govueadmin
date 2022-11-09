<template>
    <v-container fluid >
        <v-row dense>
            <v-col>
                <v-dialog v-model="dialog" max-width="600px">
                    <v-card>
                        <v-card-title class="headline primary lighten-2 pa-3" primary-title>
                            <span class="headline white--text">{{ modalTitle }}</span>
                        </v-card-title>

                        <v-form :action="addEditActionUrl" method="post"  @submit.prevent="saveMenu()">
                            <v-card-text>
                                <v-row dense v-if="containsErrors || serverValidationErrors.length > 0">
                                    <v-col>
                                        <validation-error-alert :contains-errors="containsErrors"></validation-error-alert>
                                        <server-validation-messages :validation-messages="serverValidationErrors"></server-validation-messages>
                                    </v-col>
                                </v-row>

                                <v-row dense>
                                    <v-col md="6" sm="6" xs="12">
                                        <v-text-field outlined dense label="Name*"
                                            v-model="MenuName"
                                            :error-messages="nameErrors" >
                                        </v-text-field>
                                    </v-col>

                                    <v-col md="6" sm="6" xs="12">
                                        <v-text-field outlined dense label="URL*"
                                            v-model="MenuURL"
                                            hint="e.g. users, users/create"
                                            :error-messages="urlErrors" >
                                        </v-text-field>
                                    </v-col>
                                </v-row>

                                <v-row dense>
                                    <v-col md="6" sm="6" xs="12">
                                        <v-text-field outlined dense label="Icon"
                                            v-model="MenuIcon"
                                            hint="e.g. fa-user">
                                        </v-text-field>
                                    </v-col>

                                    <v-col md="6" sm="6" xs="12">
                                        <v-autocomplete outlined dense multiple  hint="Type to search" clearable search-input v-model="MenuRoles" label="Roles*" :items="rolesList" :error-messages="menuRolesErrors" item-text="name" item-value="id">
                                        </v-autocomplete>
                                    </v-col>
                                </v-row>

                                <v-row dense>
                                    <v-col md="6" sm="6" xs="12">
                                        <v-text-field outlined dense label="Custom Class"
                                            v-model="MenuCustomClass"
                                            hint="e.g. abc xyz">
                                        </v-text-field>
                                    </v-col>
                                    <v-col md="6" sm="6" xs="12">
                                        <v-switch
                                        v-model="MenuTargetBlank"
                                        label="Open in New Tab"
                                        ></v-switch>
                                    </v-col>
                                </v-row>
                            </v-card-text>
                            <v-card-actions>
                                <v-btn color="primary" type="submit">Save</v-btn>
                                <v-spacer></v-spacer>
                                <v-btn type="button" text @click="$emit('close')">Cancel</v-btn>
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

var ServerValidationMessages = require('./ServerValidationMessages.vue').default

var ValidationErrorAlert = require('./ValidationErrorAlert.vue').default
export default {
    props: ['menuItem', 'menuItemModal',  'modalTitle'],
    data() {
        return {
            serverValidationErrors: [],
            MenuName: null,
            MenuURL: null,
            MenuIcon: null,
            MenuCustomClass: null,
            MenuTargetBlank: null,
            MenuRoles: []
        }
    },
    computed: {
        dialog() {
            this.serverValidationErrors = []
            this.MenuName = this.menuItem != null ? this.menuItem.title : ''
            this.MenuURL = this.menuItem != null ? this.menuItem.url : ''
            this.MenuIcon = this.menuItem != null ? this.menuItem.icon : ''
            this.MenuCustomClass = this.menuItem != null ? this.menuItem.custom_class : ''
            this.MenuTargetBlank = (this.menuItem != null && this.menuItem.target_blank == '1') ? true : false
            this.MenuRoles = (this.menuItem != null && this.menuItem.rolesList.length > 0) ?  this.menuItem.rolesList : []
            return this.menuItemModal
        },
        addEditActionUrl() {
            let action = this.$URLs.MENUS_LIST

            if (this.menuItem != null) action += '/' + this.menuItem.id

            return action
        },

        nameErrors() {
            const errors = [];

            if (!this.$v.MenuName.$dirty) return errors;

            !this.$v.MenuName.required && errors.push("Please enter the menu name.");

            !this.$v.MenuName.maxLength &&
                errors.push("Menu name must not be longer than 20 characters.");

            return errors;
        },

        urlErrors() {
            const errors = [];

            if (!this.$v.MenuURL.$dirty) return errors;

            !this.$v.MenuURL.required && errors.push("Please enter the menu url.");

            !this.$v.MenuURL.maxLength &&
                errors.push("Menu url must not be longer than 100 characters.");

            return errors;
        },

        containsErrors() {
            return this.nameErrors.length > 0 || this.urlErrors.length > 0 || this.menuRolesErrors.length > 0;
        },

        rolesList() {
            return this.$store.state.RolesMaster
        },

        menuRolesErrors() {
            const errors = [];

            if (!this.$v.MenuRoles.$dirty) return errors;

            !this.$v.MenuRoles.required && errors.push("Please select atleast one role.");

            return errors;
        }
    },
    methods: {
        async saveMenu() {
            if (this.$Utils.isValidForm.call(this)) {
                this.$store.dispatch('showProgress', true)

                let data = {
                    'MenuName' : this.MenuName,
                    'MenuURL' : this.MenuURL,
                    'MenuIcon' : this.MenuIcon,
                    'MenuCustomClass' : this.MenuCustomClass,
                    'MenuTargetBlank' : this.MenuTargetBlank,
                    'MenuRoles': this.MenuRoles
                };
                await this.$axios.get(this.$URLs.SANCTUM_CSRF)
                await this.$axios({
                    url: this.addEditActionUrl,
                    method: "PUT",
                    data: data
                }).then(response => {
                    this.$store.dispatch('showProgress', false);
                    this.resetDialog()
                    this.$emit('close')
                }).catch(e => {
                    this.$store.dispatch('showProgress', false)
                    this.serverValidationErrors = []
                    this.$store.dispatch('serverError', e)
                });
            }

            return false;
        },

        resetDialog() {
            this.$v.$reset()
        }
    },

    validations: {
        MenuName: {
            required,
            maxLength: maxLength(20)
        },
        MenuURL: {
            required,
            maxLength: maxLength(100)
        },
        MenuRoles: {
            required
        }
    },

    components: {
        'server-validation-messages' : ServerValidationMessages,
        'validation-error-alert': ValidationErrorAlert
    }

}
</script>
