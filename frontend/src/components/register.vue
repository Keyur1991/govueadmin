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
				<v-alert
					:value="registerResponse"
					:type="registerResponseType"
					transition="scale-transition"
				>{{ registerResponseMessage }} <router-link to="/login">Click To Login</router-link></v-alert>
				<v-card class="elevation-6" v-if="!hasRegistered">
					<v-toolbar flat>
						<v-toolbar-title>Create Your Account</v-toolbar-title>
					</v-toolbar>
					<v-progress-linear
						v-if="ShowProgress"
						class="mt-0"
						color="warning"
						height="8"
						value="15"
						:active="ShowProgress"
						:indeterminate="ShowProgress"
					></v-progress-linear>
					<v-card-text>
						<v-alert
							:value="containsErrors"
							type="error"
							transition="scale-transition"
						>There are some validation errors, please fix them first.</v-alert>
						<v-form ref="registrationForm" @submit.prevent="submitForm()" method="post">
							<v-text-field
								outlined
								name="first_name"
								label="First Name"
								type="text"
								v-model="FirstName"
								:error-messages="nameErrors"
							></v-text-field>

							<v-text-field
								outlined
								name="last_name"
								label="Last Name"
								type="text"
								v-model="LastName"
								:error-messages="nameErrors"
							></v-text-field>

							<v-text-field
								outlined
								name="email"
								id="email"
								label="Email"
								type="email"
								v-model="Email"
								:error-messages="emailErrors"
							></v-text-field>
							
							<!-- <v-select
								outlined
								name="role"
								id="role"
								label="Role"
								v-model="Role"
								:items="Roles"
								item-text="name"
								item-value="id"
								:error-messages="roleErrors"
							></v-select> -->
							
							<v-text-field
								outlined
								name="password"
								id="password"
								label="Password"
								type="password"
								v-model="Password"
								:error-messages="passwordErrors"
							></v-text-field>

							<v-text-field
								outlined
								name="password_confirmation"
								id="password-confirm"
								label="Confirm Password"
								type="password"
								v-model="ConfirmPassword"
								:error-messages="confirmPasswordErrors"
							></v-text-field>

							<v-row justify="space-between" no-gutters>
								<v-col md="9" sm="9">
									<v-btn color="primary" small outlined to="/login" text>Click Here To Login</v-btn>
								</v-col>
								<v-col md="3" sm="3">
									<v-btn
										type="submit"
										color="primary"
										:loading="ShowProgress"
										:disabled="ShowProgress"
										@click.native="loader = 'ShowProgress'"
									>Register</v-btn>
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
import { validationMixin } from "vuelidate";

import {
	required,
	email,
	minLength,
	maxLength,
	sameAs
} from "vuelidate/lib/validators";

