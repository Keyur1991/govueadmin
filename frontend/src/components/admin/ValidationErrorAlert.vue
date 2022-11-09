<template>
    <v-alert class="mt-1"
        type="error"
        :value="containsErrors"
        color="error"
        transition="scale-transition"
        outlined
    >{{ $vuetify.lang.t('$vuetify.ValidationErrorsText') }}
    </v-alert>
</template>
<script>
export default {
    props: ["containsErrors"]
}
</script>
