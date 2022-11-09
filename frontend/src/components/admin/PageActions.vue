<template>
    <div style="float:right;">
        <v-tooltip bottom v-if="hasAddAccess && (hideAdd == null || !hideAdd)">
            <template v-if="addAction != null" v-slot:activator="{ on }">
                <v-btn  v-on="on" color="primary" :to="addAction" icon text>
                    <v-icon small>fa-plus</v-icon>
                </v-btn>
            </template>
            <template v-else v-slot:activator="{ on }">
                <v-btn  v-on="on" color="primary" @click="addItem" icon text>
                    <v-icon small>fa-plus</v-icon>
                </v-btn>
            </template>
            <span>{{ addTitle }}</span>
        </v-tooltip>

        <v-tooltip bottom v-if="hasListingAccess && (hideFilter == null || !hideFilter)">
            <template v-slot:activator="{ on }">
                <v-btn v-on="on" exact-active-class color="primary" icon text @click="toggleFilter">
                    <v-icon small>fa-filter</v-icon>
                </v-btn>
            </template>
            <span>{{ $vuetify.lang.t('$vuetify.AdvanceFilter') }}</span>
            
        </v-tooltip>

        <v-tooltip bottom v-if="hasListingAccess && (hideMore == null || !hideMore)">
            <template v-slot:activator="{ on }">
                <v-menu v-on="on">
                    <template v-slot:activator="{ on }">
                        <v-btn v-on="on" color="primary" text icon>
                            <v-icon small>fa-cogs</v-icon>
                        </v-btn>
                    </template>
                    <v-list>
                        <v-list-item v-if="hasActDeactAccess && !hideActDeact" @click="activeInactive('1')">
                            <v-list-item-icon>
                                <v-icon small>fa-check</v-icon>
                            </v-list-item-icon>
                            <v-list-item-content>
                                <v-list-item-title> {{ $vuetify.lang.t('$vuetify.ActivateBtn') }}</v-list-item-title>
                            </v-list-item-content>
                        </v-list-item>
                        <v-list-item v-if="hasActDeactAccess && !hideActDeact" @click="activeInactive('0')">
                            <v-list-item-icon>
                                <v-icon small>fa-ban</v-icon>
                            </v-list-item-icon>
                            <v-list-item-content>
                                <v-list-item-title>
                                    {{ $vuetify.lang.t('$vuetify.DeactivateBtn') }}
                                </v-list-item-title>
                            </v-list-item-content>
                        </v-list-item>
                        <v-list-item v-if="hasExportAccess" @click="$emit('export')">
                            <v-list-item-icon>
                                <v-icon small>fa-download</v-icon>
                            </v-list-item-icon>
                            <v-list-item-content>
                                <v-list-item-title>{{ $vuetify.lang.t('$vuetify.ExportBtn') }}</v-list-item-title>
                            </v-list-item-content>
                        </v-list-item>
                        <v-list-item v-if="hasImportAccess" @click="$emit('import')">
                            <v-list-item-icon>
                                <v-icon small>fa-upload</v-icon>
                            </v-list-item-icon>
                            <v-list-item-content>
                                <v-list-item-title>{{ $vuetify.lang.t('$vuetify.ImportBtn') }}</v-list-item-title>
                            </v-list-item-content>
                        </v-list-item>
                    </v-list>
                </v-menu>
            </template>
            <span>{{ $vuetify.lang.t('$vuetify.MoreActions') }}</span>
            
        </v-tooltip>

        <v-tooltip bottom v-if="backLink != null && (hideBack == null || !hideBack)">
            <template v-slot:activator="{ on }">
                <v-btn v-on="on" color="primary" :to="backLink" text icon><v-icon small>fa-arrow-left</v-icon></v-btn>
            </template>
            <span>{{ $vuetify.lang.t('$vuetify.BackBtn') }}</span>
        </v-tooltip>
    </div>
</template>
<script>
export default {
    props: {
        hasAddAccess: Boolean,
        hasListingAccess: Boolean,
        addAction: String,
        addTitle: String,
        hideAdd: Boolean,
        hideMore: Boolean,
        hideFilter: Boolean,
        hideBack: Boolean,
        backLink: String,
        hasImportAccess: Boolean,
        hasExportAccess: Boolean,
        hasActDeactAccess: Boolean,
        hideActDeact: Boolean
    },
    methods: {
        activeInactive(value) {
            
            this.$emit('actinact', value)
        },
        addItem() {
            this.$emit('add')
        },
        toggleFilter() {
            this.$emit('filter')
        }
    }

}
</script>
