import { createStore } from "vuex"
import * as getters from "./getters"
import * as actions from "./actions"
import * as mutations from "./mutations"

const state = {
    loggedin: false
}
const store = createStore({
    state,
    getters,
    mutations,
    actions
})

export default store 