<template>
    <v-container fluid v-if="hasLoggedIn">
        <v-row dense justify="space-between" >
            <v-col>
                <v-card title="Menus">
                    <v-card-title primary-title>
                        <v-row dense>
                            <v-col md="10" sm="10" xs="8"><h3>Menus Management</h3></v-col>
                            <v-col md="2" sm="2" xs="4">
                                <v-btn color="primary" v-if="hasAddAccess" outlined @click="newMenu" right small><v-icon small>fa-plus</v-icon>&nbsp;&nbsp;New Menu</v-btn>
                            </v-col>
                        </v-row>
                    </v-card-title>
                    <v-card-text>
                        <v-row dense v-if="hasListingAccess">
                            <v-col md="8" offset-md="2" class="pa-2">
                                <v-card class="elevation-0">
                                    <v-card-text>
                                        <VueNestable v-model="nestableItems" :max-depth="MaximumDepth" @change="reorder">
                                            <div slot="placeholder">
                                                <b>This list is empty</b>
                                                <p>You can add items with the help of form left side.</p>
                                            </div>
                                            <template slot-scope="{ item }">
                                                <v-row dense class="menu-row">
                                                    <v-col md="8" sm="8" xs="6" class="menu-row-title">
                                                        <VueNestableHandle  :item="item">
                                                            <v-icon small>fa-bars</v-icon>&nbsp;&nbsp;<span>{{ item.text }}</span>
                                                        </VueNestableHandle>
                                                    </v-col>
                                                    <v-col md="4" sm="4" xs="6" >
                                                        <v-btn v-if="hasDeleteAccess" class="pull-right edit-menu" slot="activator" color="primary" @click="deleteMenu(item.id)"  icon  right>
                                                            <v-icon size="12">fa-trash</v-icon>
                                                        </v-btn>
                                                        <v-btn v-if="hasEditAccess" class="pull-right delete-menu" slot="activator" color="primary" @click="editMenu(item)" icon >
                                                            <v-icon size="12">fa-edit</v-icon>
                                                        </v-btn>
                                                    </v-col>
                                                </v-row>
                                            </template>
                                        </VueNestable>
                                    </v-card-text>
                                </v-card>
                            </v-col>
                        </v-row>
                        <unauthorized :display="hasListingAccess"></unauthorized>
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>
        <menu-add-edit :menu-item="MenuToEdit" :menu-item-modal="OpenMenuEditModal" @close="loadMenusTree" :modal-title="MenuModalTitle"></menu-add-edit>
        <delete-modal
			:delete-item-modal="DeleteItemModal"
			:url="DeleteItemUrl"
			@close="DeleteItemModal = false"
			@reload="loadMenusTree"
		></delete-modal>
    </v-container>
</template>

<script>
import Vue from "vue"
import VueNestable from 'vue-nestable'


Vue.use(VueNestable)

var Unauthorized = require("./Unauthorized.vue").default;

var MenuAddEditModal = require("./MenuAddEditModal.vue").default;

var DeleteModal = require("./DeleteModal.vue").default;
export default {
    data() {
        return {
            nestableItems: [],
            createActionUrl: this.$URLs.MENUS_LIST,
            Name: '',
            MaximumDepth: 3,
            hasAddAccess: null,
            hasEditAccess: null,
            hasDeleteAccess: null,
            hasListingAccess: null,
            OpenMenuEditModal: false,
            MenuToEdit: null,
            MenuModalTitle: null,
            DeleteItemModal: false,
            DeleteItemUrl: ''
        }
    },

    computed: {
        ShowProgress() {
            return this.$store.state.ShowProgress
        },
        nameErrors() {
            const errors = [];

            if (!this.$v.Name.$dirty) return errors;

            !this.$v.Name.required && errors.push("Please enter menu name.");

            !this.$v.Name.maxLength &&
                errors.push("Name must not be longer than 20 characters");

            return errors;
        },

		hasLoggedIn() {
			return this.$store.state.userHasLoggedIn
		}
    },
    async created() {
        await this.$axios.get(this.$URLs.SANCTUM_CSRF)
        await this.$Utils.checkUserLoggedIn.call(this)
        await this.getAccessDetails();
        await this.loadMenusTree();
        await this.$store.dispatch('rolesList')
    },

    methods: {
        loadMenusTree() {
            this.OpenMenuEditModal = false
            this.MenuToEdit = null
            this.$store.dispatch('showProgress', true)
            return this.$axios({
                url: this.createActionUrl,
                method: "GET",
                params: {parent_id: '0'}
            }).then(response => {
                this.$store.dispatch('showProgress', false)
                this.nestableItems = response.data.data
            }).catch(e => {
                this.$store.dispatch('showProgress', false)
                this.$store.dispatch('serverError', e)
            });
        },

        reorder() {
            this.$store.dispatch('showProgress', true)
            let options = { menus: this.nestableItems }
            this.$axios.post(this.$URLs.MENUS_REORDER, options)
                .then((response) => {
                    this.$store.dispatch('showProgress', false)
                    this.$store.dispatch("showSnackbarMessage", {
                        message: response.data.message,
                        code: "1"
                    });
                }).catch((e) => {
                    this.$store.dispatch('showProgress', false)
                    this.$store.dispatch('serverError', e)
                })
        },

		getAccessDetails() {
			this.$store.dispatch("showProgress", true);
            this.$axios
				.get(this.$URLs.MENUS_ACCESS)
				.then(response => {
					this.$store.dispatch("showProgress", false);

					this.hasAddAccess = response.data.data.canAdd;
					this.hasEditAccess = response.data.data.canEdit;
					this.hasDeleteAccess = response.data.data.canDelete;
					this.hasListingAccess = response.data.data.canViewList;
				})
				.catch(e => {
					this.$store.dispatch("serverError", e);
					this.$store.dispatch("showProgress", false);
				});
		},

        editMenu(item) {
            this.MenuToEdit = item
            this.MenuModalTitle = 'Edit Menu'
            this.OpenMenuEditModal = true
        },

        deleteMenu(id) {
            this.DeleteItemUrl = this.$URLs.MENUS_LIST + '/' + id
            this.DeleteItemModal = true
        },

        newMenu() {
            this.MenuToEdit = null
            this.MenuModalTitle = 'New Menu'
            this.OpenMenuEditModal = true
        },
    },
    components: {
        'unauthorized' : Unauthorized,
        'menu-add-edit' : MenuAddEditModal,
        'delete-modal' : DeleteModal
    }
}
</script>

<style scoped lang="css">
@import "../../styles/_nestable.css";

.show {display: flex;}

.display-hide {display: none;}

.menu-row {border: 1px solid #ddd; border-radius: 5px;}

.pull-right {float: right;}

.menu-row-title {font-size: 14px;}

.nestable-handle {margin-left: 10px;}

.edit-menu {margin: 0px; margin-left: 10px;}
.delete-menu {margin: 0px;margin-right: 10px;}
</style>
