<template>
    <v-container fluid v-if="hasLoggedIn">
        <v-row dense justify="space-between" >
            <v-col md="9" sm="9" xs="12">
				<page-title>{{ (OperationMode == 'Edit') ? $vuetify.lang.t('$vuetify.Users.Listing.EditUser') : $vuetify.lang.t('$vuetify.Users.Listing.NewBtn') }}</page-title>
			</v-col>
			<v-col md="3" sm="3" xs="12" >
				<breadcrumbs :items="breadcrumbs" />
			</v-col>
        </v-row>

        <v-row dense v-if="MakeFormVisible && (hasAddAccess || hasEditAccess)">
            <v-col sm="7" md="7" xs="12">
                <v-card class="elevation-2 pa-3 mt-3">
                    <validation-error-alert :contains-errors="containsErrors"></validation-error-alert>
                    <server-validation-messages :validation-messages="serverValidationErrors"></server-validation-messages>
                    <v-form :action="createEditActionUrl" method="post" ref="newUserForm" @submit.prevent="createOrEditUser">
                        <v-row dense>
                            <v-col>
                                <v-select
                                    outlined
                                    dense
                                    v-model="UserRoles"
                                    :items="rolesList"
                                    item-text="name"
                                    item-value="id"
                                    attach
                                    chips
                                    clearable
                                    :label="`${$vuetify.lang.t('$vuetify.Roles.Listing.PageTitle')}*`"
                                    :error-messages="userRolesErrors"
                                    multiple
                                ></v-select>
                            </v-col>
                        </v-row>
                        
                        <v-row dense>
                            <v-col>
                                <v-text-field
                                    outlined
                                    dense
                                    v-model="Name"
                                    :error-messages="nameErrors"
                                    :label="`${$vuetify.lang.t('$vuetify.Users.Fields.Name')}*`"
                                ></v-text-field>
                            </v-col>
                        </v-row>

                        <v-row dense>
                            <v-col>
                                <v-text-field
                                    outlined
                                    dense
                                    v-model="Email"
                                    :error-messages="emailErrors"
                                    :label="`${$vuetify.lang.t('$vuetify.Users.Fields.Email')}*`"
                                ></v-text-field>
                            </v-col>
                        </v-row>

                        <v-row dense>
                            <v-col>
                                <v-text-field
                                    outlined
                                    dense
                                    name="password"
                                    id="password"
                                    :label="`${$vuetify.lang.t('$vuetify.Users.Fields.Password')}*`"
                                    type="password"
                                    v-model="Password"
                                    :error-messages="passwordErrors">
                                </v-text-field>
                            </v-col>
                        </v-row>

                        <v-row dense>
                            <v-col>
                                <v-text-field
                                    outlined
                                    dense
                                    name="password_confirmation"
                                    id="password-confirm"
                                    :label="`${$vuetify.lang.t('$vuetify.Users.Fields.ConfirmPassword')}*`"
                                    type="password"
                                    v-model="ConfirmPassword"
                                    :error-messages="confirmPasswordErrors">
                                </v-text-field>
                            </v-col>
                        </v-row>

                        <v-row dense>
                            <v-col md="8" sm="8" xs="12">
                                <p class="mb-1">{{ $vuetify.lang.t('$vuetify.Users.Fields.SelectProfilePicture') }}</p>
                                <v-btn
                                    color="primary"
                                    outlined
                                    small
                                    @click="triggerFileInput">
                                    {{ (imageName != null) ? imageName : 'Browse' }}
                                    <v-icon right small>fa-folder</v-icon>
                                </v-btn>

                                <input type="file"
                                    v-on:change="handleFileSelection"
                                    ref="ProfilePicture"
                                    style="display: none;"
                                    name="ProfilePicture"
                                />
                                <p class="red--text" transition="scale-transition" v-for="(item, index) in ProfilePictureErrors" :key="index">{{ item }}</p>
                            </v-col>
                            <v-col md="4" sm="4" xs="12">
                                <v-img v-if="imageUrl != null"
                                    :src="imageUrl"
                                    width="150"
                                    height="150"
                                    class="elevation-6"
                                    >
                                </v-img>
                                <v-btn v-if="imageUrl != null"

                                    color="red"
                                    outlined
                                    small
                                    @click="resetFileSelection">
                                    Remove
                                    <v-icon right small>fa-times</v-icon>
                                </v-btn>
                            </v-col>
                        </v-row>
                        
                        <v-row dense justify="space-between">
                            <v-col md="10" sm="10">
                                <v-btn color="primary" type="submit">{{ $vuetify.lang.t('$vuetify.SaveBtn') }}</v-btn>
                            </v-col>
                            <v-col md="2" sm="2">
                                <v-btn @click="clear"  type="button" text>{{ $vuetify.lang.t('$vuetify.ResetBtn') }}</v-btn>
                            </v-col>
                        </v-row>
                    </v-form>
                </v-card>
            </v-col>
            <v-col sm="5" md="5" xs="12">
                <v-card class="elevation-2 pa-2 mt-3">
                    <v-card-text>
                        <p><v-icon small>fa-dot-circle</v-icon> {{ $vuetify.lang.t('$vuetify.Users.Validations.AllFieldsRequired') }}</p>
                        <p><v-icon small>fa-dot-circle</v-icon> {{ $vuetify.lang.t('$vuetify.Users.Validations.PasswordValidationRules') }}</p>
                        <p><v-icon small>fa-dot-circle</v-icon> {{ $vuetify.lang.t('$vuetify.Users.Validations.ProfilePictureValidationRules') }}</p>
                        <p v-if="PreviewCurrentImage != null">
                            <v-layout row wrap pa-3>
                                <v-flex xs12 sm8 md8 >
                                    <v-img
                                        :src="PreviewCurrentImage"
                                        width="200"
                                        height="200"
                                        class="elevation-2"
                                        :lazy-src="LazyLoader"
                                    >
                                    </v-img>
                                </v-flex>
                                <v-flex xs12 sm4 md4 v-if="EligibleForUnlink">
                                    <v-btn color="red"
                                        outlined
                                        small
                                        @click="removeCurrentImage"
                                        >
                                        {{ $vuetify.lang.t('$vuetify.DeleteBtn') }}
                                        <v-icon right small>fa-trash</v-icon>
                                    </v-btn>
                                </v-flex>
                            </v-layout>
                        </p>
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>
        <v-row dense>
            <v-col class="pa-2" md12 v-if="(hasAddAccess != null && hasAddAccess) || (hasEditAccess != null && !hasEditAccess)">
                <unauthorized :display="hasAddAccess || hasEditAccess"></unauthorized>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
