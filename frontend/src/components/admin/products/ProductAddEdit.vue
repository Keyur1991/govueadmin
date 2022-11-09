<template>
    <v-container fluid >
        <v-row dense>
            <v-col>
                <v-dialog v-model="dialog" max-width="600px">
                    <v-card>
                        <v-card-title class="headline primary lighten-2 pa-3" primary-title>
                            <span class="headline white--text">{{ title }}</span>
                        </v-card-title>

                        <v-form :action="createEditActionUrl" method="post" ref="createEditProductForm" @submit.prevent="createOrEditProduct()">
                            <v-card-text>
                                <v-row dense wrap v-if="containsErrors || serverValidationErrors.length > 0">
                                    <v-col>
                                        <validation-error-alert :contains-errors="containsErrors"></validation-error-alert>
                                        <server-validation-messages :validation-messages="serverValidationErrors"></server-validation-messages>
                                    </v-col>
                                </v-row>

                                
                                    <v-row dense>
                                    <v-col>
                                        <v-text-field outlined dense :label="`${$vuetify.lang.t('$vuetify.Products.Fields.Name1')}*`"
                                            v-model="Name1" :error-messages="Name1Errors">
                                        </v-text-field>
                                    </v-col>
                                </v-row>
                                <v-row dense>
                                    <v-col>
                                        <v-text-field outlined dense :label="`${$vuetify.lang.t('$vuetify.Products.Fields.Name2')}*`"
                                            v-model="Name2" :error-messages="Name2Errors">
                                        </v-text-field>
                                    </v-col>
                                </v-row>
                                <v-row dense>
                                    <v-col>
                                        <v-text-field outlined dense :label="`${$vuetify.lang.t('$vuetify.Products.Fields.Price')}*`"
                                            v-model="Price" :error-messages="PriceErrors">
                                        </v-text-field>
                                    </v-col>
                                </v-row>
                                <v-row dense>
                                    <v-col>
                                        <v-text-field outlined dense :label="`${$vuetify.lang.t('$vuetify.Products.Fields.CategoryId')}*`"
                                            v-model="CategoryId" :error-messages="CategoryIdErrors">
                                        </v-text-field>
                                    </v-col>
                                </v-row>

                                
                            </v-card-text>
                            <v-card-actions>
                                <v-btn color="primary" type="submit">{{ $vuetify.lang.t('$vuetify.SaveBtn') }}</v-btn>
                                <v-spacer></v-spacer>
                                <v-btn  type="button" text @click="dialog = false">{{ $vuetify.lang.t('$vuetify.CancelBtn') }}</v-btn>
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
  email,
  maxLength,
  minLength,
  alpha,
  alphaNum,
  numeric
} from "vuelidate/lib/validators";

var ServerValidationMessages = require('../ServerValidationMessages.vue').default

var ValidationErrorAlert = require('../ValidationErrorAlert.vue').default
export default {
    props: ['productItem', 'title', 'edit', 'productItemModal'],
    data() {
        return {
            Name1 : '',
            Name2 : '',
            Price : '',
            CategoryId : '',

            serverValidationErrors: []
        }
    },
    computed: {
        dialog: {
            set(value) {
                this.resetDialog()

                this.$emit('close')
            },

            get() {
                if (this.edit) {
                this.Name1 = this.productItem.name1
                this.Name2 = this.productItem.name2
                this.Price = this.productItem.price
                this.CategoryId = this.productItem.category_id

                } else {
                this.Name1 = ''
                this.Name2 = ''
                this.Price = ''
                this.CategoryId = ''

                }
                this.serverValidationErrors = []
                return this.productItemModal
            }
        },
        ShowProgress() {
            return this.$store.state.ShowProgress
        },
        createEditActionUrl() {
            if (this.edit == true) {
                return this.$URLs.PRODUCTS_LIST + '/' + this.productItem.id
            } else {
                return this.$URLs.PRODUCTS_LIST
            }
        },

Name1Errors() {
    const errors = [];

    if (this.$v.Name1 && !this.$v.Name1.$dirty) return errors;

    !this.$v.Name1.required && errors.push("Required.");
!this.$v.Name1.minLength && errors.push("Min. Length: 3.");
!this.$v.Name1.maxLength && errors.push("Max. Length: 255.");


    return errors;
},Name2Errors() {
    const errors = [];

    if (this.$v.Name2 && !this.$v.Name2.$dirty) return errors;

    

    return errors;
},PriceErrors() {
    const errors = [];

    if (this.$v.Price && !this.$v.Price.$dirty) return errors;

    !this.$v.Price.required && errors.push("Required.");


    return errors;
},CategoryIdErrors() {
    const errors = [];

    if (this.$v.CategoryId && !this.$v.CategoryId.$dirty) return errors;

    

    return errors;
},

        containsErrors() {
            return this.Name1Errors.length > 0 || this.Name2Errors.length > 0 || this.PriceErrors.length > 0 || this.CategoryIdErrors.length > 0;
        }
    },
    methods: {
        createOrEditProduct() {
            if (this.$Utils.isValidForm.call(this)) {
                this.$store.dispatch('showProgress', true)

                let data = {'Name1' : this.Name1,
'Name2' : this.Name2,
'Price' : this.Price,
'CategoryId' : this.CategoryId,
};

                this.$axios({
                    url: this.createEditActionUrl,
                    method: "PUT",
                    data: data
                }).then(response => {
                    this.dialog = false;

                    this.$store.dispatch('showSnackbarMessage',{message: response.data.message, code : '1'})

                    this.$store.dispatch('showProgress', false);

                    this.$emit('reload');
                }).catch(e => {
                    this.$store.dispatch('showProgress', false)

                    this.serverValidationErrors = []

                    if (typeof e.response != 'undefined' && e.response.status == 422) {
                        console.log(e)
                    } else {
                        this.$store.dispatch('serverError', e)
                    }
                });
            }

            return false;
        },

        resetDialog() {
            this.$v.$reset()
        }
    },

    validations: {
Name1: {required,
minLength: minLength(3),
maxLength: maxLength(255),
},

Name2: {},

Price: {required,
},

CategoryId: {},


    },

    components: {
        'server-validation-messages' : ServerValidationMessages,
        'validation-error-alert': ValidationErrorAlert
    }

}
</script>
