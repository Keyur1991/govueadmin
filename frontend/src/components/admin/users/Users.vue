<template>
	<v-container fluid v-if="hasLoggedIn">
		<v-row dense justify="space-between" >
			<v-col md="10" sm="10" xs="12">
				<page-title>{{ $vuetify.lang.t('$vuetify.Users.Listing.PageTitle') }}</page-title>
			</v-col>
			<v-col md="2" sm="2" xs="12" >
				<breadcrumbs :items="breadcrumbs" />
			</v-col>
		</v-row>
		<v-row dense >
			<v-col>
				<v-card>
					<v-row class="pa-2" dense v-if="hasListingAccess" >
						<v-col md="4" sm="4" xs="12" >
							<search @filter="globalSearch" />
						</v-col>
						<v-col md="8" sm="8" xs="12" >
							<page-actions :has-add-access="hasAddAccess" :has-listing-access="hasListingAccess" add-action="/admin/users/create" :add-title="`${$vuetify.lang.t('$vuetify.Users.Listing.NewBtn')}`" :has-import-access="hasImportAccess" :has-export-access="hasExportAccess" :has-act-deact-access="hasActDeactAccess" hide-back @filter="toggleFilter" @actinact="activeInactive" @export="ExportModal = true" @import="ImportModal = true" />
						</v-col>
					</v-row>
					<v-row class="pa-2" dense v-if="hasListingAccess && (NameFilter != '' || EmailFilter != '' || IsActiveFilter != '')">
						<v-col>
							<div class="text-xs-left">
								<v-fade-transition mode="out-in">
									<v-chip
										v-if="NameFilter != null && NameFilter.trim() != ''"
										close
										@click.close="clearAndRefresh('NameFilter')"
									>{{ $vuetify.lang.t('$vuetify.Users.Fields.Name') }}: {{ NameFilter }}</v-chip>
								</v-fade-transition>
								<v-fade-transition mode="out-in">
									<v-chip
										v-if="EmailFilter != null && EmailFilter.trim() != ''"
										close
										@click.close="clearAndRefresh('EmailFilter')"
									>{{ $vuetify.lang.t('$vuetify.Users.Fields.Email') }}: {{ EmailFilter }}</v-chip>
								</v-fade-transition>

								<v-fade-transition mode="out-in">
									<v-chip
										v-if="IsActiveFilter != null && IsActiveFilter.trim() != ''"
										close
										@click.close="clearAndRefresh('IsActiveFilter')"
									>{{ $vuetify.lang.t('$vuetify.Users.Fields.Active') }}: {{ IsActiveFilter == '1' ? $vuetify.lang.t('$vuetify.Yes') : $vuetify.lang.t('$vuetify.No') }}</v-chip>
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
								:items="users"
								:options.sync="pagination"
								hide-default-footer
								show-select
								:server-items-length="totalUsers"
								:loading="loading"
								class="elevation-2"
							>
								<v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>
								<template v-slot:item.is_active="{ item }">
									<v-chip small label v-if="item.is_active" color="green" text-color="white">{{ $vuetify.lang.t('$vuetify.Yes') }}</v-chip>
									<v-chip small label v-else color="red" text-color="white">{{ $vuetify.lang.t('$vuetify.No') }}</v-chip>
								</template>
								<template v-slot:item.actions="{ item }">
									<v-tooltip bottom v-if="hasEditAccess">
										<template v-slot:activator="{ on }">
											<v-btn
												v-on="on"
												color="primary"
												class="mr-2"
												:to="editLink(item)"
												icon
											>
												<v-icon small>fa-edit</v-icon>
											</v-btn>
										</template>
										
										<span>{{ $vuetify.lang.t('$vuetify.EditBtn') }}</span>
									</v-tooltip>
									
									<v-tooltip bottom v-if="hasDeleteAccess">
										<template v-slot:activator="{ on }">
											<v-btn
												v-on="on"
												color="primary"
												class="mr-2"
												@click="triggerDeleteModal(item)"
												icon
												
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
				<v-navigation-drawer
					v-if="hasListingAccess"
					v-model="filterOpened"
					absolute
					right
					width="250"
				>
					<v-row dense>
						<v-col>
							<v-card>
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
												<v-icon>fa-arrow-right</v-icon>
											</v-btn>
										</v-col>
									</v-row>
								</v-card-title>
								<v-card-text class="pt-0">
									<v-row dense>
										<v-col>
											<v-text-field
												outlined
												dense
												v-model="NameFilter"
												:label="`${$vuetify.lang.t('$vuetify.Users.Fields.Name')}` "
												append-icon="fa-times"
												@click:append="NameFilter = ''"
											></v-text-field>

											<v-text-field
												outlined
												dense
												v-model="EmailFilter"
												:label="`${$vuetify.lang.t('$vuetify.Users.Fields.Email')}` "
												append-icon="fa-times"
												@click:append="EmailFilter = ''"
											></v-text-field>

											<v-select
												outlined
												dense
												v-model="IsActiveFilter"
												:items="ActiveFilterOptions"
												item-text="text" 
												item-value="value" 
												:label="`${$vuetify.lang.t('$vuetify.Users.Fields.Active')}` "
												append-icon="fa-times"
												@click:append="IsActiveFilter = ''"
											></v-select>

											<v-checkbox v-model="RememberFiltersCheck" :label="`${$vuetify.lang.t('$vuetify.RememberFilters')}`" color="primary"></v-checkbox>
										</v-col>
									</v-row>
								</v-card-text>

								<v-card-actions class="ml-3">
									<v-btn @click="fetchData" color="primary" >
										<v-icon small>fa-search</v-icon>&nbsp;{{ $vuetify.lang.t('$vuetify.SearchBtn') }}
									</v-btn>

									<v-btn  color="secondary" @click="reset" >
										<v-icon small>fa-undo</v-icon>&nbsp;{{ $vuetify.lang.t('$vuetify.ResetBtn') }}
									</v-btn>
								</v-card-actions>
							</v-card>
						</v-col>
					</v-row>
				</v-navigation-drawer>
			</v-col>
		</v-row>
		
		<view-user
			v-if="hasViewAccess"
			:user="ViewUser"
			:has-edit-access="hasEditAccess"
			ref="ViewUserDialog"
		></view-user>
		<delete-modal
			:delete-item-modal="DeleteItemModal"
			:url="DeleteItemUrl"
			@close="DeleteItemModal = false"
			@reload="fetchData"
		></delete-modal>
		<export-users :export-modal="ExportModal" @export="ExportData" @close="ExportModal = false"></export-users>
		<import-users :import-modal="ImportModal" @import="ImportData" @close="ImportModal = false"></import-users>
	</v-container>
