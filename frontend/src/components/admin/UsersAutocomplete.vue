<template>
    <v-autocomplete
        v-model="Users" 
        :items="items" 
        :loading="isLoading" 
        :search-input.sync="search"
        hide-no-data
        multiple
        chips
        hide-selected 
        item-text="name"
        item-value="id"
        :label="`${label != null ? label : 'Assign User'}`"
        placeholder="Start typing to Search"
        prepend-icon="fa-users" 
        return-object
        @change="selected"
    >
        <template v-slot:selection="data">
            <v-chip
                :selected="data.selected"
                close
                class="chip--select-multi"
                @input="remove(data.item)"
            >
                <v-avatar>
                <img :src="data.item.profile_pic">
                </v-avatar>
                {{ data.item.name }}
            </v-chip>
        </template>
        <template v-slot:item="data">
            <v-list-tile-avatar>
                <img :src="data.item.profile_pic">
            </v-list-tile-avatar>
            <v-list-tile-content>
                <v-list-tile-title v-html="data.item.name"></v-list-tile-title>
                <v-list-tile-sub-title v-if="data.item.rolesList.length > 0" >{{ data.item.rolesList.join(",") }}</v-list-tile-sub-title>
            </v-list-tile-content>    
        </template>
    </v-autocomplete>
</template>
<script>
export default {
    props: ["label"],
    data() {
        return {
            entries: [],
            isLoading: false,
            Users: null,
            search: null
        }
    },

    computed: {
        items() {
            console.log("Entries", this.entries)
            return _.merge(this.entries, this.Users)
        }
    },

    watch: {
        search(val) {
            if (this.isLoading) return

            if (val == null || val.length < 3) {
                return false
            }
            this.isLoading = true

            const options = {
                params : {
                    gsTerm: val
                }
            }
            this.$axios.get(this.$URLs.USER_LIST, options)
                .then((response) => {
                    this.isLoading = false
                    this.entries = response.data.data
                }).catch((error) => {
                    this.isLoading = false
                })
        }
    },

    methods: {
        remove (item) {
            this.Users = _.filter(this.Users, function(rec) {
                return rec.id != item.id
            })

            this.selected()
        },

        selected() {
            this.$emit("selected", this.Users)
        }
    }
}
</script>