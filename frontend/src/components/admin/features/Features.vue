<template>
	<v-container fluid v-if="hasLoggedIn">
		<v-row dense justify="space-between" >
			<v-col md="10" sm="10" xs="12">
				<page-title>Features Management</page-title>
			</v-col>
			<v-col md="2" sm="2" xs="12" >
				<breadcrumbs :items="breadcrumbs" />
			</v-col>
		</v-row>
		<v-row dense>
			<v-col>
				<v-card >
					<v-row class="pa-2" dense v-if="hasListingAccess" >
						<v-col md="4" sm="4" xs="12" >
							<search @filter="globalSearch" />
						</v-col>
						<v-col md="8" sm="8" xs="12" >
							<page-actions hide-more hide-back :has-add-access="hasAddAccess" :has-listing-access="hasListingAccess" add-title="New Feature" @filter="toggleFilter"   @add="newFeatureModal" />
						</v-col>
					</v-row>

					<v-row class="pa-2" dense v-if="hasListingAccess && (FeatureNameFilter != '' || ModuleFilter != '')">
						<v-col>
							<div class="text-xs-left">
								<v-fade-transition mode="out-in">
									<v-chip
										v-if="ModuleFilter != null && ModuleFilter.trim() != ''"
										close
										@click:close="clearAndRefresh('ModuleFilter')"
									>Module: {{ ModuleFilter }}</v-chip>
								</v-fade-transition>

								<v-fade-transition mode="out-in">
									<v-chip
										v-if="FeatureNameFilter != null && FeatureNameFilter.trim() != ''"
										close
										@click:close="clearAndRefresh('FeatureNameFilter')"
									>Feature: {{ FeatureNameFilter }}</v-chip>
								</v-fade-transition>
							</div>
						</v-col>
					</v-row>

					<v-row class="pa-2" dense >
						<v-col>
							<v-data-table
								v-if="hasListingAccess"
								:headers="headers"
								:items="features"
								:options.sync="pagination"
								hide-default-footer
								:server-items-items="totalFeatures"
								:loading="loading"
								class="elevation-2"
							>
								<v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>
								<template v-slot:item.actions="{ item }">
									<v-tooltip bottom>
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
										<span>Edit</span>
									</v-tooltip>
									<v-tooltip bottom>
										<template v-slot:activator="{ on }">
											<v-btn
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
										<span>Delete</span>
									</v-tooltip>
								</template>
							</v-data-table>
						</v-col>
					</v-row>
					<v-row class="pa-2" dense>
						<v-col>
							<pagination v-if="hasListingAccess" :pages="pages" @update="updatePagination" />
						</v-col>
					</v-row>
				</v-card>
				<v-navigation-drawer v-model="filterOpened" absolute right width="250">
					<v-row dense v-if="hasListingAccess">
						<v-col v-if="filterOpened" >
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
												v-model="ModuleFilter"
												label="Module "
												append-icon="fa-times"
												@click:append="ModuleFilter = ''"
											></v-text-field>

											<v-text-field
												outlined
												dense
												v-model="FeatureNameFilter"
												label="Feature "
												append-icon="fa-times"
												@click:append="FeatureNameFilter = ''"
											></v-text-field>

											<v-checkbox v-model="RememberFiltersCheck" label="Remember Filters" color="primary"></v-checkbox>
										</v-col>
									</v-row>
								</v-card-text>
								<v-card-actions class="ml-3">
									<v-btn @click="fetchData" color="primary">
										<v-icon small>fa-search</v-icon>&nbsp;Search
									</v-btn>

									<v-btn color="secondary" @click="reset" depressed>
										<v-icon small>fa-undo</v-icon>&nbsp;Reset
									</v-btn>
								</v-card-actions>
							</v-card>
						</v-col>
					</v-row>
				</v-navigation-drawer>
			</v-col>
		</v-row>
		
		<feature-add-edit
			:feature-item-modal="FeatureItemModal"
			:feature-item="FeatureRecord"
			:title="FeatureModalTitle"
			:edit="FeatureEditMode"
			@close="FeatureItemModal = false"
			@reload="fetchData"
		></feature-add-edit>

		<delete-modal
			:delete-item-modal="DeleteItemModal"
			:url="DeleteItemUrl"
			@close="DeleteItemModal = false"
			@reload="fetchData"
		></delete-modal>
	</v-container>
</template>

<script>
import Vue from "vue"
import VueCookie from "vue-cookie";

Vue.use(VueCookie);