</template>

<script>
var ViewUser = require("./ViewUser").default;
var DeleteModal = require("./../DeleteModal.vue").default;
var Unauthorized = require("./../Unauthorized.vue").default;
var Breadcrumbs = require('../Breadcrumb.vue').default
var PageTitle = require('../PageTitle.vue').default
var PageActions = require('../PageActions.vue').default
var Pagination = require('../Pagination.vue').default
var Search = require('../Search.vue').default
var ExportUsers = require('./ExportUsers.vue').default;
var ImportUsers = require('./ImportUsers.vue').default;

export default {
	data() {
		return {
			
			totalUsers: 0,
			users: [],
			loading: true,
			pagination: {
				page: 1,
				rowsPerPage: 10,
				sortBy: ["created_at"],
				descending: 1
			},
			filterOpened: false,
			NameFilter: "",
			EmailFilter: "",
			IsActiveFilter: "",
			RememberFiltersCheck: false,
			ViewUser: null,
			DeleteItemModal: false,
			DeleteItemUrl: "",
			hasAddAccess: false,
			hasListingAccess: null,
			hasEditAccess: false,
			hasDeleteAccess: false,
			hasActDeactAccess: false,
			hasViewAccess: false,
			hasImportAccess: false,
			hasExportAccess: false,
			selected: [],
			gsTerm: '',
			ExportModal: false,
			ImportModal: false
		};
	},
	computed: {
		pages() {
			if (
				this.pagination.rowsPerPage == null ||
				this.totalUsers == null ||
				this.totalUsers == 0
			) return 0;

			return Math.ceil(this.totalUsers / this.pagination.rowsPerPage);
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
					text: this.$vuetify.lang.t('$vuetify.Users.Breadcrumbs.Users'),
					disabled: true
				}
			]
		},
		headers() {
			return [
				{ text: this.$vuetify.lang.t('$vuetify.Users.Fields.Name'), value: "name", align: "left", width: "10%" },
				{ text: this.$vuetify.lang.t('$vuetify.Users.Fields.Email'), value: "email", align: "left", width:"20%" },
				{ text: this.$vuetify.lang.t('$vuetify.Users.Fields.Roles'), value: "rolesList", align: "left", width: "23%" },
				{ text: this.$vuetify.lang.t('$vuetify.CreatedAt'), value: "created_at", align: "left", width:"17%" },
				{ text: this.$vuetify.lang.t('$vuetify.Active'), value: "is_active", sortable: true, align: "left", width: "10%" },
				{ text: this.$vuetify.lang.t('$vuetify.Actions'), value: "actions", sortable: false, align: "left", width: "20%" }
			]
		},

		ActiveFilterOptions() {
			return [
				{text: this.$vuetify.lang.t('$vuetify.Yes'), value: '1'},
				{text: this.$vuetify.lang.t('$vuetify.No'), value: '0'}
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
					qEmail: this.EmailFilter,
					qActive: this.IsActiveFilter,
					qRemember: this.RememberFiltersCheck,
					gsTerm: this.gsTerm
				}
			};

			this.$axios
				.get(this.$URLs.USER_LIST, options)
				.then(res => {
					let d = res.data.data.data;

					this.totalUsers = res.data.data.total;

					this.users = [];

					this.users = d;

					if (this.users.length == d.length) {
						this.loading = false;
					}
				})
				.catch(e => {
					this.users = [];

					this.loading = false;
					this.$store.dispatch("serverError", e);
				});
		},

		editLink(item) {
			return '/admin/users/edit/' + item.id
		},

		toggleFilter() {
			this.filterOpened = !this.filterOpened;
		},

		fetchCookies() {
			let options = {
				params: {
					for: ["qName", "qEmail", "qRemember"]
				}
			};

			this.$axios
				.get(this.$URLs.USER_COOKIES, options)
				.then(res => {
					this.NameFilter = res.data.Users_qName;
					this.EmailFilter = res.data.Users_qEmail;
					this.IsActiveFilter = res.data.Users.qActive;
					this.RememberFiltersCheck =
						res.data.Users_qRemember == "true" ? true : false;
				})
				.catch(e => {
					this.$store.dispatch("serverError", e);
				});
		},

		setInitialFilters() {
			this.NameFilter =
				this.$cookie.get("Users_qName") != null
					? this.$cookie.get("Users_qName")
					: "";

			this.EmailFilter =
				this.$cookie.get("Users_qEmail") != null
					? this.$cookie.get("Users_qEmail")
					: "";

			this.IsActiveFilter =
				this.$cookie.get("Users_qActive") != null
					? this.$cookie.get("Users_qActive")
					: "";

			this.RememberFiltersCheck =
				this.$cookie.get("Users_qRemember") != null &&
				this.$cookie.get("Users_qRemember") == "true"
					? true
					: false;

			let sortby =
				this.$cookie.get("Users_sortBy") != null
					? this.$cookie.get("Users_sortBy")
					: this.pagination.sortBy;
			this.pagination.sortBy = [sortby]
			this.pagination.descending =
				this.$cookie.get("Users_descending") != null &&
				this.$cookie.get("Users_descending") == "true"
					? true
					: this.pagination.descending;
		},

		clearAndRefresh(datakey) {
			if (datakey == "NameFilter") this.NameFilter = "";

			if (datakey == "EmailFilter") this.EmailFilter = "";

			if (datakey == "IsActiveFilter") this.IsActiveFilter = "";

			this.fetchData();
		},

		reset() {
			this.NameFilter = "";
			this.EmailFilter = "";
			this.IsActiveFilter = "";
			this.RememberFiltersCheck = false;
			this.filterOpened = false;
			this.fetchData();
		},

		viewRecord(item) {
			this.ViewUser = item;
			this.$refs.ViewUserDialog.UserViewDialog = true;
		},

		activeInactive(status) {
			this.$Utils.activeInactive.call(this, status, this.$URLs.USER_ACT_DEACT);
		},

		triggerDeleteModal(item) {
			this.$Utils.triggerDeleteModal.call(this, this.$URLs.USER_LIST + "/" + item.id)
		},

		getAccessDetails() {
			this.$store.dispatch("showProgress", true);

			this.$axios
				.get(this.$URLs.USER_ACCESS)
				.then(response => {
					this.$store.dispatch("showProgress", false);

					this.hasAddAccess = response.data.data.canAdd;
					this.hasEditAccess = response.data.data.canEdit;
					this.hasDeleteAccess = response.data.data.canDelete;
					this.hasListingAccess = response.data.data.canViewList;
					this.hasActDeactAccess =
						response.data.data.canActivateDeactive;
					this.hasViewAccess = response.data.data.canView;
					this.hasImportAccess = response.data.data.canImport;
					this.hasExportAccess = response.data.data.canExport;
				})
				.catch(e => {
					this.$store.dispatch("serverError", e);
					this.$store.dispatch("showProgress", false);
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
		},
		ExportData() {
			this.ExportModal = false

			this.$store.dispatch("showExportProcess", true)

			let options = {
				qName: this.NameFilter,
				qEmail: this.EmailFilter,
				qActive: this.IsActiveFilter,
				qRemember: this.RememberFiltersCheck,
				gsTerm: this.gsTerm
			}

			this.$axios.post(this.$URLs.USER_EXPORT, options)
				.then((response) => {
					window.open(this.$ApiBaseUrl + '/storage/' + response.data.data, '_blank')
					this.$store.dispatch("showExportProcess", false);
					this.$Utils.unlinkExportedFile.call(this, response.data.data);
				}).catch((e) => {
					this.$store.dispatch("serverError", e);
					this.$store.dispatch("showExportProcess", false);
				});
				
		},
		ImportData() {
			this.ImportModal = false
		}

	},

	async created() {
		await this.$axios.get(this.$URLs.SANCTUM_CSRF)
		await this.$Utils.checkUserLoggedIn.call(this)
		await this.getAccessDetails();
		await this.setInitialFilters();
	},

	components: {
		"view-user": ViewUser,
		"delete-modal": DeleteModal,
		"unauthorized": Unauthorized,
		"breadcrumbs": Breadcrumbs,
		"page-title": PageTitle,
		"page-actions": PageActions,
		"pagination": Pagination,
		"search": Search,
		"export-users": ExportUsers,
		"import-users": ImportUsers
	}
};
</script>
