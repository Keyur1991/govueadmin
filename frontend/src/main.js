import axios from "axios"
import Vue from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";
import store from "./store";
import vuetify from "./plugins/vuetify";
import { URLs, $ApiBaseUrl } from './urls';
import { Utils } from './admin';
import Vuelidate from 'vuelidate'
import VueCookie from "vue-cookie";

axios.defaults.withCredentials = true

Vue.prototype.$axios = axios
Vue.prototype.$URLs = URLs
Vue.prototype.$Utils = Utils
Vue.prototype.$ApiBaseUrl = $ApiBaseUrl
Vue.use(Vuelidate)
Vue.use(VueCookie);

Vue.prototype._ = require('lodash');

router.beforeEach((to, from, next) => {
    document.title = to.meta.title
    next()
})

Vue.config.productionTip = false;

new Vue({
    router,
    store,
    vuetify,
    render: (h) => h(App),
}).$mount("#app");