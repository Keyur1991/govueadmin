<template>
    <v-container fluid >
        <v-row dense>
            <v-dialog v-model="dialog" max-width="600px">
                <v-card>
                    <v-card-title class="headline primary lighten-2 pa-3" primary-title>
                        <span class="headline white--text">{{ title }}</span>
                    </v-card-title>
                    <v-form :action="createEditActionUrl" method="post" @submit.prevent="createOrEditFeature()">
                        <v-card-text>
                            <v-row dense wrap v-if="containsErrors">
                                <v-col>
                                    <validation-error-alert :contains-errors="containsErrors"></validation-error-alert>
                                </v-col>
                            </v-row>
                            <v-row dense>
                                <v-col>
                                    <v-text-field outlined dense label="Module*"
                                        v-model="ModuleName"
                                        :error-messages="moduleErrors" >
                                    </v-text-field>
                                </v-col>
                            </v-row>
                            <v-row dense>
                                <v-col>
                                    <v-text-field outlined dense label="Feature*"
                                        v-model="FeatureName"
                                        :error-messages="featureErrors" >
                                    </v-text-field>
                                </v-col>
                            </v-row>
                            <v-row dense>
                                <v-col>
                                    <v-textarea
                                        outlined dense
                                        v-model="FeatureDescription"
                                        auto-grow
                                        label="Description*"
                                        rows="2"
                                        :error-messages="descriptionErrors"
                                    ></v-textarea>
                                </v-col>
                            </v-row>
                        </v-card-text>
                        <v-card-actions>
                            <v-btn color="primary" type="submit">Save</v-btn>
                            <v-spacer></v-spacer>
                            <v-btn  type="button" text @click="closeModal()">Cancel</v-btn>
                        </v-card-actions>
                    </v-form>
                </v-card>
            </v-dialog>
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

var ValidationErrorAlert = require('../ValidationErrorAlert.vue').default

export default {
    props: ['featureItem', 'title', 'edit', 'featureItemModal'],
    data() {
        return {
            ModuleName: '',
            FeatureName: '',
            RouteName: '',
            FeatureDescription: ''
        }
    },
    computed: {
        dialog() {
            if (this.edit) {
                this.ModuleName = this.featureItem.module
                this.FeatureName = this.featureItem.feature_name
                this.RouteName = this.featureItem.route_name
                this.FeatureDescription = this.featureItem.description
            } else {
                this.ModuleName = ''
                this.FeatureName = ''
                this.RouteName = ''
                this.FeatureDescription = ''
            }

            return this.featureItemModal
        },
        createEditActionUrl() {
            if (this.edit == true) {
                return this.$URLs.FEATURE_LIST + '/' + this.featureItem.id
            } else {
                return this.$URLs.FEATURE_LIST
            }
        },

        moduleErrors() {
            const errors = [];

            if (!this.$v.FeatureName.$dirty) return errors;

            !this.$v.FeatureName.required && errors.push("Please enter the module name.");

            !this.$v.ModuleName.maxLength &&
                errors.push("Module name must not be longer than 255 characters.");

            return errors;
        },

        featureErrors() {
            const errors = [];

            if (!this.$v.FeatureName.$dirty) return errors;

            !this.$v.FeatureName.required && errors.push("Please enter the feature name.");

            !this.$v.FeatureName.maxLength &&
                errors.push("Feature name must not be longer than 255 characters.");

            return errors;
        },

        descriptionErrors() {
            const errors = [];

            if (!this.$v.FeatureDescription.$dirty) return errors;

            !this.$v.FeatureDescription.required && errors.push("Please enter the description.");

            !this.$v.FeatureDescription.maxLength &&
                errors.push("Description must not be longer than 500 characters.");

            return errors;
        },

        containsErrors() {
            return this.moduleErrors.length > 0 || this.featureErrors.length > 0
                || this.descriptionErrors.length > 0;
        }
    },
    methods: {
        async createOrEditFeature() {
            if (this.$Utils.isValidForm.call(this)) {
                this.$store.dispatch('showProgress', true)
                await this.$axios.get(this.$URLs.SANCTUM_CSRF)
                let data = {
                    'Module' : this.ModuleName,
                    'FeatureName': this.FeatureName,
                    'RouteName': this.RouteName,
                    'FeatureDescription': this.FeatureDescription
                };

                await this.$axios({
                    url: this.createEditActionUrl,
                    method: "PUT",
                    data: data
                }).then(response => {
                    this.closeModal()

                    this.$store.dispatch('showSnackbarMessage', {message: response.data.message, code: '1'})

                    this.$store.dispatch('showProgress', false);

                    this.$emit('reload');

                }).catch(e => {
                    this.$store.dispatch('showProgress', false)
                    this.$store.dispatch('serverError', e)
                });
            }

            return false;
        },

        resetDialog() {
            this.ModuleName = ''
            this.FeatureName = ''
            this.RouteName = ''
            this.$v.$reset()
        },

        closeModal() {
            this.resetDialog()
            this.$emit('close')
        }
    },

    validations: {
        ModuleName: {
            required,
            maxLength: maxLength(255)
        },
        FeatureName: {
            required,
            maxLength: maxLength(255)
        },
        
        FeatureDescription: {
            required,
            maxLength: maxLength(500)
        }
    },

    created() {

    },

    components: {
        'validation-error-alert': ValidationErrorAlert
    }

}
</script>
