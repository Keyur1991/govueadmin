<template>
    <v-container fluid>
        <v-row dense>
            <v-col>
                <v-dialog scrollable v-model="dialog" max-width="900px">
                    <v-card flat>
                        <v-card-text class="pl-0 pr-0">
                            <v-tabs color="primary" light slider-color="primary">
                                <v-tab ripple @change="isSettingsTab()">Settings</v-tab>
                                <v-tab ripple @change="isProcessTab()">Import</v-tab>
                                <v-tab-item >
                                    <v-spacer></v-spacer>
                                    <v-card flat>
                                        <v-card-text>
                                            <v-row dense>
                                                <v-col md="4" sm="4">
                                                    <v-select outlined dense v-model="ImportFileFormat" :items="SupportedFormats" :label="`${$vuetify.lang.t('$vuetify.Format')}`"></v-select>
                                                </v-col>
                                                <v-col md="4" sm="4" v-if="ImportFileFormat == 'CSV'" >
                                                    <v-fade-transition>
                                                        <v-select outlined dense v-model="ImportFileCsvDelimeter" :items="Delimiters" :label="`${$vuetify.lang.t('$vuetify.Delimeter')}`"></v-select>
                                                    </v-fade-transition>
                                                </v-col>
                                                <v-col md="4" sm="4" v-if="ImportFileFormat == 'CSV'" >
                                                    <v-fade-transition>
                                                        <v-select outlined dense v-model="ImportFileCsvEncloser" :items="Enclosers" :label="`${$vuetify.lang.t('$vuetify.TextEncloser')}`"></v-select>
                                                    </v-fade-transition>
                                                </v-col>
                                            </v-row>

                                            <h4>{{ $vuetify.lang.t('$vuetify.SelectColumns') }}</h4>
                                            <v-row column style="height: 300px;">
                                                <v-col style="overflow:auto;">
                                                    <v-data-table v-model="SelectColumns" item-key="field" hide-default-footer show-select :headers="headers" :items="ImportColumns" :server-total-length="ImportColumns.length">
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
                                    </v-card>
                                </v-tab-item>
                                <v-tab-item> 
                                    <v-card flat>
                                        <v-card-text>
                                            <v-row dense>
                                                <v-col md="4" sm="4" class="mr-2">
                                                    <v-btn
                                                        color="primary"
                                                        outlined
                                                        @click="triggerFileInput">
                                                        {{ (ImportFileName != '') ? ImportFileName : 'Browse' }}
                                                        <v-icon right small>fa-folder</v-icon>
                                                    </v-btn>

                                                    <input type="file"
                                                        v-on:input="handleFileSelection"
                                                        ref="ImportFileRef"
                                                        style="display: none;"
                                                        name="ImportFileInput"
                                                    />
                                                </v-col>
                                            </v-row>
                                            <v-row dense>
                                                <v-col>
                                                    <v-alert v-if="ServerErrors.length > 0" :value="true" color="error" outline class="pa-1">
                                                        {{ ServerErrors.join(" ") }}
                                                    </v-alert>
                                                </v-col>
                                            </v-row>
                                        </v-card-text>
                                    </v-card>
                                </v-tab-item>
                            </v-tabs>
                        </v-card-text>
                        <v-card-actions>
                            <v-btn color="primary" v-if="SettingsTab" type="button" @click="save">{{ $vuetify.lang.t('$vuetify.SaveBtn') }}</v-btn>
                            <v-btn v-if="ProcessTab" type="button" color="primary"  @click="downloadSample">{{ $vuetify.lang.t('$vuetify.DownloadSampleBtn') }}</v-btn>
                            <v-btn color="primary" v-if="ProcessTab" type="button" @click="importData">{{ $vuetify.lang.t('$vuetify.ImportBtn') }}</v-btn>
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
        importModal: {
            type: Boolean,
            default: false
        }
    },
    data() {
        return {
            SettingsTab: true,
            ProcessTab: false,
            SelectColumns: [],
            ImportFileFormat: 'CSV',
            ImportFileCsvDelimeter: ';',
            ImportFileCsvEncloser: '~',
            ImportColumns: [],
            ImportFileName: '',
            SupportedFormats: ['CSV', 'XLS', 'XLSX'],
            PreviewColumns: [],
            ImportFile: null,
            ServerErrors: []
        }
    },

    computed: {
        dialog: {
            get() {
                if (this.importModal) this.getSettings()
                
                return this.importModal
            },

            set() {
                this.$emit('close')
            }
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
        isSettingsTab() {
            this.ProcessTab = false
            this.SettingsTab = true
        },

        isProcessTab() {
            this.SettingsTab = false
            this.ProcessTab = true
        },
        getSettings() {
            this.$store.dispatch("showProgress", true)

            this.$axios.get(this.$URLs.POSTS_IMPORT_SETTINGS)
                .then((response) => {
                    this.$store.dispatch("showProgress", false)
                    this.ImportColumns = response.data.data.columns
                    this.ImportFileFormat = response.data.data.importFileFormat
                    this.ImportFileCsvDelimeter = response.data.data.importFileCsvDelimeter
                    this.ImportFileCsvEncloser = response.data.data.importFileCsvEncloser
                    this.SelectColumns = this._.filter(this.ImportColumns, (item) => { if (item.isSelected) return this._.map(item, 'field') })
                }).catch((e) => {
                    this.$store.dispatch("serverError", e)
					this.$store.dispatch("showProgress", false)
                })
        },

        save() {
            this.$store.dispatch("showProgress", true)

            let options = {
                ImportFileFormat: this.ImportFileFormat,
                ImportFileCsvDelimeter: this.ImportFileCsvDelimeter,
                ImportFileCsvEncloser: this.ImportFileCsvEncloser,
                ImportColumns: this._.map(this.ImportColumns, (item) => { 
                    item.isSelected = this._.indexOf(this._.map(this.SelectColumns, 'field'), item.field) !== -1
                    return item
                })
            }

            return this.$axios.post(this.$URLs.POSTS_IMPORT_SETTINGS, options)
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

        triggerFileInput() {
            this.$refs.ImportFileRef.click()
        },

        handleFileSelection(e) {
            const files = e.target.files

            if (files[0] !== undefined) {
                if (files[0].name.lastIndexOf('.') <= 0) {
                    return
                }

                this.ImportFileName = files[0].name

                const fr = new FileReader;

                fr.readAsDataURL(files[0])

                this.ImportFile = files[0]
            } else {
                this.ImportFile = null
            }
        },

        downloadSample() {
            window.location.href = '/storage/posts_import.' + this.ImportFileFormat.toLowerCase();
        },

        importData() {
            if (this.ImportFile == null) {
                this.$store.dispatch("showSnackbarMessage", {message: this.$vuetify.lang.t('$vuetify.SelectFileMsg'), code: '0'})
                return false;
            }

            this.$store.dispatch('showProgress', true);
            let data = new FormData();

            data.append('Type', 'users');
            data.append('ImportFile', this.ImportFile)

            this.$axios.post(this.$URLs.POSTS_IMPORT, data)
                .then((response) => {
                    this.$store.dispatch('showProgress', false);
                }).catch((e) => {
                    this.$store.dispatch('showProgress', false);
                    
                    if (typeof e.response.status != 'undefined' && e.response.status == 422) {
                        if (typeof e.response.data.errors.ImportFile != 'undefined') {
                            this.ServerErrors = e.response.data.errors.ImportFile
                        }
                    } else {
                        this.$store.dispatch("serverError", e)
                    }
                })

        }
    }
}
</script>