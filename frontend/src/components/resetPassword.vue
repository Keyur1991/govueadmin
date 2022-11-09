<template>
<v-content>
    <v-container fluid fill-height>
        <v-layout align-center justify-center>
            <v-flex xs12 sm8 md4>
                <flash-message :response="flashResponse" :response-type="flashResponseType" :response-message="flashResponseMessage" />
                <v-card class="elevation-6">
                    <v-toolbar dark color="secondary">
                        <v-toolbar-title class="text-xl-center text-lg-center text-md-center text-sm-center text-xs-center">Reset Password</v-toolbar-title>
                        <v-spacer></v-spacer>
                    </v-toolbar>
                    <v-progress-linear v-if="ShowProgress" class="mt-0" color="warning" height="8" value="15" :active="ShowProgress" :indeterminate="ShowProgress"></v-progress-linear>

                    <v-card-text>
                        <validation-error-alert :contains-errors="containsErrors"></validation-error-alert>
                        <server-validation-messages :validation-messages="serverValidationErrors"></server-validation-messages>
                        <v-form v-if="!PasswordReset" method="post" ref="forgotPasswordForm" @submit.prevent="submitForm()">
                            <v-text-field disabled name="email" id="email" label="Email" type="email" v-model="Email" :error-messages="emailErrors">
                            </v-text-field>

                            <v-text-field
                                name="password"
                                id="password"
                                label="Password"
                                type="password"
                                v-model="Password"
                                :error-messages="passwordErrors"
                            ></v-text-field>

                            <v-spacer></v-spacer>

                            <v-text-field
                                name="password_confirmation"
                                id="password-confirm"
                                label="Confirm Password"
                                type="password"
                                v-model="ConfirmPassword"
                                :error-messages="confirmPasswordErrors"
                            ></v-text-field>


                            <v-btn type="submit" color="secondary" block :loading="ShowProgress" :disabled="ShowProgress" @click.native="loader = 'ShowProgress'">Submit
                            </v-btn>

                            <v-spacer></v-spacer>


                        </v-form>
                        <v-btn color="primary" small outline to="/login">Back To Login</v-btn>
                    </v-card-text>
                </v-card>
            </v-flex>
        </v-layout>
    </v-container>
</v-content>
</template>

<script>
import {
    required,
    email,
    minLength,
	maxLength,
	sameAs
} from "vuelidate/lib/validators";

var ValidationErrorAlert = require('./admin/ValidationErrorAlert.vue').default
var ServerValidationMessages = require('./admin/ServerValidationMessages.vue').default
var FlashMessage = require('./admin/FlashMessage.vue').default

export default {
    data() {
        return {
            resetPasswordActionUrl: this.$URLs.RESET_PWD,
            serverValidationErrors: [],
            ShowProgress: false,
            Email: '',
            flashResponse: false,
			flashResponseMessage: "",
			flashResponseType: null,
            Password: '',
            ConfirmPassword: '',
            Token: '',
            PasswordReset: false
        }
    },

    computed: {
        containsErrors() {
            return this.serverValidationErrors.length > 0 || this.emailErrors.length > 0
        },

        emailErrors() {
            const errors = [];

            if (!this.$v.Email.$dirty) return errors;

            !this.$v.Email.required && errors.push('Please enter email address.')

                !this.$v.Email.email && errors.push('Please enter valid email address.')

            return errors;
        },

        passwordErrors: {
            get() {
                const errors = [];

                if (!this.$v.Password.$dirty) return errors;

                !this.$v.Password.required &&
                    errors.push("Please enter your password.");

                !this.$v.Password.minLength &&
                    errors.push("Password must be atlease 6 characters long.");

                !this.$v.Password.maxLength &&
                    errors.push("Password must not be longer than 10 characters.");

                return errors;
            },

            set() {

            }
		},

		confirmPasswordErrors: {
            get() {
                const errors = [];

                if (!this.$v.ConfirmPassword.$dirty) return errors;

                !this.$v.ConfirmPassword.required &&
                    errors.push("Please enter your password.");

                !this.$v.ConfirmPassword.sameAsPassword &&
                    errors.push("Password could not be matched.");

                return errors;
            },

            set() {

            }
		}
    },

    validations: {
        Email: {
            required,
            email
        },

		Password: {
			required,
			minLength: minLength(6),
			maxLength: maxLength(10)
		},

		ConfirmPassword: {
			required,
			sameAsPassword: sameAs("Password")
		}
    },

    methods: {
        submitForm() {
            if (this.$Utils.isValidForm.call(this)) {
                this.ShowProgress = true;

                let params = {
                    token: this.Token,
    				email: this.Email,
                    password: this.Password,
                    password_confirmation: this.ConfirmPassword
    			};

                this.$axios({
    				url: this.resetPasswordActionUrl,
    				method: "POST",
    				data: params
    			})
				.then(response => {
					this.ShowProgress = false;
                    this.$v.$reset();
                    this.Email = this.Password = this.ConfirmPassword = ''
                    this.$Utils.triggerFlash.call(this, "success", response.data.message, 0);
                    setTimeout(() => {
                        this.$router.push('/login')
                    }, 5000);
				})
				.catch(e => {
					this.ShowProgress = false;
					this.serverValidationErrors = [];

					if (typeof e.response !== undefined) {
						if (e.response.status == 422) {
							if (e.response.data.errors.email) {
								e.response.data.errors.email.forEach(
									(value, index, array) => {
										this.serverValidationErrors.push(value);
									}
								);
							}
						} else {
							this.$Utils.triggerFlash.call(this, "error", e.response.data.message);
						}
					}
				});
            }

            return false;
        },

        setTokenEmailData() {
            this.Token = this.$route.params.token
            this.Email = this.$route.query.email
        }
    },

    created() {
        this.setTokenEmailData()
    },

    components: {
        'validation-error-alert': ValidationErrorAlert,
        'server-validation-messages' : ServerValidationMessages,
        'flash-message': FlashMessage
    }
}
</script>
