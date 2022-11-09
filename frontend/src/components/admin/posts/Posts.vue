<template>
	<v-container fluid v-if="hasLoggedIn">
		<v-row dense justify="space-between" >
			<v-col md="10" sm="10" xs="12">
				<page-title>{{ $vuetify.lang.t('$vuetify.Posts.Listing.PageTitle') }}</page-title>
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
							<page-actions hide-back :has-add-access="hasAddAccess" :has-listing-access="hasListingAccess" :add-title="`${$vuetify.lang.t('$vuetify.Posts.Listing.NewBtn')}`" :has-import-access="hasImportAccess" :has-export-access="hasExportAccess" :has-act-deact-access="hasActDeactAccess" @filter="toggleFilter" @actinact="activeInactive" hide-act-deact @add="newPostModal" @export="ExportModal = true" @import="ImportModal = true"/>
						</v-col>
					</v-row>
					<v-row dense v-if="hasListingAccess">
						<v-col>
							<div class="text-xs-left">
	                            <v-fade-transition mode="out-in">
                                <v-chip
                                    v-if="TitleFilter != null && TitleFilter.trim() != ''"
                                    close
                                    @click.close="clearAndRefresh('TitleFilter')"
                                >{{$vuetify.lang.t('$vuetify.Posts.Fields.Title')}}: {{ TitleFilter }}</v-chip>
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
								:items="posts"
								:options.sync="pagination"
								hide-default-footer
								show-select
								:server-items-length="totalPosts"
								:loading="loading"
								class="elevation-2"
							>
								<v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>
								<template v-slot:item.actions="{ item }">
									<v-tooltip bottom v-if="hasEditAccess">
										<template v-slot:activator="{ on }">
											<v-btn
												v-on="on"
												color="primary"
												class="mr-2"
												@click="editRecord(item)"
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
												<v-icon small>fa-arrow-right</v-icon>
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
                                        v-model="TitleFilter"
                                        :label="`${$vuetify.lang.t('$vuetify.Posts.Fields.Title')}` "
                                        append-icon="fa-times"
                                        @click:append="TitleFilter = ''"
                                    ></v-text-field>

											<v-checkbox v-model="RememberFiltersCheck" :label="`${$vuetify.lang.t('$vuetify.RememberFilters')}`"  color="primary"></v-checkbox>
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
		<post-add-edit
			:post-item-modal="PostItemModal"
			:post-item="PostRecord"
			:title="PostModalTitle"
			:edit="PostEditMode"
			@close="PostItemModal = false"
			@reload="fetchData"
		></post-add-edit>

		<delete-modal
			:delete-item-modal="DeleteItemModal"
			:url="DeleteItemUrl"
			@close="DeleteItemModal = false"
			@reload="fetchData"
		></delete-modal>
		<export-post :export-modal="ExportModal" @export="ExportData" @close="ExportModal = false"></export-post>
		<import-post :import-modal="ImportModal" @import="ImportData" @close="ImportModal = false"></import-post>
	</v-container>
</template>

<script>
import Vue from "vue";
import VueCookie from "vue-cookie";

Vue.use(VueCookie);

var PostAddEdit = require("./PostAddEdit.vue").default;
var DeleteModal = require("./../DeleteModal.vue").default;
var Unauthorized = require("./../Unauthorized.vue").default;
var Breadcrumbs = require('../Breadcrumb.vue').default
var PageTitle = require('../PageTitle.vue').default
var PageActions = require('../PageActions.vue').default
var Pagination = require('../Pagination.vue').default
var Search = require('../Search.vue').default
var ExportPosts = require('./ExportPosts.vue').default;
var ImportPosts = require('./ImportPosts.vue').default;

