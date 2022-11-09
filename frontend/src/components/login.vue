<style lang="scss">
    .theme--light.v-application {
        background: $body-bg;
    	color: $body-color;
    }
</style>
<template>
	<v-container fluid fill-height>
		
		<v-row v-if="hasLoggedIn != null">
			<v-col md="6" offset-md="3" sm="6" offset-sm="3" xs="12">
				<flash-message :response="flashResponse" :response-type="flashResponseType" :response-message="flashResponseMessage"></flash-message>

				<v-alert
					:value="hasLoggedOut"
					type="warning"
					transition="scale-transition"
				>{{ loggedOutMessage }}</v-alert>

				<v-card class="elevation-6">
					<v-toolbar v-if="!hasLoggedIn" flat>
						<v-toolbar-title>Login To Your Account</v-toolbar-title>
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
					<v-card-text v-if="!hasLoggedIn">
						<v-alert
							:value="containsErrors"
							type="error"
							transition="scale-transition"
						>Email or Password is incorrect.</v-alert>

						<v-form :action="loginActionUrl" method="post" ref="loginForm" @submit.prevent="doLogin">
							
							<v-text-field
								outlined
								append-icon="fa-user"
								v-model="Email"
								name="email"
								label="Email"
								type="email"
							></v-text-field>
							<v-text-field
								outlined
								id="password"
								append-icon="fa-lock"
								v-model="Password"
								name="password"
								label="Password"
								type="password"
							></v-text-field>

							<v-row justify="space-between" no-gutters>
								<v-col md="10" sm="10">
									<v-btn color="primary" small outlined to="/forgotPassword" text>Forgot Your Password?</v-btn>
								</v-col>
								
								<v-col md="2" sm="2">
									<v-btn color="primary" type="submit" left>Login</v-btn>
								</v-col>
							</v-row>
						</v-form>

						<v-spacer></v-spacer>

						<p class="text-md-left mt-1">
							Don't have account yet?
							<v-btn color="primary" small outlined to="/register" text>Click Here To Register</v-btn>
						</p>
					</v-card-text>
					</v-card>
				</v-col>
		</v-row>
	</v-container>
</template>


<script>
var FlashMessage = require('./admin/FlashMessage.vue').default

export default {
	data() {
		return {
			loginActionUrl: this.$URLs.POST_LOGIN,
			serverErrors: [],
			ShowProgress: false,
			Email: "",
			Password: "",
			hasLoggedIn: null,
			flashResponse: false,
			flashResponseMessage: "",
			flashResponseType: null,
			hasLoggedOut: false,
			loggedOutMessage: null
		};
	},
	beforeRouteEnter(to, from, next) {
		next(vm => {
			vm.$axios.post(vm.$URLs.MY_USER)
			.then((response) => {
				next('/admin');
			}).catch((e) => {
				vm.hasLoggedIn = false
				next();
			});	
		})
	},
	computed: {
		
		containsErrors() {
			return this.serverErrors.length > 0;
		}
	},


	methods: {
		async doLogin() {
			this.ShowProgress = true;
			let params = {
				Email: this.Email,
				Password: this.Password,
				web: true
			};

			await this.$axios({
				url: this.loginActionUrl,
				method: "POST",
				data: params
			})
				.then(response => {
					this.ShowProgress = false;
					//this.$store.dispatch('setAuthTokenHeader', response.data.data.access_token)
					//Utils.setAuthorizationHeader.call(this)
					this.hasLoggedIn = true;
					this.$Utils.triggerFlash.call(this, "success", response.data.message);
					
					setTimeout(() => {
						if (this.$store.state.NextRedirection != null) {
							this.$router.push(this.$store.state.NextRedirection)
						} else {
							this.$router.push("/admin/dashboard")
						}
					}, 3000)
				})
				.catch(e => {
					this.ShowProgress = false;
					this.serverErrors = [];

					if (typeof e.response !== undefined) {
						if (e.response.status == 422) {
							if (e.response.data.errors.email) {
								e.response.data.errors.email.forEach(
									(value, index, array) => {
										this.serverErrors.push(value);
									}
								);
							}
						} else {
							this.$Utils.triggerFlash.call(this, "error", e.response.data.message);
						}
					}
				});
		},

		checkLoggedOut() {
			if (typeof this.$route.query.loggedOut != 'undefined') {
				this.hasLoggedOut = true

				switch(this.$route.query.loggedOut) {
					case '1':
						this.loggedOutMessage = 'You were logged out!';
						break;

					case '2':
						this.loggedOutMessage = 'You were logged out due to session expired!';
						break;

					case '3':
						this.loggedOutMessage = 'You need to login before you proceed';
						break;

					default:
						this.loggedOutMessage = 'You were logged out!';
						break;
				}

				this.$store.dispatch("setNextRedirection", this.$route.query.next);

				this.$router.replace('/login');


				setTimeout(() => {
					this.hasLoggedOut = false
				}, 5000);
			}
		}
	},

	created() {

		this.checkLoggedOut();
	},

	components: {
		'flash-message': FlashMessage
	}
};
</script>
