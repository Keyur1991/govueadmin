<template>
    <v-app>
        <v-navigation-drawer v-if="hasLoggedIn" v-model="drawer" absolute temporary>
            <menus></menus>
        </v-navigation-drawer>

        <v-app-bar app color="primary">
            <v-icon v-if="hasLoggedIn"  @click="drawer = !drawer" color="white">fa-bars</v-icon>

            <v-toolbar-title class="white--text ">&nbsp;&nbsp;Admin Panel</v-toolbar-title>

            <v-spacer></v-spacer>
            <v-toolbar-items v-if="hasLoggedIn">
                <v-menu offset-y transition="slide-y-reverse-transition" left>
                    <template v-slot:activator="{ on }">
                        <v-btn v-on="on" text v-if="hasLoggedIn" class="white--text">
                            <img :src="currentLangImg" :alt="currentLangText" >&nbsp;&nbsp;{{ currentLangText }}
                        </v-btn>
                    </template>
                    <v-list dense>
                        <v-list-item v-for="(lng, index) in languageOptions" :key="index" @click="switchLanguage(lng.value)">
                            <v-list-item-action>
                                <img :src="`/img/${lng.value}.png`" :alt="lng.text" >
                            </v-list-item-action>
                            <v-list-item-content>
                                {{ lng.text }}
                            </v-list-item-content>
                        </v-list-item>
                    </v-list>
                </v-menu>
                <v-menu offset-y transition="slide-y-reverse-transition">
                    <template v-slot:activator="{ on }">
                        <v-btn v-on="on" text class="white--text">
                            <v-avatar size="25px" left> <img :src="UserAvatar" alt="Avatar" >
                            </v-avatar>&nbsp;&nbsp;&nbsp;<v-icon small>fa-caret-down</v-icon>
                        </v-btn>
                    </template>
                    <v-list dense>
                        <v-list-item to="/admin/profile" >
                            <v-list-item-icon>
                                <v-icon small left>fa-user</v-icon>
                            </v-list-item-icon>
                            <v-list-item-content>
                                My Profile
                            </v-list-item-content>
                        </v-list-item>
                        <v-list-item @click="logout">
                            <v-list-item-icon>
                                <v-icon small left>fa-key</v-icon>
                            </v-list-item-icon>
                            <v-list-item-content>
                                Log Out
                            </v-list-item-content>
                        </v-list-item>
                    </v-list>
                </v-menu>
            </v-toolbar-items>
        </v-app-bar>

        <!-- Sizes your content based upon application components -->
        <v-main>
            <v-fade-transition hide-on-leave>
                <router-view></router-view>
            </v-fade-transition>
        </v-main>

        <v-footer app>
            <v-dialog v-model="ShowProgress" content-class="no-box-shadow" >
                <v-card style="background: transparent;">
                    <v-card-text class="text-center pt-1">
                        <v-progress-circular
                            indeterminate
                            color="primary lighten-2"
                        ></v-progress-circular>
                    </v-card-text>
                </v-card>
            </v-dialog>
            <v-slide-y-reverse-transition>
                <v-snackbar
                    v-model="snackbar"
                    top
                    right
                    :color="snackbarColor"
                    :timeout="snackBarTimeout"
                    class="mt-3"

                    >
                    {{ snackbarText }}
                    <template v-slot:action="{ attrs }">
                        <v-btn
                        color="white"
                        text
                        v-bind="attrs"
                        @click="resetSnackbar()"
                        >
                        <v-icon small>fa-times</v-icon>
                        </v-btn>
                    </template>
                    
                </v-snackbar>
            </v-slide-y-reverse-transition>
            <export-process :export-modal="exportLoader"></export-process>
        </v-footer>
    </v-app>
</template>

<script>
const Menus = require('./components/admin/Menus.vue').default;
const ExportModal = require('./components/ExportLoader.vue').default;

export default {
    name: "App",

    data: () => ({
        drawer: false,
        mini: false,
        snackBarTimeout: 5000
    }),

    computed: {
        snackbar: {
            get() {
                return this.$store.state.snackbar
            },

            set() {
                this.$store.dispatch('resetSnackbar')
            }
        },

        snackbarColor: {
            get() {
                return this.$store.state.snackbarColor
            },
            set() {

            }
        },

        snackbarText: {
            get() {
                return this.$store.state.snackbarText
            },
            set() {

            }
        },

        ShowProgress: {
            get() {
                return this.$store.state.ShowProgress
            },

            set() {

            }
        },

        hasLoggedIn: {
            get() {
                return this.$store.state.userHasLoggedIn
            },

            set() {

            }
        },

        UserAvatar: {
            get() {
                return this.$store.state.UserAvatar
            },

            set() {

            }
        },

        currentLangImg: {

            get() {
                return '/img/' + this.$vuetify.lang.current + '.png'
            },

            set() {

            }
        },

        currentLangText: {
            get() {
                let lng;

                switch (this.$vuetify.lang.current) {
                    case 'en':
                        lng = 'English'
                        break;

                    case 'fr':
                        lng = 'French'
                        break;

                    case 'nl':
                        lng = 'Dutch'
                        break;

                    default:
                        lng = 'English'
                        break;
                }

                return lng
            },

            set() {

            }
        },

        languageOptions() {
            return this.$store.state.LanguageOptions
        },

        exportLoader() {
            return this.$store.state.ShowExportProcess
        }
    },

    methods: {
        async logout() {
            this.$store.dispatch("showProgress", true);
            await this.$axios.post(this.$URLs.USER_LOG_OUT)
                .then(() => {
                    this.$store.dispatch("showProgress", false);
                    this.$store.dispatch("hasLoggedIn", false);
                    this.$router.push("/login?loggedOut=1")
                })
            return false;
        },


        resetSnackbar() {
            this.$store.dispatch('resetSnackbar')
        },

        toggleDrawer() {
            this.drawer = !this.drawer;
            this.mini = !this.mini
        },

        switchLanguage(lang) {
            this.$Utils.switchLanguage.call(this, lang)
        }
    },

    components: {
        'menus': Menus,
        'export-process': ExportModal
    },

    created() {
       // this.switchLanguage(this.$cookie.get('Curr_Lang'))
            //this.$store.dispatch('setAuthTokenHeader', this.$cookie.get('Auth_Token'))
            //Utils.setAuthorizationHeader.call(this)
    }
};

</script>
