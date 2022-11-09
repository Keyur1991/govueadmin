<style lang="scss">
    .theme--light.v-application {
        background: $body-bg;
    	color: $body-color;
    }
</style>
<template>
    <v-container fluid fill-height>
        <v-row>
            <v-col md="6" offset-md="3" sm="6" offset-sm="3" xs="12">
                <flash-message :response="flashResponse" :response-type="flashResponseType" :response-message="flashResponseMessage" />
                <v-card class="elevation-6">
                    <v-toolbar flat>
                        <v-toolbar-title>Forgot Your Password ?</v-toolbar-title>
                    </v-toolbar>
                    <v-progress-linear v-if="ShowProgress" class="mt-0" color="warning" height="8" value="15" :active="ShowProgress" :indeterminate="ShowProgress"></v-progress-linear>

                    <v-card-text>
                        <validation-error-alert :contains-errors="containsErrors"></validation-error-alert>
                        <server-validation-messages :validation-messages="serverValidationErrors"></server-validation-messages>
                        <v-form :action="forgotPasswordActionUrl" method="post" ref="forgotPasswordForm" @submit.prevent="submitForm()">
                            <v-text-field outlined append-icon="fa-user" name="email" id="email" label="Email" type="email" v-model="Email" :error-messages="emailErrors">
                            </v-text-field>

                            <v-row justify="space-between" no-gutters>
								<v-btn color="primary" small outlined to="/login" text>Go Back To Login</v-btn>
								
								<v-col md="2" sm="2">
									<v-btn color="primary" type="submit" :loading="ShowProgress" :disabled="ShowProgress" @click.native="loader = 'ShowProgress'">Submit</v-btn>
								</v-col>
							</v-row>
                        </v-form>
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
import {
    required,
    email
} from "vuelidate/lib/validators";

var ValidationErrorAlert = require('./admin/ValidationErrorAlert.vue').default
var ServerValidationMessages = require('./admin/ServerValidationMessages.vue').default
var FlashMessage = require('./admin/FlashMessage.vue').default

export default {
    data() {
        return {
            forgotPasswordActionUrl: this.$URLs.FORGOT_PWD,
            serverValidationErrors: [],
            ShowProgress: false,
            Email: '',
            flashResponse: false,
			flashResponseMessage: "",
			flashResponseType: null,
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
        }
    },

    validations: {
        Email: {
            required,
            email
        }
    },

    methods: {
        async submitForm() {
            if (this.$Utils.isValidForm.call(this)) {
                this.ShowProgress = true;
                await this.$axios.get(this.$URLs.SANCTUM_CSRF)
                let params = {
    				email: this.Email,
    			};

                this.$axios({
    				url: this.forgotPasswordActionUrl,
    				method: "POST",
    				data: params
    			})
				.then(response => {
					this.ShowProgress = false;
					this.$Utils.triggerFlash.call(this, "success", response.data.message, 0);
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
        }
    },

    components: {
        'validation-error-alert': ValidationErrorAlert,
        'server-validation-messages' : ServerValidationMessages,
        'flash-message': FlashMessage
    }
}
</script>
