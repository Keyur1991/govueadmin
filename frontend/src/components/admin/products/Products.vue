<template>
	<v-container fluid v-if="hasLoggedIn">
		<v-row dense justify="space-between" >
			<v-col md="10" sm="10" xs="12">
				<page-title>{{ $vuetify.lang.t('$vuetify.Products.Listing.PageTitle') }}</page-title>
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
							<page-actions hide-back :has-add-access="hasAddAccess" :has-listing-access="hasListingAccess" :add-title="`${$vuetify.lang.t('$vuetify.Products.Listing.NewBtn')}`" :has-import-access="hasImportAccess" :has-export-access="hasExportAccess" :has-act-deact-access="hasActDeactAccess" @filter="toggleFilter" @actinact="activeInactive" hide-act-deact @add="newProductModal" @export="ExportModal = true" @import="ImportModal = true"/>
						</v-col>
					</v-row>
					<v-row dense v-if="hasListingAccess">
						<v-col>
							<div class="text-xs-left">
	                            <v-fade-transition mode="out-in">
                                <v-chip
                                    v-if="Name1Filter != null && Name1Filter.trim() != ''"
                                    close
                                    @click.close="clearAndRefresh('Name1Filter')"
                                >{{$vuetify.lang.t('$vuetify.Products.Fields.Name1')}}: {{ Name1Filter }}</v-chip>
                            </v-fade-transition>
                            <v-fade-transition mode="out-in">
                                <v-chip
                                    v-if="CategoryIdFilter != null && CategoryIdFilter.trim() != ''"
                                    close
                                    @click.close="clearAndRefresh('CategoryIdFilter')"
                                >{{$vuetify.lang.t('$vuetify.Products.Fields.CategoryId')}}: {{ CategoryIdFilter }}</v-chip>
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
								:items="products"
								:options.sync="pagination"
								hide-default-footer
								show-select
								:server-items-length="totalProducts"
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
                                        v-model="Name1Filter"
                                        :label="`${$vuetify.lang.t('$vuetify.Products.Fields.Name1')}` "
                                        append-icon="fa-times"
                                        @click:append="Name1Filter = ''"
                                    ></v-text-field>
                                    <v-text-field
                                        outlined
                                        dense
                                        v-model="CategoryIdFilter"
                                        :label="`${$vuetify.lang.t('$vuetify.Products.Fields.CategoryId')}` "
                                        append-icon="fa-times"
                                        @click:append="CategoryIdFilter = ''"
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
		<product-add-edit
			:product-item-modal="ProductItemModal"
			:product-item="ProductRecord"
			:title="ProductModalTitle"
			:edit="ProductEditMode"
			@close="ProductItemModal = false"
			@reload="fetchData"
		></product-add-edit>

		<delete-modal
			:delete-item-modal="DeleteItemModal"
			:url="DeleteItemUrl"
			@close="DeleteItemModal = false"
			@reload="fetchData"
		></delete-modal>
		<export-product :export-modal="ExportModal" @export="ExportData" @close="ExportModal = false"></export-product>
		<import-product :import-modal="ImportModal" @import="ImportData" @close="ImportModal = false"></import-product>
	</v-container>
</template>

<script>
import Vue from "vue";
import VueCookie from "vue-cookie";

Vue.use(VueCookie);

var ProductAddEdit = require("./ProductAddEdit.vue").default;
var DeleteModal = require("./../DeleteModal.vue").default;
var Unauthorized = require("./../Unauthorized.vue").default;
var Breadcrumbs = require('../Breadcrumb.vue').default
var PageTitle = require('../PageTitle.vue').default
var PageActions = require('../PageActions.vue').default
var Pagination = require('../Pagination.vue').default
var Search = require('../Search.vue').default
var ExportProducts = require('./ExportProducts.vue').default;
var ImportProducts = require('./ImportProducts.vue').default;

