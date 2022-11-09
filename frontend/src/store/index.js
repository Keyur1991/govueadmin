import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
    state: {
        LazyLoader: '/img/image-loader.gif',
        snackbar: false,
        snackbarColor: '',
        snackbarText: '',
        RolesMaster: [],
        ShowProgress: false,
        AuthToken: '',
        userHasLoggedIn: null,
        UserAvatar: null,
        UserInfo: null,
        NoImgAvailable: '/img/not-available.jpg',
        LanguageOptions: [
            { value: 'en', text: 'English' },
            { value: 'fr', text: 'French' },
        ],
        SupportedExportFormats: ['XLS', 'CSV', 'XLSX', 'TSV', 'ODS'],
        Delimiters: [';', ',', '|'],
        Enclosers: ['~', '"', ':'],
        ShowExportProcess: false,
        NextRedirection: null
    },
    getters: {
        getFileUrlBaseName: (state) => (fileurl) => {
            let base = new String(fileurl).substring(fileurl.lastIndexOf('/') + 1);
            if (base.lastIndexOf(".") != -1)
                base = base.substring(0, base.lastIndexOf("."));
            return base;
        },

        isValidForm: (state) => (validator) => {
            validator.$touch();

            const invalidFields = Object.keys(validator.$params).filter(
                fieldName => validator[fieldName].$invalid
            );

            console.log(invalidFields)
            return invalidFields.length == 0;
        }
    },
    mutations: {
        showSnackbarMessage: (state, { data }) => {
            state.snackbarText = data.message

            if (data.code == '1') {
                state.snackbarColor = 'success'
            } else {
                state.snackbarColor = 'error'
            }

            state.snackbar = true
        },

        resetSnackbar: (state) => {
            state.snackbar = false
            state.snackbarColor = ''
            state.snackbarText = ''

        },

        rolesList(state) {
            let options = {
                params: {
                    qIsActive: '1'
                }
            }
            this._vm.$axios.get(this._vm.$URLs.ROLES_LIST, options)
                .then((response) => {
                    state.RolesMaster = response.data.data
                }).catch((e) => {

                })
        },

        showProgress: (state, { loader }) => {
            state.ShowProgress = loader
        },

        technicalError: (state) => {
            state.snackbarColor = 'red'

            state.snackbarText = 'There has been some technical problem, please try after some time.'

            setTimeout(() => { state.snackbar = true }, 1000)
        },

        serverError(state, { resobj }) {
            state.snackbarColor = 'red'

            state.snackbarText = ''

            if (typeof resobj.response !== undefined && typeof resobj.response.status !== undefined) {
                if (resobj.response.status == '403') {
                    state.snackbarText = resobj.response.data.message
                } else if (resobj.response.status == '401') {
                    window.location.href = this._vm.$URLs.GET_LOGIN + '?loggedOut=2'
                }

                setTimeout(() => { state.snackbar = true }, 1000)
            }

        },

        hasLoggedIn: (state, { status }) => {
            state.userHasLoggedIn = status;
        },

        setAuthTokenHeader: (state, { token }) => {
            state.AuthToken = 'bearer ' + token
        },

        setUserMeta: (state, { userInfo }) => {
            //state.UserAvatar = userInfo.profile_pic
            state.userHasLoggedIn = true;
            state.UserInfo = userInfo;
        },

        showExportProcess: (state, { loader }) => {
            state.ShowExportProcess = loader
        },

        setNextRedirection: (state, { nextUrl }) => {
            state.NextRedirection = nextUrl
        }
    },
    actions: {
        showSnackbarMessage: ({ commit }, data) => {
            commit('showSnackbarMessage', { data });
        },

        resetSnackbar: ({ commit }) => {
            commit('resetSnackbar')
        },

        rolesList: ({ commit }) => {
            commit('rolesList')
        },

        showProgress: ({ commit }, loader) => {
            commit('showProgress', { loader })
        },

        technicalError: ({ commit }) => {
            commit('technicalError')

        },

        serverError: ({ commit }, resobj) => {
            commit('serverError', { resobj })
        },

        hasLoggedIn: ({ commit }, status) => {
            commit('hasLoggedIn', { status })
        },

        setAuthTokenHeader: ({ commit }, token) => {
            commit('setAuthTokenHeader', { token })
        },

        setUserMeta: ({ commit }, userInfo) => {
            commit('setUserMeta', { userInfo })
        },

        showExportProcess: ({ commit }, loader) => {
            commit('showExportProcess', { loader })
        },

        setNextRedirection: ({ commit }, nextUrl) => {
            commit('setNextRedirection', { nextUrl })
        }
    },
    modules: {},
});