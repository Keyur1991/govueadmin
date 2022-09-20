import { $axios } from "../plugins/axios"
import urls from '../plugins/urls'

const axios = $axios
console.log(axios)
export const isLoggedIn = ({ commit }) => {
    axios.post(urls.ME)
    .then(response => {
        commit("changeLoggedIn", {status : true})
    }).catch(e => {
        commit("changeLoggedIn", {status : false})
    })
}

export const logout = ({commit}) => {
    axios.post(urls.POST_LOGOUT)
        .then(response => {
            commit("changeLoggedIn", {status : false})
        }).catch(e => {

        })
}