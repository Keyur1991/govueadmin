<style lang="scss">
    .theme--light.v-application {
        background: $body-bg;
    	color: $body-color;
    }
    .v-icon.v-icon {
        align-items: center;
        display: inline-flex;
        -webkit-font-feature-settings: "liga";
        font-feature-settings: "liga";
        font-size: $icon-size;
        justify-content: center;
        letter-spacing: normal;
        line-height: 1;
        position: relative;
        text-indent: 0;
        transition: .3s cubic-bezier(.25,.8,.5,1),visibility 0s;
        vertical-align: middle;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;
    }

</style>
<template>
   
    <v-list v-if="showWhenLoggedIn" >
        <template v-for="(item, index) in MenuItems" >
            <v-list-item v-bind:key="index"  v-if="item.children.length == 0"  :to="item.url" exact>
                <v-list-item-icon>
                    <v-icon >{{ item.icon }}</v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                    <v-list-item-title>{{ $vuetify.lang.t('$vuetify.Menus.' + item.title) }}</v-list-item-title>
                </v-list-item-content>
            </v-list-item>

            <v-list-group v-bind:key="index" v-else >
                <template v-slot:activator>
                  <v-list-item-icon>
                    <v-icon >{{ item.icon }}</v-icon>
                  </v-list-item-icon>  
                  <v-list-item-content>  
                        <v-list-item-title>{{ $vuetify.lang.t('$vuetify.Menus.' + item.title) }}</v-list-item-title>  
                  </v-list-item-content>
                </template>

                <v-list-item v-for="(child, childIndex) in item.children" :key="childIndex" :to="child.url" exact>
                    <v-list-item-icon>
                        <v-icon >{{ child.icon }}</v-icon>
                    </v-list-item-icon>
                    <v-list-item-content>
                        <v-list-item-title>{{ $vuetify.lang.t('$vuetify.Menus.' + child.title) }}</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
            </v-list-group>
        </template>
    </v-list>
</template>

<script>
    export default {
        data() {
            return {
                MenuItems: []
            }
        },

        computed: {
            showWhenLoggedIn() {
                return this.$store.state.userHasLoggedIn
            }
        },

        methods: {
            loadMenus() {
                this.$store.dispatch('showProgress', true)
                this.$axios.get(this.$URLs.MENUS_LIST, {params: { CheckPermission: true, parent_id: '0'} })
                .then((response) => {
                    this.MenuItems = response.data.data
                    this.$store.dispatch('showProgress', false)
                }).catch((error) => {
                    this.$store.dispatch('showProgress', false)
                    this.$store.dispatch('serverError', error)
                })
            }
        },

        created() {
            this.loadMenus()
        }
    }
</script>
