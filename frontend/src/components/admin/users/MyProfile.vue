<template>
<v-container fluid>
    <v-row dense>
        <v-col md="9" sm="9" xs="12">
            <page-title>{{ $vuetify.lang.t('$vuetify.Users.MyProfile.PageTitle') }}</page-title>
        </v-col>
        <v-col md="3" sm="3" xs="12">
            <breadcrumbs :items="breadcrumbs" />
        </v-col>
    </v-row>
    <v-row dense>
        <v-col>
            <v-form :action="profileUpdateActionUrl" method="post" ref="newUserForm" @submit.prevent="updateProfile">
                <v-card>
                    <v-card-text>
                        <v-row dense>
                            <v-col>
                                <validation-error-alert :contains-errors="containsErrors"></validation-error-alert>
                                <server-validation-messages :validation-messages="serverValidationErrors"></server-validation-messages>
                            </v-col>
                        </v-row>
                        <v-row dense>
                            <v-col>
                                <v-tabs>
                                    <v-tab ripple @change="isBasicTab()">{{ $vuetify.lang.t('$vuetify.Users.MyProfile.Tabs.BasicInfo') }}</v-tab>
                                    <v-tab ripple @change="isPasswordChangeTab()">{{ $vuetify.lang.t('$vuetify.Users.MyProfile.Tabs.ChangePassword') }}</v-tab>

                                    <v-tab-item>
                                        <v-card flat>
                                            <v-card-text>
                                                <v-row dense class="pa-2">
                                                    <v-col xs="12" sm="6" md="6" >
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
                                                                <v-text-field outlined dense v-model="Email"
                                                                    :disabled="disableEmailField"
                                                                    :label="`${$vuetify.lang.t('$vuetify.Users.Fields.Email')}`">
                                                                </v-text-field>
                                                            </v-col>
                                                        </v-row>

                                                        <v-row dense>
                                                            <v-col xs="12" sm="8" md="8">
                                                                <p class="mb-1">{{ $vuetify.lang.t('$vuetify.Users.Fields.SelectProfilePicture') }}</p>
                                                                <v-btn
                                                                    color="secondary"
                                                                    outlined
                                                                    small
                                                                    @click="triggerFileInput">
                                                                    {{ (imageName != '') ? imageName : 'Browse' }}
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
                                                            <v-col xs="12" sm="4" md="4" >
                                                                    <v-img v-if="imageUrl"
                                                                    :src="imageUrl"
                                                                    width="150"
                                                                    height="150"
                                                                    class="elevation-6"
                                                                    >
                                                                    </v-img>
                                                                    <v-btn v-if="imageUrl"
                                                                    color="red"
                                                                    outlined
                                                                    small
                                                                    @click="resetFileSelection">
                                                                    Remove
                                                                    <v-icon right small>fa-times</v-icon>
                                                                </v-btn>
                                                            </v-col>
                                                        </v-row>
                                                    </v-col>
                                                    <v-col sm="1" md="1" xs="12"></v-col>
                                                    <v-col xs="12" sm="5" md="5">
                                                        <p><v-icon small>fa-dot-circle</v-icon> {{ $vuetify.lang.t('$vuetify.Users.Validations.AllFieldsRequired') }}</p>
                                                        <p><v-icon small>fa-dot-circle</v-icon> {{ $vuetify.lang.t('$vuetify.Users.Validations.ProfilePictureValidationRules') }}</p>
                                                        <p v-if="PreviewCurrentImage != null">
                                                            <v-row dense>
                                                                <v-col xs="12" sm="8" md="8" >
                                                                    <v-img
                                                                        :src="PreviewCurrentImage"
                                                                        width="200"
                                                                        height="200"
                                                                        class="elevation-2"
                                                                        :lazy-src="LazyLoader"
                                                                    >
                                                                    </v-img>
                                                                </v-col>
                                                                <v-col xs="12" sm="4" md="4" v-if="EligibleForUnlink">
                                                                    <v-btn color="red"
                                                                        outlined
                                                                        small
                                                                        @click="removeCurrentImage"
                                                                        >
                                                                        {{ $vuetify.lang.t('$vuetify.DeleteBtn') }}
                                                                        <v-icon right small>fa-trash</v-icon>
                                                                    </v-btn>
                                                                </v-col>
                                                            </v-row>
                                                        </p>
                                                    </v-col>
                                                </v-row>
                                            </v-card-text>
                                        </v-card>
                                    </v-tab-item>
                                    <v-tab-item>
                                        <v-card flat>
                                            <v-card-text>
                                                <v-row dense class="pa-3">
                                                    <v-col xs="12" sm="6" md="6" >
                                                        <v-text-field
                                                            outlined
                                                            dense
                                                            name="oldPassword"
                                                            id="oldPassword"
                                                            :label="`${$vuetify.lang.t('$vuetify.Users.Fields.OldPassword')}*`"
                                                            type="password"
                                                            v-model="OldPassword"
                                                            :error-messages="oldPasswordErrors">
                                                        </v-text-field>

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
                                                    <v-col sm="1" md="1" xs="12"></v-col>
                                                    <v-col xs="12" sm="5" md="5">
                                                        <p><v-icon small>fa-dot-circle</v-icon> {{ $vuetify.lang.t('$vuetify.Users.Validations.PasswordValidationRules') }}</p>
                                                    </v-col>
                                                </v-row>
                                            </v-card-text>
                                        </v-card>
                                    </v-tab-item>
                                </v-tabs>
                            </v-col>
                        </v-row>
                    </v-card-text>
                    <v-card-actions>
                        <v-btn  color="primary" type="submit">{{ $vuetify.lang.t('$vuetify.SaveBtn') }}</v-btn>
                        <v-btn @click="clear"  type="button">{{ $vuetify.lang.t('$vuetify.ResetBtn') }}</v-btn>
                    </v-card-actions>
                </v-card>
            </v-form>
        </v-col>
        <v-flex xs12 sm12 md12>
            
            <v-layout row wrap  class="ma-2">
                <v-flex xs12 sm12 md12>
                    
                </v-flex>
            </v-layout>
        </v-flex>
    </v-row>
