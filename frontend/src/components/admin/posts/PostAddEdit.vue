<template>
    <v-container fluid >
        <v-row dense>
            <v-col>
                <v-dialog v-model="dialog" max-width="600px">
                    <v-card>
                        <v-card-title class="headline primary lighten-2 pa-3" primary-title>
                            <span class="headline white--text">{{ title }}</span>
                        </v-card-title>

                        <v-form :action="createEditActionUrl" method="post" ref="createEditPostForm" @submit.prevent="createOrEditPost()">
                            <v-card-text>
                                <v-row dense wrap v-if="containsErrors || serverValidationErrors.length > 0">
                                    <v-col>
                                        <validation-error-alert :contains-errors="containsErrors"></validation-error-alert>
                                        <server-validation-messages :validation-messages="serverValidationErrors"></server-validation-messages>
                                    </v-col>
                                </v-row>

                                
                                    <v-row dense>
                                    <v-col>
                                        <v-text-field outlined dense :label="`${$vuetify.lang.t('$vuetify.Posts.Fields.Title')}*`"
                                            v-model="Title" :error-messages="TitleErrors">
                                        </v-text-field>
                                    </v-col>
                                </v-row>
                                <v-row dense>
                                    <v-col>
                                        <v-text-field outlined dense :label="`${$vuetify.lang.t('$vuetify.Posts.Fields.Description')}*`"
                                            v-model="Description" :error-messages="DescriptionErrors">
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
    props: ['postItem', 'title', 'edit', 'postItemModal'],
    data() {
        return {
            Title : '',
            Description : '',

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
                this.Title = this.postItem.title
                this.Description = this.postItem.description

                } else {
                this.Title = ''
                this.Description = ''

                }
                this.serverValidationErrors = []
                return this.postItemModal
            }
        },
        ShowProgress() {
            return this.$store.state.ShowProgress
        },
        createEditActionUrl() {
            if (this.edit == true) {
                return this.$URLs.POSTS_LIST + '/' + this.postItem.id
            } else {
                return this.$URLs.POSTS_LIST
            }
        },

TitleErrors() {
    const errors = [];

    if (this.$v.Title && !this.$v.Title.$dirty) return errors;

    !this.$v.Title.required && errors.push("Required.");
!this.$v.Title.minLength && errors.push("Min. Length: 3.");
!this.$v.Title.maxLength && errors.push("Max. Length: 255.");
!this.$v.Title.alphaNum && errors.push("Only alphabets & numbers allowed.");


    return errors;
},DescriptionErrors() {
    const errors = [];

    if (this.$v.Description && !this.$v.Description.$dirty) return errors;

    !this.$v.Description.required && errors.push("Required.");
!this.$v.Description.minLength && errors.push("Min. Length: 3.");
!this.$v.Description.maxLength && errors.push("Max. Length: 255.");


    return errors;
},

        containsErrors() {
            return this.TitleErrors.length > 0 || this.DescriptionErrors.length > 0;
        }
    },
    methods: {
        createOrEditPost() {
            if (this.$Utils.isValidForm.call(this)) {
                this.$store.dispatch('showProgress', true)

                let data = {'Title' : this.Title,
'Description' : this.Description,
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
Title: {required,
minLength: minLength(3),
maxLength: maxLength(255),
alphaNum,
},

Description: {required,
minLength: minLength(3),
maxLength: maxLength(255),
},


    },

    components: {
        'server-validation-messages' : ServerValidationMessages,
        'validation-error-alert': ValidationErrorAlert
    }

}
</script>
