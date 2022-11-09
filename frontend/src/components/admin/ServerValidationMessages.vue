<template>
    <v-card v-if="validationMessages.length > 0" class="mx-auto" color="red darken-1">
        <v-card-text>
            <ul><li v-for="(item, index) in validationMessages" class="white--text" :key="index">{{ item }}</li></ul>
        </v-card-text>
    </v-card>
</template>

<script>
export default {
    props: ['validationMessages']
}
</script>