export default {
	data() {
		return {
			totalPosts: 0,
			posts: [],
			loading: true,
			pagination: {
				page: 1,
				rowsPerPage: 10,
				sortBy: ["created_at"],
				descending: 1
			},
			filterOpened: false,
            TitleFilter: '',

			RememberFiltersCheck: false,
			PostRecord: {},
			PostModalTitle: "New Post",
			PostEditMode: false,
			PostDialog: false,
			DeleteItemUrl: "",
			DeleteItemModal: false,
			hasAddAccess: false,
			hasEditAccess: false,
			hasListingAccess: null,
			hasActDeactAccess: false,
			hasDeleteAccess: false,
			hasImportAccess: false,
			hasExportAccess: false,
			PostItemModal: false,
			selected: [],
			gsTerm: '',
			ImportModal: false,
			ExportModal: false
		};
	},
	computed: {
		pages () {
            if (this.pagination.rowsPerPage == null ||
                this.totalPosts == null || this.totalPosts == 0
            ) return 0

            return Math.ceil(this.totalPosts / this.pagination.rowsPerPage)
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
					text: this.$vuetify.lang.t('$vuetify.Posts.Breadcrumbs.Posts'),
					disabled: true
				}
			]
		},
		headers() {
			return [
				                { text: this.$vuetify.lang.t('$vuetify.Posts.Fields.Title'), value: "title", align: "left", width: "10%" },

				{ text: this.$vuetify.lang.t('$vuetify.CreatedAt'), value: "created_at", sortable: true, align: "left", width: "5%" },
				{ text: this.$vuetify.lang.t('$vuetify.Actions'), value: "actions", sortable: false, align: "left", width: "16%" }
			]
		},
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
					qRemember: this.RememberFiltersCheck,
					gsTerm: this.gsTerm,
                    qTitle: this.TitleFilter,

				}
			};

			this.$axios
				.get(this.$URLs.POSTS_LIST, options)
				.then(res => {
					let d = res.data.data.data;

					this.totalPosts = res.data.data.total;

					this.posts = [];

					this.posts = d;

					if (this.posts.length == d.length) {
						this.loading = false;
					}
				})
				.catch(e => {
					this.posts = [];

					this.loading = false;
					this.$store.dispatch("serverError", e);
				});
		},

		editRecord(item) {
			this.PostRecord = item;
			this.PostModalTitle = this.$vuetify.lang.t('$vuetify.Posts.Listing.EditPost');
			this.PostEditMode = true;
			this.PostItemModal = true;
		},

		deleteRecord(item) {},

		toggleFilter() {
			this.filterOpened = !this.filterOpened;
		},

		fetchCookies() {
			let options = {
				params: {
					for: ["qRemember","qTitle",
]
				}
			};

			var $this = this;
			this.$axios
				.get(this.$URLs.POSTS_COOKIES, options)
				.then(res => {
					console.log(res.data);
					this.$nextTick(() => {
                        this.TitleFilter = res.data.data.Posts_qTitle;

						this.RememberFiltersCheck =
							res.data.Posts_qRemember == "true" ? true : false;
					});
				})
				.catch(e => {
					this.$store.dispatch("serverError", e);
				});
		},

		setInitialFilters() {
            this.TitleFilter =
                this.$cookie.get("Posts_qTitle") != null
                    ? this.$cookie.get("Posts_qTitle")
                    : "";

			this.RememberFiltersCheck =
				this.$cookie.get("Posts_qRemember") != null &&
				this.$cookie.get("Posts_qRemember") == "true"
					? true
					: false;

			let sortBy =
				this.$cookie.get("Posts_sortBy") != null
					? this.$cookie.get("Posts_sortBy")
					: this.pagination.sortBy;
					
			this.pagination.sortBy = [sortBy]

			this.pagination.descending =
				this.$cookie.get("Posts_descending") != null &&
				this.$cookie.get("Posts_descending") == "true"
					? true
					: this.pagination.descending;
		},

		clearAndRefresh(datakey) {
            if (datakey == "TitleFilter") {
                this.TitleFilter = "";
            }

			this.fetchData();
		},

		reset() {
			this.RememberFiltersCheck = false;
			this.filterOpened = false;
			this.fetchData();
		},

		newPostModal() {
			this.PostRecord = {};
			this.PostModalTitle = this.$vuetify.lang.t('$vuetify.Posts.Listing.NewBtn');
			this.PostEditMode = false;
			this.PostItemModal = true;
		},

		activeInactive(status) {
			this.$Utils.activeInactive.call(this, status, this.$URLs.POSTS_ACT_DEACT);
		},

		triggerDeleteModal(item) {
			this.$Utils.triggerDeleteModal.call(this, this.$URLs.POSTS_LIST + "/" + item.id)
		},

		getAccessDetails() {
			this.$store.dispatch("showProgress", true);

			this.$axios
				.get(this.$URLs.POSTS_ACCESS)
				.then(response => {
					this.$store.dispatch("showProgress", false);

					this.hasAddAccess = response.data.data.canAdd;
					this.hasEditAccess = response.data.data.canEdit;
					this.hasDeleteAccess = response.data.data.canDelete;
					this.hasListingAccess = response.data.data.canViewList;
					this.hasActDeactAccess =
						response.data.data.canActivateDeactive;
					this.hasImportAccess = response.data.data.canImport;
					this.hasExportAccess = response.data.data.canExport;
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
		},
		ExportData() {
			this.ExportModal = false

			this.$store.dispatch("showExportProcess", true)

			let options = {
				gsTerm: this.gsTerm,
				                    qTitle: this.TitleFilter,

			}

			this.$axios.post(this.$URLs.POSTS_EXPORT, options)
				.then((response) => {
					window.location.href = '/storage/' + response.data.data;
					this.$store.dispatch("showExportProcess", false);
					this.$Utils.unlinkExportedFile.call(response.data.data);
				}).catch((e) => {
					this.$store.dispatch("serverError", e);
					this.$store.dispatch("showExportProcess", false);
				});
		},
		ImportData() {
			this.ImportModal = false
		}
	},

	components: {
		"post-add-edit": PostAddEdit,
		"delete-modal": DeleteModal,
		"unauthorized": Unauthorized,
		"breadcrumbs": Breadcrumbs,
		"page-title": PageTitle,
		"page-actions": PageActions,
		"pagination": Pagination,
		"search": Search,
		"export-post": ExportPosts,
		"import-post": ImportPosts
	},

	async created() {
		await this.$Utils.checkUserLoggedIn.call(this)
		await this.getAccessDetails()
		await this.setInitialFilters()
	}
};
</script>
