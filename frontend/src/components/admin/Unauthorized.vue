<template>
    <v-alert
    :value="display != null && !display"
    type="error" 
    color="red darken-1"
    dark
    >
    You are not authorized to perform that action.
    </v-alert>
</template>

<script>
export default {
    props: ['display']
}
</script>