export default {
	data() {
		return {
			FirstName: "",
			LastName: "",
			Email: "",
			Password: "",
			ConfirmPassword: "",
			ShowProgress: false,
			PostUrl: this.$URLs.USER_REGISTER,
			registerResponse: false,
            registerResponseType: null,
			registerResponseMessage: '',
			hasRegistered: false,
			Roles: [],
			Role: null
		};
	},
	beforeRouteEnter(to, from, next) {
		next(vm => {
			vm.$axios.post(vm.$URLs.MY_USER)
			.then((response) => {
				next('/admin');
			}).catch((e) => {
				next();
			});
		})
		
	},
	
	mixins: [validationMixin],

	validations: {
		FirstName: {
			required,
			maxLength: maxLength(20)
		},

		LastName: {
			required,
			maxLength: maxLength(20)
		},

		Email: {
			required,
			email,
			maxLength: maxLength(100)
		},

		/*Role: {
			required
		},*/

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

	computed: {
		firstNameErrors: {
            get() {
                const errors = [];

                if (!this.$v.FirstName.$dirty) return errors;

                !this.$v.FirstName.required && errors.push("Please enter your first name.");

                !this.$v.FirstName.maxLength &&
                    errors.push("First name must not be longer than 20 characters");

                return errors;
            },

            set() {

            }
		},

		lastNameErrors: {
            get() {
                const errors = [];

                if (!this.$v.LastName.$dirty) return errors;

                !this.$v.LastName.required && errors.push("Please enter your last name.");

                !this.$v.LastName.maxLength &&
                    errors.push("Last name must not be longer than 20 characters");

                return errors;
            },

            set() {

            }
		},

		emailErrors: {
            get() {
                const errors = [];

                if (!this.$v.Email.$dirty) return errors;

                !this.$v.Email.required &&
                    errors.push("Please enter your email address.");

                !this.$v.Email.email &&
                    errors.push("Please enter valid email address.");

                !this.$v.Email.maxLength &&
                    errors.push("Email must not be longer than 100 characters");

                return errors;
            },

            set() {

            }
		},

		roleErrors: {
            get() {
                const errors = [];

                // if (!this.$v.Role.$dirty) return errors;

                // !this.$v.Role.required && errors.push("Please select a role.");

                return errors;
            },

            set() {

            }
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
		},

		containsErrors: {
            get() {
                return (
                    this.firstNameErrors.length > 0 ||
					this.lastNameErrors.length > 0 ||
					this.emailErrors.length > 0 ||
					this.roleErrors.length > 0 || 
                    this.passwordErrors.length > 0 ||
                    this.confirmPasswordErrors.length > 0
                );
            },

            set() {

            }
		}
	},

	methods: {
		isValidForm() {
			this.$v.$touch();

			const invalidFields = Object.keys(this.$v.$params).filter(
				fieldName => this.$v[fieldName].$invalid
			);

			return invalidFields.length == 0;
		},

		submitForm() {
			if (this.isValidForm()) {
				this.ShowProgress = true;
				let UserRoles = []
				//UserRoles.push(this.Role)

				let params = {
					FirstName: this.FirstName,
					LastName: this.LastName,
					Email: this.Email,
					Password: this.Password,
					UserRoles: UserRoles
				};

				this.$axios.post(this.PostUrl,params)
					.then((response) => {
                        this.ShowProgress = false;
						this.hasRegistered = true;
                        this.registerResponseType = 'success';
						this.registerResponseMessage = response.data.message;
						this.triggerFlash();
					})
					.catch((e) => {
						this.ShowProgress = false;
						if (typeof e.response != 'undefined') {
                            this.firstNameErrors = []
							this.lastNameErrors = []
                            this.emailErrors = []
                            this.passwordErrors = []
                            if (e.response.status == 422) {
                                if (e.response.data.errors.first_name) {
                                    e.response.data.errors.first_name.forEach(
                                        (value, index, array) => {
                                            this.nameErrors.push(value);
                                        }
                                    );
                                }

								if (e.response.data.errors.last_name) {
                                    e.response.data.errors.last_name.forEach(
                                        (value, index, array) => {
                                            this.lastNameErrors.push(value);
                                        }
                                    );
                                }

                                if (e.response.data.errors.email) {
                                    e.response.data.errors.email.forEach(
                                        (value, index, array) => {
                                            this.emailErrors.push(value);
                                        }
                                    );
                                }

                                if (e.response.data.errors.password) {
                                    e.response.data.errors.password.forEach(
                                        (value, index, array) => {
                                            this.passwordErrors.push(value);
                                        }
                                    );
                                }
                            }
                        }
					});
			}

			return false;
		},
		
		triggerFlash() {
			this.registerResponse = true;
			// setTimeout(() => {
			// 	this.registerResponse = false;
			// }, 5000);
		},


		FetchRoles() {
			this.$axios.get(this.$URLs.REG_ROLES + '/0',
				{
					params: {
						SkipAdminRoles: 1
					}
				}).then((response) => {
					this.Roles = response.data.data
				}).catch((error) => {
					console.log(error)
				})
		}
	},

	created() {
		this.FetchRoles()
	}
};
</script>