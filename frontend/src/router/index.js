import { createRouter, createWebHistory } from 'vue-router'
const Home = () => import("../views/HelloWorld.vue")
const Login = () => import("../views/Login.vue")
const Register = () => import("../views/Register.vue")
const Dashboard = () => import("../views/Dashboard.vue")

//import Home from '/src/components/HelloWorld.vue'
const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home,
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
    },
    {
        path: '/register',
        name: 'Register',
        component: Register,
    },
    {
        path: '/dashboard',
        name: 'Dashboard',
        component: Dashboard
    }
]
const router = createRouter({
    history: createWebHistory(),
    routes,
})
export default router