export const Utils = {
    async activeInactive(status, action) {
        if (this.selected.length == 0) {

            this.$store.dispatch("showSnackbarMessage", {
                message: "Please select a row to perform the operation.",
                code: "0"
            });

            return false;
        }

        this.$store.dispatch("showProgress", true);

        let options = {
            IsActive: status,
            _method: "PUT",
            ids: this._.map(this.selected, "id")
        };


        await this.$axios({
            url: action,
            method: "POST",
            data: options
        }).then(response => {
            this._.forEach(this.selected, (obj) => {
                this._.set(obj, 'is_active', status == "1" ? true : false)
            })

            this.selected = [];

            this.$store.dispatch("showProgress", false);
            this.$store.dispatch("showSnackbarMessage", {
                message: response.data.message,
                code: "1"
            });
        }).catch(e => {
            this.$store.dispatch("serverError", e);
            this.$store.dispatch("showProgress", false);
        });
    },
    triggerDeleteModal(action) {
        this.DeleteItemUrl = action;
        this.DeleteItemModal = true;
    },
    setAuthorizationHeader() {
        this.$axios.defaults.headers.common['Authorization'] = this.$store.state.AuthToken
    },
    checkUserLoggedIn() {
        this.$store.dispatch("showProgress", true);

        return this.$axios.post(this.$URLs.MY_USER)
            .then((response) => {
                this.$store.dispatch("showProgress", false);
                this.$store.dispatch('setUserMeta', response.data.data)
            }).catch((e) => {
                console.log(e)
                this.$store.dispatch("showProgress", false);
                window.location.href = '/login?loggedOut=3&next=' + this.$route.fullPath
            });

    },

    async removeCurrentImage(id) {
        if (typeof id === undefined) id = this.$route.params.id

        let params = {
            '_method': 'DELETE'
        };

        this.$store.dispatch('showProgress', true)
        return await this.$axios({
            url: this.$URLs.USER_PICTURE_REMOVE + "/" + id,
            method: "DELETE",
            data: params
        }).then(response => {
            this.$store.dispatch('showProgress', false)
            this.$store.dispatch('showSnackbarMessage', { message: response.data.message, code: '1' })
            this.loadUserEditData()
        }).catch(e => {
            this.$store.dispatch('showProgress', false)
            this.$store.dispatch('serverError', e)
        });
    },

    resetFileSelection() {
        this.imageName = ''
        this.imageUrl = ''
        this.ProfilePicture = null
    },

    handleFileSelection(e) {
        this.ProfilePictureErrors = []

        const files = e.target.files

        if (files[0] !== undefined) {
            if (files[0].name.lastIndexOf('.') <= 0) {
                return
            }

            this.imageName = files[0].name

            let extArr = this.imageName.split('.')

            let ext = extArr[extArr.length - 1].toLowerCase()

            if (ext == 'png' || ext == 'jpg' || ext == 'jpeg' || ext == 'gif') {

                if (files[0].size <= 2048000) {
                    const fr = new FileReader;

                    fr.readAsDataURL(files[0])

                    fr.addEventListener('load', () => {
                        this.imageUrl = fr.result
                        this.ProfilePicture = files[0]
                    })
                } else {
                    this.ProfilePictureErrors.push('Maximum 2mb file size is allowed.')
                    this.resetFileSelection()
                }

            } else {
                this.ProfilePictureErrors.push('Please select a valid file.')
                this.resetFileSelection()
            }
        } else {
            this.resetFileSelection()
        }
    },
    eligibleForUnlink() {
        let allow = true
        let profilePic = this.PreviewCurrentImage;

        if (profilePic != null && profilePic != '') {
            profilePic = this.$store.getters.getFileUrlBaseName(profilePic)

            if (profilePic == 'not-available') {
                allow = false
            }
        }

        return allow;
    },

    isValidForm() {
        this.$v.$touch();

        const invalidFields = Object.keys(this.$v.$params).filter(
            fieldName => this.$v[fieldName].$invalid
        );

        console.log(invalidFields)

        return invalidFields.length == 0;
    },

    triggerFlash(type, message, hideFlash) {
        this.flashResponseType = type
        this.flashResponseMessage = message
        this.flashResponse = true;

        if (typeof hideFlash == 'undefined' || hideFlash == '1') {
            setTimeout(() => {
                this.flashResponse = false;
            }, 5000);
        }
    },

    setLocaleHeader() {
        this.$axios.defaults.headers.common['Curr_Lang'] = this.$vuetify.lang.current;
    },

    async switchLanguage(lang) {
        if (typeof lang == 'undefined' || lang == null || lang == '') {
            lang = 'en'
        }
        this.$vuetify.lang.current = lang
        this.$Utils.setLocaleHeader.call(this)
        this.$store.dispatch("showProgress", true)
        await this.$axios.post(this.$URLs.LANG_SWITCH + '/' + lang)
            .then((response) => {
                this.$store.dispatch("showProgress", false);
            }).catch((e) => {
                this.$store.dispatch("showProgress", false);
            })
    },

    goToTop() {
        this.$vuetify.goTo('#app', { duration: 300, easing: 'linear' })
    },

    unlinkExportedFile(file) {
        this.$axios.post(this.$URLs.EXPORT_UNLINK, { unlink: file });
    }
};