var FeatureAddEdit = require("./FeatureAddEdit.vue").default;
var DeleteModal = require("./../DeleteModal.vue").default;
var Breadcrumbs = require('../Breadcrumb.vue').default
var PageTitle = require('../PageTitle.vue').default
var PageActions = require('../PageActions.vue').default
var Pagination = require('../Pagination.vue').default
var Search = require('../Search.vue').default
export default {
	data() {
		return {
			breadcrumbs: [
				{
					text: "Home",
					action: "/admin/dashboard",
					disabled: false
				},
				{
					text: "Features",
					disabled: true
				}
			],
			totalFeatures: 0,
			features: [],
			loading: true,
			pagination: {
				page: 1,
				rowsPerPage: 10,
				sortBy: ["created_at"],
				descending: 1
			},
			headers: [
				{
					text: "Module ",
					value: "module",
					align: "left",
					width: "17%",
					sortable: true,
				},
				{ 
					text: "Feature ", 
					value: "feature_name", 
					align: "left",
					sortable: true, 
				},
				{
					text: "Created At ",
					value: "created_at",
					align: "left",
					width: "17%",
					sortable: true,
				},
				{
					text: "Actions",
					value: "actions",
					sortable: false,
					width: "20%",
				}
			],
			filterOpened: false,
			ModuleFilter: "",
			FeatureNameFilter: "",
			RememberFiltersCheck: false,
			FeatureRecord: {},
			FeatureModalTitle: "New Feature",
			FeatureEditMode: false,
			FeatureItemModal: false,
			DeleteItemUrl: "",
            DeleteItemModal: false,
            hasAddAccess: false,
			hasEditAccess: false,
			hasListingAccess: null,
			hasActDeactAccess: false,
			hasDeleteAccess: false,
			gsTerm: ''
		};
	},

	computed: {
		pages() {
			if (
				this.pagination.rowsPerPage == null ||
				this.totalFeatures == null ||
				this.totalFeatures == 0
			)
				return 0;

			return Math.ceil(this.totalFeatures / this.pagination.rowsPerPage);
		},

		hasLoggedIn() {
			return this.$store.state.userHasLoggedIn
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
					qModule: this.ModuleFilter,
					qFeatureName: this.FeatureNameFilter,
					qRemember: this.RememberFiltersCheck,
					gsTerm: this.gsTerm
				}
			};

			this.$axios
				.get(this.$URLs.FEATURE_LIST, options)
				.then(res => {
					let d = res.data.data.data;

					this.totalFeatures = res.data.data.total;

					this.features = [];

					this.features = d;

					if (this.features.length == d.length) {
						this.loading = false;
					}
				})
				.catch(e => {
					this.features = [];
					this.loading = false;
					this.$store.dispatch("serverError", e);
				});
		},

		editRecord(item) {
			this.FeatureRecord = item;
			this.FeatureModalTitle = "Edit Feature";
			this.FeatureEditMode = true;
			this.FeatureItemModal = true;
		},

		deleteRecord(item) {},

		toggleFilter() {
			this.filterOpened = !this.filterOpened;
		},

		fetchCookies() {
			let options = {
				params: {
					for: ["qModule", "qFeatureName", "qRemember"]
				}
			};

			this.$axios
				.get(this.$URLs.FEATURE_COOKIES, options)
				.then(res => {
					console.log(res.data);
					this.$nextTick(() => {
						this.ModuleFilter = res.data.data.Features_qModule;
						this.FeatureNameFilter =
							res.data.data.Features_qFeatureName;

						this.RememberFiltersCheck =
							res.data.data.Features_qRemember == "true"
								? true
								: false;
					});
				})
				.catch(e => {
					this.$store.dispatch("serverError", e);
				});
		},

		setInitialFilters() {
			this.ModuleFilter =
				this.$cookie.get("Features_qModule") != null
					? this.$cookie.get("Features_qModule")
					: "";
			this.FeatureNameFilter =
				this.$cookie.get("Features_qFeatureName") != null
					? this.$cookie.get("Features_qFeatureName")
					: "";

			this.RememberFiltersCheck =
				this.$cookie.get("Features_qRemember") != null &&
				this.$cookie.get("Features_qRemember") == "true"
					? true
					: false;

			let sortby =
				this.$cookie.get("Features_sortBy") != null
					? this.$cookie.get("Features_sortBy")
					: this.pagination.sortBy;
			
			this.pagination.sortBy = [sortby]

			this.pagination.descending =
				this.$cookie.get("Features_descending") != null &&
				this.$cookie.get("Features_descending") == "true"
					? true
					: this.pagination.descending;
		},

		clearAndRefresh(datakey) {
			if (datakey == "ModuleFilter") {
				this.ModuleFilter = "";
			}

			if (datakey == "FeatureNameFilter") {
				this.FeatureNameFilter = "";
			}

			this.fetchData();
		},

		reset() {
			this.ModuleFilter = "";
			this.FeatureNameFilter = "";
			this.RememberFiltersCheck = false;
			this.filterOpened = false;
			this.fetchData();
		},

		newFeatureModal() {
			this.FeatureRecord = {};
			this.FeatureModalTitle = "New Feature";
			this.FeatureEditMode = false;
			this.FeatureItemModal = true;
		},

		triggerDeleteModal(item) {
			this.DeleteItemUrl = this.$URLs.FEATURE_LIST + "/" + item.id;
			this.DeleteItemModal = true;
		},

		getAccessDetails() {
			this.$store.dispatch("showProgress", true);

			this.$axios
				.get(this.$URLs.FEATURE_ACCESS)
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
		"feature-add-edit": FeatureAddEdit,
		"delete-modal": DeleteModal,
        "breadcrumbs": Breadcrumbs,
		"page-title": PageTitle,
		"page-actions": PageActions,
		"pagination": Pagination,
		"search": Search
	},

	async created() {
		await this.$axios.get(this.$URLs.SANCTUM_CSRF)
		await this.$Utils.checkUserLoggedIn.call(this)
		await this.getAccessDetails();
		await this.setInitialFilters();
		await this.fetchData()
	}
};
</script>
