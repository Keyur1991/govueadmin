<template>
    <v-container fluid>
        <v-row dense>
            <v-col>
                <v-dialog scrollable v-model="dialog" max-width="900px">
                    <v-card>
                        <v-card-title class="headline primary lighten-2 pa-3" primary-title>
                            <span class="headline white--text">Export Posts</span>
                        </v-card-title>
                        <v-card-text>
                            <v-row dense class="mt-2">
                                <v-col md="4" sm="4">
                                    <v-select outlined dense v-model="ExportFileFormat" :items="SupportedFormats" :label="`${$vuetify.lang.t('$vuetify.Format')}`"></v-select>
                                </v-col>
                                <v-col md="4" sm="4" v-if="ExportFileFormat == 'CSV'">
                                    <v-fade-transition>
                                        <v-select outlined dense v-model="ExportFileCsvDelimeter" :items="Delimiters" :label="`${$vuetify.lang.t('$vuetify.Delimeter')}`"></v-select>
                                    </v-fade-transition>
                                </v-col>
                                <v-flex md="4" sm="4" v-if="ExportFileFormat == 'CSV'">
                                    <v-fade-transition>
                                        <v-select outlined dense v-model="ExportFileCsvEncloser" :items="Enclosers" :label="`${$vuetify.lang.t('$vuetify.TextEncloser')}`"></v-select>
                                    </v-fade-transition>
                                </v-flex>
                            </v-row>    

                            <h4>{{ $vuetify.lang.t('$vuetify.SelectColumns') }}</h4>
                            <v-row column style="height: 300px;">
                                <v-col style="overflow:auto;">
                                    <v-data-table v-model="SelectColumns" item-key="field" hide-default-footer show-select :headers="headers" :items="ExportColumns" :server-total-length="ExportColumns.length">
                                        <template v-slot:item.caption="{ item }">
                                            <v-edit-dialog
                                            :return-value.sync="item.caption"
                                            lazy>{{ item.caption }}
                                                <template v-slot:input>
                                                    <v-text-field
                                                        outlined
                                                        dense
                                                        v-model="item.caption"
                                                        label="Edit"
                                                        single-line
                                                    ></v-text-field>
                                                </template>
                                            </v-edit-dialog>
                                        </template>
                                    </v-data-table>
                                </v-col>
                            </v-row>
                        </v-card-text>
                        <v-card-actions>
                            <v-btn color="primary" type="button"  @click="save">{{ $vuetify.lang.t('$vuetify.SaveBtn') }}</v-btn>
                            <v-btn color="primary" type="button"  @click="saveAndExport">{{ $vuetify.lang.t('$vuetify.SaveExpBtn') }}</v-btn>
                            <v-spacer></v-spacer>
                            <v-btn type="button" text @click="$emit('close')">{{ $vuetify.lang.t('$vuetify.CancelBtn') }}</v-btn>
                        </v-card-actions>
                    </v-card>
                </v-dialog>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
export default {
    props: {
        exportModal: {
            type: Boolean,
            default: false
        }
    },
    
    data: () => {
        return {
            ExportFileFormat: 'XLS',
            ExportFileCsvDelimeter: ';',
            ExportFileCsvEncloser: '~',
            ExportColumns: [],
            SelectColumns: []
        }
    },

    computed: {
        dialog: {
            get() {
                if (this.exportModal) this.getSettings()
                return this.exportModal
            },

            set() {
                this.$emit('close')
            }
        },

        SupportedFormats() {
            return this.$store.state.SupportedExportFormats
        },

        Delimiters() {
            return this.$store.state.Delimiters
        },

        Enclosers() {
            return this.$store.state.Enclosers
        },

        headers() {
            return [
                { text: this.$vuetify.lang.t('$vuetify.Caption'), value: "caption", align: "left", sortable: false, width: "45%"},
				{ text: this.$vuetify.lang.t('$vuetify.Column'), value: "field", align: "left", sortable: false, width: "45%"}
            ]
        }
    },

    methods: {
        getSettings() {
            this.$store.dispatch("showProgress", true)

            this.$axios.get(this.$URLs.POSTS_EXPORT_SETTINGS)
                .then((response) => {
                    this.$store.dispatch("showProgress", false)
                    this.ExportColumns = response.data.data.columns
                    this.ExportFileFormat = response.data.data.exportFileFormat
                    this.ExportFileCsvDelimeter = response.data.data.exportFileCsvDelimeter
                    this.ExportFileCsvEncloser = response.data.data.exportFileCsvEncloser
                    this.SelectColumns = this._.filter(this.ExportColumns, (item) => { if (item.isSelected) return this._.map(item, 'field') })
                }).catch((e) => {
                    this.$store.dispatch("serverError", e)
					this.$store.dispatch("showProgress", false)
                })
        },

        save() {
            this.$store.dispatch("showProgress", true)

            let options = {
                ExportFileFormat: this.ExportFileFormat,
                ExportFileCsvDelimeter: this.ExportFileCsvDelimeter,
                ExportFileCsvEncloser: this.ExportFileCsvEncloser,
                ExportColumns: this._.map(this.ExportColumns, (item) => { 
                    item.isSelected = this._.indexOf(this._.map(this.SelectColumns, 'field'), item.field) !== -1
                    return item
                })
            }

            return this.$axios.post(this.$URLs.POSTS_EXPORT_SETTINGS, options)
                .then((response) => {
                    this.$store.dispatch("showProgress", false)
                    this.$store.dispatch("showSnackbarMessage", {message: response.data.message, code: '1'})
                    return true
                }).catch((e) => {
                    this.$store.dispatch("serverError", e)
                    this.$store.dispatch("showProgress", false)
                    return false
                })
        },

        async saveAndExport() {
            let b = await this.save()

            if (b) this.$emit('export')
        }
    },

    created() {
        
    }
}
</script>