export default {
	data() {
		return {
			totalProducts: 0,
			products: [],
			loading: true,
			pagination: {
				page: 1,
				rowsPerPage: 10,
				sortBy: ["created_at"],
				descending: 1
			},
			filterOpened: false,
            Name1Filter: '',
            CategoryIdFilter: '',

			RememberFiltersCheck: false,
			ProductRecord: {},
			ProductModalTitle: "New Product",
			ProductEditMode: false,
			ProductDialog: false,
			DeleteItemUrl: "",
			DeleteItemModal: false,
			hasAddAccess: false,
			hasEditAccess: false,
			hasListingAccess: null,
			hasActDeactAccess: false,
			hasDeleteAccess: false,
			hasImportAccess: false,
			hasExportAccess: false,
			ProductItemModal: false,
			selected: [],
			gsTerm: '',
			ImportModal: false,
			ExportModal: false
		};
	},
	computed: {
		pages () {
            if (this.pagination.rowsPerPage == null ||
                this.totalProducts == null || this.totalProducts == 0
            ) return 0

            return Math.ceil(this.totalProducts / this.pagination.rowsPerPage)
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
					text: this.$vuetify.lang.t('$vuetify.Products.Breadcrumbs.Products'),
					disabled: true
				}
			]
		},
		headers() {
			return [
				                { text: this.$vuetify.lang.t('$vuetify.Products.Fields.Name1'), value: "name1", align: "left", width: "10%" },
                { text: this.$vuetify.lang.t('$vuetify.Products.Fields.Price'), value: "price", align: "left", width: "10%" },
                { text: this.$vuetify.lang.t('$vuetify.Products.Fields.CategoryId'), value: "category_id", align: "left", width: "10%" },

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
                    qName1: this.Name1Filter,
                    qCategoryId: this.CategoryIdFilter,

				}
			};

			this.$axios
				.get(this.$URLs.PRODUCTS_LIST, options)
				.then(res => {
					let d = res.data.data.data;

					this.totalProducts = res.data.data.total;

					this.products = [];

					this.products = d;

					if (this.products.length == d.length) {
						this.loading = false;
					}
				})
				.catch(e => {
					this.products = [];

					this.loading = false;
					this.$store.dispatch("serverError", e);
				});
		},

		editRecord(item) {
			this.ProductRecord = item;
			this.ProductModalTitle = this.$vuetify.lang.t('$vuetify.Products.Listing.EditProduct');
			this.ProductEditMode = true;
			this.ProductItemModal = true;
		},

		deleteRecord(item) {},

		toggleFilter() {
			this.filterOpened = !this.filterOpened;
		},

		fetchCookies() {
			let options = {
				params: {
					for: ["qRemember","qName1",
"qCategoryId",
]
				}
			};

			var $this = this;
			this.$axios
				.get(this.$URLs.PRODUCTS_COOKIES, options)
				.then(res => {
					console.log(res.data);
					this.$nextTick(() => {
                        this.Name1Filter = res.data.data.Products_qName1;
                        this.CategoryIdFilter = res.data.data.Products_qCategoryId;

						this.RememberFiltersCheck =
							res.data.Products_qRemember == "true" ? true : false;
					});
				})
				.catch(e => {
					this.$store.dispatch("serverError", e);
				});
		},

		setInitialFilters() {
            this.Name1Filter =
                this.$cookie.get("Products_qName1") != null
                    ? this.$cookie.get("Products_qName1")
                    : "";
            this.CategoryIdFilter =
                this.$cookie.get("Products_qCategoryId") != null
                    ? this.$cookie.get("Products_qCategoryId")
                    : "";

			this.RememberFiltersCheck =
				this.$cookie.get("Products_qRemember") != null &&
				this.$cookie.get("Products_qRemember") == "true"
					? true
					: false;

			let sortBy =
				this.$cookie.get("Products_sortBy") != null
					? this.$cookie.get("Products_sortBy")
					: this.pagination.sortBy;
					
			this.pagination.sortBy = [sortBy]

			this.pagination.descending =
				this.$cookie.get("Products_descending") != null &&
				this.$cookie.get("Products_descending") == "true"
					? true
					: this.pagination.descending;
		},

		clearAndRefresh(datakey) {
            if (datakey == "Name1Filter") {
                this.Name1Filter = "";
            }
            if (datakey == "CategoryIdFilter") {
                this.CategoryIdFilter = "";
            }

			this.fetchData();
		},

		reset() {
			this.RememberFiltersCheck = false;
			this.filterOpened = false;
			this.fetchData();
		},

		newProductModal() {
			this.ProductRecord = {};
			this.ProductModalTitle = this.$vuetify.lang.t('$vuetify.Products.Listing.NewBtn');
			this.ProductEditMode = false;
			this.ProductItemModal = true;
		},

		activeInactive(status) {
			this.$Utils.activeInactive.call(this, status, this.$URLs.PRODUCTS_ACT_DEACT);
		},

		triggerDeleteModal(item) {
			this.$Utils.triggerDeleteModal.call(this, this.$URLs.PRODUCTS_LIST + "/" + item.id)
		},

		getAccessDetails() {
			this.$store.dispatch("showProgress", true);

			this.$axios
				.get(this.$URLs.PRODUCTS_ACCESS)
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
				                    qName1: this.Name1Filter,
                    qCategoryId: this.CategoryIdFilter,

			}

			this.$axios.post(this.$URLs.PRODUCTS_EXPORT, options)
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

	components: {
		"product-add-edit": ProductAddEdit,
		"delete-modal": DeleteModal,
		"unauthorized": Unauthorized,
		"breadcrumbs": Breadcrumbs,
		"page-title": PageTitle,
		"page-actions": PageActions,
		"pagination": Pagination,
		"search": Search,
		"export-product": ExportProducts,
		"import-product": ImportProducts
	},

	async created() {
		await this.$Utils.checkUserLoggedIn.call(this)
		await this.getAccessDetails()
		await this.setInitialFilters()
	}
};
</script>
