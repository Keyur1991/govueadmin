import Vue from "vue";
import Vuetify from "vuetify/lib/framework";
import '@fortawesome/fontawesome-free/css/all.css'
import en from '../locale/en';
import fr from '../locale/fr';


Vue.use(Vuetify);

export default new Vuetify({
    icons: {
        iconfont: "fa",
    },
    lang: {
        locales: { en, fr },
        current: 'en',
    },
});