<template>
	<v-container fluid v-if="hasLoggedIn">
		<v-row dense justify="space-between" >
			<v-col md="10" sm="10" xs="12">
				<page-title>{{ $vuetify.lang.t('$vuetify.Roles.Listing.PageTitle') }}</page-title>
			</v-col>
			<v-col md="2" sm="2" xs="12" >
				<breadcrumbs :items="breadcrumbs" />
			</v-col>
		</v-row>
		<v-row dense >
			<v-col>
				<v-card >
					<v-row class="pa-2" dense v-if="hasListingAccess" >
						<v-col md="4" sm="4" xs="12" >
							<search @filter="globalSearch" />
						</v-col>
						<v-col md="8" sm="8" xs="12" >
							<page-actions hide-back :has-add-access="hasAddAccess" :has-listing-access="hasListingAccess" :add-title="`${$vuetify.lang.t('$vuetify.Roles.Listing.NewBtn')}`" :has-act-deact-access="hasActDeactAccess" @filter="toggleFilter" @actinact="activeInactive" @add="newRoleModal"/>
						</v-col>
					</v-row>

					<v-row class="pa-2" dense v-if="hasListingAccess && (NameFilter != '')">
						<v-col>
							<div class="text-xs-left">
								<v-fade-transition mode="out-in">
									<v-chip
										v-if="NameFilter != null && NameFilter.trim() != ''"
										close
										@click:close="clearAndRefresh('NameFilter')"
									>{{ $vuetify.lang.t('$vuetify.Roles.Fields.Name') }}: {{ NameFilter }}</v-chip>
								</v-fade-transition>
							</div>
						</v-col>
					</v-row>

					<v-row class="pa-2" dense >
						<v-col>
							<v-data-table
								v-if="hasListingAccess"
								v-model="selected"
								:headers="headers"
								:items="roles"
								:options.sync="pagination"
								hide-default-footer
								show-select
								:server-items-length="totalRoles"
								:loading="loading"
								class="elevation-2"
							>
								<v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>
								<template v-slot:item.is_active="{ item }">
									<v-chip small label v-if="item.is_active" color="green" text-color="white">{{ $vuetify.lang.t('$vuetify.Yes') }}</v-chip>
									<v-chip small label v-else color="red" text-color="white">{{ $vuetify.lang.t('$vuetify.No') }}</v-chip>
								</template>
								<template v-slot:item.actions="{ item }">
									<v-tooltip bottom v-if="hasSetPermissionsAccess">
										<template v-slot:activator="{ on }">
											<v-btn
												v-on="on"
												color="primary"
												class="mr-2"
												@click="setPermissions(item)"
												icon
												text
											>
												<v-icon small>fa-user-lock</v-icon>
											</v-btn>
										</template>
										<span>{{ $vuetify.lang.t('$vuetify.Roles.Listing.PermissionsBtn') }}</span>
									</v-tooltip>
									<v-tooltip bottom v-if="hasEditAccess">
										<template v-slot:activator="{ on }">
											<v-btn
												v-on="on"
												color="primary"
												class="mr-2"
												@click="editRecord(item)"
												icon
												text
											>
												<v-icon small>fa-edit</v-icon>
											</v-btn>
										</template>
										<span>{{ $vuetify.lang.t('$vuetify.EditBtn') }}</span>
									</v-tooltip>
									<v-tooltip bottom v-if="hasDeleteAccess">
										<template v-slot:activator="{ on }">
											<v-btn
												v-if="item.id != '1'"
												v-on="on"
												color="primary"
												class="mr-2"
												@click="triggerDeleteModal(item)"
												icon
												text
											>
												<v-icon small>fa-trash</v-icon>
											</v-btn>
										</template>
										<span>{{ $vuetify.lang.t('$vuetify.DeleteBtn') }}</span>
									</v-tooltip>
								</template>
								
							</v-data-table>
							<unauthorized :display="hasListingAccess"></unauthorized>
						</v-col>
					</v-row>
					<v-row class="pa-2" dense>
						<v-col>
							<pagination v-if="hasListingAccess" :pages="pages" @update="updatePagination" />
						</v-col>
					</v-row>
				</v-card>
				<v-navigation-drawer v-model="filterOpened" absolute  right width="250">
					<v-row dense v-if="hasListingAccess">
						<v-col v-if="filterOpened">
							<v-card class="elevation-2">
								<v-card-title class="pb-0">
									<v-row dense class="border-bottom-1">
										<v-col xs="9" sm="9" md="9" >
											<h5 class="title">{{ $vuetify.lang.t('$vuetify.Filter') }}</h5>
										</v-col>
										<v-col xs="3" sm="3" md="3" >
											<v-btn
												class="right"
												color="primary"
												icon
												@click="filterOpened = false"
												
											>
												<v-icon >fa-arrow-right</v-icon>
											</v-btn>
										</v-col>
									</v-row>
								</v-card-title>
								<v-card-text class="pt-2">
									<v-row dense>
										<v-col>
											<v-text-field
												outlined
												dense
												v-model="NameFilter"
												:label="`${$vuetify.lang.t('$vuetify.Roles.Fields.Name')}`"
												append-icon="fa-times"
												@click:append="NameFilter = ''"
												hide-details
											></v-text-field>

											<v-checkbox v-model="RememberFiltersCheck" :label="`${$vuetify.lang.t('$vuetify.RememberFilters')}`" color="primary"></v-checkbox>
										</v-col>
									</v-row>
								</v-card-text>
								<v-card-actions class="ml-3">
									<v-btn @click="fetchData" color="primary" >
										<v-icon small>fa-search</v-icon>&nbsp;{{ $vuetify.lang.t('$vuetify.SearchBtn') }}
									</v-btn>
								
									<v-btn color="secondary" @click="reset" >
										<v-icon small>fa-undo</v-icon>&nbsp;{{ $vuetify.lang.t('$vuetify.ResetBtn') }}
									</v-btn>
								</v-card-actions>
							</v-card>
						</v-col>
					</v-row>
				</v-navigation-drawer>
			</v-col>
		</v-row>

		<role-add-edit
			:role-item-modal="RoleItemModal"
			:role-item="RoleRecord"
			:title="RoleModalTitle"
			:edit="RoleEditMode"
			@close="RoleItemModal = false"
			@reload="fetchData"
		></role-add-edit>

		<delete-modal
			:delete-item-modal="DeleteItemModal"
			:url="DeleteItemUrl"
			@close="DeleteItemModal = false"
			@reload="fetchData"
		></delete-modal>

		<permissions-modal
			:permissions-save-modal="PermissionsModal"
			:role-item="RoleRecord"
			:permissions="PermissionsList"
			@close="PermissionsModal = false"
			@reload="fetchData"
		></permissions-modal>
	</v-container>
