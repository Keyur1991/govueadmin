<template>
    <v-alert
        :value="response"
        :type="responseType"
        transition="scale-transition"
    >{{ responseMessage }}</v-alert>
</template>
<script>
export default {
    props:['response', 'responseType', 'responseMessage']
}
</script>
