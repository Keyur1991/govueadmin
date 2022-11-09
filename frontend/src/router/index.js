import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const login = () =>
    import ( /* webpackChunkName: "chunks/users" */ './../components/login.vue')

const register = () =>
    import ( /* webpackChunkName: "chunks/users" */ './../components/register.vue');

const forgotPassword = () =>
    import ( /* webpackChunkName: "chunks/users" */ './../components/forgotPassword.vue')

const resetPassword = () =>
    import ( /* webpackChunkName: "chunks/users" */ './../components/resetPassword.vue')

const dashboard = () =>
    import ( /* webpackChunkName: "chunks/dashboard" */ './../components/dashboard.vue');

const roles = () =>
    import ( /* webpackChunkName: "chunks/roles" */ './../components/admin/roles/Roles.vue')

const features = () =>
    import ( /* webpackChunkName: "chunks/features" */ './../components/admin/features/Features.vue');

const menuMaker = () =>
    import ( /* webpackChunkName: "chunks/menus" */ './../components/admin/MenuMaker.vue');

const users = () =>
    import ( /* webpackChunkName: "chunks/users" */ './../components/admin/users/Users.vue')

const createEditUser = () =>
    import ( /* webpackChunkName: "chunks/users" */ './../components/admin/users/CreateEditUser.vue')

const myprofile = () =>
    import ( /* webpackChunkName: "chunks/users" */ './../components/admin/users/MyProfile.vue');

const posts = () =>
    import ( /* webpackChunkName: "chunks/users" */ './../components/admin/posts/Posts.vue')

const products = () => import('./../components/admin/products/Products.vue')

const routes = [{
        path: '/login',
        name: 'Login',
        meta: {
            title: 'Login'
        },
        component: login
    },
    {
        path: '/register',
        name: 'Register',
        meta: {
            title: 'Register'
        },
        component: register
    },
    {
        path: '/forgotPassword',
        name: 'Forgot Password',
        meta: {
            title: 'Forgot Password'
        },
        component: forgotPassword
    },
    {
        path: '/password/reset/:token',
        name: 'Reset Password',
        meta: {
            title: 'Reset Password'
        },
        component: resetPassword
    },
    {
        path: '/',
        redirect: '/login'
    },
    {
        path: '/admin',
        meta: {
            title: 'Dashboard'
        },
        redirect: '/admin/dashboard'
    },
    {
        path: '/admin/dashboard',
        meta: {
            title: 'Dashboard'
        },
        components: {
            default: dashboard
        }
    },
    {
        path: '/admin/roles',
        meta: {
            title: 'Roles Management | Roles List'
        },
        component: roles
    },
    {
        path: '/admin/features',
        meta: {
            title: 'Features Management | Features List'
        },
        component: features
    },
    {
        path: '/admin/menus',
        meta: {
            title: 'Menus Management'
        },
        component: menuMaker
    },
    {
        path: '/admin/users',
        meta: {
            title: 'Users Management | Users List'
        },
        component: users
    },
    {
        path: '/admin/users/create',
        meta: {
            title: 'Users Management | New User'
        },
        component: createEditUser,
        props: { OperationMode: 'Create' }
    },
    {
        path: '/admin/users/edit/:id',
        meta: {
            title: 'Users Management | Edit User'
        },
        component: createEditUser,
        props: { OperationMode: 'Edit' }
    },
    {
        path: '/admin/profile',
        meta: {
            title: 'My Profile'
        },
        component: myprofile
    },
    {
        path: '/admin/posts',
        meta: {
            title: 'Posts'
        },
        component: posts
    },
    {
        path: '/admin/products',
        meta: {
            title: 'Products Management'
        },
        component: products
    },
];

const router = new VueRouter({
    mode: "history",
    base: process.env.BASE_URL,
    routes,
});


export default router;