</template>

<script>
import Vue from "vue";

import VueCookie from "vue-cookie";

Vue.use(VueCookie);

var RoleAddEdit = require("./RoleAddEdit.vue").default;
var DeleteModal = require("./../DeleteModal.vue").default;
var PermissionsModal = require("./PermissionsModal.vue").default;
var Unauthorized = require("./../Unauthorized.vue").default;
var Breadcrumbs = require('../Breadcrumb.vue').default
var PageTitle = require('../PageTitle.vue').default
var PageActions = require('../PageActions.vue').default
var Pagination = require('../Pagination.vue').default
var Search = require('../Search.vue').default

export default {
	data() {
		return {
			totalRoles: 0,
			roles: [],
			loading: true,
			pagination: {
				page: 1,
				rowsPerPage: 10,
				sortBy: ["created_at"],
				descending: 1
			},
			filterOpened: false,
			NameFilter: "",
			RememberFiltersCheck: false,
			RoleRecord: {},
			RoleModalTitle: "New Role",
			RoleEditMode: false,
			RoleDialog: false,
			DeleteItemUrl: "",
			DeleteItemModal: false,
			PermissionsModal: false,
			PermissionsList: [],
			SelectedPermissions: [],
			hasAddAccess: false,
			hasEditAccess: false,
			hasListingAccess: null,
			hasActDeactAccess: false,
			hasDeleteAccess: false,
			hasSetPermissionsAccess: false,
			RoleItemModal: false,
			selected: [],
			gsTerm: ''
		};
	},
	computed: {
		pages () {
			if (this.pagination.rowsPerPage == null ||
                this.totalRoles == null || this.totalRoles == 0
            ) return 0

			return Math.ceil(this.totalRoles / this.pagination.rowsPerPage)
        },

		hasLoggedIn() {
			return this.$store.state.userHasLoggedIn
		},
		breadcrumbs() {
			return [
				{
					text: this.$vuetify.lang.t('$vuetify.Home'),
					action: "/admin/dashboard",
					disabled: false
				},
				{
					text: this.$vuetify.lang.t('$vuetify.Roles.Breadcrumbs.Roles'),
					disabled: true
				}
			]
		},
		headers() {
			return [
				{ text: this.$vuetify.lang.t('$vuetify.Roles.Fields.Name'), value: "name", align: "left"},
				{ text: this.$vuetify.lang.t('$vuetify.CreatedAt'), value: "created_at", align: "left", width:"17%" },
				{ text: this.$vuetify.lang.t('$vuetify.Active'), value: "is_active", sortable: true, align: "left", width: "10%" },
				{ text: this.$vuetify.lang.t('$vuetify.Actions'), value: "actions", sortable: false, align: "left", width: "20%" }
			]
		}
	},
	watch: {
		pagination: {
			handler() {
				this.fetchData();
			}
		}
	},

	methods: {
		fetchData() {
			this.filterOpened = false;
			this.loading = true;

			let options = {
				params: {
					page: this.pagination.page,
					perPage: this.pagination.rowsPerPage,
					sortBy: this.pagination.sortBy[0],
					descending: this.pagination.descending,
					qName: this.NameFilter,
					qRemember: this.RememberFiltersCheck,
					gsTerm: this.gsTerm
				}
			};

			this.$axios
				.get(this.$URLs.ROLES_LIST, options)
				.then(res => {
					let d = res.data.data.data;

					this.totalRoles = res.data.data.total;

					this.roles = [];

					this.roles = d;

					if (this.roles.length == d.length) {
						this.loading = false;
					}
				})
				.catch(e => {
					this.roles = [];

					this.loading = false;
					this.$store.dispatch("serverError", e);
				});
		},

		editRecord(item) {
			this.RoleRecord = item;
			this.RoleModalTitle = "Edit Role";
			this.RoleEditMode = true;
			this.RoleItemModal = true;
		},

		deleteRecord(item) {},

		toggleFilter() {
			this.filterOpened = !this.filterOpened;
		},

		fetchCookies() {
			let options = {
				params: {
					for: ["qName", "qRemember"]
				},
				withCredentials: true,
			};

			var $this = this;
			this.$axios
				.get(this.$URLs.ROLES_COOKIES, options)
				.then(res => {
					console.log(res.data);
					this.$nextTick(() => {
						this.NameFilter = res.data.Roles_qName;
						this.RememberFiltersCheck =
							res.data.Roles_qRemember == "true" ? true : false;
					});
				})
				.catch(e => {
					this.$store.dispatch("serverError", e);
				});
		},

		setInitialFilters() {
			this.NameFilter =
				this.$cookie.get("Roles_qName") != null
					? this.$cookie.get("Roles_qName")
					: "";

			this.RememberFiltersCheck =
				this.$cookie.get("Roles_qRemember") != null &&
				this.$cookie.get("Roles_qRemember") == "true"
					? true
					: false;

			let sortby =
				this.$cookie.get("Roles_sortBy") != null
					? this.$cookie.get("Roles_sortBy")
					: this.pagination.sortBy;

			this.pagination.sortBy = [sortby]
			this.pagination.descending =
				this.$cookie.get("Roles_descending") != null &&
				this.$cookie.get("Roles_descending") == "true"
					? true
					: this.pagination.descending;
		},

		clearAndRefresh(datakey) {
			if (datakey == "NameFilter") {
				this.NameFilter = "";
			}

			this.fetchData();
		},

		reset() {
			this.NameFilter = "";
			this.EmailFilter = "";
			this.RememberFiltersCheck = false;
			this.filterOpened = false;
			this.fetchData();
		},

		newRoleModal() {
			this.RoleRecord = {};
			this.RoleModalTitle = "New Role";
			this.RoleEditMode = false;
			this.RoleItemModal = true;
		},

		activeInactive(status) {
			this.$Utils.activeInactive.call(this, status, this.$URLs.ROLE_ACT_DEACT);
		},

		triggerDeleteModal(item) {
			this.$Utils.triggerDeleteModal.call(this, this.$URLs.ROLES_LIST + "/" + item.id)
		},

		setPermissions(item) {
			this.RoleRecord = item;
			this.$store.dispatch("showProgress", true);

			this.$axios({
				url: this.$URLs.ROLES_PERMISSIONS + '/' + item.id,
				method: "GET"
			})
				.then(response => {
					this.$store.dispatch("showProgress", false);

					this.PermissionsList = response.data.data.permissions;

					this.PermissionsModal = true;
				})
				.catch(e => {
					this.$store.dispatch("showProgress", false);

					this.$store.dispatch("serverError", e);
				});
		},

		getAccessDetails() {
			this.$store.dispatch("showProgress", true);

			this.$axios
				.get(this.$URLs.ROLE_ACCESS)
				.then(response => {
					this.$store.dispatch("showProgress", false);

					this.hasAddAccess = response.data.data.canAdd;
					this.hasEditAccess = response.data.data.canEdit;
					this.hasDeleteAccess = response.data.data.canDelete;
					this.hasListingAccess = response.data.data.canViewList;
					this.hasActDeactAccess =
						response.data.data.canActivateDeactive;
					this.hasSetPermissionsAccess =
						response.data.data.canSetPermissions;
				})
				.catch(e => {
					this.$store.dispatch("showProgress", false);
					this.$store.dispatch("serverError", e);
				});
		},
		updatePagination(currentPage, rowsPerPage) {
			this.pagination.page = currentPage
			this.pagination.rowsPerPage = rowsPerPage
			this.fetchData()
		},
		globalSearch(term) {
			this.gsTerm = term
			this.pagination.page = 1
			this.fetchData()
		}
	},

	components: {
		"role-add-edit": RoleAddEdit,
		"delete-modal": DeleteModal,
		"permissions-modal": PermissionsModal,
		"unauthorized": Unauthorized,
		"breadcrumbs": Breadcrumbs,
		"page-title": PageTitle,
		"page-actions": PageActions,
		"pagination": Pagination,
		"search": Search
	},

	async created() {
		await this.$axios.get(this.$URLs.SANCTUM_CSRF)
		await this.$Utils.checkUserLoggedIn.call(this)
		await this.getAccessDetails()
		await this.setInitialFilters()
	}
};
</script>