import Vue from 'vue'
import Vuelidate from 'vuelidate'

Vue.use(Vuelidate)

import {
  required,
  email,
  minLength,
  maxLength,
  sameAs,
  requiredIf,
  and
} from "vuelidate/lib/validators";

const StrictPassword = function(value) {
    let regx = new RegExp(/^(?=.*[a-zA-Z])(?=.*\d)(?=.*[!@#$%^&*()_+])[A-Za-z\d][A-Za-z\d!@#$%^&*()_+]{7,19}$/);

    return regx.test(value)
}

var Unauthorized = require("../Unauthorized.vue").default;
var ValidationErrorAlert = require('../ValidationErrorAlert.vue').default
var ServerValidationMessages = require('../ServerValidationMessages.vue').default
var Breadcrumbs = require('../Breadcrumb.vue').default
var PageTitle = require('../PageTitle.vue').default

export default {
    props: ['OperationMode'],
    data() {
        return {
            createEditActionUrl: this.$URLs.USER_LIST,
            Name: '',
            Email: '',
            Password: '',
            ConfirmPassword: '',
            ProfilePicture: null,
            imageUrl: null,
            imageName: null,
            ProfilePictureErrors: [],
            EditFormReady: false,
            PreviewCurrentImage: null,
            disableEmailField: false,
            User: null,
            EligibleForValidation: this.OperationMode == 'Create',
            UserRoles: [],
            hasAddAccess: null,
            hasEditAccess: null,
            serverValidationErrors: []
        }
    },

    computed: {
        LazyLoader() {
            return this.$store.state.LazyLoader
        },

        breadcrumbs() { 
            return [
                {
                    text: this.$vuetify.lang.t('$vuetify.Home'),
                    action: '/admin/dashboard',
                    disabled: false
                },
                {
                    text: this.$vuetify.lang.t('$vuetify.Users.Breadcrumbs.Users'),
                    action: '/admin/users',
                    disabled: false
                },
                {
                    text: this.$vuetify.lang.t('$vuetify.Users.Breadcrumbs.Create'),
                    action: '/admin/users/create',
                    disabled: true
                }
            ]
        },

        PasswordCaption() {
            return this.$vuetify.lang.t('$vuetify.Users.Fields.Password') + '*';
        },

        ConfirmPasswordCaption() {
            return this.$vuetify.lang.t('$vuetify.Users.Fields.ConfirmPassword') + '*';
        },
            
        nameErrors() {
            const errors = [];

            if (!this.$v.Name.$dirty) return errors;

            !this.$v.Name.required && errors.push(this.$vuetify.lang.t('$vuetify.Users.Validations.NameRequired'));

            !this.$v.Name.maxLength && errors.push(this.$vuetify.lang.t('$vuetify.Users.Validations.NameMaxLength'));

            return errors;
        },

        emailErrors() {
            const errors = [];

            if (!this.$v.Email.$dirty) return errors;

            !this.$v.Email.required && errors.push(this.$vuetify.lang.t('$vuetify.Users.Validations.EmailRequired'));

            !this.$v.Email.email && errors.push(this.$vuetify.lang.t('$vuetify.Users.Validations.EmailInvalid'));

            return errors;
        },

        passwordErrors() {
            const errors = [];

            if (!this.$v.Password.$dirty) return errors;

            !this.$v.Password.required && errors.push(this.$vuetify.lang.t('$vuetify.Users.Validations.PasswordRequired'));

            !this.$v.Password.minLength && errors.push(`${this.$vuetify.lang.t('$vuetify.Users.Validations.PasswordText1')} ${this.$v.Password.$params.minLength.min} ${this.$vuetify.lang.t('$vuetify.Users.Validations.PasswordText2')}`);

            !this.$v.Password.maxLength &&
                errors.push(`${this.$vuetify.lang.t('$vuetify.Users.Validations.PasswordText3')} ${this.$v.Password.$params.maxLength.max} ${this.$vuetify.lang.t('$vuetify.Users.Validations.PasswordText4')}`);

            this.Password != '' && !this.$v.Password.StrictPassword && errors.push(this.$vuetify.lang.t('$vuetify.Users.Validations.PasswordValidationRules'));
            return errors;
        },

        confirmPasswordErrors() {
            const errors = [];

            if (!this.$v.ConfirmPassword.$dirty) return errors;

            !this.$v.ConfirmPassword.required && errors.push(this.$vuetify.lang.t('$vuetify.Users.Validations.PasswordRequired'));

            !this.$v.ConfirmPassword.sameAsPassword && errors.push(this.$vuetify.lang.t('$vuetify.Users.Validations.PasswordNotMatch'));

            return errors;
        },

        containsErrors() {
            return (
                this.nameErrors.length > 0 ||
                this.emailErrors.length > 0 ||
                this.passwordErrors.length > 0 ||
                this.confirmPasswordErrors.length > 0 ||
                this.userRolesErrors.length > 0
            );
        },
        MakeFormVisible() {
            return this.OperationMode == 'Create' || this.EditFormReady
        },

        EligibleForUnlink() {
            return this.$Utils.eligibleForUnlink.call(this)
        },

        rolesList() {
            return this.$store.state.RolesMaster
        },

        userRolesErrors() {
            const errors = [];

            if (!this.$v.UserRoles.$dirty) return errors;

            !this.$v.UserRoles.required && errors.push(this.$vuetify.lang.t('$vuetify.Users.Validations.RoleRequired'));

            return errors
        },

		hasLoggedIn() {
			return this.$store.state.userHasLoggedIn
		}
    },

    validations: {
        Name: {
            required,
            maxLength: maxLength(20)
        },
        Email: {
            required: requiredIf((vueInstance) => {
                return vueInstance.EligibleForValidation
            }),
            email
        },
        Password: {
            required: requiredIf((vueInstance) => {
                return vueInstance.EligibleForValidation
            }),
            minLength: minLength(6),
            maxLength: maxLength(15),
            StrictPassword (value) {
                return and('required', value => StrictPassword(value))
            }
        },
        ConfirmPassword: {
            required: requiredIf((vueInstance) => {
                return vueInstance.EligibleForValidation
            }),
            sameAsPassword: sameAs('Password')
        },

        UserRoles: {
            required
        }

    },

    methods: {
        async createOrEditUser() {
            if (this.$Utils.isValidForm.call(this)) {
                window.scrollTo(0,0);
                
                this.$store.dispatch('showProgress', true)
                await this.$axios.get(this.$URLs.SANCTUM_CSRF)
                let data = new FormData();

                data.append('Name', this.Name)

                data.append('Email', this.Email)

                if (this.Password != '') {
                    data.append('Password', this.Password)
                }

                if (this.ProfilePicture != null) {
                    data.append('ProfilePicture', this.ProfilePicture)
                }

                data.append('UserRoles', this.UserRoles)

                data.append('_method', 'PUT')

                await this.$axios({
                    url: this.createEditActionUrl,
                    method: "POST",
                    data: data
                }).then(response => {
                    this.$store.dispatch('showProgress', false)

                    this.$store.dispatch('showSnackbarMessage',{message: response.data.message, code : '1'})

                    this.clear()

                    this.loadUserEditData()
                }).catch(e => {
                    this.$store.dispatch('showProgress', false)

                    this.serverValidationErrors = [];

                    if (typeof e.response != 'undefined' && e.response.status == 422) {
                        if (typeof e.response.data.errors.Name != 'undefined') {
                            e.response.data.errors.Name.forEach((value, index, array) => {
                                this.serverValidationErrors.push(value);
                            });
                        }

                        if (typeof e.response.data.errors.Email != 'undefined') {
                            e.response.data.errors.Email.forEach((value, index, array) => {
                                this.serverValidationErrors.push(value);
                            });
                        }

                        if (typeof e.response.data.errors.Password != 'undefined') {
                            e.response.data.errors.Password.forEach((value, index, array) => {
                                this.serverValidationErrors.push(value);
                            });
                        }
                    } else {
                        this.$store.dispatch('serverError', e)
                    }
                });
            } else {
                window.scrollTo(0,0);
            }

            return false;
        },

        clear() {
            this.$v.$reset();
            this.resetFileSelection()
            this.ProfilePictureErrors = []

            if (this.OperationMode == 'Create') {
                this.$refs.newUserForm.reset()
            } else {
                this.setData()
            }
        },

        triggerFileInput() {
            this.$refs.ProfilePicture.click()
        },

        handleFileSelection(e) {
            this.$Utils.handleFileSelection.call(this, e)
        },

        resetFileSelection() {
            this.imageName= null
            this.imageUrl = null
            this.ProfilePicture = null
            this.PreviewCurrentImage = null
            
        },

        loadUserEditData() {
            if (this.OperationMode == 'Edit') {
                this.$store.dispatch('showProgress', true)

                let action = this.$URLs.USER_LIST + '/' + this.$route.params.id
                this.$axios
                    .get(action)
                    .then(res => {
                        let d = res.data.data;

                        this.User = d
                        this.setData()

                        this.$store.dispatch('showProgress', false)
                    })
                    .catch(eres => {
                        this.$store.dispatch('showProgress', false)
                        this.$store.dispatch('serverError', eres)
                    });
            }
        },

        setData() {

            if (this.OperationMode == 'Edit') {
                this.Name = this.User.name;
                this.Email = this.User.email;
                this.PreviewCurrentImage = this.User.profile_pic;
                this.EditFormReady = true
                this.PasswordCaption = 'Password ';
                this.ConfirmPasswordCaption = 'Confirm Password ';
                this.createEditActionUrl = this.$URLs.USER_LIST + '/' + this.$route.params.id
                this.disableEmailField = true
                this.Password = this.ConfirmPassword = ''
                this.UserRoles = this.User.rolesList
            }
        },

        async removeCurrentImage() {
            this.$store.dispatch('showProgress', true)

            await this.$axios.get(this.$URLs.SANCTUM_CSRF)
            let params = {
                '_method': 'DELETE'
            };

            await this.$axios({
                url: this.$URLs.USER_PICTURE_REMOVE + "/" + this.$route.params.id,
                method: "DELETE",
                data: params
            }).then(response => {
                this.$store.dispatch('showProgress', false)

                this.$store.dispatch('showSnackbarMessage',{message: response.data.message, code: '1'})

                this.loadUserEditData()
            }).catch(e => {
                this.$store.dispatch('showProgress', false)
                this.$store.dispatch('serverError', e)
            });
        },

        getAccessDetails() {
            this.$store.dispatch("showProgress", true);

            this.$axios
                .get(this.$URLs.USER_ACCESS)
                .then(response => {
                    this.$store.dispatch("showProgress", false);
                    this.hasAddAccess = response.data.data.canAdd;
                    this.hasEditAccess = response.data.data.canEdit;
                }).catch(e => {
                    this.$store.dispatch("serverError", e);
                    this.$store.dispatch("showProgress", false);
                });
        }
    },

    components: {
        'unauthorized' : Unauthorized,
        'validation-error-alert': ValidationErrorAlert,
        'server-validation-messages' : ServerValidationMessages,
		"breadcrumbs": Breadcrumbs,
		"page-title": PageTitle
    },

    async created() {
        await this.$axios.get(this.$URLs.SANCTUM_CSRF)
        await this.$Utils.checkUserLoggedIn.call(this)
        await this.$store.dispatch('rolesList')
        await this.loadUserEditData();
        await this.getAccessDetails();
    }
}
</script>
