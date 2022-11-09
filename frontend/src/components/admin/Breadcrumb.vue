<template>
    <v-breadcrumbs :items="items">
        <template v-slot:item="{ item }">
            <v-breadcrumbs-item
                exact
                :disabled="item.disabled"
                :to="item.action"
            >{{ item.text }}
            </v-breadcrumbs-item>
        </template>
    </v-breadcrumbs>
</template>

<script>
export default {
    props: ['items']
}
</script>
