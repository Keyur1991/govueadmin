<template>
  <v-container fluid fill-height>
    <v-row>
      <v-col md="6" offset-md="3" sm="6" offset-sm="3" xs="12">
        <v-card class="elevation-6">
          <v-toolbar flat>
            <v-toolbar-title>Register Your Account</v-toolbar-title>
          </v-toolbar>
          <v-snackbar v-model="snackbar">
            {{ message }}

            <template v-slot:actions>
                <v-btn
                color="pink"
                variant="text"
                @click="snackbar = false"
                >
                Close
                </v-btn>
            </template>
          </v-snackbar>
          <v-card-text>
            <v-form method="post" ref="loginForm" @submit.prevent="doRegister">
              <v-text-field
                outlined
                v-model="FirstName"
                name="first_name"
                label="First Name*"
                :error-messages="errors.FirstName"
              ></v-text-field>
              <v-text-field
                outlined
                v-model="LastName"
                name="last_name"
                label="Last Name*"
                :error-messages="errors.LastName"
              ></v-text-field>
              <v-text-field
                outlined
                v-model="Email"
                name="email"
                label="Email"
                type="email"
                :error-messages="errors.Email"
              ></v-text-field>
              <v-text-field
                outlined
                id="password"
                v-model="Password"
                name="password"
                label="Password"
                type="password"
                :error-messages="errors.Password"
              ></v-text-field>
              <v-row justify="space-between" no-gutters>
                <v-col md="10" sm="10">
                  <p>
                    Already have an account ?
                    <router-link to="/login">Login Now</router-link>
                  </p>
                </v-col>

                <v-col md="2" sm="2">
                  <v-btn
                    color="primary"
                    :disabled="isSubmitting"
                    type="submit"
                    left
                    >Register</v-btn
                  >
                </v-col>
              </v-row>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { useField, useForm } from "vee-validate";
import { string, object } from "yup";
import { inject, ref } from "vue";

const $axios = inject("axios");
const $urls = inject("urls");
const snackbar = ref(null)
const message = ref(null)

const schema = object({
  FirstName: string().required(),
  LastName: string().required(),
  Email: string().required().email(),
  Password: string().required().min(8),
});

const { errors, handleSubmit, isSubmitting, resetForm } = useForm({
  validationSchema: schema,
});

const { value: FirstName } = useField("FirstName");
const { value: LastName } = useField("LastName");
const { value: Email } = useField("Email");
const { value: Password } = useField("Password");

const doRegister = handleSubmit(async (values, actions) => {
  $axios
    .post($urls.POST_REGISTER, values)
    .then((response) => {
      snackbar.value = true
      message.value = "Your account has been successfully registered."
      resetForm({
          values: {
            FirstName: '',
            LastName: '',
            Email: '',
            Password: ''
        },
      })
    })
    .catch((e) => {
      console.log(e);
      if (e.status == 422 && e.data.errors) {

        if (e.data.errors.first_name) {
          actions.setFieldError("FirstName", e.data.errors.first_name.join(","));
        }

        if (e.data.errors.last_name) {
          actions.setFieldError("LastName", e.data.errors.last_name.join(","));
        }

        if (e.data.errors.email) {
          actions.setFieldError("Email", e.data.errors.email.join(","));
        }

        if (e.data.errors.password) {
          actions.setFieldError("Password", e.data.errors.password.join(","));
        }
      }
    });
});
</script>
