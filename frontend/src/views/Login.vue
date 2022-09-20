<template>
  <v-container fluid fill-height>
    <v-row>
			<v-col md="6" offset-md="3" sm="6" offset-sm="3" xs="12">
				<v-card class="elevation-6">
					<v-toolbar flat>
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
					<v-alert
						v-if="serverErrors != null"
						:value="serverErrors"
						type="error"
						transition="scale-transition"
					>Email or Password is incorrect.</v-alert>
					<v-card-text >
						
						<v-form method="post" ref="loginForm" @submit.prevent="doLogin">
							<v-text-field
								outlined
								append-icon="fa-user"
								v-model="Email"
								name="email"
								label="Email"
								type="email"
								:error-messages="errors.Email"
							></v-text-field>
							<v-text-field
								outlined
								id="password"
								append-icon="fa-lock"
								v-model="Password"
								name="password"
								label="Password"
								type="password"
								:error-messages="errors.Password"
							></v-text-field>
							<v-row justify="space-between" no-gutters>
								<v-col md="10" sm="10">
									<p>Forgot your password ? <router-link to="/forgotPassword">Reset Now</router-link></p>
								</v-col>
								
								<v-col md="2" sm="2">
									<v-btn color="primary" :disabled="submitting" type="submit" left>Login</v-btn>
								</v-col>
							</v-row>
							<v-row class="mt-3"> 
								<v-col>
									<p>Don't have account yet ? <router-link to="/register">Register Now</router-link></p>
								</v-col>
							</v-row>
						</v-form>
					</v-card-text>
				</v-card>
			</v-col>
		</v-row>
		<v-snackbar v-model="showToastr" :color="`${typeToastr == 'error' ? 'red' : 'green'}-darken-2`" location="top left">
            {{ textToastr }}

            <template v-slot:actions>
                <v-btn
                color="white"
                variant="text"
                @click="showToastr = false"
                >
                Close
                </v-btn>
            </template>
          </v-snackbar>
  </v-container>
</template>

<script setup>
import { useField, useForm } from "vee-validate";
import { string, object } from "yup"
import { inject, ref } from "vue"
import { useRouter } from "vue-router"

const $axios = inject('axios')
const $urls = inject('urls')
const submitting = ref(false)
const serverErrors = ref(null)
const showToastr = ref(false)
const typeToastr = ref(null)
const textToastr = ref(null)
const $router = useRouter()

const schema = object({
	Email: string().required().email(),
	Password: string().required().min(8)
})

const { errors, handleSubmit } = useForm({
	validationSchema: schema
})

const { value: Email } = useField("Email")
const { value: Password } = useField("Password")

const doLogin = handleSubmit(values => {
	submitting.value = true
	$axios.post($urls.POST_LOGIN, values)
		.then(response => {
			submitting.value = false
			$router.push("/")
		})
		.catch(e => {
			submitting.value = false
			showToastr.value = true
			typeToastr.value = "error"
			textToastr.value = e.data.message
			
		})
})
</script>