</v-container>

</template>

<script>
import Vue from 'vue'
import Vuelidate from 'vuelidate'
var Breadcrumbs = require('../Breadcrumb.vue').default
var PageTitle = require('../PageTitle.vue').default
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

var ValidationErrorAlert = require('../ValidationErrorAlert.vue').default

var ServerValidationMessages = require('../ServerValidationMessages.vue').default

export default {
    props: ['OperationMode'],
    data() {
        return {
            
            profileUpdateActionUrl: this.$URLs.MY_PROFILE,
            Name: '',
            Email: '',
            OldPassword: '',
            Password: '',
            ConfirmPassword: '',
            ProfilePicture: null,
            imageUrl: '',
            imageName: '',
            ProfilePictureErrors: [],
            PreviewCurrentImage: null,
            PasswordCaption: 'Password *',
            ConfirmPasswordCaption: 'Confirm Password *',
            OldPasswordCaption: 'Old Password *',
            disableEmailField: true,
            User: null,
            serverValidationErrors: [],
            IsBasicTab: true,
            IsPasswordChangeTab: false
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
                    text: this.$vuetify.lang.t('$vuetify.Users.Breadcrumbs.MyProfile'),
                    action: '/admin/users/mine',
                    disabled: true
                }
            ]
        },

        nameErrors() {
            const errors = [];

            if (!this.$v.Name.$dirty) return errors;

            !this.$v.Name.required && errors.push(this.$vuetify.lang.t('$vuetify.Users.Validations.NameRequired'));

            !this.$v.Name.maxLength &&
                errors.push(this.$vuetify.lang.t('$vuetify.Users.Validations.NameMaxLength'));

            return errors;
        },

        oldPasswordErrors() {
            const errors = [];

            if (!this.$v.OldPassword.$dirty) return errors;

            !this.$v.OldPassword.required && errors.push(this.$vuetify.lang.t('$vuetify.Users.Validations.PasswordRequired'));

            return errors;
        },

        passwordErrors() {
            const errors = [];

            if (!this.$v.Password.$dirty) return errors;

            !this.$v.Password.required && errors.push(this.$vuetify.lang.t('$vuetify.Users.Validations.PasswordRequired'));

            !this.$v.Password.minLength &&
                errors.push(`${this.$vuetify.lang.t('$vuetify.Users.Validations.PasswordText1')} ${this.$v.Password.$params.minLength.min} ${this.$vuetify.lang.t('$vuetify.Users.Validations.PasswordText2')}`);

            !this.$v.Password.maxLength &&
                errors.push(`${this.$vuetify.lang.t('$vuetify.Users.Validations.PasswordText3')} ${this.$v.Password.$params.maxLength.max} ${this.$vuetify.lang.t('$vuetify.Users.Validations.PasswordText4')}`);

            this.Password != '' && !this.$v.Password.StrictPassword && errors.push(this.$vuetify.lang.t('$vuetify.Users.Validations.PasswordValidationRules'));
            return errors;
        },

        confirmPasswordErrors() {
            const errors = [];

            if (!this.$v.ConfirmPassword.$dirty) return errors;

            !this.$v.ConfirmPassword.required &&
                errors.push(this.$vuetify.lang.t('$vuetify.Users.Validations.PasswordRequired'));

            !this.$v.ConfirmPassword.sameAsPassword &&
                errors.push(this.$vuetify.lang.t('$vuetify.Users.Validations.PasswordNotMatch'));

            return errors;
        },

        containsErrors() {
            return (
                this.nameErrors.length > 0 ||
                this.oldPasswordErrors.length > 0 ||
                this.passwordErrors.length > 0 ||
                this.confirmPasswordErrors.length > 0
            );
        },

        EligibleForUnlink() {
            return this.$Utils.eligibleForUnlink.call(this)
        }
    },

    validations: {
        Name: {
            required: requiredIf((vueInstance) => {
                return vueInstance.IsBasicTab
            }),
            maxLength: maxLength(20)
        },
        OldPassword: {
            required: requiredIf((vueInstance) => {
                return vueInstance.IsPasswordChangeTab
            })
        },
        Password: {
            required: requiredIf((vueInstance) => {
                return vueInstance.IsPasswordChangeTab
            }),
            minLength: minLength(6),
            maxLength: maxLength(15),
            StrictPassword (value) {
                return and('required', value => StrictPassword(value))
            }
        },
        ConfirmPassword: {
            required: requiredIf((vueInstance) => {
                return vueInstance.IsPasswordChangeTab
            }),
            sameAsPassword: sameAs('Password')
        }
    },

    methods: {
        async updateProfile() {
            if (this.$Utils.isValidForm.call(this)) {
                window.scrollTo(0,0);
                this.serverValidationErrors = [];
                this.$store.dispatch('showProgress', true)
                await this.$axios.get(this.$URLs.SANCTUM_CSRF)
                let data = new FormData();

                data.append('Name', this.Name)

                if (this.OldPassword != '') {
                    data.append('OldPassword', this.OldPassword)
                }
                if (this.Password != '') {
                    data.append('Password', this.Password)
                }

                if (this.ProfilePicture != null) {
                    data.append('ProfilePicture', this.ProfilePicture)
                }

                data.append('_method', 'PUT')

                await this.$axios({
                    url: this.profileUpdateActionUrl,
                    method: "POST",
                    data: data
                }).then(response => {
                    this.$store.dispatch('showProgress', false)
                    this.$store.dispatch('setUserMeta', response.data.data)

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

                        if (typeof e.response.data.errors.OldPassword != 'undefined') {
                            e.response.data.errors.OldPassword.forEach((value, index, array) => {
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
            this.setData()
        },

        triggerFileInput() {
            this.$refs.ProfilePicture.click()
        },

        handleFileSelection(e) {
            this.$Utils.handleFileSelection.call(this, e)
        },

        resetFileSelection() {
            this.$Utils.resetFileSelection.call(this)
        },

        loadUserEditData() {
            this.User = this.$store.state.UserInfo
            this.setData()
        },

        setData() {
            this.Name = this.User.name;
            this.Email = this.User.email;
            this.PreviewCurrentImage = this.User.profile_pic;
            this.PasswordCaption = 'Password ';
            this.ConfirmPasswordCaption = 'Confirm Password ';
            this.disableEmailField = true
            this.Password = this.ConfirmPassword = this.OldPassword = ''
        },

        async removeCurrentImage() {
            let b = () => { this.PreviewCurrentImage = this.$store.state.NoImgAvailable }
            await this.$Utils.removeCurrentImage.call(this, 'self')
            await b()
        },
        isBasicTab() {
            this.$v.$reset();
            this.setData();
            this.IsBasicTab = true
            this.IsPasswordChangeTab = false
        },

        isPasswordChangeTab() {
            this.$v.$reset();
            this.IsBasicTab = false
            this.IsPasswordChangeTab = true
        }
    },

    components: {
        'validation-error-alert': ValidationErrorAlert,
        'server-validation-messages' : ServerValidationMessages,
        "breadcrumbs": Breadcrumbs,
		"page-title": PageTitle,
    },

    async created() {
        await this.$axios.get(this.$URLs.SANCTUM_CSRF)
        await this.$Utils.checkUserLoggedIn.call(this)
        await this.loadUserEditData()
    }
}
</script>
